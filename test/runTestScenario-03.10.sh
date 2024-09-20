#!  /usr/bin/ksh

#   test scenatio #03.10
export TEST_SCENARIO='03.10'
export DESCRIPTION='command add plugin without parameters'

export TARGET_OPTIONS='add plugin'
export EXPECTED_EXIT_STATUS=255
export EXPECTED_RESULT='[error] fail sending add plugin command to Kong: 400 Bad Request'
export EXPECTED_RESULT_TYPE='string'

performFunctionalTestScenario
