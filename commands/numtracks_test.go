/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/numtracks_test.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package commands

import (
	"testing"

	"github.com/matthieugrieger/mumbledj/bot"
	"github.com/stretchr/testify/suite"
)

type NumTracksCommandTestSuite struct {
	Command NumTracksCommand
	suite.Suite
}

func (suite *NumTracksCommandTestSuite) SetupSuite() {
	DJ = bot.NewMumbleDJ()

	DJ.BotConfig.Aliases.NumTracks = []string{"numtracks", "num"}
	DJ.BotConfig.Descriptions.NumTracks = "numtracks"
	DJ.BotConfig.Permissions.NumTracks = false
}

func (suite *NumTracksCommandTestSuite) SetupTest() {
	DJ.Queue = bot.NewQueue()
}

func (suite *NumTracksCommandTestSuite) TestAliases() {
	suite.Equal([]string{"numtracks", "num"}, suite.Command.Aliases())
}

func (suite *NumTracksCommandTestSuite) TestDescription() {
	suite.Equal("numtracks", suite.Command.Description())
}

func (suite *NumTracksCommandTestSuite) TestIsAdminCommand() {
	suite.False(suite.Command.IsAdminCommand())
}

func (suite *NumTracksCommandTestSuite) TestExecuteWhenZeroTracksAreInQueue() {
	message, isPrivateMessage, err := suite.Command.Execute(nil)

	suite.NotEqual("", message, "A message should be returned.")
	suite.Contains(message, "<b>0</b> tracks", "The returned message should state that there are no tracks in the queue.")
	suite.True(isPrivateMessage, "This should be a private message.")
	suite.Nil(err, "No error should be returned.")
}

func (suite *NumTracksCommandTestSuite) TestExecuteWhenOneTrackIsInQueue() {
	track := new(bot.Track)
	track.Title = "test"
	track.Submitter = "test"

	DJ.Queue.Queue = append(DJ.Queue.Queue, track)

	message, isPrivateMessage, err := suite.Command.Execute(nil)

	suite.NotEqual("", message, "A message should be returned.")
	suite.Contains(message, "<b>1</b> track", "The returned message should state that there is one track in the queue.")
	suite.True(isPrivateMessage, "This should be a private message.")
	suite.Nil(err, "No error should be returned.")
}

func (suite *NumTracksCommandTestSuite) TestExecuteWhenTwoOrMoreTracksAreInQueue() {
	track1 := new(bot.Track)
	track1.Title = "test"
	track1.Submitter = "test"

	track2 := new(bot.Track)
	track2.Title = "test"
	track2.Submitter = "test"

	DJ.Queue.Queue = append(DJ.Queue.Queue, track1)
	DJ.Queue.Queue = append(DJ.Queue.Queue, track2)

	message, isPrivateMessage, err := suite.Command.Execute(nil)

	suite.NotEqual("", "A message should be returned.")
	suite.Contains(message, "tracks", "The returned message should use the plural form of the word track.")
	suite.True(isPrivateMessage, "This should be a private message.")
	suite.Nil(err, "No error should be returned.")
}

func TestNumTracksCommandTestSuite(t *testing.T) {
	suite.Run(t, new(NumTracksCommandTestSuite))
}
