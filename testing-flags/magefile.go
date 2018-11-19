// +build mage

package main

import (
	"log"
	"os"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

// Default target to run when none is specified
// If not set, running mage will list available targets
// var Default = Test

func init() {
	if err := os.Setenv(mg.VerboseEnv, "1"); err != nil {
		log.Fatal(err)
	}
}

// Test is run all test
func Test() {
	mg.Deps(TestIntegrationDocker, TestIntegrationWithoutDocker, TestNormal)
}

// TestIntegrationDocker is integration-test with docker
func TestIntegrationDocker() error {
	return sh.Run("go", "test", "-tags=integration", "-docker")
}

// TestIntegrationWithoutDocker is integration-test without docker
func TestIntegrationWithoutDocker() error {
	return sh.Run("go", "test", "-tags=integration")
}

// TestNormal is normal golang test
func TestNormal() error {
	return sh.Run("go", "test")
}

