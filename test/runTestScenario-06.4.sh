#!  /usr/bin/ksh

#   test scenatio #06.4
export TEST_SCENARIO='06.4'
export DESCRIPTION='command delete route'

export TARGET_OPTIONS="-verbose delete route --id=${ROUTE_GUID}"
export EXPECTED_EXIT_STATUS=0
export EXPECTED_RESULT='^http response status code: 204 No Content$'
export EXPECTED_RESULT_TYPE='regex'

performFunctionalTestScenario
