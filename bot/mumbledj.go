/*
 * MumbleDJ
 * By Matthieu Grieger
 * bot/mumbledj.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package bot

import (
	"crypto/tls"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/layeh/gumble/gumble"
	"github.com/layeh/gumble/gumbleffmpeg"
	"github.com/layeh/gumble/gumbleutil"
)

// MumbleDJ is a struct that keeps track of all aspects of the bot's state.
type MumbleDJ struct {
	Client       *gumble.Client
	GumbleConfig *gumble.Config
	AudioStream  *gumbleffmpeg.Stream
	BotConfig    *Config
	Queue        *Queue
	Cache        *Cache
	Skips        *SkipTracker
	Commander    *Commander
	Log          *log.Logger
	KeepAlive    chan bool
}

// DJ is a struct that keeps track of all aspects of MumbleDJ's environment.
var DJ *MumbleDJ

// NewMumbleDJ initializes and returns a MumbleDJ type.
func NewMumbleDJ() *MumbleDJ {
	dj := new(MumbleDJ)

	dj.Commander = NewCommander()

	// TODO: Load from config file if necessary.
	dj.BotConfig = NewConfig()
	dj.Queue = NewQueue()
	dj.Cache = NewCache()
	dj.Skips = NewSkipTracker()
	// TODO: Allow for redirection of log output.
	dj.Log = log.New(os.Stderr, "MumbleDJ", 0)

	if err := dj.CheckDependencies(); err != nil {
		dj.Log.Fatalln(err.Error())
	}

	return dj
}

// OnConnect event. First moves MumbleDJ into the default channel if one exists.
// The configuration is loaded and the audio stream is initialized.
func (dj *MumbleDJ) OnConnect(e *gumble.ConnectEvent) {
	defaultChannel := strings.Split(dj.BotConfig.General.DefaultChannel, "/")
	dj.Client.Self.Move(dj.Client.Channels.Find(defaultChannel...))

	dj.AudioStream = nil
	dj.AudioStream.Volume = dj.BotConfig.Volume.Default

	dj.Client.Self.SetComment(dj.BotConfig.General.DefaultComment)

	if dj.BotConfig.Cache.Enabled {
		dj.Cache.UpdateStatistics()
		go dj.Cache.CleanPeriodically()
	}
}

// OnDisconnect event. Terminates MumbleDJ process or retries connection if
// automatic connection retries are enabled.
func (dj *MumbleDJ) OnDisconnect(e *gumble.DisconnectEvent) {
	if dj.BotConfig.Connection.RetryEnabled &&
		(e.Type == gumble.DisconnectError || e.Type == gumble.DisconnectKicked) {
		dj.Log.Printf("Disconnected from server. Retrying connection every %d seconds %d times.\n",
			dj.BotConfig.Connection.RetryInterval,
			dj.BotConfig.Connection.RetryAttempts)

		success := false
		for retries := 0; retries < dj.BotConfig.Connection.RetryAttempts; retries++ {
			dj.Log.Println("Retrying connection...")
			if err := dj.Client.Connect(); err == nil {
				dj.Log.Println("Successfully reconnected to the server!")
				success = true
				break
			}
			time.Sleep(time.Duration(dj.BotConfig.Connection.RetryInterval) * time.Second)
		}
		if !success {
			dj.KeepAlive <- true
			dj.Log.Fatalln("Could not reconnect to server. Exiting...")
		}
	} else {
		dj.KeepAlive <- true
		dj.Log.Fatalln("Disconnected from server. No reconnect attempts will be made.")
	}
}

// OnTextMessage event. Checks for command prefix and passes it to the Commander
// if it exists. Ignores the incoming message otherwise.
func (dj *MumbleDJ) OnTextMessage(e *gumble.TextMessageEvent) {
	plainMessage := gumbleutil.PlainText(&e.TextMessage)
	if len(plainMessage) != 0 {
		if plainMessage[0] == dj.BotConfig.General.CommandPrefix[0] &&
			plainMessage != dj.BotConfig.General.CommandPrefix {
			message, isPrivateMessage, err := dj.Commander.FindAndExecuteCommand(e.Sender, plainMessage[1:])
			if err != nil {
				dj.SendPrivateMessage(e.Sender, fmt.Sprintf("An error occurred while executing your command: %s", err.Error()))
			} else {
				if isPrivateMessage {
					dj.SendPrivateMessage(e.Sender, message)
				} else {
					dj.Client.Self.Channel.Send(message, false)
				}
			}
		}
	}
}

// OnUserChange event. Checks UserChange type and adjusts skip trackers to
// reflect the current status of the users on the server.
func (dj *MumbleDJ) OnUserChange(e *gumble.UserChangeEvent) {
	if e.Type.Has(gumble.UserChangeDisconnected) || e.Type.Has(gumble.UserChangeChannel) {
		dj.Skips.RemoveTrackSkip(e.User)
		dj.Skips.RemovePlaylistSkip(e.User)
	}
}

// SendPrivateMessage sends a private message to the specified user. This method
// verifies that the targeted user is still present in the server before attempting
// to send the message.
func (dj *MumbleDJ) SendPrivateMessage(user *gumble.User, message string) {
	if targetUser := dj.Client.Self.Channel.Users.Find(user.Name); targetUser != nil {
		targetUser.Send(message)
	}
}

// CheckDependencies checks whether or not the dependencies for MumbleDJ
// (most notably youtube-dl) are installed and executable. Returns nil if
// no dependencies are misconfigured/missing, returns an error otherwise
// describing the issue.
func (dj *MumbleDJ) CheckDependencies() error {
	return nil
}

// Connect starts the process for connecting to a Mumble server.
func (dj *MumbleDJ) Connect() error {
	// Create Gumble config.
	dj.GumbleConfig = &gumble.Config{
		Username: dj.BotConfig.Connection.Username,
		Password: dj.BotConfig.Connection.Password,
		Address:  dj.BotConfig.Connection.Address + ":" + dj.BotConfig.Connection.Port,
		Tokens:   strings.Split(dj.BotConfig.Connection.AccessTokens, " "),
	}

	// Create Gumble client.
	dj.Client = gumble.NewClient(dj.GumbleConfig)

	// Initialize key pair if needed.
	dj.GumbleConfig.TLSConfig.InsecureSkipVerify = true
	if !dj.BotConfig.Connection.Insecure {
		gumbleutil.CertificateLockFile(dj.Client, fmt.Sprintf("%s/.mumbledjcert.lock", os.Getenv("HOME")))
	}
	if dj.BotConfig.Connection.Cert != "" {
		if dj.BotConfig.Connection.Key != "" {
			dj.BotConfig.Connection.Key = dj.BotConfig.Connection.Cert
		}

		if certificate, err := tls.LoadX509KeyPair(dj.BotConfig.Connection.Cert, dj.BotConfig.Connection.Key); err == nil {
			dj.GumbleConfig.TLSConfig.Certificates = append(dj.GumbleConfig.TLSConfig.Certificates, certificate)
		} else {
			return err
		}
	}

	dj.Client.Attach(gumbleutil.Listener{
		Connect:     dj.OnConnect,
		Disconnect:  dj.OnDisconnect,
		TextMessage: dj.OnTextMessage,
		UserChange:  dj.OnUserChange,
	})
	dj.Client.Attach(gumbleutil.AutoBitrate)

	if err := dj.Client.Connect(); err != nil {
		return err
	}

	<-dj.KeepAlive

	return nil
}
