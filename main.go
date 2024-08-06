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

		jsonOutput bool
	)

	//	CLI arguments
	flag.BoolVar(&version, "version", false, "show kconf version")

	flag.StringVar(&kongAddress, "kong-address", "localhost", "Kong configuration address")
	flag.IntVar(&kongPort, "port", 8001, "Kong configuration port")
	flag.BoolVar(&jsonOutput, "json-output", false, "use json output for every command")

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

	err := kconf(kongServer, flag.Args(), jsonOutput)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[error] fail attempting to send command to Kong server: %s\n", err.Error())
		os.Exit(-1)
	}
}

// kconf utility
func kconf(myKongServer *KongServer, command []string, jsonOutput bool) error {

	//	compile all regex required to extract parameters for commands
	nameRegEx, err := regexp.Compile(`^--name\s*=\s*(\S+)\s*$`)
	if err != nil {
		return err
	}

	urlRegEx, err := regexp.Compile(`^--url\s*=\s*(\S.*)\s*$`)
	if err != nil {
		return err
	}

	idRegEx, err := regexp.Compile(`^--id\s*=\s*(\S.*)\s*$`)
	if err != nil {
		return err
	}

	//	command to get Kong status
	if len(command) == 1 && command[0] == "status" {
		return myKongServer.CheckStatus(jsonOutput)
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

			return myKongServer.AddService(newService, jsonOutput)
		}

		return errors.New("missing entity for command add")
	}

	//	command query
	if len(command) >= 1 && command[0] == "query" {
		if len(command) >= 2 && command[1] == "service" {
			var id string

			for i := 2; i < len(command); i++ {
				match := idRegEx.FindAllStringSubmatch(command[i], -1)
				if len(match) == 1 {
					id = match[0][1]
				}
			}

			if len(id) == 0 {
				return errors.New("missing service id: option --id={id} required for this command")
			}

			return myKongServer.QueryService(id, jsonOutput)
		}

		return errors.New("missing entity for command add")
	}

	//	command list
	if len(command) >= 1 && command[0] == "list" {
		if len(command) >= 2 && command[1] == "service" {
			return myKongServer.ListServices(jsonOutput)
		}

		return errors.New("missing entity for command add")
	}

	return nil
}
