# kconf

[![Go Build](https://github.com/aldebap/kconf/actions/workflows/go.yml/badge.svg)](https://github.com/aldebap/kconf/actions/workflows/go.yml)

kconf is just a simple utility to configure [Kong Gateway](https://konghq.com/products/kong-gateway) using **cli** (Command Line Interface).
To achieve this, kconf was planned to implement calls to all [Kong APIs](https://docs.konghq.com/gateway/api/admin-oss/latest/).
The initial idea for kconf was inpired by [GCP gcloud](https://cloud.google.com/sdk/gcloud/).

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

- <font color="orange">**-help**</font> - show kconf help
- <font color="orange">**-version**</font> - show kconf version
- <font color="orange">**-kong-address**</font> - set Kong configuration address (default "localhost")
- <font color="orange">**-port**</font> - set Kong configuration port (default 8001)
- <font color="orange">**-json-output**</font> - use json output for every command
- <font color="orange">**-verbose**</font> - run in verbose mode

The available commands are: status, add, query, list, update, delete and status.
The Kong entities are: service, route, consumer and plugin.

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

- <font color="green">**plugin**</font> - add a new consumer.
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

### Command <font color="green">list</font>

- <font color="green">**service**</font> - list all services.
This command doesn't have options.
If there are services in **Kong**, `kconf` will return a list of all services.

```sh
$ kconf list service
3302f59b-4bb0-410c-988b-d7e4e02a8c6e: Consulta-Bin --> https://api.pagar.me:443/bin/v1/499577
```

- <font color="green">**route**</font> - list all route.
This command doesn't have options.
If there are routes in **Kong**, `kconf` will return a list of all routes.

```sh
$ kconf list route
0ee7a361-0ac0-4468-b7b9-fc041d9c8ed7: Consulta-Bin - [GET] [http]:[/api/v1/bin/499577] --> Service Id: 3302f59b-4bb0-410c-988b-d7e4e02a8c6e
```

### Command <font color="green">update</font>

### Command <font color="green">delete</font>

### Consumer Plugins

```sh
kconf add consumer-basic-auth --id=${CUSTOMER_GUID} --user-name=guest --password=1234"
```

```sh
kconf add consumer-key-auth --id=${CUSTOMER_GUID} --key=d5a37fa6-b033-4107-a29f-ebf51b443968 --ttl=0"
```

```sh
kconf add plugin --name=key-auth --route-id=${ROUTE_GUID} --enabled=true
```

## kconf backlog

### Features backlog (for v0.3 release)

- [ ] Endpoint to add a new upstream
- [ ] Endpoint to get a upstream
- [ ] Endpoint to get a list of upstreams
- [ ] Endpoint to update a upstream
- [ ] Endpoint to delete a upstream
- [ ] Endpoint to add a new target
- [ ] Endpoint to get a target
- [ ] Endpoint to get a list of targets
- [ ] Endpoint to update a target
- [ ] Endpoint to delete a target
- [ ] Endpoint to add a IP Restriction pluging for a consumer
- [ ] Endpoint to add a Rate Limit pluging for a consumer
- [ ] Endpoint to add a Request Size Limit pluging for a consumer
- [ ] Endpoint to add a Syslog pluging for a consumer

### Features backlog (for v0.4 release)

- [ ] Endpoint to add a LDAP pluging for a consumer
- [ ] Endpoint to add a OAuth2 pluging for a consumer
- [ ] Endpoint to add a HMAC Auth pluging for a consumer
- [ ] Endpoint to add a Kong Functions (Pre) pluging for a consumer
- [ ] Endpoint to add a Kong Functions (Post) pluging for a consumer
- [ ] Endpoint to add a OpenTelemetry pluging for a consumer
- [ ] Endpoint to add a Correlation ID pluging for a consumer
- [ ] Endpoint to add a Request Transformer pluging for a consumer
- [ ] Endpoint to add a Response Transformer pluging for a consumer
- [ ] Endpoint to add a gRPC Web pluging for a consumer
- [ ] Endpoint to add a File Log pluging for a consumer
- [ ] Endpoint to add a HTTP Log pluging for a consumer
- [ ] Add parameter to Add Plugin command to specify plugin config
