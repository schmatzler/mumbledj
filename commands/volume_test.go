/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/volume_test.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package commands

import (
	"testing"

	"github.com/layeh/gumble/gumble"
	"github.com/matthieugrieger/mumbledj/bot"
	"github.com/stretchr/testify/suite"
)

type VolumeCommandTestSuite struct {
	Command VolumeCommand
	suite.Suite
}

func (suite *VolumeCommandTestSuite) SetupSuite() {
	DJ = bot.NewMumbleDJ()

	DJ.BotConfig.Aliases.Volume = []string{"volume", "vol"}
	DJ.BotConfig.Descriptions.Volume = "volume"
	DJ.BotConfig.Permissions.Volume = false
	DJ.BotConfig.Volume.Lowest = 0.2
	DJ.BotConfig.Volume.Highest = 1
	DJ.BotConfig.Volume.Default = 0.4
}

func (suite *VolumeCommandTestSuite) SetupTest() {
	DJ.Volume = DJ.BotConfig.Volume.Default
}

func (suite *VolumeCommandTestSuite) TestAliases() {
	suite.Equal([]string{"volume", "vol"}, suite.Command.Aliases())
}

func (suite *VolumeCommandTestSuite) TestDescription() {
	suite.Equal("volume", suite.Command.Description())
}

func (suite *VolumeCommandTestSuite) TestIsAdminCommand() {
	suite.False(suite.Command.IsAdminCommand())
}

func (suite *VolumeCommandTestSuite) TestExecuteWithNoArgs() {
	message, isPrivateMessage, err := suite.Command.Execute(nil)

	suite.NotEqual("", message, "A message should be returned.")
	suite.True(isPrivateMessage, "This should be a private message.")
	suite.Nil(err, "No error should be returned.")
	suite.Contains(message, "0.4", "The returned string should contain the current volume.")
}

func (suite *VolumeCommandTestSuite) TestExecuteWithValidArg() {
	dummyUser := &gumble.User{
		Name: "test",
	}
	message, isPrivateMessage, err := suite.Command.Execute(dummyUser, "0.6")

	suite.NotEqual("", message, "A message should be returned.")
	suite.False(isPrivateMessage, "This should not be a private message.")
	suite.Nil(err, "No error should be returned.")
	suite.Contains(message, "0.6", "The returned string should contain the new volume.")
	suite.Contains(message, "test", "The returned string should contain the username of whomever changed the volume.")
}

func (suite *VolumeCommandTestSuite) TestExecuteWithArgOutOfRange() {
	message, isPrivateMessage, err := suite.Command.Execute(nil, "1.4")

	suite.Equal("", message, "No message should be returned as an error occurred.")
	suite.True(isPrivateMessage, "This should be a private message.")
	suite.NotNil(err, "An error should be returned since the provided argument was outside of the valid range.")
}

func (suite *VolumeCommandTestSuite) TestExecuteWithInvalidArg() {
	message, isPrivateMessage, err := suite.Command.Execute(nil, "test")

	suite.Equal("", message, "No message should be returned as an error occurred.")
	suite.True(isPrivateMessage, "This should be a private message.")
	suite.NotNil(err, "An error should be returned as a non-floating-point argument was provided.")
}

func TestVolumeCommandTestSuite(t *testing.T) {
	suite.Run(t, new(VolumeCommandTestSuite))
}
