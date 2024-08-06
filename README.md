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
