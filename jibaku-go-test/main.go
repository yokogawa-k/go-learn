package main

import (
	"fmt"
	"time"

	"github.com/garicchi/jibaku-go/jibaku"
)

func main() {
	c := 0

	for {
		jibaku.Check(time.Second, 10)
		fmt.Println(c)
		c++
	}

}
