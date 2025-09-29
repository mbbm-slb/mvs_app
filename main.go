package main

import (
	"fmt"
	"github.com/mbbm-slb/mvs_lib_alpha"
)

var gitVersionInfo string

func printVersion() {
	fmt.Println(gitVersionInfo)
}

func main() {
    fmt.Printf("Main using Alpha: %s", mvslibalpha.Alpha())
}
