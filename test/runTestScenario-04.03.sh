#!  /usr/bin/ksh

#   test scenatio #04.3
export TEST_SCENARIO='04.3'
export DESCRIPTION='command query service without parameters'

export TARGET_OPTIONS='query service'
export EXPECTED_EXIT_STATUS=255
export EXPECTED_RESULT='[error] missing service id: option --id={id} required for this command'
export EXPECTED_RESULT_TYPE='string'

performFunctionalTestScenario
