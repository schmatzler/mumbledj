/*
 * MumbleDJ
 * By Matthieu Grieger
 * bot/config_test.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package bot

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ConfigTestSuite struct {
	suite.Suite
}

func (suite *ConfigTestSuite) TestNewConfig() {
	config := NewConfig()

	assert.NotNil(suite.T(), config, "New config should not be nil")
	assert.NotNil(suite.T(), config.General, "New config's general section should not be nil")
	assert.NotNil(suite.T(), config.Connection, "New config's connection section should not be nil")
	assert.NotNil(suite.T(), config.Volume, "New config's volume section should not be nil")
	assert.NotNil(suite.T(), config.Cache, "New config's cache section should not be nil")
	assert.NotNil(suite.T(), config.Aliases, "New config's aliases section should not be nil")
	assert.NotNil(suite.T(), config.Permissions, "New config's permissions section should not be nil")
	assert.NotNil(suite.T(), config.Descriptions, "New config's descriptions section should not be nil")
}

func (suite *ConfigTestSuite) TestLoadFromConfigFile() {
	config := new(Config)
	err := config.LoadFromConfigFile("../config.yaml")

	assert.Nil(suite.T(), err, "No error should be returned")

	assert.NotNil(suite.T(), config.General, "New config's general section should not be nil")
	assert.NotNil(suite.T(), config.Connection, "New config's connection section should not be nil")
	assert.NotNil(suite.T(), config.Volume, "New config's volume section should not be nil")
	assert.NotNil(suite.T(), config.Cache, "New config's cache section should not be nil")
	assert.NotNil(suite.T(), config.Aliases, "New config's aliases section should not be nil")
	assert.NotNil(suite.T(), config.Permissions, "New config's permissions section should not be nil")
	assert.NotNil(suite.T(), config.Descriptions, "New config's descriptions section should not be nil")
}

func TestConfigTestSuite(t *testing.T) {
	suite.Run(t, new(ConfigTestSuite))
}
