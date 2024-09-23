#!  /usr/bin/ksh

#   test scenatio #05.3
export TEST_SCENARIO='05.3'
export DESCRIPTION='command list service'

export TARGET_OPTIONS="-verbose list service"
export EXPECTED_EXIT_STATUS=0
export EXPECTED_RESULT='^http response status code: 200 OK$'
export EXPECTED_RESULT_TYPE='regex'

performFunctionalTestScenario
