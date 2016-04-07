/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/reset.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package commands

import (
	"errors"
	"fmt"

	"github.com/layeh/gumble/gumble"
)

// ResetCommand is a command that resets the queue and cache.
type ResetCommand struct{}

// Aliases returns the current aliases for the command.
func (c *ResetCommand) Aliases() []string {
	return DJ.BotConfig.Aliases.Reset
}

// Description returns the description for the command.
func (c *ResetCommand) Description() string {
	return DJ.BotConfig.Descriptions.Reset
}

// IsAdminCommand returns true if the command is only for admin use, and
// returns false otherwise.
func (c *ResetCommand) IsAdminCommand() bool {
	return DJ.BotConfig.Permissions.Reset
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
func (c *ResetCommand) Execute(user *gumble.User, args ...string) (string, bool, error) {
	if len(DJ.Queue.Queue) == 0 {
		return "", true, errors.New("The queue is already empty")
	}

	DJ.Queue.Queue = DJ.Queue.Queue[:0]
	if DJ.AudioStream != nil {
		DJ.AudioStream.Stop()
		DJ.AudioStream = nil
	}

	if err := DJ.Cache.DeleteAll(); err != nil {
		return "", true, err
	}

	return fmt.Sprintf("<b>%s</b> has reset the queue.", user.Name), false, nil
}
