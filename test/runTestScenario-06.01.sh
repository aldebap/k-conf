#!  /usr/bin/ksh

#   test scenatio #06.1
export TEST_SCENARIO='06.1'
export DESCRIPTION='command update without entity'

export TARGET_OPTIONS='update'
export EXPECTED_EXIT_STATUS=255
export EXPECTED_RESULT='[error] missing entity for command update: available entities: service, route'
export EXPECTED_RESULT_TYPE='string'

performFunctionalTestScenario
