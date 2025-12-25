package main

import (
	"fmt"
	"runtime"
)

var (
	Version   = "dev"
	BuildTime = "unknown"
)

func printVersion() {
	fmt.Printf("ccpm-tui %s\n", Version)
	fmt.Printf("  Built:    %s\n", BuildTime)
	fmt.Printf("  Go:       %s\n", runtime.Version())
	fmt.Printf("  OS/Arch:  %s/%s\n", runtime.GOOS, runtime.GOARCH)
}
