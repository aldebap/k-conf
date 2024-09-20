#!  /usr/bin/ksh

#   test scenatio #03.11
export TEST_SCENARIO='03.11'
export DESCRIPTION='command add plugin with invalid value for option --enabled'

export TARGET_OPTIONS="add plugin --enabled=maybe"
export EXPECTED_EXIT_STATUS=255
export EXPECTED_RESULT='[error] wrong value for option --enabled: maybe'
export EXPECTED_RESULT_TYPE='string'

performFunctionalTestScenario
