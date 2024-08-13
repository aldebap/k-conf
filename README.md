# kconf

This is just a **cli** utility to configure [Kong Gateway](https://konghq.com/products/kong-gateway).
The inspiration came from [GCP gcloud](https://cloud.google.com/sdk/gcloud/).

The goal of kconf is to implement calls to all [Kong APIs](https://docs.konghq.com/gateway/api/admin-oss/latest/)
using a **cli** interface.

## Building kconf

xpto

```sh
cmd/build.sh
```

## Using kconf

xpto

```sh
kconf -json-output add service --name=Products --url=http://localhost:8080/api/v1/products
```

### Features backlog (for v0.1 release)

- [ ] Endpoint to delete a service
- [X] ~~Endpoint to add a new route~~
- [X] ~~Endpoint to get a route~~
- [X] ~~Endpoint to get a list of routes~~
- [ ] Endpoint to delete a route
- [X] ~~Set ENV variable with the ID from command add (canceled)~~
- [X] ~~--silent option (replaced by --verbose)~~
- [ ] Test automation
- [X] ~~Action to build and test~~
- [ ] Create a dev branch
