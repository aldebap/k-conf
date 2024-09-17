#!  /usr/bin/ksh

#   test scenatio #06.7
export TEST_SCENARIO='06.7'
export DESCRIPTION='command update route without parameters'

export TARGET_OPTIONS='update route'
export EXPECTED_EXIT_STATUS=255
export EXPECTED_RESULT='[error] missing route id: option --id={id} required for this command'
export EXPECTED_RESULT_TYPE='string'

performFunctionalTestScenario
