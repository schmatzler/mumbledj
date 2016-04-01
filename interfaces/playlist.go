/*
 * MumbleDJ
 * By Matthieu Grieger
 * interfaces/playlist.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package interfaces

// Playlist is an interface of methods that must be implemented by playlists.
type Playlist interface {
	ID() string
	Title() string
	Author() string
	Submitter() string
	Service() string
	Duration() string
	NumTracks() int
}
