#!  /usr/bin/ksh

#   test scenatio #06.8
export TEST_SCENARIO='06.8'
export DESCRIPTION='command update non existing route'

export TARGET_OPTIONS="-verbose update route --id=00000-00000 --protocols=http,https --methods=GET --paths=/api/v1/test_scenario-06.8"
export EXPECTED_EXIT_STATUS=255
export EXPECTED_RESULT='[error] route not found'
export EXPECTED_RESULT_TYPE='regex'

performFunctionalTestScenario
