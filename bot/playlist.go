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
	pID        string
	pTitle     string
	pAuthor    string
	pSubmitter string
	pService   string
	pDuration  time.Duration
	pNumTracks int
}

// ID returns the ID of the playlist.
func (p *Playlist) ID() string {
	return p.pID
}

// Title returns the title of the playlist.
func (p *Playlist) Title() string {
	return p.pTitle
}

// Author returns the author of the playlist.
func (p *Playlist) Author() string {
	return p.pAuthor
}

// Submitter returns the submitter of the playlist.
func (p *Playlist) Submitter() string {
	return p.pSubmitter
}

// Service returns the name of the service from which the playlist was retrieved from.
func (p *Playlist) Service() string {
	return p.pService
}

// Duration returns the duration of the playlist.
func (p *Playlist) Duration() time.Duration {
	return p.pDuration
}

// NumTracks returns the number of tracks in the playlist.
func (p *Playlist) NumTracks() int {
	return p.pNumTracks
}
