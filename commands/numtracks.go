/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/numtracks.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package commands

import (
	"fmt"

	"github.com/layeh/gumble/gumble"
)

// NumTracksCommand is a command that outputs the current number of tracks
// in the queue.
type NumTracksCommand struct{}

// Aliases returns the current aliases for the command.
func (c *NumTracksCommand) Aliases() []string {
	return DJ.BotConfig.Aliases.NumTracks
}

// Description returns the description for the command.
func (c *NumTracksCommand) Description() string {
	return DJ.BotConfig.Descriptions.NumTracks
}

// IsAdminCommand returns true if the command is only for admin use, and
// returns false otherwise.
func (c *NumTracksCommand) IsAdminCommand() bool {
	return DJ.BotConfig.Permissions.NumTracks
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
func (c *NumTracksCommand) Execute(user *gumble.User, args ...string) (string, bool, error) {
	if len(DJ.Queue.Queue) == 1 {
		return "There is currently <b>1</b> track in the queue", true, nil
	}

	return fmt.Sprintf("There are currently <b>%d</b> tracks in the queue.", len(DJ.Queue.Queue)), true, nil
}
