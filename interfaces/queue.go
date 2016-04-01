/*
 * MumbleDJ
 * By Matthieu Grieger
 * interfaces/queue.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package interfaces

// Queue is an interface of methods that must be implemented by queue manangers.
type Queue interface {
	AddTracks(t ...Track) error
	CurrentTrack() (Track, error)
	PeekNextTrack() (Track, error)
	Traverse(func(int, Track))
	ShuffleTracks()
	NextTrack()
	RandomNextTrack(bool)
	Skip()
	SkipPlaylist()
}
