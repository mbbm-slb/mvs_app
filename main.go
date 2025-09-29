package main

import (
	"fmt"
	"github.com/mbbm-slb/mvs_lib_alpha"
	"github.com/mbbm-slb/mvs_lib_beta"
)

var gitVersionInfo string

func printVersion() {
	fmt.Println(gitVersionInfo)
}

func main() {
    fmt.Printf("Main using Alpha: %s\n", mvslibalpha.Alpha())
	fmt.Printf("Main using Beta: %s\n", mvslibbeta.Beta())
}
