#!  /usr/bin/ksh

#   test scenatio #07.9
export TEST_SCENARIO='07.9'
export DESCRIPTION='command delete plugin without parameters'

export TARGET_OPTIONS='delete plugin'
export EXPECTED_EXIT_STATUS=255
export EXPECTED_RESULT='[error] missing plugin id: option --id={id} required for this command'
export EXPECTED_RESULT_TYPE='string'

performFunctionalTestScenario
