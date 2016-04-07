/*
 * MumbleDJ
 * By Matthieu Grieger
 * main.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package main

import (
	"log"
	"os"

	"github.com/codegangsta/cli"
	"github.com/matthieugrieger/mumbledj/bot"
	"github.com/matthieugrieger/mumbledj/commands"
)

// DJ is a global variable that holds various details about the bot's state.
var DJ *bot.MumbleDJ

// Warn is a global logger that logs warn messages.
var Warn *log.Logger

// Error is a global logger that logs error messages.
var Error *log.Logger

// Info is a global logger that logs info messages.
var Info *log.Logger

func main() {
	DJ = bot.NewMumbleDJ()

	app := cli.NewApp()
	app.Name = "MumbleDJ"
	app.Usage = "A Mumble bot that plays audio from various media sites."
	app.Version = "3.0.0-alpha"
	DJ.Version = app.Version

	// TODO: Allow user to redirect log output.
	Warn = log.New(os.Stdout, "MumbleDJ WARN:", 0)
	Error = log.New(os.Stdout, "MumbleDJ ERROR:", 0)
	Info = log.New(os.Stdout, "MumbleDJ INFO:", 0)

	// Injection into sub-packages.
	commands.DJ = DJ
	commands.Warn = Warn
	commands.Error = Error
	commands.Info = Info

	bot.Warn = Warn
	bot.Error = Error
	bot.Info = Info

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "server, s",
			Value:       "127.0.0.1",
			Usage:       "address of Mumble server to connect to",
			Destination: &DJ.BotConfig.Connection.Address,
		},
		cli.StringFlag{
			Name:        "port, o",
			Value:       "64738",
			Usage:       "port of Mumble server to connect to",
			Destination: &DJ.BotConfig.Connection.Port,
		},
		cli.StringFlag{
			Name:        "username, u",
			Value:       "MumbleDJ",
			Usage:       "username for the bot",
			Destination: &DJ.BotConfig.Connection.Username,
		},
		cli.StringFlag{
			Name:        "password, p",
			Value:       "",
			Usage:       "password for the Mumble server",
			Destination: &DJ.BotConfig.Connection.Password,
		},
		cli.StringFlag{
			Name:        "channel, c",
			Value:       "",
			Usage:       "channel the bot enters after connecting to the Mumble server",
			Destination: &DJ.BotConfig.General.DefaultChannel,
		},
		cli.StringFlag{
			Name:        "cert, e",
			Value:       "",
			Usage:       "path to PEM certificate",
			Destination: &DJ.BotConfig.Connection.Cert,
		},
		cli.StringFlag{
			Name:        "key, k",
			Value:       "",
			Usage:       "path to PEM key",
			Destination: &DJ.BotConfig.Connection.Key,
		},
		cli.StringFlag{
			Name:        "accesstokens, a",
			Value:       "",
			Usage:       "list of access tokens separated by spaces",
			Destination: &DJ.BotConfig.Connection.AccessTokens,
		},
		cli.BoolFlag{
			Name:        "insecure, i",
			Usage:       "if present, the bot will not check Mumble certs for consistency",
			Destination: &DJ.BotConfig.Connection.Insecure,
		},
	}
	app.Run(os.Args)

	if err := DJ.Connect(); err != nil {
		Error.Fatalf("\nA fatal error occurred: %s", err.Error())
	}
}
