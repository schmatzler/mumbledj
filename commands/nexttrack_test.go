/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/nexttrack_test.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package commands

import (
	"testing"

	"github.com/matthieugrieger/mumbledj/bot"
	"github.com/stretchr/testify/suite"
)

type NextTrackCommandTestSuite struct {
	Command NextTrackCommand
	suite.Suite
}

func (suite *NextTrackCommandTestSuite) SetupSuite() {
	DJ = bot.NewMumbleDJ()

	DJ.BotConfig.Aliases.NextTrack = []string{"nexttrack", "next"}
	DJ.BotConfig.Descriptions.NextTrack = "nexttrack"
	DJ.BotConfig.Permissions.NextTrack = false
}

func (suite *NextTrackCommandTestSuite) SetupTest() {
	DJ.Queue = bot.NewQueue()
}

func (suite *NextTrackCommandTestSuite) TestAliases() {
	suite.Equal([]string{"nexttrack", "next"}, suite.Command.Aliases())
}

func (suite *NextTrackCommandTestSuite) TestDescription() {
	suite.Equal("nexttrack", suite.Command.Description())
}

func (suite *NextTrackCommandTestSuite) TestIsAdminCommand() {
	suite.False(suite.Command.IsAdminCommand())
}

func (suite *NextTrackCommandTestSuite) TestExecuteWhenQueueIsEmpty() {
	message, isPrivateMessage, err := suite.Command.Execute(nil)

	suite.Equal("", message, "No message should be returned.")
	suite.True(isPrivateMessage, "This should be a private message.")
	suite.NotNil(err, "An error should be returned due to the queue being empty.")
}

func (suite *NextTrackCommandTestSuite) TestExecuteWhenQueueHasOneTrack() {
	track := new(bot.Track)
	track.Title = "test"
	track.Submitter = "test"

	DJ.Queue.Queue = append(DJ.Queue.Queue, track)

	message, isPrivateMessage, err := suite.Command.Execute(nil)

	suite.Equal("", message, "No message should be returned.")
	suite.True(isPrivateMessage, "This should be a private message.")
	suite.NotNil(err, "An error should be returned due to the queue having only one track.")
}

func (suite *NextTrackCommandTestSuite) TestExecuteWhenQueueHasTwoOrMoreTracks() {
	track1 := new(bot.Track)
	track1.Title = "first"
	track1.Submitter = "test"

	track2 := new(bot.Track)
	track2.Title = "second"
	track2.Submitter = "test"

	DJ.Queue.Queue = append(DJ.Queue.Queue, track1)
	DJ.Queue.Queue = append(DJ.Queue.Queue, track2)

	message, isPrivateMessage, err := suite.Command.Execute(nil)

	suite.NotEqual("", "A message containing information for the next track should be returned.")
	suite.Contains(message, "second", "The returned message should contain information about the second track in the queue.")
	suite.True(isPrivateMessage, "This should be a private message.")
	suite.Nil(err, "No error should be returned.")
}

func TestNextTrackCommandTestSuite(t *testing.T) {
	suite.Run(t, new(NextTrackCommandTestSuite))
}
