#!  /usr/bin/ksh

#   test scenatio #05.4
export TEST_SCENARIO='05.4'
export DESCRIPTION='command list route'

export TARGET_OPTIONS="-verbose list route --id=${ROUTE_GUID}"
export EXPECTED_EXIT_STATUS=0
export EXPECTED_RESULT='^http response status code: 200 OK$'
export EXPECTED_RESULT_TYPE='regex'

performFunctionalTestScenario
