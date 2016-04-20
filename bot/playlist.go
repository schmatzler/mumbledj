/*
 * MumbleDJ
 * By Matthieu Grieger
 * bot/playlist.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package bot

import "time"

// Playlist stores all metadata related to a playlist of tracks.
type Playlist struct {
	ID        string
	Title     string
	Author    string
	Submitter string
	Service   string
	Duration  time.Duration
	NumTracks int
}

// GetID returns the ID of the playlist.
func (p *Playlist) GetID() string {
	return p.ID
}

// GetTitle returns the title of the playlist.
func (p *Playlist) GetTitle() string {
	return p.Title
}

// GetAuthor returns the author of the playlist.
func (p *Playlist) GetAuthor() string {
	return p.Author
}

// GetSubmitter returns the submitter of the playlist.
func (p *Playlist) GetSubmitter() string {
	return p.Submitter
}

// GetService returns the name of the service from which the playlist was retrieved from.
func (p *Playlist) GetService() string {
	return p.Service
}

// GetDuration returns the duration of the playlist.
func (p *Playlist) GetDuration() time.Duration {
	return p.Duration
}

// GetNumTracks returns the number of tracks in the playlist.
func (p *Playlist) GetNumTracks() int {
	return p.NumTracks
}
