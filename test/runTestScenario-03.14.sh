#!  /usr/bin/ksh

#   test scenatio #03.14
export TEST_SCENARIO='03.14'
export DESCRIPTION='command add consumer-basic-auth with non existing consumer'

export TARGET_OPTIONS="add consumer-basic-auth --id=00000-00000"
export EXPECTED_EXIT_STATUS=255
export EXPECTED_RESULT='[error] consumer not found'
export EXPECTED_RESULT_TYPE='string'

performFunctionalTestScenario
