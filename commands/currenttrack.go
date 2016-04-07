/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/currenttrack.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package commands

import (
	"errors"
	"fmt"

	"github.com/layeh/gumble/gumble"
)

// CurrentTrackCommand is a command that outputs information related to
// the track that is currently playing (if one exists).
type CurrentTrackCommand struct{}

// Aliases returns the current aliases for the command.
func (c *CurrentTrackCommand) Aliases() []string {
	return DJ.BotConfig.Aliases.CurrentTrack
}

// Description returns the description for the command.
func (c *CurrentTrackCommand) Description() string {
	return DJ.BotConfig.Descriptions.CurrentTrack
}

// IsAdminCommand returns true if the command is only for admin use, and
// returns false otherwise.
func (c *CurrentTrackCommand) IsAdminCommand() bool {
	return DJ.BotConfig.Permissions.CurrentTrack
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
func (c *CurrentTrackCommand) Execute(user *gumble.User, args ...string) (string, bool, error) {
	if len(DJ.Queue.Queue) == 0 {
		return "", true, errors.New("There are no tracks in the queue")
	}

	currentTrack := DJ.Queue.Queue[0]

	return fmt.Sprintf("The current track is <b>%s</b>, added by <b>%s</b>.",
		currentTrack.Title(), currentTrack.Submitter()), true, nil
}
