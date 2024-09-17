#!  /usr/bin/ksh

#   test scenatio #06.10
export TEST_SCENARIO='06.10'
export DESCRIPTION='command update consumer without parameters'

export TARGET_OPTIONS='update consumer'
export EXPECTED_EXIT_STATUS=255
export EXPECTED_RESULT='[error] missing consumer id: option --id={id} required for this command'
export EXPECTED_RESULT_TYPE='string'

performFunctionalTestScenario
