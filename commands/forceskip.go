/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/forceskip.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package commands

import (
	"errors"
	"fmt"

	"github.com/layeh/gumble/gumble"
)

// ForceSkipCommand is a command that immediately skips the current track.
type ForceSkipCommand struct{}

// Aliases returns the current aliases for the command.
func (c *ForceSkipCommand) Aliases() []string {
	return DJ.BotConfig.Aliases.ForceSkip
}

// Description returns the description for the command.
func (c *ForceSkipCommand) Description() string {
	return DJ.BotConfig.Descriptions.ForceSkip
}

// IsAdminCommand returns true if the command is only for admin use, and
// returns false otherwise.
func (c *ForceSkipCommand) IsAdminCommand() bool {
	return DJ.BotConfig.Permissions.ForceSkip
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
func (c *ForceSkipCommand) Execute(user *gumble.User, args ...string) (string, bool, error) {
	if len(DJ.Queue.Queue) == 0 {
		return "", true, errors.New("The queue is currently empty. There are no tracks to skip")
	}

	DJ.Queue.Skip()

	return fmt.Sprintf("The current track has been forcibly skipped by <b>%s</b>.",
		user.Name), false, nil
}
