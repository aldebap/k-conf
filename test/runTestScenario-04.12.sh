#!  /usr/bin/ksh

#   test scenatio #04.12
export TEST_SCENARIO='04.12'
export DESCRIPTION='command query plugin without parameters'

export TARGET_OPTIONS='query plugin'
export EXPECTED_EXIT_STATUS=255
export EXPECTED_RESULT='[error] missing plugin id: option --id={id} required for this command'
export EXPECTED_RESULT_TYPE='string'

performFunctionalTestScenario
