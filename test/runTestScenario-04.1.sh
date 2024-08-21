#!  /usr/bin/ksh

#   test scenatio #04.1
export TEST_SCENARIO='04.1'
export DESCRIPTION='command query without entity'

export TARGET_OPTIONS='query'
export EXPECTED_EXIT_STATUS=255
export EXPECTED_RESULT='[error] missing entity for command query: available entities: service, route'
export EXPECTED_RESULT_TYPE='string'

performFunctionalTestScenario
