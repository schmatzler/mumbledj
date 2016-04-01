/*
 * MumbleDJ
 * By Matthieu Grieger
 * config.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package main

// GeneralConfig holds general configuration values.
type GeneralConfig struct {
	CommandPrefix        string
	SkipRatio            float32
	PlaylistSkipRatio    float32
	DefaultComment       string
	DefaultChannel       []string
	MaxTrackDuration     int
	MaxTracksPerPlaylist int
	AutomaticShuffleOn   bool
	PlayerCommand        string
}

// ConnectionConfig holds connection configuration values.
type ConnectionConfig struct {
	Address       string
	Port          string
	Password      string
	Username      string
	Insecure      bool
	Cert          string
	Key           string
	AccessTokens  []string
	RetryEnabled  bool
	RetryAttempts int
	RetryInterval int
}

// VolumeConfig holds volume configuration values.
type VolumeConfig struct {
	Default float32
	Lowest  float32
	Highest float32
}

// CacheConfig holds cache configuration values.
type CacheConfig struct {
	Enabled       bool
	MaximumSize   int
	ExpireTime    int
	CheckInterval int
	Directory     string
}

// AliasesConfig holds command alias configuration values.
type AliasesConfig struct {
	Add               []string
	Skip              []string
	SkipPlaylist      []string
	ForceSkip         []string
	ForceSkipPlaylist []string
	Help              []string
	Volume            []string
	Move              []string
	Reload            []string
	Reset             []string
	NumTracks         []string
	NextTrack         []string
	CurrentTrack      []string
	SetComment        []string
	NumCached         []string
	CacheSize         []string
	Kill              []string
	Shuffle           []string
	ToggleShuffle     []string
	ListTracks        []string
	Version           []string
}

// PermissionsConfig holds command permission configuration values.
type PermissionsConfig struct {
	Enabled           bool
	Add               bool
	AddPlaylist       bool
	Skip              bool
	SkipPlaylist      bool
	ForceSkip         bool
	ForceSkipPlaylist bool
	Help              bool
	Volume            bool
	Move              bool
	Reload            bool
	Reset             bool
	NumTracks         bool
	NextTrack         bool
	CurrentTrack      bool
	SetComment        bool
	NumCached         bool
	CacheSize         bool
	Kill              bool
	Shuffle           bool
	ToggleShuffle     bool
	ListTracks        bool
	Version           bool
}

// DescriptionsConfig holds command description configuration values.
type DescriptionsConfig struct {
	Add               string
	Skip              string
	SkipPlaylist      string
	ForceSkip         string
	ForceSkipPlaylist string
	Help              string
	Volume            string
	Move              string
	Reload            string
	Reset             string
	NumTracks         string
	NextTrack         string
	CurrentTrack      string
	SetComment        string
	NumCached         string
	CacheSize         string
	Kill              string
	Shuffle           string
	ToggleShuffle     string
	ListTracks        string
	Version           string
}

// Config gathers all logic related to configuration via commandline arguments
// and configuration files.
type Config struct {
	ConfigFileLocation string
	General            GeneralConfig
	Connection         ConnectionConfig
	Volume             VolumeConfig
	Cache              CacheConfig
	Aliases            AliasesConfig
	Permissions        PermissionsConfig
	Descriptions       DescriptionsConfig
}

// NewConfig returns a new config populated with default values.
func NewConfig() *Config {
	generalConfig := GeneralConfig{
		CommandPrefix:        "!",
		SkipRatio:            0.5,
		PlaylistSkipRatio:    0.5,
		DefaultComment:       "Hello! I am a bot. Type !help for a list of commands.",
		DefaultChannel:       []string{""},
		MaxTrackDuration:     0,
		MaxTracksPerPlaylist: 50,
		AutomaticShuffleOn:   false,
		PlayerCommand:        "ffmpeg",
	}
	connectionConfig := ConnectionConfig{
		Address:       "127.0.0.1",
		Port:          "64738",
		Password:      "",
		Username:      "MumbleDJ",
		Insecure:      false,
		Cert:          "",
		Key:           "",
		AccessTokens:  []string{""},
		RetryEnabled:  true,
		RetryAttempts: 10,
		RetryInterval: 5,
	}
	volumeConfig := VolumeConfig{
		Default: 0.4,
		Lowest:  0.01,
		Highest: 0.8,
	}
	cacheConfig := CacheConfig{
		Enabled:       false,
		MaximumSize:   512,
		ExpireTime:    24,
		CheckInterval: 5,
		Directory:     "~/.mumbledj/cache", // TODO: Set to $XDG_CACHE_HOME
	}
	aliasesConfig := AliasesConfig{
		Add:               []string{"add", "a"},
		Skip:              []string{"skip", "s"},
		SkipPlaylist:      []string{"skipplaylist", "sp"},
		ForceSkip:         []string{"forceskip", "fs"},
		ForceSkipPlaylist: []string{"forceskipplaylist", "fsp"},
		Help:              []string{"help", "h"},
		Volume:            []string{"volume", "v"},
		Move:              []string{"move", "m"},
		Reload:            []string{"reload", "r"},
		Reset:             []string{"reset", "re"},
		NumTracks:         []string{"numtracks", "numsongs", "nt"},
		NextTrack:         []string{"nexttrack", "nextsong", "next"},
		CurrentTrack:      []string{"currenttrack", "currentsong", "current"},
		SetComment:        []string{"setcomment", "comment", "sc"},
		NumCached:         []string{"numcached", "nc"},
		CacheSize:         []string{"cachesize", "cs"},
		Kill:              []string{"kill", "k"},
		Shuffle:           []string{"shuffle", "shuf", "sh"},
		ToggleShuffle:     []string{"toggleshuffle", "toggleshuf", "togshuf", "tsh"},
		ListTracks:        []string{"listtracks", "listsongs", "list", "l"},
		Version:           []string{"version", "v"},
	}
	permissionsConfig := PermissionsConfig{
		Enabled:           true,
		Add:               false,
		AddPlaylist:       false,
		Skip:              false,
		SkipPlaylist:      false,
		ForceSkip:         true,
		ForceSkipPlaylist: true,
		Help:              false,
		Volume:            false,
		Move:              true,
		Reload:            true,
		Reset:             true,
		NumTracks:         false,
		NextTrack:         false,
		CurrentTrack:      false,
		SetComment:        true,
		NumCached:         true,
		CacheSize:         true,
		Kill:              true,
		Shuffle:           true,
		ToggleShuffle:     true,
		ListTracks:        false,
		Version:           false,
	}
	descriptionsConfig := DescriptionsConfig{
		Add:               "Adds a track or playlist from a media site to the audio queue.",
		Skip:              "Places a vote to skip the current track.",
		SkipPlaylist:      "Places a vote to skip the current playlist.",
		ForceSkip:         "Immediately skips the current track.",
		ForceSkipPlaylist: "Immediately skips the current playlist.",
		Help:              "Outputs this list of commands.",
		Volume:            "Changes the volume if an argument is provided, outputs the current volume otherwise.",
		Move:              "Moves the bot into the Mumble channel provided via argument.",
		Reload:            "Reloads the configuration file.",
		Reset:             "Resets the audio queue by removing all queue items.",
		NumTracks:         "Outputs the number of tracks currently in the audio queue.",
		NextTrack:         "Outputs information about the next track in the queue if one exists.",
		CurrentTrack:      "Outputs information about the current track in the queue if one exists.",
		SetComment:        "Sets the comment displayed next to MumbleDJ's username in Mumble.",
		NumCached:         "Outputs the number of tracks cached on disk if caching is enabled.",
		CacheSize:         "Outputs the file size of the cache in MiB if caching is enabled.",
		Kill:              "Stops the bot and cleans its cache directory.",
		Shuffle:           "Randomizes the tracks currently in the audio queue.",
		ToggleShuffle:     "Toggles permanent track shuffling on/off.",
		ListTracks:        "Outputs a list of the tracks currently in the queue.",
		Version:           "Outputs the current version of MumbleDJ.",
	}

	return &Config{
		ConfigFileLocation: "~/.mumbledj/config.yaml", // TODO: Set to $XDG_CONFIG_HOME
		General:            generalConfig,
		Connection:         connectionConfig,
		Volume:             volumeConfig,
		Cache:              cacheConfig,
		Aliases:            aliasesConfig,
		Permissions:        permissionsConfig,
		Descriptions:       descriptionsConfig,
	}
}

// LoadFromConfigFile loads configuration values from the filepath specified via
// the filepath argument.
func (c *Config) LoadFromConfigFile(filepath string) error {
	return nil
}

// LoadFromCommandline loads configuration values from the commandline.
func (c *Config) LoadFromCommandline() {

}
