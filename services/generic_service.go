/*
 * MumbleDJ
 * By Matthieu Grieger
 * services/generic_service.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package services

import "regexp"

// GenericService is a generic struct that should be embedded
// in other service structs, as it provides useful helper
// methods and properties.
type GenericService struct {
	ReadableName  string
	TrackRegex    []string
	PlaylistRegex []string
}

// CheckURL matches the passed URL with a list of regex patterns
// for valid URLs associated with this service. Returns true if a
// match is found, false otherwise.
func (gs *GenericService) CheckURL(url string) bool {
	if gs.isTrack(url) || gs.isPlaylist(url) {
		return true
	}
	return false
}

func (gs *GenericService) isTrack(url string) bool {
	for _, regex := range gs.TrackRegex {
		re, _ := regexp.Compile(regex)
		if re.MatchString(url) {
			return true
		}
	}
	return false
}

func (gs *GenericService) isPlaylist(url string) bool {
	for _, regex := range gs.PlaylistRegex {
		re, _ := regexp.Compile(regex)
		if re.MatchString(url) {
			return true
		}
	}
	return false
}

// TODO: Implement this method!
func (gs *GenericService) getID(url string) (string, error) {
	allRegex := append(gs.TrackRegex, gs.PlaylistRegex...)
	return "", nil
}
