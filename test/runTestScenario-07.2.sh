#!  /usr/bin/ksh

#   test scenatio #07.2
export TEST_SCENARIO='07.2'
export DESCRIPTION='command delete with invalid entity'

export TARGET_OPTIONS='delete bug'
export EXPECTED_EXIT_STATUS=255
export EXPECTED_RESULT='[error] invalid entity for command delete: bug'
export EXPECTED_RESULT_TYPE='string'

performFunctionalTestScenario
