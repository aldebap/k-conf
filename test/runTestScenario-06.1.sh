#!  /usr/bin/ksh

#   test scenatio #06.1
export TEST_SCENARIO='06.1'
export DESCRIPTION='command delete without entity'

export TARGET_OPTIONS='delete'
export EXPECTED_EXIT_STATUS=255
export EXPECTED_RESULT='[error] missing entity for command delete: available entities: service, route'
export EXPECTED_RESULT_TYPE='string'

performFunctionalTestScenario
