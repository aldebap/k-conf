#!  /usr/bin/ksh

#   test scenatio #03.2
export TEST_SCENARIO='03.2'
export DESCRIPTION='command add with invalid entity'

export TARGET_OPTIONS='add bug'
export EXPECTED_EXIT_STATUS=255
export EXPECTED_RESULT='[error] invalid entity for command add: bug'
export EXPECTED_RESULT_TYPE='string'

performFunctionalTestScenario
