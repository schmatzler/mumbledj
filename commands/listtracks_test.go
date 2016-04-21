/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/listtracks_test.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package commands

import (
	"testing"

	"github.com/matthieugrieger/mumbledj/bot"
	"github.com/stretchr/testify/suite"
)

type ListTracksCommandTestSuite struct {
	Command ListTracksCommand
	suite.Suite
}

func (suite *ListTracksCommandTestSuite) SetupSuite() {
	DJ = bot.NewMumbleDJ()

	DJ.BotConfig.Aliases.ListTracks = []string{"listtracks", "list"}
	DJ.BotConfig.Descriptions.ListTracks = "listtracks"
	DJ.BotConfig.Permissions.ListTracks = false
}

func (suite *ListTracksCommandTestSuite) SetupTest() {
	DJ.Queue = bot.NewQueue()
}

func (suite *ListTracksCommandTestSuite) TestAliases() {
	suite.Equal([]string{"listtracks", "list"}, suite.Command.Aliases())
}

func (suite *ListTracksCommandTestSuite) TestDescription() {
	suite.Equal("listtracks", suite.Command.Description())
}

func (suite *ListTracksCommandTestSuite) TestIsAdminCommand() {
	suite.False(suite.Command.IsAdminCommand())
}

func (suite *ListTracksCommandTestSuite) TestExecuteWithNoTracks() {
	message, isPrivateMessage, err := suite.Command.Execute(nil)

	suite.Equal("", message, "No message should be returned.")
	suite.True(isPrivateMessage, "This should be a private message.")
	suite.NotNil(err, "An error should be returned as there are no tracks to list.")
}

func (suite *ListTracksCommandTestSuite) TestExecuteWithNoArg() {
	track := new(bot.Track)
	track.Title = "title"
	track.Submitter = "test"

	// This results in a nil-pointer error and I have no idea why.
	// An identical setup works in currenttrack_test.go.
	DJ.Queue.AddTrack(track)

	message, isPrivateMessage, err := suite.Command.Execute(nil)

	suite.NotEqual("", message, "A message containing track information should be returned.")
	suite.Contains(message, "title", "The returned message should contain the track title.")
	suite.Contains(message, "test", "The returned message should contain the track submitter.")
	suite.True(isPrivateMessage, "This should be a private message.")
	suite.Nil(err, "No error should be returned.")
}

func (suite *ListTracksCommandTestSuite) TestExecuteWithValidArg() {

}

func (suite *ListTracksCommandTestSuite) TestExecuteWithArgLargerThanQueueLength() {

}

func (suite *ListTracksCommandTestSuite) TestExecuteWithInvalidArg() {

}

func TestListTracksCommandTestSuite(t *testing.T) {
	suite.Run(t, new(ListTracksCommandTestSuite))
}
