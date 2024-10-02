#!  /usr/bin/ksh

#   test scenatio #03.14
export TEST_SCENARIO='03.14'
export DESCRIPTION='command add consumer-basic-auth without parameters'

export TARGET_OPTIONS='add consumer-basic-auth'
export EXPECTED_EXIT_STATUS=255
export EXPECTED_RESULT='[error] missing consumer id: option --id={id} required for this command'
export EXPECTED_RESULT_TYPE='string'

performFunctionalTestScenario
