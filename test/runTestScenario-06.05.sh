#!  /usr/bin/ksh

#   test scenatio #06.5
export TEST_SCENARIO='06.5'
export DESCRIPTION='command update non existing service'

export TARGET_OPTIONS="-verbose update service --id=00000-00000 --enabled=false"
export EXPECTED_EXIT_STATUS=255
export EXPECTED_RESULT='[error] service not found'
export EXPECTED_RESULT_TYPE='regex'

performFunctionalTestScenario
