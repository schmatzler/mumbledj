/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/numcached.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package commands

import (
	"errors"
	"fmt"

	"github.com/layeh/gumble/gumble"
)

// NumCachedCommand is a command that outputs the number of tracks that
// are currently cached on disk (if caching is enabled).
type NumCachedCommand struct{}

// Aliases returns the current aliases for the command.
func (c *NumCachedCommand) Aliases() []string {
	return DJ.BotConfig.Aliases.NumCached
}

// Description returns the description for the command.
func (c *NumCachedCommand) Description() string {
	return DJ.BotConfig.Descriptions.NumCached
}

// IsAdminCommand returns true if the command is only for admin use, and
// returns false otherwise.
func (c *NumCachedCommand) IsAdminCommand() bool {
	return DJ.BotConfig.Permissions.NumCached
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
func (c *NumCachedCommand) Execute(user *gumble.User, args ...string) (string, bool, error) {
	if !DJ.BotConfig.Cache.Enabled {
		return "", true, errors.New("Caching is currently disabled")
	}

	DJ.Cache.UpdateStatistics()
	return fmt.Sprintf("There are currently <b>%d</b> items stored in the cache.",
		DJ.Cache.NumAudioFiles), true, nil
}
