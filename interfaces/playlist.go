/*
 * MumbleDJ
 * By Matthieu Grieger
 * interfaces/playlist.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package interfaces

import "time"

// Playlist is an interface of methods that must be implemented by playlists.
type Playlist interface {
	GetID() string
	GetTitle() string
	GetAuthor() string
	GetSubmitter() string
	GetService() string
	GetDuration() time.Duration
	GetNumTracks() int
}
