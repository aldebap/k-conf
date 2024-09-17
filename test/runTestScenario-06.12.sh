#!  /usr/bin/ksh

#   test scenatio #06.12
export TEST_SCENARIO='06.12'
export DESCRIPTION='command update consumer'

export TARGET_OPTIONS="-verbose update consumer --id=${CUSTOMER_GUID} --custom-id=test-scenario-03.11 --user-name=root,guest --tags=gold-tier"
export EXPECTED_EXIT_STATUS=0
export EXPECTED_RESULT='^http response status code: 200 OK'
export EXPECTED_RESULT_TYPE='regex'

performFunctionalTestScenario
