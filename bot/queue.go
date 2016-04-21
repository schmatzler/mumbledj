/*
 * MumbleDJ
 * By Matthieu Grieger
 * bot/queue.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package bot

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/matthieugrieger/mumbledj/interfaces"
)

// Queue holds the audio queue itself along with useful methods for
// performing actions on the queue.
type Queue struct {
	Queue []interfaces.Track
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// NewQueue initializes a new queue and returns it.
func NewQueue() *Queue {
	return &Queue{
		Queue: make([]interfaces.Track, 0),
	}
}

// AddTrack adds a track to the back of the queue.
func (q *Queue) AddTrack(t interfaces.Track) error {
	beforeLen := len(q.Queue)

	// An error should never occur here since maxTrackDuration is restricted to
	// ints. Any error in the configuration will be caught during yaml load.
	maxTrackDuration, _ := time.ParseDuration(fmt.Sprintf("%ds",
		DJ.BotConfig.General.MaxTrackDuration))

	if DJ.BotConfig.General.MaxTrackDuration == 0 ||
		t.GetDuration() <= maxTrackDuration {
		q.Queue = append(q.Queue, t)
	} else {
		return errors.New("The track is too long to add to the queue")
	}
	if len(q.Queue) == beforeLen+1 {
		return nil
	}
	return errors.New("Could not add track to queue")
}

// CurrentTrack returns the current Track.
func (q *Queue) CurrentTrack() (interfaces.Track, error) {
	if len(q.Queue) != 0 {
		return q.Queue[0], nil
	}
	return nil, errors.New("There are no tracks currently in the queue")
}

// PeekNextTrack peeks at the next track and returns it.
func (q *Queue) PeekNextTrack() (interfaces.Track, error) {
	if len(q.Queue) > 1 {
		if DJ.BotConfig.General.AutomaticShuffleOn {
			q.RandomNextTrack(false)
		}
		return q.Queue[1], nil
	}
	return nil, errors.New("There is no track coming up next")
}

// Traverse is a traversal function for Queue. Allows a visit function to
// be passed in which performs the specified action on each queue item.
func (q *Queue) Traverse(visit func(i int, t interfaces.Track)) {
	if len(q.Queue) > 0 {
		for queueIndex, queueTrack := range q.Queue {
			visit(queueIndex, queueTrack)
		}
	}
}

// ShuffleTracks shuffles the queue using an inside-out algorithm.
func (q *Queue) ShuffleTracks() {
	// Skip the first track, as it is likely playing.
	for i := range q.Queue[1:] {
		j := rand.Intn(i + 1)
		q.Queue[i+1], q.Queue[j+1] = q.Queue[j+1], q.Queue[i+1]
	}
}

// RandomNextTrack sets a random track as the next track to be played.
func (q *Queue) RandomNextTrack(queueWasEmpty bool) {
	if len(q.Queue) > 1 {
		nextTrackIndex := 1
		if queueWasEmpty {
			nextTrackIndex = 0
		}
		swapIndex := nextTrackIndex + rand.Intn(len(q.Queue)-1)
		q.Queue[nextTrackIndex], q.Queue[swapIndex] = q.Queue[swapIndex], q.Queue[nextTrackIndex]
	}
}

// Skip performs the necessary actions that take place when a track is skipped
// via a command.
func (q *Queue) Skip() {
	if len(q.Queue) > 1 {
		q.nextTrack()
	} else {
		q.Queue = make([]interfaces.Track, 0)
	}
}

// SkipPlaylist performs the necessary actions that take place when a playlist
// is skipped via a command.
func (q *Queue) SkipPlaylist() {
	if playlist, err := q.Queue[0].GetPlaylist(); err == nil {
		currentPlaylistID := playlist.GetID()

		// We must loop backwards to prevent missing any elements after deletion.
		for i := len(q.Queue) - 1; i >= 0; i-- {
			if otherTrackPlaylist, err := q.Queue[i].GetPlaylist(); err == nil {
				if otherTrackPlaylist.GetID() == currentPlaylistID {
					q.Queue = append(q.Queue[:i], q.Queue[i+1:]...)
				}
			}
		}
	}
}

func (q *Queue) nextTrack() {
	q.Queue = q.Queue[1:]
}
