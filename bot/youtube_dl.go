/*
 * MumbleDJ
 * By Matthieu Grieger
 * bot/youtube_dl.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package bot

import (
	"errors"
	"fmt"
	"os/exec"
	"sync"
	"time"

	"github.com/antonholmquist/jason"
	"github.com/layeh/gumble/gumble"
)

// YouTubeDL is a struct that gathers all methods related to the youtube-dl
// software.
// youtube-dl: https://rg3.github.io/youtube-dl/
type YouTubeDL struct {
}

// CheckInstallation attempts to execute the version command for youtube-dl.
// If this command fails then youtube-dl is either not installed or is not
// accessible in the user's path.
func (yt *YouTubeDL) CheckInstallation() error {
	command := exec.Command("youtube-dl", "--version")
	if err := command.Run(); err != nil {
		return errors.New("youtube-dl is not properly installed")
	}
	return nil
}

// GetTracks returns track objects retrieved from the provided URL.
func (yt *YouTubeDL) GetTracks(url string, user *gumble.User) ([]*Track, error) {
	var (
		jsonBytes  []byte
		err        error
		json       *jason.Object
		extractor  string
		isPlaylist bool
		tracks     []*Track
	)

	command := exec.Command("youtube-dl", "--dump-json", "--flat-playlist", url)
	jsonBytes, err = command.Output()
	if err != nil {
		return nil, err
	}

	json, err = jason.NewObjectFromBytes(jsonBytes)
	if err != nil {
		return nil, err
	}

	extractor, err = json.GetString("extractor")
	if err != nil {
		// We may have a playlist, which has a different json key; check for it.
		// If successful, we can set isPlaylist to true.
		extractor, err = json.GetString("ie_key")
		if err != nil {
			return nil, errors.New("The service you attempted to use is not supported by youtube-dl")
		}
		isPlaylist = true
	}

	if !yt.isWhitelisted(extractor) {
		return nil, errors.New("The service you attempted to use is not whitelisted")
	}

	if isPlaylist {
		// Multiple webpages must be fetched and parsed. To speed this up, separate
		// goroutines are spawned to handle each track of the playlist.
		var trackObjects []*jason.Object
		var waitGroup sync.WaitGroup

	} else {
		// Only one webpage must be fetched and parsed, no goroutines are necessary.
		track := yt.createTrack(json, user, nil)
		tracks = append(tracks, track)
	}

	return tracks, nil
}

func (yt *YouTubeDL) isWhitelisted(extractor string) bool {
	for _, whitelistedExtractor := range DJ.BotConfig.Services.Whitelist {
		if extractor == whitelistedExtractor {
			return true
		}
	}
	return false
}

func (yt *YouTubeDL) createTrack(json *jason.Object, user *gumble.User, playlist *Playlist) *Track {
	// TODO: Perform better error handling here.
	id, _ := json.GetString("id")
	title, _ := json.GetString("title")
	author, _ := json.GetString("uploader")
	submitter := user.Name
	service, _ := json.GetString("extractor")
	thumbnailURL, _ := json.GetString("thumbnail")
	durationSeconds, _ := json.GetInt64("duration")
	duration, _ := time.ParseDuration(fmt.Sprintf("%ds", durationSeconds))

	return &Track{
		ID:           id,
		Title:        title,
		Author:       author,
		Submitter:    submitter,
		Service:      service,
		Filename:     "",
		ThumbnailURL: thumbnailURL,
		Duration:     duration,
		Playlist:     playlist,
	}
}
