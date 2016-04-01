/*
 * MumbleDJ
 * By Matthieu Grieger
 * interfaces/cache.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package interfaces

// Cache is an interface of methods that must be implemented by cache managers.
type Cache interface {
	CheckDirectorySize()
	UpdateStatistics()
	CleanPeriodically()
	DeleteOldest() error
	DeleteAll() error
}
