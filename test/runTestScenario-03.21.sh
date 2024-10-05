#!  /usr/bin/ksh

#   test scenatio #03.21
export TEST_SCENARIO='03.21'
export DESCRIPTION='command add consumer-jwt with non existing consumer'

export TARGET_OPTIONS="add consumer-jwt --id=00000-00000"
export EXPECTED_EXIT_STATUS=255
export EXPECTED_RESULT='[error] consumer not found'
export EXPECTED_RESULT_TYPE='string'

performFunctionalTestScenario
