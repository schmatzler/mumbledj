/*
 * MumbleDJ
 * By Matthieu Grieger
 * interfaces/config.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package interfaces

// Config is an interface of methods that must be implemented by config managers.
type Config interface {
	LoadFromConfigFile(string) error
	LoadFromCommandline()
}
