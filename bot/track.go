/*
 * MumbleDJ
 * By Matthieu Grieger
 * bot/track.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package bot

import (
	"errors"
	"time"

	"github.com/matthieugrieger/mumbledj/interfaces"
)

// Track stores all metadata related to an audio track.
type Track struct {
	ID           string
	Title        string
	Author       string
	Submitter    string
	Service      string
	Filename     string
	ThumbnailURL string
	Duration     time.Duration
	Playlist     interfaces.Playlist
}

// GetID returns the ID of the track.
func (t *Track) GetID() string {
	return t.ID
}

// GetTitle returns the title of the track.
func (t *Track) GetTitle() string {
	return t.Title
}

// GetAuthor returns the author of the track.
func (t *Track) GetAuthor() string {
	return t.Author
}

// GetSubmitter returns the submitter of the track.
func (t *Track) GetSubmitter() string {
	return t.Submitter
}

// GetService returns the name of the service from which the track was retrieved from.
func (t *Track) GetService() string {
	return t.Service
}

// GetFilename returns the name of the file stored on disk, if it exists. If no
// file on disk exists an empty string and error are returned.
func (t *Track) GetFilename() (string, error) {
	if t.Filename != "" {
		return t.Filename, nil
	}
	return "", errors.New("This track is not currently stored on disk")
}

// GetThumbnailURL returns the URL to the thumbnail for the track. If no thumbnail
// exists an empty string and error are returned.
func (t *Track) GetThumbnailURL() (string, error) {
	if t.ThumbnailURL != "" {
		return t.ThumbnailURL, nil
	}
	return "", errors.New("This track does not have a thumbnail")
}

// GetDuration returns the duration of the track.
func (t *Track) GetDuration() time.Duration {
	return t.Duration
}

// GetPlaylist returns the playlist the track is associated with, if it exists. If
// the track is not associated with a playlist a nil playlist and error are returned.
func (t *Track) GetPlaylist() (interfaces.Playlist, error) {
	if t.Playlist != nil {
		return t.Playlist, nil
	}
	return nil, errors.New("This track is not associated with a playlist")
}
