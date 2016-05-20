/*
 * MumbleDJ
 * By Matthieu Grieger
 * services/mixcloud.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package services

// Mixcloud is a wrapper around the Mixcloud API.
// https://www.mixcloud.com/developers/
type Mixcloud struct {
}

// CheckAPIKey performs a test API call with the API key
// provided in the configuration file to determine if the
// service should be enabled.
func (mc *Mixcloud) CheckAPIKey() error {
	return nil
}
