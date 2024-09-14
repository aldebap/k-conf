#!  /usr/bin/ksh

#   test scenatio #04.7
export TEST_SCENARIO='04.7'
export DESCRIPTION='command query non existing route'

export TARGET_OPTIONS="-verbose query route --id=00000-00000"
export EXPECTED_EXIT_STATUS=255
export EXPECTED_RESULT='[error] route not found'
export EXPECTED_RESULT_TYPE='string'

performFunctionalTestScenario
