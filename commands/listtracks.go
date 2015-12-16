/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/listtracks.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package commands

import (
	"github.com/layeh/gumble/gumble"
	"github.com/matthieugrieger/mumbledj/state"
	"github.com/spf13/viper"
)

// ListTracksCommand is a command that lists the tracks that are currently in the queue.
type ListTracksCommand struct{}

// Aliases is a method that returns the current aliases for the command.
func (c *ListTracksCommand) Aliases() []string {
	return viper.GetStringSlice("aliases.listtracks")
}

// IsAdmin is a command that returns a bool that determines if a command is an
// admin command or not.
func (c *ListTracksCommand) IsAdmin() bool {
	return viper.GetBool("permissions.listtracks")
}

// Execute executes the command with the given bot state, user, and arguments.
func (c *ListTracksCommand) Execute(state *state.BotState, user *gumble.User, args ...string) (*state.BotState, string, bool, error) {
	return nil, "", false, nil
}