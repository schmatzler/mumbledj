/*
 * MumbleDJ
 * By Matthieu Grieger
 * bot/queue_test.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package bot

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/matthieugrieger/mumbledj/interfaces"
	"github.com/stretchr/testify/suite"
)

type QueueTestSuite struct {
	suite.Suite
	FirstTrack  *Track
	SecondTrack *Track
	ThirdTrack  *Track
}

func (suite *QueueTestSuite) SetupSuite() {
	DJ = NewMumbleDJ()

	DJ.BotConfig.General.AutomaticShuffleOn = false

	suite.FirstTrack = &Track{ID: "first"}
	suite.SecondTrack = &Track{ID: "second"}
	suite.ThirdTrack = &Track{ID: "third"}
}

func (suite *QueueTestSuite) SetupTest() {
	DJ.Queue = NewQueue()
	DJ.BotConfig.General.MaxTrackDuration = 0

	// Override the initialized seed for consistent test results.
	rand.Seed(1)
}

func (suite *QueueTestSuite) TestNewQueue() {
	suite.Zero(len(DJ.Queue.Queue), "The new queue should be empty.")
}

func (suite *QueueTestSuite) TestAddTrackWhenTrackIsValid() {
	suite.Zero(len(DJ.Queue.Queue), "The queue should be empty.")
	err := DJ.Queue.AddTrack(suite.FirstTrack)

	suite.Equal(1, len(DJ.Queue.Queue), "There should now be one track in the queue.")
	suite.Nil(err, "No error should be returned.")
}

func (suite *QueueTestSuite) TestAddTrackWhenTrackIsTooLong() {
	// Set max track duration to 5 seconds
	DJ.BotConfig.General.MaxTrackDuration = 5

	// Create duration longer than 5 seconds
	duration, _ := time.ParseDuration("6s")

	longTrack := &Track{}

	longTrack.Duration = duration

	suite.Zero(len(DJ.Queue.Queue), "The queue should be empty.")
	err := DJ.Queue.AddTrack(longTrack)

	suite.Zero(len(DJ.Queue.Queue), "The queue should still be empty.")
	suite.NotNil(err, "An error should be returned due to the track being too long.")
}

func (suite *QueueTestSuite) TestCurrentTrackWhenOneExists() {
	DJ.Queue.AddTrack(suite.FirstTrack)

	track, err := DJ.Queue.CurrentTrack()

	suite.NotNil(track, "The returned track should be non-nil.")
	suite.Equal("first", track.GetID(), "The returned track should be the one just added to the queue.")
	suite.Nil(err, "No error should be returned.")
}

func (suite *QueueTestSuite) TestCurrentTrackWhenOneDoesNotExist() {
	track, err := DJ.Queue.CurrentTrack()

	suite.Nil(track, "The returned track should be nil.")
	suite.NotNil(err, "An error should be returned because there are no tracks in the queue.")
}

func (suite *QueueTestSuite) TestPeekNextTrackWhenOneExists() {
	DJ.Queue.AddTrack(suite.FirstTrack)
	DJ.Queue.AddTrack(suite.SecondTrack)

	track, err := DJ.Queue.PeekNextTrack()

	suite.NotNil(track, "The returned track should be non-nil.")
	suite.Equal("second", track.GetID(), "The returned track should be the second one added to the queue.")
	suite.Nil(err, "No error should be returned.")
}

func (suite *QueueTestSuite) TestPeekNextTrackWhenOneDoesNotExist() {
	track, err := DJ.Queue.PeekNextTrack()

	suite.Nil(track, "The returned track should be nil.")
	suite.NotNil(err, "An error should be returned because there are no tracks in the queue.")

	DJ.Queue.AddTrack(suite.FirstTrack)

	track, err = DJ.Queue.PeekNextTrack()

	suite.Nil(track, "The returned track should be nil.")
	suite.NotNil(err, "An error should be returned because there is only one track in the queue.")
}

func (suite *QueueTestSuite) TestTraverseWhenNoTracksExist() {
	trackString := ""

	DJ.Queue.Traverse(func(i int, t interfaces.Track) {
		trackString += t.GetID() + ", "
	})

	suite.Equal("", trackString, "No tracks should be traversed as there are none in the queue.")
}

func (suite *QueueTestSuite) TestTraverseWhenTracksExist() {
	trackString := ""
	DJ.Queue.AddTrack(suite.FirstTrack)
	DJ.Queue.AddTrack(suite.SecondTrack)

	DJ.Queue.Traverse(func(i int, t interfaces.Track) {
		trackString += t.GetID() + ", "
	})

	suite.NotEqual("", trackString, "The trackString should not be empty as there were tracks to traverse.")
	suite.Contains(trackString, "first", "The traverse method should have visited the first track.")
	suite.Contains(trackString, "second", "The traverse method should have visited the second track.")
}

func (suite *QueueTestSuite) TestShuffleTracks() {
	DJ.Queue.AddTrack(suite.FirstTrack)

	DJ.Queue.ShuffleTracks()
	suite.Equal(suite.FirstTrack, DJ.Queue.Queue[0], "Shuffle shouldn't do anything when only one track is in the queue.")

	DJ.Queue.AddTrack(suite.SecondTrack)

	DJ.Queue.ShuffleTracks()
	suite.Equal(suite.FirstTrack, DJ.Queue.Queue[0], "Shuffle shouldn't do anything when only two tracks are in the queue.")
	suite.Equal(suite.SecondTrack, DJ.Queue.Queue[1], "Shuffle shouldn't do anything when only two tracks are in the queue.")

	DJ.Queue.AddTrack(suite.ThirdTrack)

	for i := 0; i < 10; i++ {
		DJ.Queue.AddTrack(&Track{ID: fmt.Sprintf("%d", i+4)})
	}

	originalQueue := make([]interfaces.Track, len(DJ.Queue.Queue))
	copy(originalQueue, DJ.Queue.Queue)
	DJ.Queue.ShuffleTracks()
	suite.NotEqual(originalQueue, DJ.Queue.Queue, "The shuffled queue should not be the same as the original queue.")
}

func (suite *QueueTestSuite) TestRandomNextTrackWhenQueueWasEmpty() {
	DJ.Queue.AddTrack(suite.FirstTrack)

	DJ.Queue.RandomNextTrack(true)

	suite.Equal(suite.FirstTrack, DJ.Queue.Queue[0], "RandomNextTrack shouldn't do anything when there is only one track in the queue.")

	for i := 0; i < 25; i++ {
		DJ.Queue.AddTrack(&Track{ID: fmt.Sprintf("%d", i+1)})
	}

	DJ.Queue.RandomNextTrack(true)

	suite.NotEqual(suite.FirstTrack, DJ.Queue.Queue[0], "The first track should no longer be the same.")
}

func (suite *QueueTestSuite) TestRandomNextTrackWhenQueueWasNotEmpty() {
	DJ.Queue.AddTrack(suite.FirstTrack)
	DJ.Queue.RandomNextTrack(false)

	suite.Equal(suite.FirstTrack, DJ.Queue.Queue[0], "RandomNextTrack shouldn't do anything when there is only one track in the queue.")

	DJ.Queue.AddTrack(suite.SecondTrack)
	DJ.Queue.RandomNextTrack(false)

	suite.Equal(suite.FirstTrack, DJ.Queue.Queue[0], "RandomNextTrack shouldn't do anything when there is only two tracks in the queue and the queue was not previously empty.")
	suite.Equal(suite.SecondTrack, DJ.Queue.Queue[1], "RandomNextTrack shouldn't do anything when there is only two tracks in the queue and the queue was not previously empty.")

	for i := 0; i < 25; i++ {
		DJ.Queue.AddTrack(&Track{ID: fmt.Sprintf("%d", i+2)})
	}

	DJ.Queue.RandomNextTrack(false)

	suite.Equal(suite.FirstTrack, DJ.Queue.Queue[0], "Since the queue was not previously empty the first track should not be touched.")
	suite.NotEqual(suite.SecondTrack, DJ.Queue.Queue[1], "The next track should be randomized.")
}

func (suite *QueueTestSuite) TestSkipWhenQueueHasLessThanTwoTracks() {
	DJ.Queue.AddTrack(suite.FirstTrack)
	suite.Equal(1, len(DJ.Queue.Queue), "There should be one item in the queue.")

	DJ.Queue.Skip()
	suite.Zero(len(DJ.Queue.Queue), "There should now be zero items in the queue.")
}

func (suite *QueueTestSuite) TestSkipWhenQueueHasTwoOrMoreTracks() {
	DJ.Queue.AddTrack(suite.FirstTrack)
	DJ.Queue.AddTrack(suite.SecondTrack)

	suite.Equal(suite.FirstTrack, DJ.Queue.Queue[0], "The track added first should be at the front of the queue.")
	suite.Equal(2, len(DJ.Queue.Queue), "There should be two items in the queue.")

	DJ.Queue.Skip()

	suite.Equal(suite.SecondTrack, DJ.Queue.Queue[0], "The track added second should be at the front of the queue.")
	suite.Equal(1, len(DJ.Queue.Queue), "There should be one item in the queue.")
}

func (suite *QueueTestSuite) TestSkipPlaylistWhenFirstTrackIsNotPartOfPlaylist() {
	DJ.Queue.AddTrack(suite.FirstTrack)
	DJ.Queue.AddTrack(suite.SecondTrack)
	DJ.Queue.AddTrack(suite.ThirdTrack)

	DJ.Queue.SkipPlaylist()

	suite.Equal(3, len(DJ.Queue.Queue), "No tracks should be skipped.")
}

func (suite *QueueTestSuite) TestSkipPlaylistWhenFirstTrackIsPartOfPlaylist() {
	playlist := &Playlist{ID: "playlist"}
	track1 := &Track{Playlist: playlist}
	track2 := &Track{Playlist: playlist}
	track3 := &Track{}

	DJ.Queue.AddTrack(track1)
	DJ.Queue.AddTrack(track2)
	DJ.Queue.AddTrack(track3)

	suite.Equal(3, len(DJ.Queue.Queue), "There should be three tracks in the queue.")
	DJ.Queue.SkipPlaylist()
	suite.Equal(1, len(DJ.Queue.Queue), "There should be one track remaining in the queue.")
}

func (suite *QueueTestSuite) TestSkipPlaylistWhenPlaylistIsShuffled() {
	playlist := &Playlist{ID: "playlist"}
	otherPlaylist := &Playlist{ID: "otherplaylist"}
	track1 := &Track{Playlist: playlist}
	track2 := &Track{}
	track3 := &Track{Playlist: otherPlaylist}
	track4 := &Track{Playlist: playlist}

	DJ.Queue.AddTrack(track1)
	DJ.Queue.AddTrack(track2)
	DJ.Queue.AddTrack(track3)
	DJ.Queue.AddTrack(track4)

	suite.Equal(4, len(DJ.Queue.Queue), "There should be four tracks in the queue.")
	DJ.Queue.SkipPlaylist()
	suite.Equal(2, len(DJ.Queue.Queue), "There should be two tracks remaining in the queue.")
}

func TestQueueTestSuite(t *testing.T) {
	suite.Run(t, new(QueueTestSuite))
}
