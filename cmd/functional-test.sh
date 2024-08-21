#!  /usr/bin/ksh

#   color constants
export  RED='\033[0;31m'
export  GREEN='\033[0;32m'
export  LIGHTGRAY='\033[0;37m'
export  NOCOLOR='\033[0m'

#   cli arguments parsing
export  SINGLE_TEST_SCRIPT=''
export  START_KONG='false'

if [ "${1}" == "--test-script" ]
then
    shift

    export  SINGLE_TEST_SCRIPT="${1}"

    if [ -z "${SINGLE_TEST_SCRIPT}" ]
    then
	    echo -e "${RED}[error] option --test-script requires an argument with test script name${NOCOLOR}"
	    exit 1
    fi

    shift
fi

if [ "${1}" == "--start-kong" ]
then
    export  START_KONG='true'
fi

#   set environment
export  VERBOSE='true'

export  TEST_DIRECTORY=test
export  OUTPUT=testScenario.out
export  ERROR=testScenario.err

#   function to perform a functional test scenario
function performFunctionalTestScenario {

    echo -e "[run-test] ${TARGET}: ${GREEN}running test scenario: ${LIGHTGRAY}#${TEST_SCENARIO}: ${DESCRIPTION}${NOCOLOR}"

    bin/${PROJECT_TARGET} ${TARGET_OPTIONS} > ${OUTPUT} 2> ${ERROR}
    EXIT_STATUS=$?

    if [ ${EXIT_STATUS} -ne ${EXPECTED_EXIT_STATUS} ]
    then
	    echo -e "${RED}[error] unexpected exit status:${LIGHTGRAY} ${EXIT_STATUS}: expected: ${EXPECTED_EXIT_STATUS}${NOCOLOR}"
	    exit 1
    fi

    if [ ${EXIT_STATUS} -eq 0 ]
    then
        SUCCESS='false'

        if [ "${EXPECTED_RESULT_TYPE}" == 'string' ]
        then
            if [ "$( cat ${OUTPUT} )" == "${EXPECTED_RESULT}" ]
            then
                SUCCESS='true'
            fi
        else if [ "${EXPECTED_RESULT_TYPE}" == 'regex' ]
        then
            SUCCESS=$( cat ${OUTPUT} | perl -n -e "if( /${EXPECTED_RESULT}/ ) { print qq/true/; }" )
        else if [ "${EXPECTED_RESULT_TYPE}" == 'regex_id' ]
        then
            SUCCESS=$( cat ${OUTPUT} | perl -n -e "if( /${EXPECTED_RESULT}/ ) { print qq/true/; }" )
            export  REGEX_RESULT=$( cat ${OUTPUT} | perl -n -e "if( /${EXPECTED_RESULT}/ ) { print qq/\$1\n/; }" )
        fi
        fi
        fi

        if [ "${SUCCESS}" == 'true' ]
        then
            echo -e "[run-test] ${GREEN}   --- PASS${NOCOLOR}"
        else
	        echo -e "${RED}[error] unexpected result:${LIGHTGRAY} '$( cat ${OUTPUT} )' should be '${EXPECTED_RESULT}'${NOCOLOR}"
	        exit 1
        fi
    else
        if [ "$( cat ${ERROR} )" == "${EXPECTED_RESULT}" ]
        then
            echo -e "[run-test] ${GREEN}   --- PASS${NOCOLOR}"
        else
	        echo -e "${RED}[error] unexpected error:${LIGHTGRAY} '$( cat ${ERROR} )' should be '${EXPECTED_RESULT}'${NOCOLOR}"
	        exit 1
        fi
    fi
}

#   function to execute the "functional-test" target action
function functionalTestTarget {

    echo -e "[build] ${TARGET}: ${GREEN}running functional tests on target ${LIGHTGRAY}${PROJECT_TARGET}${NOCOLOR}"

    #   first start Kong
    if [ "${START_KONG}" == 'true' ]
    then
        . cmd/startKong.sh
        sleep 10s
    fi

    for TEST_SCRIPT in ${TEST_DIRECTORY}/runTestScenario*
    do
        if [ -z "${SINGLE_TEST_SCRIPT}" -o "${TEST_SCRIPT}" == "${SINGLE_TEST_SCRIPT}" ]
        then
            . ${TEST_SCRIPT}
        fi
    done

    #   stop Kong afterwards
    if [ "${START_KONG}" == 'true' ]
    then
        . cmd/stopKong.sh
    fi
}

export  TARGET=functional-test
export  PROJECT_TARGET=kconf

functionalTestTarget

rm -f ${OUTPUT} ${ERROR}
