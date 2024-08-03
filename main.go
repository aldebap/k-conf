////////////////////////////////////////////////////////////////////////////////
//	main.go  -  Jul-5-2024  -  aldebap
//
//	Entry point for Kong Configuration tool
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"regexp"
)

const (
	versionInfo string = "kconf 0.1"
)

// main entry point for kconf
func main() {
	var (
		version bool

		//	Kong server configuration
		kongAddress string
		kongPort    int
	)

	//	CLI arguments
	flag.BoolVar(&version, "version", false, "show kconf version")

	flag.StringVar(&kongAddress, "kong-address", "localhost", "Kong configuration address")
	flag.IntVar(&kongPort, "port", 8001, "Kong configuration port")

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

	err := kconf(kongServer, flag.Args())
	if err != nil {
		fmt.Fprintf(os.Stderr, "[error] fail attempting to send command to Kong server: %s\n", err.Error())
		os.Exit(-1)
	}
}

// kconf utility
func kconf(myKongServer *KongServer, command []string) error {

	//	compile all regex required to extract parameters for commands
	nameRegEx, err := regexp.Compile(`^--name\s*=\s*(\S+)\s*$`)
	if err != nil {
		return err
	}

	urlRegEx, err := regexp.Compile(`^--url\s*=\s*(\S.*)\s*$`)
	if err != nil {
		return err
	}

	//	command to get Kong status
	if len(command) == 1 && command[0] == "status" {
		return myKongServer.CheckStatus()
	}

	//	command add
	if len(command) >= 1 && command[0] == "add" {
		if len(command) >= 2 && command[1] == "service" {
			var name string
			var url string
			var enabled bool = true

			for i := 2; i < len(command); i++ {
				match := nameRegEx.FindAllStringSubmatch(command[i], -1)
				if len(match) == 1 {
					name = match[0][1]
				}

				match = urlRegEx.FindAllStringSubmatch(command[i], -1)
				if len(match) == 1 {
					url = match[0][1]
				}
			}
			newService := NewKongService(name, url, enabled)

			return myKongServer.AddService(newService)
		}

		return errors.New("missing entity for command add")
	}

	return nil
}
