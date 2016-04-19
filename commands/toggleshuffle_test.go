/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/toggleshuffle_test.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package commands

import (
	"testing"

	"github.com/matthieugrieger/mumbledj/bot"
	"github.com/stretchr/testify/suite"
)

type ToggleShuffleCommandTestSuite struct {
	Command ToggleShuffleCommand
	suite.Suite
}

func (suite *ToggleShuffleCommandTestSuite) SetupSuite() {
	DJ = bot.NewMumbleDJ()

	DJ.BotConfig.Aliases.ToggleShuffle = []string{"toggleshuffle", "ts"}
	DJ.BotConfig.Descriptions.ToggleShuffle = "toggleshuffle"
	DJ.BotConfig.Permissions.ToggleShuffle = true
}

func (suite *ToggleShuffleCommandTestSuite) TestAliases() {
	suite.Equal([]string{"toggleshuffle", "ts"}, suite.Command.Aliases())
}

func (suite *ToggleShuffleCommandTestSuite) TestDescription() {
	suite.Equal("toggleshuffle", suite.Command.Description())
}

func (suite *ToggleShuffleCommandTestSuite) TestIsAdminCommand() {
	suite.True(suite.Command.IsAdminCommand())
}

func (suite *ToggleShuffleCommandTestSuite) TestExecuteWhenShuffleIsOff() {
	DJ.BotConfig.General.AutomaticShuffleOn = false

	message, isPrivateMessage, err := suite.Command.Execute(nil)

	suite.NotEqual("", message, "A message should be returned.")
	suite.False(isPrivateMessage, "This should not be a private message.")
	suite.Nil(err, "No error should be returned.")
	suite.True(DJ.BotConfig.General.AutomaticShuffleOn, "Automatic shuffling should now be on.")
}

func (suite *ToggleShuffleCommandTestSuite) TestExecuteWhenShuffleIsOn() {
	DJ.BotConfig.General.AutomaticShuffleOn = true

	message, isPrivateMessage, err := suite.Command.Execute(nil)

	suite.NotEqual("", message, "A message should be returned.")
	suite.False(isPrivateMessage, "This should not be a private message.")
	suite.Nil(err, "No error should be returned.")
	suite.False(DJ.BotConfig.General.AutomaticShuffleOn, "Automatic shuffling should now be off.")
}

func TestToggleShuffleCommandTestSuite(t *testing.T) {
	suite.Run(t, new(ToggleShuffleCommandTestSuite))
}
