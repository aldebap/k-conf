#!  /usr/bin/ksh

#   test scenatio #04.14
export TEST_SCENARIO='04.14'
export DESCRIPTION='command query plugin'

export TARGET_OPTIONS="-verbose query plugin --id=${PLUGIN_GUID}"
export EXPECTED_EXIT_STATUS=0
export EXPECTED_RESULT='^http response status code: 200 OK$'
export EXPECTED_RESULT_TYPE='regex'

performFunctionalTestScenario
