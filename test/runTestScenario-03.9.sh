#!  /usr/bin/ksh

#   test scenatio #03.9
export TEST_SCENARIO='03.9'
export DESCRIPTION='command add consumer'

export TARGET_OPTIONS="-verbose add consumer --custom-id=test-scenario-03.9 --user-name=guest --tags=silver-tier"
export EXPECTED_EXIT_STATUS=0
export EXPECTED_RESULT='^new consumer ID: (\S+)$'
export EXPECTED_RESULT_TYPE='regex_id'

performFunctionalTestScenario

export CUSTOMER_GUID="${REGEX_RESULT}"
