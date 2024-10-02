#!  /usr/bin/ksh

#   test scenatio #03.16
export TEST_SCENARIO='03.16'
export DESCRIPTION='command add consumer-basic-auth'

export TARGET_OPTIONS="-verbose add consumer-basic-auth --id=${CUSTOMER_GUID} --user-name=guest --password=1234"
export EXPECTED_EXIT_STATUS=0
export EXPECTED_RESULT='^new plugin ID: (\S+)$'
export EXPECTED_RESULT_TYPE='regex_id'

performFunctionalTestScenario

#export PLUGIN_GUID="${REGEX_RESULT}"
