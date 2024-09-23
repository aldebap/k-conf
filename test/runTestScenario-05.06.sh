#!  /usr/bin/ksh

#   test scenatio #05.6
export TEST_SCENARIO='05.6'
export DESCRIPTION='command list plugin'

export TARGET_OPTIONS="-verbose list plugin"
export EXPECTED_EXIT_STATUS=0
export EXPECTED_RESULT='^http response status code: 200 OK$'
export EXPECTED_RESULT_TYPE='regex'

performFunctionalTestScenario
