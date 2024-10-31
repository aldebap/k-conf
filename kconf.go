////////////////////////////////////////////////////////////////////////////////
//	kconf.go  -  Aug-13-2024  -  aldebap
//
//	kconf cli parser
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

var (
	nameRegEx       *regexp.Regexp
	urlRegEx        *regexp.Regexp
	enabledRegEx    *regexp.Regexp
	protocolsRegEx  *regexp.Regexp
	methodsRegEx    *regexp.Regexp
	pathsRegEx      *regexp.Regexp
	serviceIdRegEx  *regexp.Regexp
	customIdRegEx   *regexp.Regexp
	userNameRegEx   *regexp.Regexp
	tagsRegEx       *regexp.Regexp
	routeIdRegEx    *regexp.Regexp
	idRegEx         *regexp.Regexp
	passwordRegEx   *regexp.Regexp
	algorithmRegEx  *regexp.Regexp
	keyRegEx        *regexp.Regexp
	secretRegEx     *regexp.Regexp
	ttlRegEx        *regexp.Regexp
	upstreamIdRegEx *regexp.Regexp
	targetRegEx     *regexp.Regexp
)

func compileRegExp() error {

	var err error

	//	compile all regex required to extract parameters for command add
	nameRegEx, err = regexp.Compile(`^--name\s*=\s*(\S+)\s*$`)
	if err != nil {
		return err
	}

	urlRegEx, err = regexp.Compile(`^--url\s*=\s*(\S.*)\s*$`)
	if err != nil {
		return err
	}

	enabledRegEx, err = regexp.Compile(`^--enabled\s*=\s*(\S.*)\s*$`)
	if err != nil {
		return err
	}

	protocolsRegEx, err = regexp.Compile(`^--protocols\s*=\s*(\S.*)\s*$`)
	if err != nil {
		return err
	}

	methodsRegEx, err = regexp.Compile(`^--methods\s*=\s*(\S.*)\s*$`)
	if err != nil {
		return err
	}

	pathsRegEx, err = regexp.Compile(`^--paths\s*=\s*(\S.*)\s*$`)
	if err != nil {
		return err
	}

	serviceIdRegEx, err = regexp.Compile(`^--service-id\s*=\s*(\S.*)\s*$`)
	if err != nil {
		return err
	}

	customIdRegEx, err = regexp.Compile(`^--custom-id\s*=\s*(\S.*)\s*$`)
	if err != nil {
		return err
	}

	userNameRegEx, err = regexp.Compile(`^--user-name\s*=\s*(\S.*)\s*$`)
	if err != nil {
		return err
	}

	tagsRegEx, err = regexp.Compile(`^--tags\s*=\s*(\S.*)\s*$`)
	if err != nil {
		return err
	}

	routeIdRegEx, err = regexp.Compile(`^--route-id\s*=\s*(\S.*)\s*$`)
	if err != nil {
		return err
	}

	idRegEx, err = regexp.Compile(`^--id\s*=\s*(\S.*)\s*$`)
	if err != nil {
		return err
	}

	passwordRegEx, err = regexp.Compile(`^--password\s*=\s*(\S.*)\s*$`)
	if err != nil {
		return err
	}

	algorithmRegEx, err = regexp.Compile(`^--algorithm\s*=\s*(\S.*)\s*$`)
	if err != nil {
		return err
	}

	keyRegEx, err = regexp.Compile(`^--key\s*=\s*(\S.*)\s*$`)
	if err != nil {
		return err
	}

	secretRegEx, err = regexp.Compile(`^--secret\s*=\s*(\S.*)\s*$`)
	if err != nil {
		return err
	}

	ttlRegEx, err = regexp.Compile(`^--ttl\s*=\s*(\S.*)\s*$`)
	if err != nil {
		return err
	}

	upstreamIdRegEx, err = regexp.Compile(`^--upstream-id\s*=\s*(\S.*)\s*$`)
	if err != nil {
		return err
	}

	targetRegEx, err = regexp.Compile(`^--target\s*=\s*(\S.*)\s*$`)
	if err != nil {
		return err
	}

	return nil
}

