////////////////////////////////////////////////////////////////////////////////
//	main.go  -  Jul-5-2024  -  aldebap
//
//	Entry point for Kong Configuration tool
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	versionInfo string = "kconf 0.2"
)

// execution options
type Options struct {
	jsonOutput bool
	verbose    bool
}

// main entry point for kconf
func main() {
	var (
		version bool

		//	Kong server configuration
		kongAddress string
		kongPort    int

		options Options
	)

	//	CLI arguments
	flag.BoolVar(&version, "version", false, "show kconf version")

	flag.StringVar(&kongAddress, "kong-address", "localhost", "Kong configuration address")
	flag.IntVar(&kongPort, "port", 8001, "Kong configuration port")
	flag.BoolVar(&options.jsonOutput, "json-output", false, "use json output for every command")
	flag.BoolVar(&options.verbose, "verbose", false, "run in verbose mode")

	flag.Parse()

	//	version option
	if version {
		fmt.Printf("%s\n", versionInfo)
		return
	}

	//	connect and send command
	kongServer := NewKongServer(kongAddress, kongPort)
	if kongServer == nil {
		fmt.Fprintf(os.Stderr, "[error] fail attempting to alocate Kong server\n")
		os.Exit(-1)
	}

	err := kconf(kongServer, flag.Args(), options)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[error] %s\n", err.Error())
		os.Exit(-1)
	}
}
