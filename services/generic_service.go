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
