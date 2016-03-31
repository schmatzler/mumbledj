/*
 * MumbleDJ
 * By Matthieu Grieger
 * playlist.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package main

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
