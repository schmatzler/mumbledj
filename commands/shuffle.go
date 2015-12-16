/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/shuffle.go
 * Copyright (c) 2014, 2015 Matthieu Grieger (MIT License)
 */

package commands

import (
	"errors"

	"github.com/layeh/gumble/gumble"
	"github.com/matthieugrieger/mumbledj/state"
	"github.com/spf13/viper"
)

// ShuffleCommand is a command that shuffles the audio queue.
type ShuffleCommand struct{}

// Aliases is a method that returns the current aliases for the command.
func (c *ShuffleCommand) Aliases() []string {
	return viper.GetStringSlice("aliases.shuffle")
}

// IsAdmin is a command that returns a bool that determines if a command is an
// admin command or not.
func (c *ShuffleCommand) IsAdmin() bool {
	return viper.GetBool("permissions.shuffle")
}

// Execute executes the command with the given bot state, user, and arguments.
func (c *ShuffleCommand) Execute(state *state.BotState, user *gumble.User, args ...string) (*state.BotState, string, bool, error) {
	if len(state.Queue.Queue) == 0 {
		return nil, "", true, errors.New("There are no tracks currently in the queue.")
	}
	if len(state.Queue.Queue) <= 2 {
		return nil, "", true, errors.New("There are not enough tracks to execute a shuffle operation.")
	}

	state.Queue.ShuffleTracks()

	return state, "The audio queue has been shuffled.", false, nil
}