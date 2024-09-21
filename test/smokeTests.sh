#!  /usr/bin/ksh

export  OUTPUT=smokeTest.out
export  ERROR=smokeTest.err

./bin/kconf status

#   services
./bin/kconf add service --name=Consulta-Bin --url=https://api.pagar.me/bin/v1 > ${OUTPUT} 2> ${ERROR}

export SERVICE_GUID=$( cat ${OUTPUT} )

echo "service ${SERVICE_GUID} created"

#   routes
./bin/kconf add route --name=Consulta-Bin --protocols=http --methods=GET --paths=/api/v1/bin/499577 --service-id=${SERVICE_GUID} > ${OUTPUT} 2> ${ERROR}

export ROUTE_GUID=$( cat ${OUTPUT} )

echo "route ${ROUTE_GUID} created"

#   key-auth plugin
./bin/kconf -verbose add plugin --name=key-auth --route-id=${ROUTE_GUID} --enabled=true

#   clean up
#./bin/kconf delete route --id=${ROUTE_GUID}
#./bin/kconf delete service --id=${SERVICE_GUID}

rm -f ${OUTPUT} ${ERROR}
