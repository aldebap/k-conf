#!  /usr/bin/ksh

#   test scenatio #07.6
export TEST_SCENARIO='07.6'
export DESCRIPTION='command delete service'

export TARGET_OPTIONS="-verbose delete service --id=${SERVICE_GUID}"
export EXPECTED_EXIT_STATUS=0
export EXPECTED_RESULT='^http response status code: 204 No Content$'
export EXPECTED_RESULT_TYPE='regex'

performFunctionalTestScenario
