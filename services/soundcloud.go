/*
 * MumbleDJ
 * By Matthieu Grieger
 * services/soundcloud.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package services

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/matthieugrieger/mumbledj/bot"
)

// SoundCloud is a wrapper around the SoundCloud API.
// https://developers.soundcloud.com/docs/api/reference
type SoundCloud struct {
	*GenericService
}

// NewSoundCloudService returns an initialized SoundCloud service object.
func NewSoundCloudService() *SoundCloud {
	return &SoundCloud{
		&GenericService{
			ReadableName: "SoundCloud",
			TrackRegex: []string{
				`https?:\/\/(www\.)?soundcloud\.com\/([\w-]+)\/([\w-]+)(#t=\n\n?(:\n\n)*)?`,
			},
			PlaylistRegex: []string{
				`https?:\/\/(www\.)?soundcloud\.com\/([\w-]+)\/sets\/([\w-]+)`,
			},
		},
	}
}

// CheckAPIKey performs a test API call with the API key
// provided in the configuration file to determine if the
// service should be enabled.
func (sc *SoundCloud) CheckAPIKey() error {
	if DJ.BotConfig.API.SoundCloud == "" {
		return errors.New("No SoundCloud API key has been provided")
	}
	url := "http://api.soundcloud.com/tracks/vjflzpbkmerb?client_id=%s"
	_, err := http.Get(fmt.Sprintf(url, DJ.BotConfig.API.SoundCloud))
	return err
}

// CheckURL matches the passed URL with a list of regex patterns
// for valid URLs associated with this service. Returns true if a
// match is found, false otherwise.
func (sc *SoundCloud) CheckURL(url string) bool {
	if sc.isTrack(url) || sc.isPlaylist(url) {
		return true
	}
	return false
}

// GetTracks uses the passed URL to find and return
// tracks associated with the URL. An error is returned
// if any error occurs during the API call.
func (sc *SoundCloud) GetTracks(url string) ([]bot.Track, error) {
	return nil, nil
}
