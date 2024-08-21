#!  /usr/bin/ksh

#   test scenatio #03.1
export TEST_SCENARIO='03.1'
export DESCRIPTION='command add without entity'

export TARGET_OPTIONS='add'
export EXPECTED_EXIT_STATUS=255
export EXPECTED_RESULT='[error] missing entity for command add: available entities: service, route'
export EXPECTED_RESULT_TYPE='string'

performFunctionalTestScenario
