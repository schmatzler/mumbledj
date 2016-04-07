/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/reload.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package commands

import "github.com/layeh/gumble/gumble"

// ReloadCommand is a command that reloads the configuration values for the bot
// from a config file.
type ReloadCommand struct{}

// Aliases returns the current aliases for the command.
func (c *ReloadCommand) Aliases() []string {
	return DJ.BotConfig.Aliases.Reload
}

// Description returns the description for the command.
func (c *ReloadCommand) Description() string {
	return DJ.BotConfig.Descriptions.Reload
}

// IsAdminCommand returns true if the command is only for admin use, and
// returns false otherwise.
func (c *ReloadCommand) IsAdminCommand() bool {
	return DJ.BotConfig.Permissions.Reload
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
func (c *ReloadCommand) Execute(user *gumble.User, args ...string) (string, bool, error) {
	// configFileLocation equal to "" just uses the default config file location.
	configFileLocation := ""
	if len(args) >= 1 {
		configFileLocation = args[0]
	}

	if err := DJ.BotConfig.LoadFromConfigFile(configFileLocation); err != nil {
		return "", true, err
	}

	return "The configuration in the configuration file has been loaded successfully.",
		true, nil
}
