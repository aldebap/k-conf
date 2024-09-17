////////////////////////////////////////////////////////////////////////////////
//	kconf.go  -  Aug-13-2024  -  aldebap
//
//	kconf cli parser
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"errors"
	"regexp"
	"strings"
)

// kconf utility
func kconf(myKongServer KongServer, command []string, options Options) error {

	if len(command) == 0 {
		return errors.New("missing command: available commands: status, add, query, list")
	}

	//	command to get Kong status
	switch command[0] {
	case "status":
		return myKongServer.CheckStatus(options)

	case "add":
		return commandAdd(myKongServer, command[1:], options)

	case "query":
		return commandQuery(myKongServer, command[1:], options)

	case "list":
		return commandList(myKongServer, command[1:], options)

	case "update":
		return commandUpdate(myKongServer, command[1:], options)

	case "delete":
		return commandDelete(myKongServer, command[1:], options)
	}

	return errors.New("invalid command: " + command[0])
}

// command add
func commandAdd(myKongServer KongServer, command []string, options Options) error {

	if len(command) == 0 {
		return errors.New("missing entity for command add: available entities: service, route")
	}

	//	compile all regex required to extract parameters for command add
	nameRegEx, err := regexp.Compile(`^--name\s*=\s*(\S+)\s*$`)
	if err != nil {
		return err
	}

	urlRegEx, err := regexp.Compile(`^--url\s*=\s*(\S.*)\s*$`)
	if err != nil {
		return err
	}

	enabledRegEx, err := regexp.Compile(`^--enabled\s*=\s*(\S.*)\s*$`)
	if err != nil {
		return err
	}

	protocolsRegEx, err := regexp.Compile(`^--protocols\s*=\s*(\S.*)\s*$`)
	if err != nil {
		return err
	}

	methodsRegEx, err := regexp.Compile(`^--methods\s*=\s*(\S.*)\s*$`)
	if err != nil {
		return err
	}

	pathsRegEx, err := regexp.Compile(`^--paths\s*=\s*(\S.*)\s*$`)
	if err != nil {
		return err
	}

	serviceIdRegEx, err := regexp.Compile(`^--service-id\s*=\s*(\S.*)\s*$`)
	if err != nil {
		return err
	}

	customIdRegEx, err := regexp.Compile(`^--custom-id\s*=\s*(\S.*)\s*$`)
	if err != nil {
		return err
	}

	userNameRegEx, err := regexp.Compile(`^--user-name\s*=\s*(\S.*)\s*$`)
	if err != nil {
		return err
	}

	tagsRegEx, err := regexp.Compile(`^--tags\s*=\s*(\S.*)\s*$`)
	if err != nil {
		return err
	}

	switch command[0] {
	case "service":
		var name string
		var url string
		var enabled bool = true

		for i := 1; i < len(command); i++ {
			match := nameRegEx.FindAllStringSubmatch(command[i], -1)
			if len(match) == 1 {
				name = match[0][1]
			}

			match = urlRegEx.FindAllStringSubmatch(command[i], -1)
			if len(match) == 1 {
				url = match[0][1]
			}

			match = enabledRegEx.FindAllStringSubmatch(command[i], -1)
			if len(match) == 1 {
				switch match[0][1] {
				case "false":
					enabled = false

				case "true":
					enabled = true

				default:
					return errors.New("wrong value for option --enabled: " + match[0][1])
				}
			}
		}
		newService := NewKongService(name, url, enabled)

		return myKongServer.AddService(newService, options)

	case "route":
		const valuesDelim = ","
		var name string
		var protocols []string
		var methods []string
		var paths []string
		var serviceId string

		for i := 1; i < len(command); i++ {
			match := nameRegEx.FindAllStringSubmatch(command[i], -1)
			if len(match) == 1 {
				name = match[0][1]
			}

			match = protocolsRegEx.FindAllStringSubmatch(command[i], -1)
			if len(match) == 1 {
				protocols = strings.Split(match[0][1], valuesDelim)
			}

			match = methodsRegEx.FindAllStringSubmatch(command[i], -1)
			if len(match) == 1 {
				methods = strings.Split(match[0][1], valuesDelim)
			}

			match = pathsRegEx.FindAllStringSubmatch(command[i], -1)
			if len(match) == 1 {
				paths = strings.Split(match[0][1], valuesDelim)
			}

			match = serviceIdRegEx.FindAllStringSubmatch(command[i], -1)
			if len(match) == 1 {
				serviceId = match[0][1]
			}
		}
		newKongRoute := NewKongRoute(name, protocols, methods, paths, serviceId)

		return myKongServer.AddRoute(newKongRoute, options)

	case "consumer":
		const valuesDelim = ","
		var customId string
		var userName string
		var tags []string

		for i := 1; i < len(command); i++ {
			match := customIdRegEx.FindAllStringSubmatch(command[i], -1)
			if len(match) == 1 {
				customId = match[0][1]
			}

			match = userNameRegEx.FindAllStringSubmatch(command[i], -1)
			if len(match) == 1 {
				userName = match[0][1]
			}

			match = tagsRegEx.FindAllStringSubmatch(command[i], -1)
			if len(match) == 1 {
				tags = strings.Split(match[0][1], valuesDelim)
			}
		}
		newKongConsumer := NewKongConsumer(customId, userName, tags)

		return myKongServer.AddConsumer(newKongConsumer, options)
	}

	return errors.New("invalid entity for command add: " + command[0])
}

