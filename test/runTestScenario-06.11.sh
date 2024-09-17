#!  /usr/bin/ksh

#   test scenatio #06.11
export TEST_SCENARIO='06.11'
export DESCRIPTION='command update non existing consumer'

export TARGET_OPTIONS="-verbose update consumer --id=00000-00000 --custom-id=test-scenario-03.11 --user-name=root,guest --tags=gold-tier"
export EXPECTED_EXIT_STATUS=255
export EXPECTED_RESULT='[error] consumer not found'
export EXPECTED_RESULT_TYPE='regex'

performFunctionalTestScenario
