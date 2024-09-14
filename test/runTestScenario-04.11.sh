#!  /usr/bin/ksh

#   test scenatio #04.11
export TEST_SCENARIO='04.11'
export DESCRIPTION='command query consumer'

export TARGET_OPTIONS="-verbose query consumer --id=${CUSTOMER_GUID}"
export EXPECTED_EXIT_STATUS=0
export EXPECTED_RESULT='^http response status code: 200 OK$'
export EXPECTED_RESULT_TYPE='regex'

performFunctionalTestScenario
