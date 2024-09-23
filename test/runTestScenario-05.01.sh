#!  /usr/bin/ksh

#   test scenatio #05.1
export TEST_SCENARIO='05.1'
export DESCRIPTION='command list without entity'

export TARGET_OPTIONS='list'
export EXPECTED_EXIT_STATUS=255
export EXPECTED_RESULT='[error] missing entity for command list: available entities: service, route'
export EXPECTED_RESULT_TYPE='string'

performFunctionalTestScenario
