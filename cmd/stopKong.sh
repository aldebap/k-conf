#!  /usr/bin/sh

#   stop Postgres and Kong containers
docker kill kong-db kong-gateway konga-admin
docker container rm kong-db kong-gateway konga-admin
docker network rm kong-net
