# kconf

This is just a **cli** utility to configure [Kong Gateway](https://konghq.com/products/kong-gateway).
The inspiration came from [GCP gcloud](https://cloud.google.com/sdk/gcloud/)

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
- [ ] Set ENV variable with the ID from command add
- [ ] --silent option
- [ ] Test automation
- [ ] Action to build and test
- [ ] Create a dev branch
