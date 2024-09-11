#!  /usr/bin/ksh

#   test scenatio #06.7
export TEST_SCENARIO='06.7'
export DESCRIPTION='command update service'

export TARGET_OPTIONS="-verbose update service --id=${SERVICE_GUID} --enabled=false"
export EXPECTED_EXIT_STATUS=0
export EXPECTED_RESULT='^http response status code: 200 OK'
export EXPECTED_RESULT_TYPE='regex'

performFunctionalTestScenario
