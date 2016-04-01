/*
 * MumbleDJ
 * By Matthieu Grieger
 * track.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package main

import (
	"errors"
	"time"

	"github.com/matthieugrieger/mumbledj/interfaces"
)

// Track stores all metadata related to an audio track.
type Track struct {
	tID           string
	tTitle        string
	tAuthor       string
	tSubmitter    string
	tService      string
	tFilename     string
	tThumbnailURL string
	tDuration     time.Duration
	tPlaylist     interfaces.Playlist
}

// ID returns the ID of the track.
func (t *Track) ID() string {
	return t.tID
}

// Title returns the title of the track.
func (t *Track) Title() string {
	return t.tTitle
}

// Author returns the author of the track.
func (t *Track) Author() string {
	return t.tAuthor
}

// Submitter returns the submitter of the track.
func (t *Track) Submitter() string {
	return t.tSubmitter
}

// Service returns the name of the service from which the track was retrieved from.
func (t *Track) Service() string {
	return t.tService
}

// Filename returns the name of the file stored on disk, if it exists. If no
// file on disk exists an empty string and error are returned.
func (t *Track) Filename() (string, error) {
	if t.tFilename != "" {
		return t.tFilename, nil
	}
	return "", errors.New("This track is not currently stored on disk")
}

// ThumbnailURL returns the URL to the thumbnail for the track. If no thumbnail
// exists an empty string and error are returned.
func (t *Track) ThumbnailURL() (string, error) {
	if t.tThumbnailURL != "" {
		return t.tThumbnailURL, nil
	}
	return "", errors.New("This track does not have a thumbnail")
}

// Duration returns the duration of the track.
func (t *Track) Duration() time.Duration {
	return t.tDuration
}

// Playlist returns the playlist the track is associated with, if it exists. If
// the track is not associated with a playlist a nil playlist and error are returned.
func (t *Track) Playlist() (interfaces.Playlist, error) {
	if t.tPlaylist != nil {
		return t.tPlaylist, nil
	}
	return nil, errors.New("This track is not associated with a playlist")
}
