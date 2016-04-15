/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/forceskipplaylist.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package commands

import (
	"errors"
	"fmt"

	"github.com/layeh/gumble/gumble"
)

// ForceSkipPlaylistCommand is a command that immediately skips the current
// playlist.
type ForceSkipPlaylistCommand struct{}

// Aliases returns the current aliases for the command.
func (c *ForceSkipPlaylistCommand) Aliases() []string {
	return DJ.BotConfig.Aliases.ForceSkipPlaylist
}

// Description returns the description for the command.
func (c *ForceSkipPlaylistCommand) Description() string {
	return DJ.BotConfig.Descriptions.ForceSkipPlaylist
}

// IsAdminCommand returns true if the command is only for admin use, and
// returns false otherwise.
func (c *ForceSkipPlaylistCommand) IsAdminCommand() bool {
	return DJ.BotConfig.Permissions.ForceSkipPlaylist
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
func (c *ForceSkipPlaylistCommand) Execute(user *gumble.User, args ...string) (string, bool, error) {
	if len(DJ.Queue.Queue) == 0 {
		return "", true, errors.New("The queue is currently empty. There are no playlists to skip")
	}
	if _, err := DJ.Queue.Queue[0].GetPlaylist(); err != nil {
		return "", true, errors.New("The current track is not part of a playlist")
	}

	DJ.Queue.SkipPlaylist()

	return fmt.Sprintf("The current playlist has been forcibly skipped by <b>%s</b>.",
		user.Name), false, nil
}
