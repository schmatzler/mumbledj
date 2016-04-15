/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/currenttrack_test.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package commands

import (
	"github.com/matthieugrieger/mumbledj/bot"
	"github.com/stretchr/testify/suite"
)

type CurrentTrackCommandTestSuite struct {
	Command CurrentTrackCommand
	suite.Suite
}

func (suite *CurrentTrackCommandTestSuite) SetupSuite() {
	DJ = bot.NewMumbleDJ()

	DJ.BotConfig.Aliases.CurrentTrack = []string{"currenttrack", "current"}
	DJ.BotConfig.Descriptions.CurrentTrack = "currenttrack"
	DJ.BotConfig.Permissions.CurrentTrack = false
}

func (suite *CurrentTrackCommandTestSuite) SetupTest() {
	DJ.Queue = bot.NewQueue()
}

func (suite *CurrentTrackCommandTestSuite) TestAliases() {
	suite.Equal([]string{"currenttrack", "current"}, suite.Command.Aliases())
}

func (suite *CurrentTrackCommandTestSuite) TestDescription() {
	suite.Equal("currenttrack", suite.Command.Description())
}

func (suite *CurrentTrackCommandTestSuite) TestIsAdminCommand() {
	suite.False(suite.Command.IsAdminCommand())
}

func (suite *CurrentTrackCommandTestSuite) TestExecuteWhenQueueIsEmpty() {
	message, isPrivateMessage, err := suite.Command.Execute(nil)

	suite.Equal("", message, "No message should be returned since an error occurred.")
	suite.True(isPrivateMessage, "This should be a private message.")
	suite.NotNil(err, "An error should be returned since the queue is empty.")
}

func (suite *CurrentTrackCommandTestSuite) TestExecuteWhenQueueNotEmpty() {
	track := new(bot.Track)
	track.Submitter = "test"
	track.Title = "test"

	DJ.Queue.AddTrack(track)

	message, isPrivateMessage, err := suite.Command.Execute(nil)

	suite.NotEqual("", message, "A message should be returned with the current track information.")
	suite.True(isPrivateMessage, "This should be a private message.")
	suite.Nil(err, "No error should be returned.")
}
