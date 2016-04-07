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

// DJ is an injected MumbleDJ struct.
var DJ *bot.MumbleDJ

// Warn is an injected logger that logs warn messages.
var Warn *log.Logger

// Error is an injected logger that logs error messages.
var Error *log.Logger

// Info is an injected logger that logs info messages.
var Info *log.Logger
