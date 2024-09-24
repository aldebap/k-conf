#!  /usr/bin/ksh

#   test scenatio #06.15
export TEST_SCENARIO='06.15'
export DESCRIPTION='command update plugin'

export TARGET_OPTIONS="-verbose update plugin --id=${PLUGIN_GUID} --enabled=false"
export EXPECTED_EXIT_STATUS=0
export EXPECTED_RESULT='^http response status code: 200 OK'
export EXPECTED_RESULT_TYPE='regex'

performFunctionalTestScenario
