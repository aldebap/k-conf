# kconf

[![Go Build](https://github.com/aldebap/kconf/actions/workflows/go.yml/badge.svg)](https://github.com/aldebap/kconf/actions/workflows/go.yml)

kconf is just a simple utility to configure [Kong Gateway](https://konghq.com/products/kong-gateway) using **cli** (Command Line Interface).
To achieve this, kconf was planned to implement calls to all [Kong APIs](https://docs.konghq.com/gateway/api/admin-oss/latest/).
The initial idea for kconf was inpired by [GCP gcloud](https://cloud.google.com/sdk/gcloud/).

## Compatibility

kconf is compatible with Kong Gateway >= 3.8.x.35.

## Installation

**macOS**

If you are on macOS, install decK using brew:

```sh
brew tap aldebap/kconf
brew install kconf
```

**Linux**

If you are Linux, you can either use the Debian or RPM archive from the [GitHub release page](https://github.com/aldebap/kconf/releases/tag/v0.2.0) or install by downloading the binary:

```sh
curl -sL https://github.com/aldebap/kconf/releases/download/v0.2.0/kconf_Linux_x86_64.tar.gz -o kconf.tar.gz
tar -xf kconf.tar.gz -C /tmp
sudo cp /tmp/kconf /usr/local/bin/
```

**Windows**

If you are on Windows, you can download the binary from the [GitHub release page](https://github.com/aldebap/kconf/releases/tag/v0.2.0) or via PowerShell:

```sh
curl -sL https://github.com/aldebap/kconf/releases/download/v0.2.0/kconf_Windows_x86_64.zip -o kconf.zip
```

## Building kconf

kconf is 100% written in Golang and this repo provides some scripts (ksh) to build it, run unit tests and run functional tests:

- <font color="green">**cmd/build.sh**</font> build kconf from source
- <font color="green">**cmd/unit-test.sh**</font> run all kconf unit tests
- <font color="green">**cmd/functional-test.sh**</font> run all kconf functional test scenarios

## Using kconf

This is the general way to invoke kconf:

```sh
kconf [generic-options] command [entity] [entity-command-options]
```

General options are the following:

- <font color="orange">`-help`</font> - show kconf help
- <font color="orange">`-version`</font> - show kconf version
- <font color="orange">`-kong-address`</font> - set Kong configuration address (default "localhost")
- <font color="orange">`-port`</font> - set Kong configuration port (default 8001)
- <font color="orange">`-json-output`</font> - use json output for every command
- <font color="orange">`-verbose`</font> - run in verbose mode

The available commands are: status, add, query, list, update, delete and status.
The Kong entities are: service, route, consumer, plugin and upstream.

### Command <font color="green">status</font>

This command just check the status of Kong.

```sh
$ kconf -verbose status
200 OK
```

### Command <font color="green">add</font>

- <font color="green">**service**</font> - add a new service.
This command have the following options:
  - <font color="orange">`--name={service name}`</font> specify service name
  - <font color="orange">`--url={service URL}`</font> specify service URL
  - <font color="orange">`--enabled=[true|false]`</font> specify enable status of the service

If the service is successfully added to **Kong**, `kconf` will return the ID for the new service.

```sh
$ kconf add service --name=Consulta-Bin --url=https://api.pagar.me/bin/v1/499577
3302f59b-4bb0-410c-988b-d7e4e02a8c6e
```

- <font color="green">**route**</font> - add a new route.
This command have the following options:
  - <font color="orange">`--name={route name}`</font> specify route name
  - <font color="orange">`--prococols=[http,https]`</font> specify a comma separated list of protocols available for the route
  - <font color="orange">`--methods=[post,get, put, patch, delete]`</font> specify a comma separated list of HTTP methods available for the route
  - <font color="orange">`--paths={paths}`</font> specify the path for exposed route
  - <font color="orange">`--service-id={paths}`</font> specify the ID of the service that will be invoked from the route

If the route is successfully added to **Kong**, `kconf` will return the ID for the new route.

```sh
$ kconf add route --name=Consulta-Bin --protocols=http --methods=GET --paths=/api/v1/bin/499577 --service-id=3302f59b-4bb0-410c-988b-d7e4e02a8c6e
0ee7a361-0ac0-4468-b7b9-fc041d9c8ed7
```

- <font color="green">**consumer**</font> - add a new consumer.
This command have the following options:
  - <font color="orange">`--custom-id={route name}`</font> specify custom id for the consumer
  - <font color="orange">`--user-name={route name}`</font> specify user name for the consumer
  - <font color="orange">`--tags={tags}`</font> specify a comma separated list of tags associated to the consumer

If the consumer is successfully added to **Kong**, `kconf` will return the ID for the new consumer.

```sh
$ kconf add consumer --custom-id=auth-consumer --user-name=guest --tags=bronze-tier
e5c22534-371d-42f8-af44-0a87e11e5752
```

- <font color="green">**plugin**</font> - add a new plugin.
This command have the following options:
  - <font color="orange">`--name={plugin name}`</font> specify plugin name
  - <font color="orange">`--service-id={paths}`</font> specify the ID of the service that plugin will be applied
  - <font color="orange">`--route-id={paths}`</font> specify the ID of the route that plugin will be applied
  - <font color="orange">`--enabled=[true|false]`</font> specify enable status of the service

If the plugin is successfully added to **Kong**, `kconf` will return the ID for the new plugin.

```sh
$ kconf add plugin --name=basic-auth --route-id=0ee7a361-0ac0-4468-b7b9-fc041d9c8ed7 --enabled=true
590ac321-5061-4f9b-a88a-380209407cff
```

- <font color="green">**upstream**</font> - add a new upstream.
This command have the following options:
  - <font color="orange">`--name={upstream name}`</font> specify upstream name
  - <font color="orange">`--algorithm={algorithm}`</font> specify algorithm for the upstream
  - <font color="orange">`--tags={tags}`</font> specify a comma separated list of tags associated to the upstream

If the upstream is successfully added to **Kong**, `kconf` will return the ID for the new upstream.

```sh
$ kconf add upstream --name=Pedidos --algorithm=round-robin
a4775f39-0ddf-4d43-a9ee-31451419b812
```

- <font color="green">**upstream-target**</font> - add a new upstream target.
This command have the following options:
  - <font color="orange">`--id={upstream id}`</font> specify upstream id for the add
  - <font color="orange">`--target={target address}`</font> specify upstream target address

If the upstream is successfully added to **Kong**, `kconf` will return the ID for the new upstream target.

```sh
$ kconf add upstream-target --id=a4775f39-0ddf-4d43-a9ee-31451419b812 --target=192.168.68.107:8080
a0110455-2652-4e83-9202-9ca212277abc
```

### Command <font color="green">query</font>

- <font color="green">**service**</font> - query a service by id.
This command have the following options:
  - <font color="orange">`--id={service id}`</font> specify service id for the query

If the service id exists in **Kong**, `kconf` will return service name and URL.

```sh
$ kconf query service --id=3302f59b-4bb0-410c-988b-d7e4e02a8c6e
service: Consulta-Bin --> https://api.pagar.me:443/bin/v1/499577
```

- <font color="green">**route**</font> - query a route by id.
This command have the following options:
  - <font color="orange">`--id={route id}`</font> specify route id for the query

If the route id exists in **Kong**, `kconf` will return route name, HTTP methods, protocols, path and

```sh
$ kconf query route --id=0ee7a361-0ac0-4468-b7b9-fc041d9c8ed7
route: Consulta-Bin - [GET] [http]:[/api/v1/bin/499577] --> Service Id: 3302f59b-4bb0-410c-988b-d7e4e02a8c6e
```

- <font color="green">**consumer**</font> - query a consumer by id.
This command have the following options:
  - <font color="orange">`--id={consumer id}`</font> specify consumer id for the query

If the consumer id exists in **Kong**, `kconf` will return custom id, user name and tags.

```sh
$ kconf query consumer --id=e5c22534-371d-42f8-af44-0a87e11e5752
consumer:  --> external-customer ([bronze-tier])
```

- <font color="green">**plugin**</font> - query a plugin by id.
This command have the following options:
  - <font color="orange">`--id={plugin id}`</font> specify plugin id for the query

If the plugin id exists in **Kong**, `kconf` will return plugin name, protocols, the service id, the route id and the consumer id.

```sh
$ kconf query plugin --id=590ac321-5061-4f9b-a88a-380209407cff
590ac321-5061-4f9b-a88a-380209407cff: basic-auth - [grpc grpcs http https ws wss]: serviceId:  ; routeId: 0ee7a361-0ac0-4468-b7b9-fc041d9c8ed7 ; consumerId:
```

- <font color="green">**upstream**</font> - query an upstream by id.
This command have the following options:
  - <font color="orange">`--id={upstream id}`</font> specify upstream id for the query

If the upstream id exists in **Kong**, `kconf` will return upstream name, algorithm and tags.

```sh
$ kconf query upstream --id=a4775f39-0ddf-4d43-a9ee-31451419b812
upstream: Pedidos --> round-robin ([])
```

- <font color="green">**upstream-target**</font> - query an upstream target by id.
This command have the following options:
  - <font color="orange">`--upstream-id={upstream id}`</font> specify upstream id for the query
  - <font color="orange">`--id={upstream target id}`</font> specify upstream target id for the query

If the upstream id exists in **Kong**, `kconf` will return target.

```sh
$ kconf query upstream-target --upstream-id=a4775f39-0ddf-4d43-a9ee-31451419b812 --id=a0110455-2652-4e83-9202-9ca212277abc
192.168.68.107:8080
```

### Command <font color="green">list</font>

- <font color="green">**service**</font> - list all services.
This command doesn't have options.
If there are services in **Kong**, `kconf` will return a list of all services.

```sh
$ kconf list service
3302f59b-4bb0-410c-988b-d7e4e02a8c6e: Consulta-Bin --> https://api.pagar.me:443/bin/v1/499577
```

- <font color="green">**route**</font> - list all routes.
This command doesn't have options.
If there are routes in **Kong**, `kconf` will return a list of all routes.

```sh
$ kconf list route
0ee7a361-0ac0-4468-b7b9-fc041d9c8ed7: Consulta-Bin - [GET] [http]:[/api/v1/bin/499577] --> Service Id: 3302f59b-4bb0-410c-988b-d7e4e02a8c6e
```

- <font color="green">**consumer**</font> - list all consumers.
This command doesn't have options.
If there are consumers in **Kong**, `kconf` will return a list of all consumers.

```sh
$ kconf list consumer
e5c22534-371d-42f8-af44-0a87e11e5752: () external-customer [bronze-tier]
```

- <font color="green">**plugin**</font> - list all plugins.
This command doesn't have options.
If there are plugins in **Kong**, `kconf` will return a list of all plugins.

```sh
$ kconf list plugin
plugin: 590ac321-5061-4f9b-a88a-380209407cff: basic-auth - [grpc grpcs http https ws wss]: serviceId:  ; routeId: 0ee7a361-0ac0-4468-b7b9-fc041d9c8ed7 ; consumerId:
```

- <font color="green">**upstream**</font> - list all upstreams.
This command doesn't have options.
If there are upstreams in **Kong**, `kconf` will return a list of all upstreams.

```sh
$ kconf list upstream
a4775f39-0ddf-4d43-a9ee-31451419b812: Pedidos --> round-robin ([])
```

- <font color="green">**upstream-target**</font> - list all targets for an upstream by id.
This command have the following options:
  - <font color="orange">`--upstream-id={upstream id}`</font> specify upstream id for the query

If there are targets in **Kong** for specified upstream, `kconf` will return a list of all targets.

```sh
$ kconf list upstream-targets --upstream-id=a4775f39-0ddf-4d43-a9ee-31451419b812
192.168.68.107:8080
```

### Command <font color="green">update</font>

- <font color="green">**service**</font> - update a service by id.
This command have the following options:
  - <font color="orange">`--id={service id}`</font> specify service id to be updated
  - <font color="orange">`--name={service name}`</font> specify service name
  - <font color="orange">`--url={service URL}`</font> specify service URL
  - <font color="orange">`--enabled=[true|false]`</font> specify enable status of the service

If the service is successfully updated in **Kong**, `kconf` will return the ID for the service.

```sh
$ kconf update service --id=3302f59b-4bb0-410c-988b-d7e4e02a8c6e --enabled=false
service: Consulta-Bin --> https://api.pagar.me:443/bin/v1/499577
```

- <font color="green">**route**</font> - update a route by id.
This command have the following options:
  - <font color="orange">`--id={route id}`</font> specify route id to be updated
  - <font color="orange">`--name={route name}`</font> specify route name
  - <font color="orange">`--prococols=[http,https]`</font> specify a comma separated list of protocols available for the route
  - <font color="orange">`--methods=[post,get, put, patch, delete]`</font> specify a comma separated list of HTTP methods available for the route
  - <font color="orange">`--paths={paths}`</font> specify the path for exposed route
  - <font color="orange">`--service-id={paths}`</font> specify the ID of the service that will be invoked from the route

If the route is successfully updated in **Kong**, `kconf` will return the ID for the route.

```sh
$ kconf update route --id=0ee7a361-0ac0-4468-b7b9-fc041d9c8ed7 --protocols=http,https
route: Consulta-Bin - [GET] [http,https]:[/api/v1/bin/499577] --> Service Id: 3302f59b-4bb0-410c-988b-d7e4e02a8c6e
```

- <font color="green">**consumer**</font> - update a consumer by id.
This command have the following options:
  - <font color="orange">`--id={consumer id}`</font> specify consumer id to be updated
  - <font color="orange">`--custom-id={route name}`</font> specify custom id for the consumer
  - <font color="orange">`--user-name={route name}`</font> specify user name for the consumer
  - <font color="orange">`--tags={tags}`</font> specify a comma separated list of tags associated to the consumer

If the consumer is successfully updated in **Kong**, `kconf` will return custom id, user name and tags.

```sh
$ kconf update consumer --id=e5c22534-371d-42f8-af44-0a87e11e5752 --custom-id=auth-consulta-bin
consumer: auth-consulta-bin --> external-customer ([bronze-tier])
```

- <font color="green">**plugin**</font> - update a plugin by id.
This command have the following options:
  - <font color="orange">`--id={plugin id}`</font> specify plugin id to be updated
  - <font color="orange">`--service-id={paths}`</font> specify the ID of the service that plugin will be applied
  - <font color="orange">`--route-id={paths}`</font> specify the ID of the route that plugin will be applied
  - <font color="orange">`--enabled=[true|false]`</font> specify enable status of the service

If the plugin is successfully updated in **Kong**, `kconf` will return plugin name, protocols, the service id, the route id and the consumer id.

```sh
$ kconf update plugin --id=590ac321-5061-4f9b-a88a-380209407cff --enabled=false
590ac321-5061-4f9b-a88a-380209407cff: basic-auth - [grpc grpcs http https ws wss]: serviceId:  ; routeId: 0ee7a361-0ac0-4468-b7b9-fc041d9c8ed7 ; consumerId:
```

- <font color="green">**upstream**</font> - update an upstream by id.
This command have the following options:
  - <font color="orange">`--id={upstream id}`</font> specify upstream id to be updated
  - <font color="orange">`--name={upstream name}`</font> specify upstream name
  - <font color="orange">`--algorithm={algorithm}`</font> specify algorithm for the upstream
  - <font color="orange">`--tags={tags}`</font> specify a comma separated list of tags associated to the upstream

If the upstream is successfully updated in **Kong**, `kconf` will return upstream name, algorithm and tags.

```sh
$ kconf update upstream --id=a4775f39-0ddf-4d43-a9ee-31451419b812 --tags=silver-tier
upstream: Pedidos --> round-robin ([silver-tier])
```

### Command <font color="green">delete</font>

- <font color="green">**service**</font> - delete a service by id.
This command have the following options:
  - <font color="orange">`--id={service id}`</font> specify service id to be deleted

```sh
kconf delete service --id=3302f59b-4bb0-410c-988b-d7e4e02a8c6e
```

- <font color="green">**route**</font> - delete a route by id.
This command have the following options:
  - <font color="orange">`--id={route id}`</font> specify route id to be deleted

```sh
kconf delete route --id=0ee7a361-0ac0-4468-b7b9-fc041d9c8ed7
```

- <font color="green">**consumer**</font> - delete a consumer by id.
This command have the following options:
  - <font color="orange">`--id={consumer id}`</font> specify consumer id to be deleted

```sh
kconf delete consumer --id=e5c22534-371d-42f8-af44-0a87e11e5752
```

- <font color="green">**plugin**</font> - delete a plugin by id.
This command have the following options:
  - <font color="orange">`--id={plugin id}`</font> specify plugin id to be deleted

```sh
kconf delete plugin --id=590ac321-5061-4f9b-a88a-380209407cff
```

- <font color="green">**upstream**</font> - delete an upstream by id.
This command have the following options:
  - <font color="orange">`--id={upstream id}`</font> specify upstream id to be deleted

```sh
kconf delete upstream --id=a4775f39-0ddf-4d43-a9ee-31451419b812
```

- <font color="green">**upstream-target**</font> - delete an upstream target by id.
This command have the following options:
  - <font color="orange">`--upstream-id={upstream id}`</font> specify upstream id the target belongs to
  - <font color="orange">`--id={upstream target id}`</font> specify upstream target id to be delete

```sh
kconf delete upstream-target --upstream-id=a4775f39-0ddf-4d43-a9ee-31451419b812 --id=a0110455-2652-4e83-9202-9ca212277abc
```

### Consumer Plugins

- <font color="green">**add consumer-basic-auth**</font> - add basic-auth plugin for a cunsumer.
This command have the following options:
  - <font color="orange">`--id={consumer id}`</font> specify consumer id to add the plugin
  - <font color="orange">`--user-name={user name}`</font> specify user name for the plugin
  - <font color="orange">`--password={password}`</font> specify password for the plugin

If the plugin is successfully added to **Kong**, `kconf` will return the ID for the new plugin.

```sh
$ kconf add consumer-basic-auth --id=7cab7e0b-3d6a-4079-aeaa-d51ab8fd2cab --user-name=guest --password=kong1234
a0229bef-dafe-4060-87ed-8a02d746d425
```

- <font color="green">**add consumer-key-auth**</font> - add key-auth plugin for a cunsumer.
This command have the following options:
  - <font color="orange">`--id={consumer id}`</font> specify consumer id to add the plugin
  - <font color="orange">`--key={auth key}`</font> specify auth key for the plugin
  - <font color="orange">`--ttl={ttl}`</font> specify ttl for the plugin

If the plugin is successfully added to **Kong**, `kconf` will return the ID for the new plugin.

```sh
$ kconf add consumer-key-auth --id=7cab7e0b-3d6a-4079-aeaa-d51ab8fd2cab --key=d5a37fa6-b033-4107-a29f-ebf51b443968 --ttl=0
bd90e0bc-0ebd-428d-925e-081ff0503a4d
```

- <font color="green">**add consumer-jwt**</font> - add JWT plugin for a cunsumer.
This command have the following options:
  - <font color="orange">`--id={consumer id}`</font> specify consumer id to add the plugin
  - <font color="orange">`--algorithm={algorithm}`</font> specify algorithm for the plugin
  - <font color="orange">`--key={auth key}`</font> specify auth key for the plugin
  - <font color="orange">`--secret={secret}`</font> specify secret for the plugin

If the plugin is successfully added to **Kong**, `kconf` will return the ID for the new plugin.

```sh
$ kconf add consumer-jwt --id=7cab7e0b-3d6a-4079-aeaa-d51ab8fd2cab --algorithm=HS256 --key=5ab5ae42-6227-4f49-a354-6eda3e19ff99 --secret=ff6d73d4-5f53-405a-8a5d-b2f03f405b14
845ffee6-cd9e-4149-b2bc-13a251770306
```

## kconf backlog

### Features backlog (for v0.3 release)

- [X] ~~Endpoint to add a new upstream~~
- [X] ~~Endpoint to get a upstream~~
- [X] ~~Endpoint to get a list of upstreams~~
- [X] ~~Endpoint to update a upstream~~
- [X] ~~Endpoint to delete a upstream~~
- [X] ~~Endpoint to add a new target~~
- [X] ~~Endpoint to get a target~~
- [X] ~~Endpoint to get a list of targets~~
- [X] ~~Endpoint to update a target~~
- [X] ~~Endpoint to delete a target~~
- [ ] Endpoint to add a IP Restriction plugin for a consumer
- [ ] Endpoint to add a Rate Limit plugin for a consumer
- [ ] Endpoint to add a Request Size Limit plugin for a consumer
- [ ] Endpoint to add a Syslog plugin for a consumer

### Features backlog (for v0.4 release)

- [ ] Endpoint to add a LDAP plugin for a consumer
- [ ] Endpoint to add a OAuth2 plugin for a consumer
- [ ] Endpoint to add a HMAC Auth plugin for a consumer
- [ ] Endpoint to add a Kong Functions (Pre) plugin for a consumer
- [ ] Endpoint to add a Kong Functions (Post) plugin for a consumer
- [ ] Endpoint to add a OpenTelemetry plugin for a consumer
- [ ] Endpoint to add a Correlation ID plugin for a consumer
- [ ] Endpoint to add a Request Transformer plugin for a consumer
- [ ] Endpoint to add a Response Transformer plugin for a consumer
- [ ] Endpoint to add a gRPC Web plugin for a consumer
- [ ] Endpoint to add a File Log plugin for a consumer
- [ ] Endpoint to add a HTTP Log plugin for a consumer
- [ ] Add parameter to Add Plugin command to specify plugin config
