#!  /usr/bin/ksh

#   color constants
export  RED='\033[0;31m'
export  GREEN='\033[0;32m'
export  LIGHTGRAY='\033[0;37m'
export  NOCOLOR='\033[0m'

#   cli arguments parsing
export  SINGLE_TEST_SCRIPT=''

if [ "${1}" == "--test-script" ]
then
    export  SINGLE_TEST_SCRIPT="${2}"

    if [ -z "${SINGLE_TEST_SCRIPT}" ]
    then
	    echo -e "${RED}[error] option --test-script requires an argument with test script name${NOCOLOR}"
	    exit 1
    fi
fi

#   set environment
export  VERBOSE='true'

export  TEST_DIRECTORY=test
export  OUTPUT=testScenario.out
export  ERROR=testScenario.err

#   function to execute the "functional-test" target action
function functionalTestTarget {

    echo -e "[build] ${TARGET}: ${GREEN}running functional tests on target ${LIGHTGRAY}${PROJECT_TARGET}${NOCOLOR}"

    #   first start Kong
    . cmd/startKong.sh
    sleep 10s

    for TEST_SCRIPT in ${TEST_DIRECTORY}/runTestScenario*
    do
        if [ -z "${SINGLE_TEST_SCRIPT}" -o "${TEST_SCRIPT}" == "${SINGLE_TEST_SCRIPT}" ]
        then
            . ${TEST_SCRIPT}
        fi
    done

    #   stop Kong afterwards
    . cmd/stopKong.sh
}

export  TARGET=functional-test
export  PROJECT_TARGET=kconf

functionalTestTarget

rm -f ${OUTPUT} ${ERROR}
