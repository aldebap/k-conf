#!  /usr/bin/ksh

#   test scenatio #03.4
export TEST_SCENARIO='03.4'
export DESCRIPTION='command add service'

export TARGET_OPTIONS='-verbose add service --name=test-scenario-03.4 --url=http://localhost:8080/api/v1/test'
export EXPECTED_EXIT_STATUS=0
export EXPECTED_RESULT='^new service ID: (\S+)$'
export EXPECTED_RESULT_TYPE='regex_id'

performFunctionalTestScenario

export SERVICE_GUID="${REGEX_RESULT}"
