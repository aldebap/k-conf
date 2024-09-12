#!  /usr/bin/ksh

#   test scenatio #03.4
export TEST_SCENARIO='03.4'
export DESCRIPTION='command add service with invalid value for option --enabled'

export TARGET_OPTIONS="add service --enabled=maybe"
export EXPECTED_EXIT_STATUS=255
export EXPECTED_RESULT='[error] wrong value for option --enabled: maybe'
export EXPECTED_RESULT_TYPE='string'

performFunctionalTestScenario
