/*
 * MumbleDJ
 * By Matthieu Grieger
 * bot/skiptracker.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package bot

import (
	"fmt"

	"github.com/layeh/gumble/gumble"
)

// SkipTracker keeps track of the list of users who have skipped the current
// track or playlist.
type SkipTracker struct {
	TrackSkips    []*gumble.User
	PlaylistSkips []*gumble.User
}

// NewSkipTracker returns an empty SkipTracker.
func NewSkipTracker() *SkipTracker {
	return &SkipTracker{
		TrackSkips:    make([]*gumble.User, 0),
		PlaylistSkips: make([]*gumble.User, 0),
	}
}

// AddTrackSkip adds a skip to the SkipTracker for the current track.
func (s *SkipTracker) AddTrackSkip(skipper *gumble.User) error {
	for _, user := range s.TrackSkips {
		if user.Name == skipper.Name {
			return fmt.Errorf("%s has already voted to skip the track", skipper.Name)
		}
	}
	s.TrackSkips = append(s.TrackSkips, skipper)
	return nil
}

// AddPlaylistSkip adds a skip to the SkipTracker for the current playlist.
func (s *SkipTracker) AddPlaylistSkip(skipper *gumble.User) error {
	for _, user := range s.PlaylistSkips {
		if user.Name == skipper.Name {
			return fmt.Errorf("%s has already voted to skip the playlist", skipper.Name)
		}
	}
	s.PlaylistSkips = append(s.PlaylistSkips, skipper)
	return nil
}

// RemoveTrackSkip removes a skip from the SkipTracker for the current track.
func (s *SkipTracker) RemoveTrackSkip(skipper *gumble.User) error {
	for i, user := range s.TrackSkips {
		if user.Name == skipper.Name {
			s.TrackSkips = append(s.TrackSkips[:i], s.TrackSkips[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("%s did not previously vote to skip the track", skipper.Name)
}

// RemovePlaylistSkip removes a skip from the SkipTracker for the current playlist.
func (s *SkipTracker) RemovePlaylistSkip(skipper *gumble.User) error {
	for i, user := range s.PlaylistSkips {
		if user.Name == skipper.Name {
			s.PlaylistSkips = append(s.PlaylistSkips[:i], s.PlaylistSkips[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("%s did not previously vote to skip the playlist", skipper.Name)
}

// ResetTrackSkips resets the skip slice for the current track.
func (s *SkipTracker) ResetTrackSkips() {
	s.TrackSkips = s.TrackSkips[:0]
}

// ResetPlaylistSkips resets the skip slice for the current playlist.
func (s *SkipTracker) ResetPlaylistSkips() {
	s.PlaylistSkips = s.PlaylistSkips[:0]
}
