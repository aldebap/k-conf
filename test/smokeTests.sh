#!  /usr/bin/ksh

#   color constants
export  RED='\033[0;31m'
export  GREEN='\033[0;32m'
export  LIGHTGRAY='\033[0;37m'
export  NOCOLOR='\033[0m'

export  OUTPUT=smokeTest.out
export  ERROR=smokeTest.err

./bin/kconf status

#   services
./bin/kconf add service --name=Consulta-Bin --url=https://api.pagar.me/bin/v1/499577 > ${OUTPUT} 2> ${ERROR}

export SERVICE_GUID=$( cat ${OUTPUT} )

echo "[info] service ${SERVICE_GUID} created"

#   routes
./bin/kconf add route --name=Consulta-Bin --protocols=http --methods=GET --paths=/api/v1/bin/499577 --service-id=${SERVICE_GUID} > ${OUTPUT} 2> ${ERROR}

export ROUTE_GUID=$( cat ${OUTPUT} )

echo "[info] route ${ROUTE_GUID} created"

#   key-auth plugin
./bin/kconf add plugin --name=basic-auth --route-id=${ROUTE_GUID} --enabled=true > ${OUTPUT} 2> ${ERROR}

export PLUGIN_GUID=$( cat ${OUTPUT} )

echo "[info] basic-auth plugin for the route ${PLUGIN_GUID} created"

#   consumer
./bin/kconf add consumer --user-name=external-customer --tags=bronze-tier > ${OUTPUT} 2> ${ERROR}

export CONSUMER_GUID=$( cat ${OUTPUT} )

echo "[info] consumer ${CONSUMER_GUID} created"

#   basic auth
./bin/kconf add consumer-basic-auth --id=${CUSTOMER_GUID} --user-name='guest' --password='1234' > ${OUTPUT} 2> ${ERROR}

#   attempt to call endpoint without authentication
export EXPECTED_RESULT='{"message":"Unauthorized"}'

curl --no-progress-meter localhost:8000/api/v1/bin/499577 > ${OUTPUT} 2> ${ERROR}

if [ "$( cat ${OUTPUT} )" == "${EXPECTED_RESULT}" ]
then
    echo -e "[run-test] ${GREEN}   --- PASS${NOCOLOR}"
else
	echo -e "${RED}[error] unexpected error:${LIGHTGRAY} '$( cat ${OUTPUT} )'${NOCOLOR}"
	exit 1
fi

#   attempt to call endpoint with invalid authentication credentials
export EXPECTED_RESULT='{"message":"Invalid authentication credentials"}'

curl --no-progress-meter --basic --user 'guest:kong' localhost:8000/api/v1/bin/499577 > ${OUTPUT} 2> ${ERROR}

if [ "$( cat ${OUTPUT} )" == "${EXPECTED_RESULT}" ]
then
    echo -e "[run-test] ${GREEN}   --- PASS${NOCOLOR}"
else
	echo -e "${RED}[error] unexpected error:${LIGHTGRAY} '$( cat ${OUTPUT} )'${NOCOLOR}"
	exit 1
fi

#   attempt to call endpoint with valid authentication credentials
export EXPECTED_RESULT='{"brand":"visa","brandName":"Visa","gaps":[4,8,12],"lenghts":[13,16],"mask":"/(\\d{1,4})/g","cvv":3,"brandImage":"https://dashboard.mundipagg.com/emb/images/brands/visa.jpg","possibleBrands":["visa"]}'

curl --no-progress-meter --basic --user 'guest:1234' localhost:8000/api/v1/bin/499577 > ${OUTPUT} 2> ${ERROR}

if [ "$( cat ${OUTPUT} )" == "${EXPECTED_RESULT}" ]
then
    echo -e "[run-test] ${GREEN}   --- PASS${NOCOLOR}"
else
	echo -e "${RED}[error] unexpected error:${LIGHTGRAY} '$( cat ${OUTPUT} )'${NOCOLOR}"
	exit 1
fi

#   key-auth plugin
#./bin/kconf -verbose add plugin --name=key-auth --route-id=${ROUTE_GUID} --enabled=true

#   clean up
#./bin/kconf delete route --id=${ROUTE_GUID}
#./bin/kconf delete service --id=${SERVICE_GUID}

rm -f ${OUTPUT} ${ERROR}
