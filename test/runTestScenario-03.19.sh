#!  /usr/bin/ksh

#   test scenatio #03.19
export TEST_SCENARIO='03.19'
export DESCRIPTION='command add consumer-key-auth'

export TARGET_OPTIONS="-verbose add consumer-key-auth --id=${CUSTOMER_GUID} --key=d5a37fa6-b033-4107-a29f-ebf51b443968 --ttl=0"
export EXPECTED_EXIT_STATUS=0
export EXPECTED_RESULT='^new plugin ID: (\S+)$'
export EXPECTED_RESULT_TYPE='regex_id'

performFunctionalTestScenario

#export PLUGIN_GUID="${REGEX_RESULT}"
