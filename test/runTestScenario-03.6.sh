#!  /usr/bin/ksh

#   test scenatio #03.6
export TEST_SCENARIO='03.6'
export DESCRIPTION='command add route without parameters'

export TARGET_OPTIONS='add route'
export EXPECTED_EXIT_STATUS=255
export EXPECTED_RESULT='[error] fail sending add route command to Kong: 400 Bad Request'
export EXPECTED_RESULT_TYPE='string'

performFunctionalTestScenario
