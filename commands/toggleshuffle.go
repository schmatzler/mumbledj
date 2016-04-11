/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/toggleshuffle.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package commands

import "github.com/layeh/gumble/gumble"

// ToggleShuffleCommand is a command that changes the Mumble comment of the bot.
type ToggleShuffleCommand struct{}

// Aliases returns the current aliases for the command.
func (c *ToggleShuffleCommand) Aliases() []string {
	return DJ.BotConfig.Aliases.ToggleShuffle
}

// Description returns the description for the command.
func (c *ToggleShuffleCommand) Description() string {
	return DJ.BotConfig.Descriptions.ToggleShuffle
}

// IsAdminCommand returns true if the command is only for admin use, and
// returns false otherwise.
func (c *ToggleShuffleCommand) IsAdminCommand() bool {
	return DJ.BotConfig.Permissions.ToggleShuffle
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
func (c *ToggleShuffleCommand) Execute(user *gumble.User, args ...string) (string, bool, error) {
	if DJ.BotConfig.General.AutomaticShuffleOn {
		DJ.BotConfig.General.AutomaticShuffleOn = false
		return "Automatic shuffling has been toggled off.", false, nil
	}
	DJ.BotConfig.General.AutomaticShuffleOn = true
	return "Automatic shuffling has been toggled on.", false, nil
}
