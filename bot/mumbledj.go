/*
 * MumbleDJ
 * By Matthieu Grieger
 * bot/mumbledj.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package bot

import (
	"crypto/tls"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/layeh/gumble/gumble"
	"github.com/layeh/gumble/gumbleffmpeg"
	"github.com/layeh/gumble/gumbleutil"
	"github.com/matthieugrieger/mumbledj/interfaces"
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
	Commands     []interfaces.Command
	KeepAlive    chan bool
	Version      string
	Volume       float32
	YouTubeDL    *YouTubeDL
}

// DJ is a struct that keeps track of all aspects of MumbleDJ's environment.
var DJ *MumbleDJ

// NewMumbleDJ initializes and returns a MumbleDJ type.
func NewMumbleDJ() *MumbleDJ {
	dj := new(MumbleDJ)

	dj.Commands = make([]interfaces.Command, 0)

	// TODO: Load from config file if necessary.
	dj.BotConfig = NewConfig()
	dj.Queue = NewQueue()
	dj.Cache = NewCache()
	dj.Skips = NewSkipTracker()

	if err := dj.CheckDependencies(); err != nil {
		Error.Fatalln(err.Error())
	}

	return dj
}

// OnConnect event. First moves MumbleDJ into the default channel if one exists.
// The configuration is loaded and the audio stream is initialized.
func (dj *MumbleDJ) OnConnect(e *gumble.ConnectEvent) {
	if dj.BotConfig.General.DefaultChannel != "" {
		defaultChannel := strings.Split(dj.BotConfig.General.DefaultChannel, "/")
		dj.Client.Self.Move(dj.Client.Channels.Find(defaultChannel...))
	}

	dj.AudioStream = nil
	dj.Volume = dj.BotConfig.Volume.Default

	dj.Client.Self.SetComment(dj.BotConfig.General.DefaultComment)

	if dj.BotConfig.Cache.Enabled {
		Info.Println("Caching enabled.")
		dj.Cache.UpdateStatistics()
		go dj.Cache.CleanPeriodically()
	} else {
		Info.Println("Caching disabled.")
	}
}

// OnDisconnect event. Terminates MumbleDJ process or retries connection if
// automatic connection retries are enabled.
func (dj *MumbleDJ) OnDisconnect(e *gumble.DisconnectEvent) {
	if dj.BotConfig.Connection.RetryEnabled &&
		(e.Type == gumble.DisconnectError || e.Type == gumble.DisconnectKicked) {
		Warn.Printf("Disconnected from server. Retrying connection every %d seconds %d times.\n",
			dj.BotConfig.Connection.RetryInterval,
			dj.BotConfig.Connection.RetryAttempts)

		success := false
		for retries := 0; retries < dj.BotConfig.Connection.RetryAttempts; retries++ {
			Info.Println("Retrying connection...")
			if err := dj.Client.Connect(); err == nil {
				Info.Println("Successfully reconnected to the server!")
				success = true
				break
			}
			time.Sleep(time.Duration(dj.BotConfig.Connection.RetryInterval) * time.Second)
		}
		if !success {
			dj.KeepAlive <- true
			Error.Fatalln("Could not reconnect to server. Exiting...")
		}
	} else {
		dj.KeepAlive <- true
		Error.Fatalln("Disconnected from server. No reconnect attempts will be made.")
	}
}

// OnTextMessage event. Checks for command prefix and passes it to the Commander
// if it exists. Ignores the incoming message otherwise.
func (dj *MumbleDJ) OnTextMessage(e *gumble.TextMessageEvent) {
	plainMessage := gumbleutil.PlainText(&e.TextMessage)
	if len(plainMessage) != 0 {
		if plainMessage[0] == dj.BotConfig.General.CommandPrefix[0] &&
			plainMessage != dj.BotConfig.General.CommandPrefix {
			message, isPrivateMessage, err := dj.FindAndExecuteCommand(e.Sender, plainMessage[1:])
			if err != nil {
				Warn.Printf("Sending error message (%s) to %s...\n", err.Error(), e.Sender.Name)
				dj.SendPrivateMessage(e.Sender, fmt.Sprintf("<b>Error:</b> %s.", err.Error()))
			} else {
				if isPrivateMessage {
					Info.Printf("Sending private message to %s...\n", e.Sender.Name)
					dj.SendPrivateMessage(e.Sender, message)
				} else {
					Info.Printf("Sending message to %s...\n", dj.Client.Self.Channel.Name)
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
		Info.Printf("%s has disconnected or changed channels, updating skip trackers...\n", e.User.Name)
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

// IsAdmin checks whether a particular Mumble user is a MumbleDJ admin.
// Returns true if the user is an admin, and false otherwise.
func (dj *MumbleDJ) IsAdmin(user *gumble.User) bool {
	// TODO: This currently hangs. Need to figure out why this is happening.
	userGroups := <-gumbleutil.UserGroups(dj.Client, user, dj.Client.Self.Channel)
	for _, userGroup := range userGroups {
		if userGroup == dj.BotConfig.Permissions.UserGroup {
			return true
		}
	}
	return false
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
	//dj.Client.Attach(gumbleutil.AutoBitrate)

	Info.Printf("Attempting connection to %s:%s...\n", dj.BotConfig.Connection.Address, dj.BotConfig.Connection.Port)
	if err := dj.Client.Connect(); err != nil {
		return err
	}

	Info.Println("Connected to server!")

	return nil
}

// FindAndExecuteCommand attempts to find a reference to a command in an
// incoming message. If found, the command is executed and the resulting
// message/error is returned.
func (dj *MumbleDJ) FindAndExecuteCommand(user *gumble.User, message string) (string, bool, error) {
	command, err := dj.findCommand(message)
	if err != nil {
		return "", true, errors.New("No command was found in this message")
	}
	return dj.executeCommand(user, message, command)
}

func (dj *MumbleDJ) findCommand(message string) (interfaces.Command, error) {
	var possibleCommand string
	if strings.Contains(message, " ") {
		possibleCommand = strings.ToLower(message[:strings.Index(message, " ")])
	} else {
		possibleCommand = strings.ToLower(message)
	}
	for _, command := range dj.Commands {
		for _, alias := range command.Aliases() {
			if possibleCommand == alias {
				return command, nil
			}
		}
	}
	return nil, errors.New("No command was found in this message")
}

func (dj *MumbleDJ) executeCommand(user *gumble.User, message string, command interfaces.Command) (string, bool, error) {
	canExecute := false
	if dj.BotConfig.Permissions.Enabled && command.IsAdminCommand() {
		canExecute = dj.IsAdmin(user)
	} else {
		canExecute = true
	}

	if canExecute {
		return command.Execute(user, strings.Split(message, " ")[1:]...)
	}
	return "", true, errors.New("You do not have permission to execute this command")
}
