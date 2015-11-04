/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/setcomment.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package commands

import (
	"github.com/layeh/gumble/gumble"
	"github.com/matthieugrieger/mumbledj/objects"
	"github.com/spf13/viper"
)

// SetCommentCommand is a command that changes the Mumble comment of the bot.
type SetCommentCommand struct{}

// Aliases is a method that returns the current aliases for the add command.
func (c *SetCommentCommand) Aliases() []string {
	return viper.GetStringSlice("aliases.setcomment")
}

// IsAdmin is a command that returns a bool that determines if a command is an
// admin command or not.
func (c *SetCommentCommand) IsAdmin() bool {
	return viper.GetBool("permissions.setcomment")
}

// Execute executes the command with the given bot state, user, and arguments.
func (c *SetCommentCommand) Execute(state *objects.BotState, user *gumble.User, args ...string) (*objects.BotState, string, error) {

}