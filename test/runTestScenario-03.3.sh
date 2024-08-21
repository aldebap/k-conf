#!  /usr/bin/ksh

#   test scenatio #03.3
export TEST_SCENARIO='03.3'
export DESCRIPTION='command add service without parameters'

export TARGET_OPTIONS='add service'
export EXPECTED_EXIT_STATUS=255
export EXPECTED_RESULT='[error] fail sending add service command to Kong: 400 Bad Request'
export EXPECTED_RESULT_TYPE='string'

performFunctionalTestScenario
