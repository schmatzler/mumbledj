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

	"github.com/matthieugrieger/mumbledj/bot"
)

// YouTube is a wrapper around the YouTube Data API.
// https://developers.google.com/youtube/v3/docs/
type YouTube struct {
	*GenericService
}

// NewYouTubeService returns an initialized YouTube service object.
func NewYouTubeService() *YouTube {
	return &YouTube{
		&GenericService{
			ReadableName: "YouTube",
			TrackRegex: []string{
				`https?:\/\/www.youtube.com\/watch\?v=(?P<id>[\w-]+)(?P<timestamp>\&t=\d*m?\d*s?)?`,
				`https?:\/\/youtube.com\/watch\?v=(?P<id>[\w-]+)(?P<timestamp>\&t=\d*m?\d*s?)?`,
				`https?:\/\/youtu.be\/(?P<id>[\w-]+)(?P<timestamp>\?t=\d*m?\d*s?)?`,
				`https?:\/\/youtube.com\/v\/(?P<id>[\w-]+)(?P<timestamp>\?t=\d*m?\d*s?)?`,
				`https?:\/\/www.youtube.com\/v\/(?P<id>[\w-]+)(?P<timestamp>\?t=\d*m?\d*s?)?`,
			},
			PlaylistRegex: []string{
				`https?:\/\/www\.youtube\.com\/playlist\?list=(?P<id>[\w-]+)`,
			},
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

// GetTracks uses the passed URL to find and return
// tracks associated with the URL. An error is returned
// if any error occurs during the API call.
func (yt *YouTube) GetTracks(url string) ([]bot.Track, error) {
	var (
		videoURL         string
		playlistURL      string
		playlistItemsURL string
		id               string
		err              error
		resp             *http.Response
		v                *jason.Object
	)

	videoURL = "https://www.googleapis.com/youtube/v3/videos?part=snippet,contentDetails&id=%s&key=%s"
	playlistURL = "https://www.googleapis.com/youtube/v3/playlists?part=snippet&id=%s&key=%s"
	playlistItemsURL = "https://www.googleapis.com/youtube/v3/playlistItems?part=snippet,contentDetails&playlistId=%s&key=%s"

	if yt.isPlaylist(url) {
		id = yt.getID(url)

		resp, err = http.Get(fmt.Sprintf(playlistURL, id, DJ.BotConfig.API.YouTube))
		defer resp.Close()
		if err != nil {
			return nil, err
		}

		v, err = jason.NewObjectFromReader(resp.Body)
		if err != nil {
			return nil, err
		}

		title, _ := v.GetString("items", "0", "snippet", "title")
	}
}
