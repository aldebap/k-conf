#!  /usr/bin/ksh

#   test scenatio #03.8
export TEST_SCENARIO='03.8'
export DESCRIPTION='command add consumer without parameters'

export TARGET_OPTIONS='add consumer'
export EXPECTED_EXIT_STATUS=255
export EXPECTED_RESULT='[error] fail sending add consumer command to Kong: 400 Bad Request'
export EXPECTED_RESULT_TYPE='string'

performFunctionalTestScenario
