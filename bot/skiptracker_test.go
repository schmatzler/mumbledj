/*
 * MumbleDJ
 * By Matthieu Grieger
 * bot/skiptracker_test.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package bot

import (
	"testing"

	"github.com/layeh/gumble/gumble"
	"github.com/stretchr/testify/suite"
)

type SkipTrackerTestSuite struct {
	suite.Suite
	User1 *gumble.User
	User2 *gumble.User
}

func (suite *SkipTrackerTestSuite) SetupSuite() {
	DJ = NewMumbleDJ()

	suite.User1 = new(gumble.User)
	suite.User1.Name = "User1"
	suite.User2 = new(gumble.User)
	suite.User2.Name = "User2"
}

func (suite *SkipTrackerTestSuite) SetupTest() {
	DJ.Skips = NewSkipTracker()
}

func (suite *SkipTrackerTestSuite) TestNewSkipTracker() {
	suite.Zero(len(DJ.Skips.TrackSkips), "The track skip slice should be empty upon initialization.")
	suite.Zero(len(DJ.Skips.PlaylistSkips), "The playlist skip slice should be empty upon initialization.")
}

func (suite *SkipTrackerTestSuite) TestAddTrackSkip() {
	err := DJ.Skips.AddTrackSkip(suite.User1)

	suite.Equal(1, len(DJ.Skips.TrackSkips), "There should now be one user in the track skip slice.")
	suite.Zero(0, len(DJ.Skips.PlaylistSkips), "The playlist skip slice should be unaffected.")
	suite.Nil(err, "No error should be returned.")

	err = DJ.Skips.AddTrackSkip(suite.User2)

	suite.Equal(2, len(DJ.Skips.TrackSkips), "There should now be two users in the track skip slice.")
	suite.Zero(0, len(DJ.Skips.PlaylistSkips), "The playlist skip slice should be unaffected.")
	suite.Nil(err, "No error should be returned.")

	err = DJ.Skips.AddTrackSkip(suite.User1)

	suite.Equal(2, len(DJ.Skips.TrackSkips), "This is a duplicate skip, so the track skip slice should still only have two users.")
	suite.Zero(0, len(DJ.Skips.PlaylistSkips), "The playlist skip slice should be unaffected.")
	suite.NotNil(err, "An error should be returned since this user has already voted to skip the current track.")
}

func (suite *SkipTrackerTestSuite) TestAddPlaylistSkip() {
	err := DJ.Skips.AddPlaylistSkip(suite.User1)

	suite.Zero(len(DJ.Skips.TrackSkips), "The track skip slice should be unaffected.")
	suite.Equal(1, len(DJ.Skips.PlaylistSkips), "There should now be one user in the playlist skip slice.")
	suite.Nil(err, "No error should be returned.")

	err = DJ.Skips.AddPlaylistSkip(suite.User2)

	suite.Zero(len(DJ.Skips.TrackSkips), "The track skip slice should be unaffected.")
	suite.Equal(2, len(DJ.Skips.PlaylistSkips), "There should now be two users in the playlist skip slice.")
	suite.Nil(err, "No error should be returned.")

	err = DJ.Skips.AddPlaylistSkip(suite.User1)

	suite.Zero(len(DJ.Skips.TrackSkips), "The track skip slice should be unaffected.")
	suite.Equal(2, len(DJ.Skips.PlaylistSkips), "This is a duplicate skip, so the playlist skip slice should still only have two users.")
	suite.NotNil(err, "An error should be returned since this user has already voted to skip the current playlist.")
}

func (suite *SkipTrackerTestSuite) TestRemoveTrackSkip() {
	DJ.Skips.AddTrackSkip(suite.User1)
	err := DJ.Skips.RemoveTrackSkip(suite.User2)

	suite.Equal(1, len(DJ.Skips.TrackSkips), "User2 has not skipped the track so the track skip slice should be unaffected.")
	suite.Zero(len(DJ.Skips.PlaylistSkips), "The playlist skip slice should be unaffected.")
	suite.NotNil(err, "An error should be returned since User2 has not skipped the track yet.")

	err = DJ.Skips.RemoveTrackSkip(suite.User1)

	suite.Zero(len(DJ.Skips.TrackSkips), "User1 skipped the track, so their skip should be removed.")
	suite.Zero(len(DJ.Skips.PlaylistSkips), "The playlist skip slice should be unaffected.")
	suite.Nil(err, "No error should be returned.")
}

func (suite *SkipTrackerTestSuite) TestRemovePlaylistSkip() {
	DJ.Skips.AddPlaylistSkip(suite.User1)
	err := DJ.Skips.RemovePlaylistSkip(suite.User2)

	suite.Zero(len(DJ.Skips.TrackSkips), "The track skip slice should be unaffected.")
	suite.Equal(1, len(DJ.Skips.PlaylistSkips), "User2 has not skipped the playlist so the playlist skip slice should be unaffected.")
	suite.NotNil(err, "An error should be returned since User2 has not skipped the playlist yet.")

	err = DJ.Skips.RemovePlaylistSkip(suite.User1)

	suite.Zero(len(DJ.Skips.TrackSkips), "The track skip slice should be unaffected.")
	suite.Zero(len(DJ.Skips.PlaylistSkips), "User1 skipped the playlist, so their skip should be removed.")
	suite.Nil(err, "No error should be returned.")
}

func (suite *SkipTrackerTestSuite) TestResetTrackSkips() {
	DJ.Skips.AddTrackSkip(suite.User1)
	DJ.Skips.AddTrackSkip(suite.User2)
	DJ.Skips.AddPlaylistSkip(suite.User1)
	DJ.Skips.AddPlaylistSkip(suite.User2)

	suite.Equal(2, len(DJ.Skips.TrackSkips), "There should be two users in the track skip slice.")
	suite.Equal(2, len(DJ.Skips.PlaylistSkips), "There should be two users in the playlist skip slice.")

	DJ.Skips.ResetTrackSkips()

	suite.Zero(len(DJ.Skips.TrackSkips), "The track skip slice has been reset, so the length should be zero.")
	suite.Equal(2, len(DJ.Skips.PlaylistSkips), "The playlist skip slice should be unaffected.")
}

func (suite *SkipTrackerTestSuite) TestResetPlaylistSkips() {
	DJ.Skips.AddTrackSkip(suite.User1)
	DJ.Skips.AddTrackSkip(suite.User2)
	DJ.Skips.AddPlaylistSkip(suite.User1)
	DJ.Skips.AddPlaylistSkip(suite.User2)

	suite.Equal(2, len(DJ.Skips.TrackSkips), "There should be two users in the track skip slice.")
	suite.Equal(2, len(DJ.Skips.PlaylistSkips), "There should be two users in the playlist skip slice.")

	DJ.Skips.ResetPlaylistSkips()

	suite.Equal(2, len(DJ.Skips.TrackSkips), "The track skip slice should be unaffected.")
	suite.Zero(len(DJ.Skips.PlaylistSkips), "The playlist skip slice has been reset, so the length should be zero.")
}

func TestSkipTrackerTestSuite(t *testing.T) {
	suite.Run(t, new(SkipTrackerTestSuite))
}
