#!  /usr/bin/ksh

#   test scenatio #07.5
export TEST_SCENARIO='07.5'
export DESCRIPTION='command delete service without parameters'

export TARGET_OPTIONS='delete service'
export EXPECTED_EXIT_STATUS=255
export EXPECTED_RESULT='[error] missing service id: option --id={id} required for this command'
export EXPECTED_RESULT_TYPE='string'

performFunctionalTestScenario
