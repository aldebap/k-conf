#!  /usr/bin/ksh

#   test scenatio #07.8
export TEST_SCENARIO='07.8'
export DESCRIPTION='command delete consumer'

export TARGET_OPTIONS="-verbose delete consumer --id=${CUSTOMER_GUID}"
export EXPECTED_EXIT_STATUS=0
export EXPECTED_RESULT='^http response status code: 204 No Content$'
export EXPECTED_RESULT_TYPE='regex'

performFunctionalTestScenario