// kconf utility
func kconf(myKongServer KongServer, command []string, options Options) error {

	if len(command) == 0 {
		return errors.New("missing command: available commands: status, add, query, list")
	}

	err := compileRegExp()
	if err != nil {
		return errors.New("internal error compiling regexp: " + err.Error())
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
		return errors.New("missing entity for command add: available entities: service, route, consumer, plugin, upstream")
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

	case "consumer-basic-auth":
		var id string
		var userName string
		var password string

		for i := 1; i < len(command); i++ {
			match := idRegEx.FindAllStringSubmatch(command[i], -1)
			if len(match) == 1 {
				id = match[0][1]
			}

			match = userNameRegEx.FindAllStringSubmatch(command[i], -1)
			if len(match) == 1 {
				userName = match[0][1]
			}

			match = passwordRegEx.FindAllStringSubmatch(command[i], -1)
			if len(match) == 1 {
				password = match[0][1]
			}
		}
		if len(id) == 0 {
			return errors.New("missing consumer id: option --id={id} required for this command")
		}

		newKongBasicAuthConfig := NewKongBasicAuthConfig(userName, password)

		return myKongServer.AddConsumerBasicAuth(id, newKongBasicAuthConfig, options)

	case "consumer-key-auth":
		var id string
		var key string
		var ttl int
		var err error

		for i := 1; i < len(command); i++ {
			match := idRegEx.FindAllStringSubmatch(command[i], -1)
			if len(match) == 1 {
				id = match[0][1]
			}

			match = keyRegEx.FindAllStringSubmatch(command[i], -1)
			if len(match) == 1 {
				key = match[0][1]
			}

			match = ttlRegEx.FindAllStringSubmatch(command[i], -1)
			if len(match) == 1 {
				ttl, err = strconv.Atoi(match[0][1])
				if err != nil {
					return errors.New("Value for option --ttl must be an integer: " + err.Error())
				}
			}
		}
		if len(id) == 0 {
			return errors.New("missing consumer id: option --id={id} required for this command")
		}

		newKongKeyAuthConfig := NewKongKeyAuthConfig(key, int64(ttl))

		return myKongServer.AddConsumerKeyAuth(id, newKongKeyAuthConfig, options)

	case "consumer-jwt":
		var id string
		var algorithm string
		var key string
		var secret string

		for i := 1; i < len(command); i++ {
			match := idRegEx.FindAllStringSubmatch(command[i], -1)
			if len(match) == 1 {
				id = match[0][1]
			}

			match = algorithmRegEx.FindAllStringSubmatch(command[i], -1)
			if len(match) == 1 {
				algorithm = match[0][1]
			}

			match = keyRegEx.FindAllStringSubmatch(command[i], -1)
			if len(match) == 1 {
				key = match[0][1]
			}

			match = secretRegEx.FindAllStringSubmatch(command[i], -1)
			if len(match) == 1 {
				secret = match[0][1]
			}
		}
		if len(id) == 0 {
			return errors.New("missing consumer id: option --id={id} required for this command")
		}

		newKongJWTConfig := NewKongJWTConfig(algorithm, key, secret)

		return myKongServer.AddConsumerJWT(id, newKongJWTConfig, options)

	case "plugin":
		var name string
		var serviceId string
		var routeId string
		var enabled bool = true

		for i := 1; i < len(command); i++ {
			match := nameRegEx.FindAllStringSubmatch(command[i], -1)
			if len(match) == 1 {
				name = match[0][1]
			}

			match = serviceIdRegEx.FindAllStringSubmatch(command[i], -1)
			if len(match) == 1 {
				serviceId = match[0][1]
			}

			match = routeIdRegEx.FindAllStringSubmatch(command[i], -1)
			if len(match) == 1 {
				routeId = match[0][1]
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
		newKongPlugin := NewKongPlugin(name, serviceId, routeId, []KongPluginConfig{}, enabled)

		return myKongServer.AddPlugin(newKongPlugin, options)

	case "upstream":
		const valuesDelim = ","
		var name string
		var algorithm string
		var tags []string

		for i := 1; i < len(command); i++ {
			match := nameRegEx.FindAllStringSubmatch(command[i], -1)
			if len(match) == 1 {
				name = match[0][1]
			}

			match = algorithmRegEx.FindAllStringSubmatch(command[i], -1)
			if len(match) == 1 {
				algorithm = match[0][1]
			}

			match = tagsRegEx.FindAllStringSubmatch(command[i], -1)
			if len(match) == 1 {
				tags = strings.Split(match[0][1], valuesDelim)
			}
		}
		newKongUpstream := NewKongUpstream(name, algorithm, tags)

		return myKongServer.AddUpstream(newKongUpstream, options)

	case "upstream-target":
		var id string
		var target string

		for i := 1; i < len(command); i++ {
			match := idRegEx.FindAllStringSubmatch(command[i], -1)
			if len(match) == 1 {
				id = match[0][1]
			}

			match = targetRegEx.FindAllStringSubmatch(command[i], -1)
			if len(match) == 1 {
				target = match[0][1]
			}
		}

		if len(id) == 0 {
			return errors.New("missing upstream id: option --id={id} required for this command")
		}
		newKongUpstreamTarget := NewKongUpstreamTarget(target)

		return myKongServer.AddUpstreamTarget(id, newKongUpstreamTarget, options)
	}

	return errors.New("invalid entity for command add: " + command[0])
}

// command query
func commandQuery(myKongServer KongServer, command []string, options Options) error {

	if len(command) == 0 {
		return errors.New("missing entity for command query: available entities: service, route")
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

	case "plugin":
		var id string

		for i := 1; i < len(command); i++ {
			match := idRegEx.FindAllStringSubmatch(command[i], -1)
			if len(match) == 1 {
				id = match[0][1]
			}
		}

		if len(id) == 0 {
			return errors.New("missing plugin id: option --id={id} required for this command")
		}

		return myKongServer.QueryPlugin(id, options)

	case "upstream":
		var id string

		for i := 1; i < len(command); i++ {
			match := idRegEx.FindAllStringSubmatch(command[i], -1)
			if len(match) == 1 {
				id = match[0][1]
			}
		}

		if len(id) == 0 {
			return errors.New("missing upstream id: option --id={id} required for this command")
		}

		return myKongServer.QueryUpstream(id, options)

	case "upstream-target":
		var upstreamId string
		var id string

		for i := 1; i < len(command); i++ {

			match := upstreamIdRegEx.FindAllStringSubmatch(command[i], -1)
			if len(match) == 1 {
				upstreamId = match[0][1]
			}

			match = idRegEx.FindAllStringSubmatch(command[i], -1)
			if len(match) == 1 {
				id = match[0][1]
			}
		}

		if len(upstreamId) == 0 {
			return errors.New("missing upstream id: option --upstream-id={id} required for this command")
		}

		if len(id) == 0 {
			return errors.New("missing upstream target id: option --id={id} required for this command")
		}

		return myKongServer.QueryUpstreamTarget(upstreamId, id, options)
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

	case "plugin":
		return myKongServer.ListPlugins(options)

	case "upstream":
		return myKongServer.ListUpstreams(options)
	}

	return errors.New("invalid entity for command list: " + command[0])
}

// command update
func commandUpdate(myKongServer KongServer, command []string, options Options) error {

	if len(command) == 0 {
		return errors.New("missing entity for command update: available entities: service, route")
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

	case "plugin":
		if len(id) == 0 {
			return errors.New("missing plugin id: option --id={id} required for this command")
		}

		var serviceId string
		var routeId string
		var enabled bool = true

		for i := 1; i < len(command); i++ {

			match := serviceIdRegEx.FindAllStringSubmatch(command[i], -1)
			if len(match) == 1 {
				serviceId = match[0][1]
			}

			match = routeIdRegEx.FindAllStringSubmatch(command[i], -1)
			if len(match) == 1 {
				routeId = match[0][1]
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
		updatedKongPlugin := NewKongPlugin("", serviceId, routeId, nil, enabled)

		return myKongServer.UpdatePlugin(id, updatedKongPlugin, options)

	case "upstream":
		if len(id) == 0 {
			return errors.New("missing upstream id: option --id={id} required for this command")
		}

		const valuesDelim = ","
		var name string
		var algorithm string
		var tags []string

		for i := 1; i < len(command); i++ {
			match := nameRegEx.FindAllStringSubmatch(command[i], -1)
			if len(match) == 1 {
				name = match[0][1]
			}

			match = algorithmRegEx.FindAllStringSubmatch(command[i], -1)
			if len(match) == 1 {
				algorithm = match[0][1]
			}

			match = tagsRegEx.FindAllStringSubmatch(command[i], -1)
			if len(match) == 1 {
				tags = strings.Split(match[0][1], valuesDelim)
			}
		}
		updatedKongUpstream := NewKongUpstream(name, algorithm, tags)

		return myKongServer.UpdateUpstream(id, updatedKongUpstream, options)
	}

	return errors.New("invalid entity for command update: " + command[0])
}

// command delete
func commandDelete(myKongServer KongServer, command []string, options Options) error {

	if len(command) == 0 {
		return errors.New("missing entity for command delete: available entities: service, route")
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

	case "consumer":
		if len(id) == 0 {
			return errors.New("missing consumer id: option --id={id} required for this command")
		}

		return myKongServer.DeleteConsumer(id, options)

	case "plugin":
		if len(id) == 0 {
			return errors.New("missing plugin id: option --id={id} required for this command")
		}

		return myKongServer.DeletePlugin(id, options)

	case "upstream":
		if len(id) == 0 {
			return errors.New("missing upstream id: option --id={id} required for this command")
		}

		return myKongServer.DeleteUpstream(id, options)
	}

	return errors.New("invalid entity for command delete: " + command[0])
}
