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
	GetID() string
	GetTitle() string
	GetAuthor() string
	GetSubmitter() string
	GetService() string
	GetFilename() (string, error)
	GetThumbnailURL() (string, error)
	GetDuration() time.Duration
	GetPlaylist() (Playlist, error)
}
