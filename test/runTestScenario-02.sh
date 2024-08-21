#!  /usr/bin/ksh

#   test scenatio #02
export TEST_SCENARIO='02'
export DESCRIPTION='command status'

export TARGET_OPTIONS='status'
export EXPECTED_EXIT_STATUS=0
export EXPECTED_RESULT='200 OK'
export EXPECTED_RESULT_TYPE='string'

performFunctionalTestScenario
