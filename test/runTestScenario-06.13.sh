#!  /usr/bin/ksh

#   test scenatio #06.13
export TEST_SCENARIO='06.13'
export DESCRIPTION='command update plugin without parameters'

export TARGET_OPTIONS='update plugin'
export EXPECTED_EXIT_STATUS=255
export EXPECTED_RESULT='[error] missing plugin id: option --id={id} required for this command'
export EXPECTED_RESULT_TYPE='string'

performFunctionalTestScenario
