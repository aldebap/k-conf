#!  /usr/bin/ksh

#   test scenatio #04.10
export TEST_SCENARIO='04.10'
export DESCRIPTION='command query non existing consumer'

export TARGET_OPTIONS="-verbose query consumer --id=00000-00000"
export EXPECTED_EXIT_STATUS=255
export EXPECTED_RESULT='[error] consumer not found'
export EXPECTED_RESULT_TYPE='string'

performFunctionalTestScenario
