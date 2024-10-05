#!  /usr/bin/ksh

#   test scenatio #03.20
export TEST_SCENARIO='03.20'
export DESCRIPTION='command add consumer-jwt without parameters'

export TARGET_OPTIONS='add consumer-jwt'
export EXPECTED_EXIT_STATUS=255
export EXPECTED_RESULT='[error] missing consumer id: option --id={id} required for this command'
export EXPECTED_RESULT_TYPE='string'

performFunctionalTestScenario
