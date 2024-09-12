#!  /usr/bin/ksh

#   test scenatio #06.7
export TEST_SCENARIO='06.7'
export DESCRIPTION='command update route'

export TARGET_OPTIONS="-verbose update route --id=${ROUTE_GUID} --protocols=http,https --methods=GET --paths=/api/v1/test_scenario-06.7"
export EXPECTED_EXIT_STATUS=0
export EXPECTED_RESULT='^http response status code: 200 OK'
export EXPECTED_RESULT_TYPE='regex'

performFunctionalTestScenario
