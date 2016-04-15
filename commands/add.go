/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/add.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package commands

import (
	"errors"
	"fmt"

	"github.com/layeh/gumble/gumble"
	"github.com/matthieugrieger/mumbledj/interfaces"
)

// AddCommand is a command that adds an audio track associated with a supported
// URL to the queue.
type AddCommand struct{}

// Aliases returns the current aliases for the command.
func (c *AddCommand) Aliases() []string {
	return DJ.BotConfig.Aliases.Add
}

// Description returns the description for the command.
func (c *AddCommand) Description() string {
	return DJ.BotConfig.Descriptions.Add
}

// IsAdminCommand returns true if the command is only for admin use, and
// returns false otherwise.
func (c *AddCommand) IsAdminCommand() bool {
	return DJ.BotConfig.Permissions.Add
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
func (c *AddCommand) Execute(user *gumble.User, args ...string) (string, bool, error) {
	var allTracks []interfaces.Track
	tracksTooLong := ""

	if len(args) == 0 {
		return "", true, errors.New("A URL must be supplied with the add command")
	}

	for _, arg := range args {
		tracks, err := DJ.YouTubeDL.GetTracks(arg)
		if err == nil {
			allTracks = append(allTracks, tracks...)
		}
	}

	if len(allTracks) == 0 {
		return "", true, errors.New("No valid tracks were found with the provided URL(s)")
	}

	addString := fmt.Sprintf("<b>%s</b> added <b>%d</b> tracks to the queue:</br>",
		user.Name, len(allTracks))

	for _, track := range allTracks {
		if err := DJ.Queue.AddTrack(track); err != nil {
			tracksTooLong += fmt.Sprintf("\"%s\" from %s </br>",
				track.GetTitle(), track.GetService())
		} else {
			addString += fmt.Sprintf("\"%s\" from %s </br>",
				track.GetTitle(), track.GetService())
		}
	}

	if len(tracksTooLong) != 0 {
		addString += "</br>The following tracks could not be added due to error or because they are too long:</br>" + tracksTooLong
	}

	return addString, false, nil
}
