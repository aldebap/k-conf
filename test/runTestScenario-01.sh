#!  /usr/bin/ksh

#   test scenatio #01
export TEST_SCENARIO='01'
export DESCRIPTION='version option'

export TARGET_OPTIONS='-version'
export EXPECTED_EXIT_STATUS=0
export EXPECTED_RESULT='kconf 0.2'
export EXPECTED_RESULT_TYPE='string'

performFunctionalTestScenario
