/*
 * MumbleDJ
 * By Matthieu Grieger
 * mumbledj_suite_test.go
 * Copyright(c) 2016 Matthieu Grieger (MIT License)
 */

package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestMumbledj(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Mumbledj Suite")
}
