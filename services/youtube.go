/*
 * MumbleDJ
 * By Matthieu Grieger
 * services/youtube.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package services

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"

	"github.com/matthieugrieger/mumbledj/bot"
)

// YouTube is a wrapper around the YouTube Data API.
// https://developers.google.com/youtube/v3/docs/
type YouTube struct {
	ReadableName  string
	TrackRegex    []string
	PlaylistRegex []string
}

// NewYouTubeService returns an initialized YouTube service object.
func NewYouTubeService() *YouTube {
	return &YouTube{
		ReadableName: "YouTube",
		TrackRegex: []string{
			`https?:\/\/www\.youtube\.com\/watch\?v=([\w-]+)(\&t=\d*m?\d*s?)?`,
			`https?:\/\/youtube\.com\/watch\?v=([\w-]+)(\&t=\d*m?\d*s?)?`,
			`https?:\/\/youtu.be\/([\w-]+)(\?t=\d*m?\d*s?)?`,
			`https?:\/\/youtube.com\/v\/([\w-]+)(\?t=\d*m?\d*s?)?`,
			`https?:\/\/www.youtube.com\/v\/([\w-]+)(\?t=\d*m?\d*s?)?`,
		},
		PlaylistRegex: []string{
			`https?:\/\/www\.youtube\.com\/playlist\?list=([\w-]+)`,
		},
	}
}

// CheckAPIKey performs a test API call with the API key
// provided in the configuration file to determine if the
// service should be enabled.
func (yt *YouTube) CheckAPIKey() error {
	if DJ.BotConfig.API.YouTube == "" {
		return errors.New("No YouTube API key has been provided")
	}
	url := "https://www.googleapis.com/youtube/v3/videos?part=snippet&id=KQY9zrjPBjo&key=%s"
	_, err := http.Get(fmt.Sprintf(url, DJ.BotConfig.API.YouTube))
	return err
}

// CheckURL matches the passed URL with a list of regex patterns
// for valid URLs associated with this service. Returns true if a
// match is found, false otherwise.
func (yt *YouTube) CheckURL(url string) bool {
	if yt.isTrack(url) || yt.isPlaylist(url) {
		return true
	}
	return false
}

// GetTracks uses the passed URL to find and return
// tracks associated with the URL. An error is returned
// if any error occurs during the API call.
func (yt *YouTube) GetTracks(url string) ([]bot.Track, error) {
	return nil, nil
}

func (yt *YouTube) isTrack(url string) bool {
	for _, regex := range yt.TrackRegex {
		re, _ := regexp.Compile(regex)
		if re.MatchString(url) {
			return true
		}
	}
	return false
}

func (yt *YouTube) isPlaylist(url string) bool {
	for _, regex := range yt.PlaylistRegex {
		re, _ := regexp.Compile(regex)
		if re.MatchString(url) {
			return true
		}
	}
	return false
}
