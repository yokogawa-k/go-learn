package foo_test

import (
	"fmt"
	"os"
	"testing"
)

func TestNormal(T *testing.T) {

	fmt.Println("Normal test")
	os.Exit(0)
}
