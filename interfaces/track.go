/*
 * MumbleDJ
 * By Matthieu Grieger
 * interfaces/track.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package interfaces

import "time"

// Track is an interface of methods that must be implemented by tracks.
type Track interface {
	ID() string
	Title() string
	Author() string
	Submitter() string
	Service() string
	Filename() (string, error)
	ThumbnailURL() (string, error)
	Duration() time.Duration
	Playlist() (Playlist, error)
}
