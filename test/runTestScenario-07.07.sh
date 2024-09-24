#!  /usr/bin/ksh

#   test scenatio #07.7
export TEST_SCENARIO='07.7'
export DESCRIPTION='command delete consumer without parameters'

export TARGET_OPTIONS='delete consumer'
export EXPECTED_EXIT_STATUS=255
export EXPECTED_RESULT='[error] missing consumer id: option --id={id} required for this command'
export EXPECTED_RESULT_TYPE='string'

performFunctionalTestScenario
