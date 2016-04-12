/*
 * MumbleDJ
 * By Matthieu Grieger
 * bot/youtube_dl.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package bot

import "github.com/matthieugrieger/mumbledj/interfaces"

// YouTubeDL is a struct that gathers all methods related to the youtube-dl
// software.
// youtube-dl: https://rg3.github.io/youtube-dl/
type YouTubeDL struct {
}

// GetTracks returns track objects retrieved from the provided URL.
func (yt *YouTubeDL) GetTracks(url string) ([]interfaces.Track, error) {
	return nil, nil
}
