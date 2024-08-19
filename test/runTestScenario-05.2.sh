#!  /usr/bin/ksh

#   test scenatio #05.2
export TEST_SCENARIO='05.2'
export DESCRIPTION='command list with invalid entity'

export TARGET_OPTIONS='list bug'
export EXPECTED_EXIT_STATUS=255
export EXPECTED_RESULT='[error] invalid entity for command list: bug'

function compareExpectedWithResult {
    return ${EXIT_STATUS} -ne ${EXPECTED_EXIT_STATUS}
}

performFunctionalTestScenario

unset -f compareExpectedWithResult
