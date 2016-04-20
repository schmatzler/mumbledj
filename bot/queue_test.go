/*
 * MumbleDJ
 * By Matthieu Grieger
 * bot/queue_test.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package bot

import (
	"testing"

	"github.com/matthieugrieger/mumbledj/bot"
	"github.com/stretchr/testify/suite"
)

type MockedQueueItem struct {
	Track
	Identifier string
}

type QueueTestSuite struct {
	suite.Suite
	FirstItem  MockedQueueItem
	SecondItem MockedQueueItem
	ThirdItem  MockedQueueItem
}

func (suite *QueueTestSuite) SetupSuite() {
	DJ = bot.NewMumbleDJ()

	DJ.BotConfig.General.AutomaticShuffleOn = false

	suite.FirstItem = MockedQueueItem{Identifier: "first"}
	suite.SecondItem = MockedQueueItem{Identifier: "second"}
	suite.ThirdItem = MockedQueueItem{Identifier: "third"}
}

func (suite *QueueTestSuite) SetupTest() {
	DJ.Queue = bot.NewQueue()
}

func (suite *QueueTestSuite) TestNewQueue() {
	suite.Zero(len(DJ.Queue.Queue), "The new queue should be empty.")
}

func (suite *QueueTestSuite) TestAddTrack() {
	suite.Zero(len(DJ.Queue.Queue), "The queue should be empty.")
	DJ.Queue.AddTrack(suite.FirstItem)
	suite.Equal(1, len(DJ.Queue.Queue), "There should now be one track in the queue.")
}

// TODO: Implement this test.
func (suite *QueueTestSuite) TestCurrentTrackWhenQueueEmpty() {

}

// TODO: Implement this test.
func (suite *QueueTestSuite) TestCurrentTrackWhenQueueNotEmpty() {

}

func TestQueueTestSuite(t *testing.T) {
	suite.Run(t, new(QueueTestSuite))
}
