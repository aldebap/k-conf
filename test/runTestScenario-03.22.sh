#!  /usr/bin/ksh

#   test scenatio #03.22
export TEST_SCENARIO='03.22'
export DESCRIPTION='command add consumer-jwt'

export TARGET_OPTIONS="-verbose add consumer-jwt --id=${CUSTOMER_GUID} --algorithm=HS256 --key=5ab5ae42-6227-4f49-a354-6eda3e19ff99 --secret=ff6d73d4-5f53-405a-8a5d-b2f03f405b14"
export EXPECTED_EXIT_STATUS=0
export EXPECTED_RESULT='^new plugin ID: (\S+)$'
export EXPECTED_RESULT_TYPE='regex_id'

performFunctionalTestScenario

#export PLUGIN_GUID="${REGEX_RESULT}"
