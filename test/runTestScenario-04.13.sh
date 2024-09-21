#!  /usr/bin/ksh

#   test scenatio #04.13
export TEST_SCENARIO='04.13'
export DESCRIPTION='command query non existing plugin'

export TARGET_OPTIONS="-verbose query plugin --id=00000-00000"
export EXPECTED_EXIT_STATUS=255
export EXPECTED_RESULT='[error] plugin not found'
export EXPECTED_RESULT_TYPE='string'

performFunctionalTestScenario
