package main

import (
	"fmt"
	"os"
)

func main(){
    fmt.Printf("uid=%d\n", os.Getuid())
    fmt.Printf("euid=%d\n", os.Geteuid())
}
