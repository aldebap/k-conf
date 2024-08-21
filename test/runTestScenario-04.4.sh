#!  /usr/bin/ksh

#   test scenatio #04.4
export TEST_SCENARIO='04.4'
export DESCRIPTION='command query service'

export TARGET_OPTIONS="-verbose query service --id=${SERVICE_GUID}"
export EXPECTED_EXIT_STATUS=0
export EXPECTED_RESULT='^http response status code: 200 OK$'
export EXPECTED_RESULT_TYPE='regex'

performFunctionalTestScenario
