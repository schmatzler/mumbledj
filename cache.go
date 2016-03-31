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

// GetCurrentStatistics retrieves the total file size and number of files
// stored in the cache and updates the member variables accordingly.
func (c *Cache) GetCurrentStatistics() (int, int64) {
	return 0, 0
}
