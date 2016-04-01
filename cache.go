/*
 * MumbleDJ
 * By Matthieu Grieger
 * cache.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package main

import (
	"os"
	"time"
)

// SortFilesByAge is a type that holds file information for cached items for
// sorting.
type SortFilesByAge []os.FileInfo

// Len returns the length of the file slice.
func (a SortFilesByAge) Len() int {
	return len(a)
}

// Swap swaps two elements in the file slice.
func (a SortFilesByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// Less compares two file modification times to determine if one is less than
// the other. Returns true if the item in index i is older than the item in
// index j, false otherwise.
func (a SortFilesByAge) Less(i, j int) bool {
	return time.Since(a[i].ModTime()) < time.Since(a[j].ModTime())
}

// Cache keeps track of the filesize of the audio cache and
// provides methods for pruning the cache.
type Cache struct {
	NumAudioFiles int
	TotalFileSize int64
}

// NewCache creates an empty Cache and returns it.
func NewCache() *Cache {
	return &Cache{
		NumAudioFiles: 0,
		TotalFileSize: 0,
	}
}

// CheckDirectorySize checks the cache directory to determine if the filesize
// of the files within exceed the user-specified size limit. If so, the oldest
// files are cleared until it is no longer exceeding the limit.
// TODO: Implement after Config.
func (c *Cache) CheckDirectorySize() {

}

// UpdateStatistics updates the statistics relevant to the cache (number of
// audio files cached, total current size of the cache).
func (c *Cache) UpdateStatistics() {
	c.NumAudioFiles, c.TotalFileSize = c.getCurrentStatistics()
}

// CleanPeriodically loops forever, deleting expired cached audio files as necessary.
// TODO: Implement after Config.
func (c *Cache) CleanPeriodically() {

}

// DeleteOldest deletes the oldest file in the cache.
// TODO: Implement after Config.
func (c *Cache) DeleteOldest() error {
	return nil
}

// DeleteAll deletes all cached audio files.
// TODO: Implement after Config.
func (c *Cache) DeleteAll() error {
	return nil
}

func (c *Cache) getCurrentStatistics() (int, int64) {
	return 0, 0
}
