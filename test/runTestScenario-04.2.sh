#!  /usr/bin/ksh

#   test scenatio #04.2
export TEST_SCENARIO='04.2'
export DESCRIPTION='command query with invalid entity'

export TARGET_OPTIONS='query bug'
export EXPECTED_EXIT_STATUS=255
export EXPECTED_RESULT='[error] invalid entity for command query: bug'
export EXPECTED_RESULT_TYPE='string'

performFunctionalTestScenario
