#!  /usr/bin/ksh

#   test scenatio #03.13
export TEST_SCENARIO='03.13'
export DESCRIPTION='command add plugin'

export TARGET_OPTIONS="-verbose add plugin --name=key-auth --route-id=${ROUTE_GUID} --enabled=true"
export EXPECTED_EXIT_STATUS=0
export EXPECTED_RESULT='^new plugin ID: (\S+)$'
export EXPECTED_RESULT_TYPE='regex_id'

performFunctionalTestScenario

export PLUGIN_GUID="${REGEX_RESULT}"
