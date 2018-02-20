// Copyright (c) 2018 Zededa, Inc.
// All rights reserved.

package main

import (
	"flag"
	"fmt"
	"github.com/zededa/go-provision/hardware"
	"os"
)

// Set from Makefile
var Version = "No version specified"

func main() {
	versionPtr := flag.Bool("v", false, "Version")
	flag.Parse()
	if *versionPtr {
		fmt.Printf("%s: %s\n", os.Args[0], Version)
		return
	}
	fmt.Println(hardware.GetHardwareModel())
}