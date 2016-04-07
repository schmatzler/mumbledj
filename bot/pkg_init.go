/*
 * MumbleDJ
 * By Matthieu Grieger
 * bot/pkg_init.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package bot

import "log"

// Warn is an injected logger that logs warn messages.
var Warn *log.Logger

// Error is an injected logger that logs error messages.
var Error *log.Logger

// Info is an injected logger that logs info messages.
var Info *log.Logger
