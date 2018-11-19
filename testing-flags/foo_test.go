// +build integration

package foo_test

import (
	"flag"
	"fmt"
	"os"
	"testing"
)

var (
	docker = flag.Bool("docker", false, "run in docker")
)

func TestMain(m *testing.M) {
	flag.Parse()
	if *docker {
		fmt.Println("integration test: docker")
	}
	if ! *docker {
		fmt.Println("integration test: machine")
	}

	os.Exit(0)
}
