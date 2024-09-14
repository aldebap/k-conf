#!  /usr/bin/ksh

#   test scenatio #04.9
export TEST_SCENARIO='04.9'
export DESCRIPTION='command query consumer without parameters'

export TARGET_OPTIONS='query consumer'
export EXPECTED_EXIT_STATUS=255
export EXPECTED_RESULT='[error] missing consumer id: option --id={id} required for this command'
export EXPECTED_RESULT_TYPE='string'

performFunctionalTestScenario
