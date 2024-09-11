#!  /usr/bin/ksh

#   test scenatio #06.6
export TEST_SCENARIO='06.6'
export DESCRIPTION='command update service without invalid value for option --enabled'

export TARGET_OPTIONS="update service --id=${SERVICE_GUID} --enabled=maybe"
export EXPECTED_EXIT_STATUS=255
export EXPECTED_RESULT='[error] wrong value for option --enabled: maybe'
export EXPECTED_RESULT_TYPE='string'

performFunctionalTestScenario
