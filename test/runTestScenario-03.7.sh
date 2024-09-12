#!  /usr/bin/ksh

#   test scenatio #03.7
export TEST_SCENARIO='03.7'
export DESCRIPTION='command add route'

export TARGET_OPTIONS="-verbose add route --name=test-scenario-03.7 --protocols=http --methods=GET,POST --paths=/api/v1/test_scenario-03.7 --service-id=${SERVICE_GUID}"
export EXPECTED_EXIT_STATUS=0
export EXPECTED_RESULT='^new route ID: (\S+)$'
export EXPECTED_RESULT_TYPE='regex_id'

performFunctionalTestScenario

export ROUTE_GUID="${REGEX_RESULT}"
