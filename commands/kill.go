/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/currenttrack.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package commands

import (
	"os"

	"github.com/layeh/gumble/gumble"
)

// KillCommand is a command that safely kills the bot.
type KillCommand struct{}

// Aliases returns the current aliases for the command.
func (c *KillCommand) Aliases() []string {
	return DJ.BotConfig.Aliases.Kill
}

// Description returns the description for the command.
func (c *KillCommand) Description() string {
	return DJ.BotConfig.Descriptions.Kill
}

// IsAdminCommand returns true if the command is only for admin use, and
// returns false otherwise.
func (c *KillCommand) IsAdminCommand() bool {
	return DJ.BotConfig.Permissions.Kill
}

// Execute executes the command with the given user and arguments.
// Return value descriptions:
//    string: A message to be returned to the user upon successful execution.
//    bool:   Whether the message should be private or not. true = private,
//            false = public (sent to whole channel).
//    error:  An error message to be returned upon unsuccessful execution.
//            If no error has occurred, pass nil instead.
// Example return statement:
//    return "This is a private message!", true, nil
func (c *KillCommand) Execute(user *gumble.User, args ...string) (string, bool, error) {
	if err := DJ.Cache.DeleteAll(); err != nil {
		return "", true, err
	}
	if err := DJ.Client.Disconnect(); err != nil {
		return "", true, err
	}

	os.Exit(0)
	return "", true, nil
}
