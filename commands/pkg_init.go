/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/pkg_init.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package commands

import (
	"log"

	"github.com/matthieugrieger/mumbledj/bot"
)

// DJ is an injected MumbleDJ struct. It is initialized here in order to
// prevent init panics. After init is complete, the DJ var is replaced with
// the injected one.
var DJ = bot.NewMumbleDJ()

// Warn is an injected logger that logs warn messages.
var Warn *log.Logger

// Error is an injected logger that logs error messages.
var Error *log.Logger

// Info is an injected logger that logs info messages.
var Info *log.Logger

func init() {
	DJ.RegisterCommand(new(VersionCommand))
}
