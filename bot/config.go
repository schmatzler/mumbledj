/*
 * MumbleDJ
 * By Matthieu Grieger
 * bot/config.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package bot

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// GeneralConfig holds general configuration values.
type GeneralConfig struct {
	CommandPrefix        string  `yaml:"command_prefix"`
	SkipRatio            float32 `yaml:"skip_ratio"`
	PlaylistSkipRatio    float32 `yaml:"playlist_skip_ratio"`
	DefaultComment       string  `yaml:"default_comment"`
	DefaultChannel       string  `yaml:"default_channel"`
	MaxTrackDuration     int     `yaml:"max_track_duration"`
	MaxTracksPerPlaylist int     `yaml:"max_tracks_per_playlist"`
	AutomaticShuffleOn   bool    `yaml:"automatic_shuffle_on"`
	AnnounceNewTrack     bool    `yaml:"announce_new_track"`
	PlayerCommand        string  `yaml:"player_command"`
}

// ConnectionConfig holds connection configuration values.
type ConnectionConfig struct {
	Address       string `yaml:"address"`
	Port          string `yaml:"port"`
	Password      string `yaml:"password"`
	Username      string `yaml:"username"`
	Insecure      bool   `yaml:"insecure"`
	Cert          string `yaml:"cert"`
	Key           string `yaml:"key"`
	AccessTokens  string `yaml:"access_tokens"`
	RetryEnabled  bool   `yaml:"retry_enabled"`
	RetryAttempts int    `yaml:"retry_attempts"`
	RetryInterval int    `yaml:"retry_interval"`
}

// VolumeConfig holds volume configuration values.
type VolumeConfig struct {
	Default float32 `yaml:"default"`
	Lowest  float32 `yaml:"lowest"`
	Highest float32 `yaml:"highest"`
}

// CacheConfig holds cache configuration values.
type CacheConfig struct {
	Enabled       bool   `yaml:"enabled"`
	MaximumSize   int    `yaml:"maximum_size"`
	ExpireTime    int    `yaml:"expire_time"`
	CheckInterval int    `yaml:"check_interval"`
	Directory     string `yaml:"directory"`
}

// AliasesConfig holds command alias configuration values.
type AliasesConfig struct {
	Add               []string `yaml:"add"`
	AddNext           []string `yaml:"add_next"`
	Skip              []string `yaml:"skip"`
	SkipPlaylist      []string `yaml:"skip_playlist"`
	ForceSkip         []string `yaml:"force_skip"`
	ForceSkipPlaylist []string `yaml:"force_skip_playlist"`
	Help              []string `yaml:"help"`
	Volume            []string `yaml:"volume"`
	Move              []string `yaml:"move"`
	Reload            []string `yaml:"reload"`
	Reset             []string `yaml:"reset"`
	NumTracks         []string `yaml:"num_tracks"`
	NextTrack         []string `yaml:"next_track"`
	CurrentTrack      []string `yaml:"current_track"`
	SetComment        []string `yaml:"set_comment"`
	NumCached         []string `yaml:"num_cached"`
	CacheSize         []string `yaml:"cache_size"`
	Kill              []string `yaml:"kill"`
	Shuffle           []string `yaml:"shuffle"`
	ToggleShuffle     []string `yaml:"toggle_shuffle"`
	ListTracks        []string `yaml:"list_tracks"`
	Version           []string `yaml:"version"`
}

// PermissionsConfig holds command permission configuration values.
type PermissionsConfig struct {
	Enabled           bool   `yaml:"enabled"`
	UserGroup         string `yaml:"user_group"`
	Add               bool   `yaml:"add"`
	AddNext           bool   `yaml:"add_next"`
	AddPlaylist       bool   `yaml:"add_playlist"`
	Skip              bool   `yaml:"skip"`
	SkipPlaylist      bool   `yaml:"skip_playlist"`
	ForceSkip         bool   `yaml:"force_skip"`
	ForceSkipPlaylist bool   `yaml:"force_skip_playlist"`
	Help              bool   `yaml:"help"`
	Volume            bool   `yaml:"volume"`
	Move              bool   `yaml:"move"`
	Reload            bool   `yaml:"reload"`
	Reset             bool   `yaml:"reset"`
	NumTracks         bool   `yaml:"num_tracks"`
	NextTrack         bool   `yaml:"next_track"`
	CurrentTrack      bool   `yaml:"current_track"`
	SetComment        bool   `yaml:"set_comment"`
	NumCached         bool   `yaml:"num_cached"`
	CacheSize         bool   `yaml:"cache_size"`
	Kill              bool   `yaml:"kill"`
	Shuffle           bool   `yaml:"shuffle"`
	ToggleShuffle     bool   `yaml:"toggle_shuffle"`
	ListTracks        bool   `yaml:"list_tracks"`
	Version           bool   `yaml:"version"`
}

// DescriptionsConfig holds command description configuration values.
type DescriptionsConfig struct {
	Add               string `yaml:"add"`
	AddNext           string `yaml:"add_next"`
	Skip              string `yaml:"skip"`
	SkipPlaylist      string `yaml:"skip_playlist"`
	ForceSkip         string `yaml:"force_skip"`
	ForceSkipPlaylist string `yaml:"force_skip_playlist"`
	Help              string `yaml:"help"`
	Volume            string `yaml:"volume"`
	Move              string `yaml:"move"`
	Reload            string `yaml:"reload"`
	Reset             string `yaml:"reset"`
	NumTracks         string `yaml:"num_tracks"`
	NextTrack         string `yaml:"next_track"`
	CurrentTrack      string `yaml:"current_track"`
	SetComment        string `yaml:"set_comment"`
	NumCached         string `yaml:"num_cached"`
	CacheSize         string `yaml:"cache_size"`
	Kill              string `yaml:"kill"`
	Shuffle           string `yaml:"shuffle"`
	ToggleShuffle     string `yaml:"toggle_shuffle"`
	ListTracks        string `yaml:"list_tracks"`
	Version           string `yaml:"version"`
}

// Config gathers all logic related to configuration via commandline arguments
// and configuration files.
type Config struct {
	ConfigFileLocation string
	General            GeneralConfig      `yaml:"general"`
	Connection         ConnectionConfig   `yaml:"connection"`
	Volume             VolumeConfig       `yaml:"volume"`
	Cache              CacheConfig        `yaml:"cache"`
	Aliases            AliasesConfig      `yaml:"aliases"`
	Permissions        PermissionsConfig  `yaml:"permissions"`
	Descriptions       DescriptionsConfig `yaml:"descriptions"`
}

// NewConfig returns a new config populated with default values.
func NewConfig() *Config {
	generalConfig := GeneralConfig{
		CommandPrefix:        "!",
		SkipRatio:            0.5,
		PlaylistSkipRatio:    0.5,
		DefaultComment:       "Hello! I am a bot. Type !help for a list of commands.",
		DefaultChannel:       "",
		MaxTrackDuration:     0,
		MaxTracksPerPlaylist: 50,
		AutomaticShuffleOn:   false,
		AnnounceNewTrack:     true,
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
		AccessTokens:  "",
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
		AddNext:           []string{"addnext", "an"},
		Skip:              []string{"skip", "s"},
		SkipPlaylist:      []string{"skipplaylist", "sp"},
		ForceSkip:         []string{"forceskip", "fs"},
		ForceSkipPlaylist: []string{"forceskipplaylist", "fsp"},
		Help:              []string{"help", "h"},
		Volume:            []string{"volume", "vol", "v"},
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
		UserGroup:         "admins",
		Add:               false,
		AddNext:           true,
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
		AddNext:           "Adds a track or playlist from a media site as the next item in the audio queue.",
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
	var (
		data []byte
		err  error
	)

	if data, err = ioutil.ReadFile(filepath); err != nil {
		return err
	}

	err = yaml.Unmarshal(data, &c)
	if err != nil {
		return err
	}

	return nil
}
