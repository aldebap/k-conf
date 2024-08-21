#!  /usr/bin/ksh

#   test scenatio #04.6
export TEST_SCENARIO='04.6'
export DESCRIPTION='command query route'

export TARGET_OPTIONS="-verbose query route --id=${ROUTE_GUID}"
export EXPECTED_EXIT_STATUS=0
export EXPECTED_RESULT='^http response status code: 200 OK$'
export EXPECTED_RESULT_TYPE='regex'

performFunctionalTestScenario
