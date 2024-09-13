# kconf

[![Go Build](https://github.com/aldebap/kconf/actions/workflows/go.yml/badge.svg)](https://github.com/aldebap/kconf/actions/workflows/go.yml)

kconf is just a simple utility to configure [Kong Gateway](https://konghq.com/products/kong-gateway) using **cli** (Command Line Interface).
To achieve this, kconf was planned to implement calls to all [Kong APIs](https://docs.konghq.com/gateway/api/admin-oss/latest/).
The initial idea for kconf was based on [GCP gcloud](https://cloud.google.com/sdk/gcloud/).

## Building kconf

kconf is 100% written in Goland and this repo provides a simple script (ksh) to build it by just typing the following:

```sh
cmd/build.sh
```

## Using kconf

kconf **cli** accepts a command followed by an entity.
The available commands are: add, query, list, update, delete and status.
The Kong entities are: service, route and consumer.

For some commands, like command add, there are options to describe the entity to be added:

```sh
kconf add service --name=Products --url=http://localhost:8080/api/v1/products
```

```sh
kconf add route --name=Products --protocols=http --methods=GET,POST --paths=/api/v1/products --service-id=27168276282768
```

### Features backlog (for v0.2 release)

- [X] ~~Endpoint to update a service~~
- [X] ~~Endpoint to update a route~~
- [X] ~~Endpoint to add a new consumer~~
- [ ] Endpoint to get a consumer
- [ ] Endpoint to get a list of consumers
- [ ] Endpoint to update a consumer
- [ ] Endpoint to delete a consumer

### Features backlog (for v0.1 release)

- [X] ~~Endpoint to delete a service~~
- [X] ~~Endpoint to add a new route~~
- [X] ~~Endpoint to get a route~~
- [X] ~~Endpoint to get a list of routes~~
- [X] ~~Endpoint to delete a route~~
- [X] ~~Set ENV variable with the ID from command add (canceled)~~
- [X] ~~--silent option (replaced by --verbose)~~
- [X] ~~Test automation~~
- [X] ~~Action to build and test~~
- [X] ~~Create a dev branch~~
