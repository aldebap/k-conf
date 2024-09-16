#!  /usr/bin/ksh

#   test scenatio #05.5
export TEST_SCENARIO='05.5'
export DESCRIPTION='command list consumer'

export TARGET_OPTIONS="-verbose list consumer"
export EXPECTED_EXIT_STATUS=0
export EXPECTED_RESULT='^http response status code: 200 OK$'
export EXPECTED_RESULT_TYPE='regex'

performFunctionalTestScenario