// command query
func commandQuery(myKongServer KongServer, command []string, options Options) error {

	if len(command) == 0 {
		return errors.New("missing entity for command query: available entities: service, route")
	}

	//	compile all regex required to extract parameters for command query
	idRegEx, err := regexp.Compile(`^--id\s*=\s*(\S.*)\s*$`)
	if err != nil {
		return err
	}

	switch command[0] {
	case "service":
		var id string

		for i := 1; i < len(command); i++ {
			match := idRegEx.FindAllStringSubmatch(command[i], -1)
			if len(match) == 1 {
				id = match[0][1]
			}
		}

		if len(id) == 0 {
			return errors.New("missing service id: option --id={id} required for this command")
		}

		return myKongServer.QueryService(id, options)

	case "route":
		var id string

		for i := 1; i < len(command); i++ {
			match := idRegEx.FindAllStringSubmatch(command[i], -1)
			if len(match) == 1 {
				id = match[0][1]
			}
		}

		if len(id) == 0 {
			return errors.New("missing route id: option --id={id} required for this command")
		}

		return myKongServer.QueryRoute(id, options)

	case "consumer":
		var id string

		for i := 1; i < len(command); i++ {
			match := idRegEx.FindAllStringSubmatch(command[i], -1)
			if len(match) == 1 {
				id = match[0][1]
			}
		}

		if len(id) == 0 {
			return errors.New("missing consumer id: option --id={id} required for this command")
		}

		return myKongServer.QueryConsumer(id, options)
	}

	return errors.New("invalid entity for command query: " + command[0])
}

// command list
func commandList(myKongServer KongServer, command []string, options Options) error {

	if len(command) == 0 {
		return errors.New("missing entity for command list: available entities: service, route")
	}

	switch command[0] {
	case "service":
		return myKongServer.ListServices(options)

	case "route":
		return myKongServer.ListRoutes(options)

	case "consumer":
		return myKongServer.ListConsumers(options)
	}

	return errors.New("invalid entity for command list: " + command[0])
}

