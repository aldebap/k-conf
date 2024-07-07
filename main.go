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
	versionInfo string = "kconf 0.1"
)

// main entry point for kconf
func main() {
	var (
		version bool
	)

	//	CLI arguments
	flag.BoolVar(&version, "version", false, "show kconf version")

	flag.Parse()

	//	version option
	if version {
		fmt.Printf("%s\n", versionInfo)
		return
	}

	//	Kong server configuration
	var kongAddress string
	var kongPort int

	//	CLI arguments
	flag.StringVar(&kongAddress, "kong-address", "localhost", "Kong configuration address")
	flag.IntVar(&kongPort, "port", 8001, "Kong configuration port")

	flag.Parse()

	err := kconf(NewKongServer(kongAddress, kongPort), flag.Args())
	if err != nil {
		fmt.Fprintf(os.Stderr, "[error] fail attempting to connect to Kong server: %s", err.Error())
		os.Exit(-1)
	}
}

// kconf utility
func kconf(myKongServer *KongServer, command []string) error {

	//	command to get Kong status
	if len(command) == 1 && command[0] == "status" {
		return myKongServer.CheckStatus()
	}

	return nil
}
