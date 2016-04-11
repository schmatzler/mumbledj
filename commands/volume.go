/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/volume.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package commands

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/layeh/gumble/gumble"
)

// VolumeCommand is a command that changes the volume of the audio output.
type VolumeCommand struct{}

// Aliases returns the current aliases for the command.
func (c *VolumeCommand) Aliases() []string {
	return DJ.BotConfig.Aliases.Volume
}

// Description returns the description for the command.
func (c *VolumeCommand) Description() string {
	return DJ.BotConfig.Descriptions.Volume
}

// IsAdminCommand returns true if the command is only for admin use, and
// returns false otherwise.
func (c *VolumeCommand) IsAdminCommand() bool {
	return DJ.BotConfig.Permissions.Volume
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
func (c *VolumeCommand) Execute(user *gumble.User, args ...string) (string, bool, error) {
	if len(args) == 0 {
		// Send the user the current volume level.
		return fmt.Sprintf("The current volume is <b>%.2f</b>.", DJ.Volume), true, nil
	}

	newVolume, err := strconv.ParseFloat(args[0], 32)
	if err != nil {
		return "", true, errors.New("An error occurred while parsing the requested volume")
	}

	newVolume32 := float32(newVolume)

	if newVolume32 < DJ.BotConfig.Volume.Lowest || newVolume32 > DJ.BotConfig.Volume.Highest {
		return "", true, fmt.Errorf("Volumes must be between the values <b>%.2f</b> and <b>%.2f</b>",
			DJ.BotConfig.Volume.Lowest, DJ.BotConfig.Volume.Highest)
	}

	if DJ.AudioStream != nil {
		DJ.AudioStream.Volume = newVolume32
	}
	DJ.Volume = newVolume32

	return fmt.Sprintf("<b>%s</b> has changed the volume to <b>%.2f</b>.",
		user.Name, newVolume32), false, nil
}
