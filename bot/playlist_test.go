/*
 * MumbleDJ
 * By Matthieu Grieger
 * bot/playlist_test.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package bot

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type PlaylistTestSuite struct {
	Playlist Playlist
	suite.Suite
}

func (suite *PlaylistTestSuite) SetupTest() {
	duration, _ := time.ParseDuration("1s")
	suite.Playlist = Playlist{
		ID:        "id",
		Title:     "title",
		Author:    "author",
		Submitter: "submitter",
		Service:   "service",
		Duration:  duration,
		NumTracks: 1,
	}
}

func (suite *PlaylistTestSuite) TestGetID() {
	suite.Equal("id", suite.Playlist.GetID())
}

func (suite *PlaylistTestSuite) TestGetTitle() {
	suite.Equal("title", suite.Playlist.GetTitle())
}

func (suite *PlaylistTestSuite) TestGetAuthor() {
	suite.Equal("author", suite.Playlist.GetAuthor())
}

func (suite *PlaylistTestSuite) TestGetSubmitter() {
	suite.Equal("submitter", suite.Playlist.GetSubmitter())
}

func (suite *PlaylistTestSuite) TestGetService() {
	suite.Equal("service", suite.Playlist.GetService())
}

func (suite *PlaylistTestSuite) TestGetDuration() {
	duration, _ := time.ParseDuration("1s")
	suite.Equal(duration, suite.Playlist.GetDuration())
}

func (suite *PlaylistTestSuite) TestGetNumTracks() {
	suite.Equal(1, suite.Playlist.GetNumTracks())
}

func TestPlaylistTestSuite(t *testing.T) {
	suite.Run(t, new(PlaylistTestSuite))
}