// command update
func commandUpdate(myKongServer KongServer, command []string, options Options) error {

	if len(command) == 0 {
		return errors.New("missing entity for command update: available entities: service, route")
	}

	//	compile all regex required to extract parameters for command update
	idRegEx, err := regexp.Compile(`^--id\s*=\s*(\S.*)\s*$`)
	if err != nil {
		return err
	}

	nameRegEx, err := regexp.Compile(`^--name\s*=\s*(\S+)\s*$`)
	if err != nil {
		return err
	}

	urlRegEx, err := regexp.Compile(`^--url\s*=\s*(\S.*)\s*$`)
	if err != nil {
		return err
	}

	enabledRegEx, err := regexp.Compile(`^--enabled\s*=\s*(\S.*)\s*$`)
	if err != nil {
		return err
	}

	protocolsRegEx, err := regexp.Compile(`^--protocols\s*=\s*(\S.*)\s*$`)
	if err != nil {
		return err
	}

	methodsRegEx, err := regexp.Compile(`^--methods\s*=\s*(\S.*)\s*$`)
	if err != nil {
		return err
	}

	pathsRegEx, err := regexp.Compile(`^--paths\s*=\s*(\S.*)\s*$`)
	if err != nil {
		return err
	}

	serviceIdRegEx, err := regexp.Compile(`^--service-id\s*=\s*(\S.*)\s*$`)
	if err != nil {
		return err
	}

	customIdRegEx, err := regexp.Compile(`^--custom-id\s*=\s*(\S.*)\s*$`)
	if err != nil {
		return err
	}

	userNameRegEx, err := regexp.Compile(`^--user-name\s*=\s*(\S.*)\s*$`)
	if err != nil {
		return err
	}

	tagsRegEx, err := regexp.Compile(`^--tags\s*=\s*(\S.*)\s*$`)
	if err != nil {
		return err
	}

	var id string

	for i := 1; i < len(command); i++ {
		match := idRegEx.FindAllStringSubmatch(command[i], -1)
		if len(match) == 1 {
			id = match[0][1]
		}
	}

	switch command[0] {
	case "service":
		if len(id) == 0 {
			return errors.New("missing service id: option --id={id} required for this command")
		}

		var name string
		var url string
		var enabled bool = true

		for i := 1; i < len(command); i++ {
			match := nameRegEx.FindAllStringSubmatch(command[i], -1)
			if len(match) == 1 {
				name = match[0][1]
			}

			match = urlRegEx.FindAllStringSubmatch(command[i], -1)
			if len(match) == 1 {
				url = match[0][1]
			}

			match = enabledRegEx.FindAllStringSubmatch(command[i], -1)
			if len(match) == 1 {
				switch match[0][1] {
				case "false":
					enabled = false

				case "true":
					enabled = true

				default:
					return errors.New("wrong value for option --enabled: " + match[0][1])
				}
			}
		}
		updatedService := NewKongService(name, url, enabled)

		return myKongServer.UpdateService(id, updatedService, options)

	case "route":
		if len(id) == 0 {
			return errors.New("missing route id: option --id={id} required for this command")
		}

		const valuesDelim = ","
		var name string
		var protocols []string
		var methods []string
		var paths []string
		var serviceId string

		for i := 1; i < len(command); i++ {
			match := nameRegEx.FindAllStringSubmatch(command[i], -1)
			if len(match) == 1 {
				name = match[0][1]
			}

			match = protocolsRegEx.FindAllStringSubmatch(command[i], -1)
			if len(match) == 1 {
				protocols = strings.Split(match[0][1], valuesDelim)
			}

			match = methodsRegEx.FindAllStringSubmatch(command[i], -1)
			if len(match) == 1 {
				methods = strings.Split(match[0][1], valuesDelim)
			}

			match = pathsRegEx.FindAllStringSubmatch(command[i], -1)
			if len(match) == 1 {
				paths = strings.Split(match[0][1], valuesDelim)
			}

			match = serviceIdRegEx.FindAllStringSubmatch(command[i], -1)
			if len(match) == 1 {
				serviceId = match[0][1]
			}
		}
		updatedRoute := NewKongRoute(name, protocols, methods, paths, serviceId)

		return myKongServer.UpdateRoute(id, updatedRoute, options)

	case "consumer":
		if len(id) == 0 {
			return errors.New("missing consumer id: option --id={id} required for this command")
		}

		const valuesDelim = ","
		var customId string
		var userName string
		var tags []string

		for i := 1; i < len(command); i++ {
			match := customIdRegEx.FindAllStringSubmatch(command[i], -1)
			if len(match) == 1 {
				customId = match[0][1]
			}

			match = userNameRegEx.FindAllStringSubmatch(command[i], -1)
			if len(match) == 1 {
				userName = match[0][1]
			}

			match = tagsRegEx.FindAllStringSubmatch(command[i], -1)
			if len(match) == 1 {
				tags = strings.Split(match[0][1], valuesDelim)
			}
		}
		updatedKongConsumer := NewKongConsumer(customId, userName, tags)

		return myKongServer.UpdateConsumer(id, updatedKongConsumer, options)
	}

	return errors.New("invalid entity for command update: " + command[0])
}

// command delete
func commandDelete(myKongServer KongServer, command []string, options Options) error {

	if len(command) == 0 {
		return errors.New("missing entity for command delete: available entities: service, route")
	}

	//	compile all regex required to extract parameters for command delete
	idRegEx, err := regexp.Compile(`^--id\s*=\s*(\S.*)\s*$`)
	if err != nil {
		return err
	}

	var id string

	for i := 1; i < len(command); i++ {
		match := idRegEx.FindAllStringSubmatch(command[i], -1)
		if len(match) == 1 {
			id = match[0][1]
		}
	}

	switch command[0] {
	case "service":
		if len(id) == 0 {
			return errors.New("missing service id: option --id={id} required for this command")
		}

		return myKongServer.DeleteService(id, options)

	case "route":
		if len(id) == 0 {
			return errors.New("missing route id: option --id={id} required for this command")
		}

		return myKongServer.DeleteRoute(id, options)
	}

	return errors.New("invalid entity for command delete: " + command[0])
}
