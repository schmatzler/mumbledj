/*
 * MumbleDJ
 * By Matthieu Grieger
 * bot/commander.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package bot

import (
	"errors"
	"strings"

	"github.com/layeh/gumble/gumble"
	"github.com/layeh/gumble/gumbleutil"
	"github.com/matthieugrieger/mumbledj/interfaces"
)

// Commander holds all available commands and determines which commands should
// be executed.
type Commander struct {
	Commands []interfaces.Command
}

// NewCommander returns a new commander with an initialized command list.
func NewCommander() *Commander {
	return nil
}

// FindAndExecuteCommand attempts to find a reference to a command in an
// incoming message. If found, the command is executed and the resulting
// message/error is returned.
func (c *Commander) FindAndExecuteCommand(user *gumble.User, message string) (string, bool, error) {
	command, err := c.findCommand(message)
	if err != nil {
		return "", true, errors.New("No command was found in this message")
	}
	return c.executeCommand(user, message, command)
}

func (c *Commander) findCommand(message string) (interfaces.Command, error) {
	possibleCommand := strings.ToLower(message[0:strings.Index(message, " ")])
	for _, command := range c.Commands {
		for _, alias := range command.Aliases() {
			if possibleCommand == alias {
				return command, nil
			}
		}
	}
	return nil, errors.New("No command was found in this message")
}

func (c *Commander) executeCommand(user *gumble.User, message string, command interfaces.Command) (string, bool, error) {
	canExecute := false
	if DJ.BotConfig.Permissions.Enabled && command.IsAdminCommand() {
		userGroups := <-gumbleutil.UserGroups(DJ.Client, user, DJ.Client.Self.Channel)
		for _, userGroup := range userGroups {
			if userGroup == DJ.BotConfig.Permissions.UserGroup {
				canExecute = true
			}
		}
	} else {
		canExecute = true
	}

	if canExecute {
		return command.Execute(user, strings.Split(message, " ")[1:]...)
	}
	return "", true, errors.New("You do not have permission to execute this command")
}
