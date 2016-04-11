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
	"github.com/matthieugrieger/mumbledj/interfaces"
)

// DJ is an injected MumbleDJ struct.
var DJ *bot.MumbleDJ

// Warn is an injected logger that logs warn messages.
var Warn *log.Logger

// Error is an injected logger that logs error messages.
var Error *log.Logger

// Info is an injected logger that logs info messages.
var Info *log.Logger

// Commands is a slice of all enabled commands.
var Commands []interfaces.Command

func init() {
	Commands = []interfaces.Command{
		new(CacheSizeCommand),
		new(CurrentTrackCommand),
		new(ForceSkipCommand),
		new(ForceSkipPlaylistCommand),
		new(HelpCommand),
		new(KillCommand),
		new(ListTracksCommand),
		new(MoveCommand),
		new(NextTrackCommand),
		new(NumCachedCommand),
		new(NumTracksCommand),
		new(ReloadCommand),
		new(ResetCommand),
		new(SetCommentCommand),
		new(ShuffleCommand),
		new(SkipCommand),
		new(SkipPlaylistCommand),
		new(ToggleShuffleCommand),
		new(VersionCommand),
	}
}
