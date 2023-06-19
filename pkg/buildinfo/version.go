package buildinfo

import (
	"flag"
	"fmt"
	"os"
)

var (
	Version   string
	GitCommit string
	BuildTime string
)

func PrintVersionOrContinue() {
	versionFlag := flag.Bool("v", false, "Print the current version and exit")

	flag.Parse()

	switch {
	case *versionFlag:
		fmt.Printf("version: %s (%s) | %s", Version, GitCommit, BuildTime)
		os.Exit(0)
	}
}
