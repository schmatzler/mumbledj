/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/setcomment.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package commands

import (
	"fmt"
	"strings"

	"github.com/layeh/gumble/gumble"
)

// SetCommentCommand is a command that changes the Mumble comment of the bot.
type SetCommentCommand struct{}

// Aliases returns the current aliases for the command.
func (c *SetCommentCommand) Aliases() []string {
	return DJ.BotConfig.Aliases.SetComment
}

// Description returns the description for the command.
func (c *SetCommentCommand) Description() string {
	return DJ.BotConfig.Descriptions.SetComment
}

// IsAdminCommand returns true if the command is only for admin use, and
// returns false otherwise.
func (c *SetCommentCommand) IsAdminCommand() bool {
	return DJ.BotConfig.Permissions.SetComment
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
func (c *SetCommentCommand) Execute(user *gumble.User, args ...string) (string, bool, error) {
	if len(args) == 0 {
		DJ.Client.Self.SetComment("")
		return "The comment for the bot has been successfully removed.", true, nil
	}

	var newComment string
	for _, arg := range args {
		newComment += arg + " "
	}
	strings.TrimSpace(newComment)
	DJ.Client.Self.SetComment(newComment)

	return fmt.Sprintf("The comment for the bot has been successfully changed to the following: %s",
		newComment), true, nil
}
