/*
 * MumbleDJ
 * By Matthieu Grieger
 * config_test.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package main_test

import (
	"reflect"

	. "github.com/matthieugrieger/mumbledj"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Config", func() {
	var (
		defaultConfig    *Config
		configFileConfig *Config
	)

	BeforeEach(func() {
		defaultConfig = NewConfig()
		configFileConfig.LoadFromConfigFile("config.yaml")
	})

	Describe("Initializing default config via NewConfig()", func() {
		Context("General config", func() {
			It("should be non-nil", func() {
				Expect(defaultConfig.General).NotTo(Equal(nil))
			})
		})
		Context("Connection config", func() {
			It("should be non-nil", func() {
				Expect(defaultConfig.Connection).NotTo(Equal(nil))
			})
		})
		Context("Volume config", func() {
			It("should be non-nil", func() {
				Expect(defaultConfig.Volume).NotTo(Equal(nil))
			})
		})
		Context("Cache config", func() {
			It("should be non-nil", func() {
				Expect(defaultConfig.Cache).NotTo(Equal(nil))
			})
		})
		Context("Aliases config", func() {
			It("should be non-nil", func() {
				Expect(defaultConfig.Aliases).NotTo(Equal(nil))
			})
		})
		Context("Permissions config", func() {
			It("should be non-nil", func() {
				Expect(defaultConfig.Permissions).NotTo(Equal(nil))
			})
		})
		Context("Descriptions config", func() {
			It("should be non-nil", func() {
				Expect(defaultConfig.Descriptions).NotTo(Equal(nil))
			})
		})
	})

	Describe("Poplating config via config file", func() {
		Context("that exists", func() {
			It("should be non-nil", func() {
				Expect(configFileConfig.General).NotTo(Equal(nil))
				Expect(configFileConfig.Connection).NotTo(Equal(nil))
				Expect(configFileConfig.Volume).NotTo(Equal(nil))
				Expect(configFileConfig.Cache).NotTo(Equal(nil))
				Expect(configFileConfig.Aliases).NotTo(Equal(nil))
				Expect(configFileConfig.Permissions).NotTo(Equal(nil))
				Expect(configFileConfig.Descriptions).NotTo(Equal(nil))
			})
			It("should be equal to the default config from NewConfig()", func() {
				Expect(reflect.DeepEqual(defaultConfig, configFileConfig)).To(BeTrue())
			})
		})
	})
})
