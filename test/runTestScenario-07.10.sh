#!  /usr/bin/ksh

#   test scenatio #07.10
export TEST_SCENARIO='07.10'
export DESCRIPTION='command delete plugin'

export TARGET_OPTIONS="-verbose delete plugin --id=${PLUGIN_GUID}"
export EXPECTED_EXIT_STATUS=0
export EXPECTED_RESULT='^http response status code: 204 No Content$'
export EXPECTED_RESULT_TYPE='regex'

performFunctionalTestScenario
