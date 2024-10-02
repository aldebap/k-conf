#!  /usr/bin/ksh

#   test scenatio #03.17
export TEST_SCENARIO='03.17'
export DESCRIPTION='command add consumer-key-auth without parameters'

export TARGET_OPTIONS='add consumer-key-auth'
export EXPECTED_EXIT_STATUS=255
export EXPECTED_RESULT='[error] missing consumer id: option --id={id} required for this command'
export EXPECTED_RESULT_TYPE='string'

performFunctionalTestScenario
