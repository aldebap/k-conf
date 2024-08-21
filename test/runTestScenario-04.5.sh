#!  /usr/bin/ksh

#   test scenatio #04.5
export TEST_SCENARIO='04.5'
export DESCRIPTION='command query route without parameters'

export TARGET_OPTIONS='query route'
export EXPECTED_EXIT_STATUS=255
export EXPECTED_RESULT='[error] missing route id: option --id={id} required for this command'
export EXPECTED_RESULT_TYPE='string'

performFunctionalTestScenario
