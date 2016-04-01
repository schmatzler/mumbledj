/*
 * MumbleDJ
 * By Matthieu Grieger
 * interfaces/skiptracker.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package interfaces

import "github.com/layeh/gumble/gumble"

// SkipTracker is an interface of methods that must be implemented in a
// skip tracker.
type SkipTracker interface {
	AddTrackSkip(*gumble.User) error
	AddPlaylistSkip(*gumble.User) error
	RemoveTrackSkip(*gumble.User) error
	RemovePlaylistSkip(*gumble.User) error
	ResetTrackSkips()
	ResetPlaylistSkips()
}
