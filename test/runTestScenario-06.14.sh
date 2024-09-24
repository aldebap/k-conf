#!  /usr/bin/ksh

#   test scenatio #06.14
export TEST_SCENARIO='06.14'
export DESCRIPTION='command update non existing plugin'

export TARGET_OPTIONS="-verbose update plugin --id=00000-00000 --enabled=false"
export EXPECTED_EXIT_STATUS=255
export EXPECTED_RESULT='[error] plugin not found'
export EXPECTED_RESULT_TYPE='regex'

performFunctionalTestScenario
