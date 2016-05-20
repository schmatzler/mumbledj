/*
 * MumbleDJ
 * By Matthieu Grieger
 * bot/youtube_dl.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package bot

import (
	"errors"
	"os/exec"
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
