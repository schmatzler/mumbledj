/*
 * MumbleDJ
 * By Matthieu Grieger
 * services/youtube.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package services

// YouTube is a wrapper around the YouTube Data API.
// https://developers.google.com/youtube/v3/docs/
type YouTube struct {
}

// CheckAPIKey performs a test API call with the API key
// provided in the configuration file to determine if the
// service should be enabled.
func (yt *YouTube) CheckAPIKey() error {
	return nil
}
