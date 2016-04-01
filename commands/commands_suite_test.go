/*
 * MumbleDJ
 * By Matthieu Grieger
 * commmands_suite_test.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package commands_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestCommands(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Commands Suite")
}
