/*
 * MumbleDJ
 * By Matthieu Grieger
 * services/mixcloud.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package services

import (
	"regexp"

	"github.com/matthieugrieger/mumbledj/bot"
)

// Mixcloud is a wrapper around the Mixcloud API.
// https://www.mixcloud.com/developers/
type Mixcloud struct {
	ReadableName  string
	TrackRegex    []string
	PlaylistRegex []string
}

// NewMixcloudService returns an initialized Mixcloud service object.
func NewMixcloudService() *Mixcloud {
	return &Mixcloud{
		ReadableName: "Mixcloud",
		TrackRegex: []string{
			`https?:\/\/(www\.)?mixcloud\.com\/([\w-]+)\/([\w-]+)(#t=\n\n?(:\n\n)*)?`,
		},
		PlaylistRegex: []string{
			"",
		},
	}
}

// CheckAPIKey performs a test API call with the API key
// provided in the configuration file to determine if the
// service should be enabled.
func (mc *Mixcloud) CheckAPIKey() error {
	// Mixcloud (at the moment) does not require an API key,
	// so we can just return nil.
	return nil
}

// CheckURL matches the passed URL with a list of regex patterns
// for valid URLs associated with this service. Returns true if a
// match is found, false otherwise.
func (mc *Mixcloud) CheckURL(url string) bool {
	if mc.isTrack(url) || mc.isPlaylist(url) {
		return true
	}
	return false
}

// GetTracks uses the passed URL to find and return
// tracks associated with the URL. An error is returned
// if any error occurs during the API call.
func (mc *Mixcloud) GetTracks(url string) ([]bot.Track, error) {
	return nil, nil
}

func (mc *Mixcloud) isTrack(url string) bool {
	for _, regex := range mc.TrackRegex {
		re, _ := regexp.Compile(regex)
		if re.MatchString(url) {
			return true
		}
	}
	return false
}

func (mc *Mixcloud) isPlaylist(url string) bool {
	for _, regex := range mc.PlaylistRegex {
		re, _ := regexp.Compile(regex)
		if re.MatchString(url) {
			return true
		}
	}
	return false
}
