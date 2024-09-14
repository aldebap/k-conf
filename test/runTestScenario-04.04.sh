#!  /usr/bin/ksh

#   test scenatio #04.4
export TEST_SCENARIO='04.4'
export DESCRIPTION='command query non existing service'

export TARGET_OPTIONS="-verbose query service --id=00000-00000"
export EXPECTED_EXIT_STATUS=255
export EXPECTED_RESULT='[error] service not found'
export EXPECTED_RESULT_TYPE='string'

performFunctionalTestScenario
