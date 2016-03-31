/*
 * MumbleDJ
 * By Matthieu Grieger
 * track.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package main

import "time"

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
	Playlist     *Playlist
}
