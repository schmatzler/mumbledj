/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/skipplaylist.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package commands

import (
	"errors"
	"fmt"

	"github.com/layeh/gumble/gumble"
)

// SkipPlaylistCommand is a command that places a vote to skip the current
// playlist.
type SkipPlaylistCommand struct{}

// Aliases returns the current aliases for the command.
func (c *SkipPlaylistCommand) Aliases() []string {
	return DJ.BotConfig.Aliases.SkipPlaylist
}

// Description returns the description for the command.
func (c *SkipPlaylistCommand) Description() string {
	return DJ.BotConfig.Descriptions.SkipPlaylist
}

// IsAdminCommand returns true if the command is only for admin use, and
// returns false otherwise.
func (c *SkipPlaylistCommand) IsAdminCommand() bool {
	return DJ.BotConfig.Permissions.SkipPlaylist
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
func (c *SkipPlaylistCommand) Execute(user *gumble.User, args ...string) (string, bool, error) {
	if len(DJ.Queue.Queue) == 0 {
		return "", true, errors.New("The queue is currently empty. There is no playlist to skip")
	}
	if _, err := DJ.Queue.Queue[0].GetPlaylist(); err != nil {
		return "", true, errors.New("The current track is not part of a playlist")
	}
	if err := DJ.Skips.AddPlaylistSkip(user); err != nil {
		return "", true, errors.New("You have already voted to skip this playlist")
	}

	return fmt.Sprintf("<b>%s</b> has voted to skip the current playlist.", user.Name), false, nil
}
