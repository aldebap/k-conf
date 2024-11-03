#!  /usr/bin/ksh

#   color constants
export  RED='\033[0;31m'
export  GREEN='\033[0;32m'
export  LIGHTGRAY='\033[0;37m'
export  NOCOLOR='\033[0m'

#   cli arguments parsing
export  SINGLE_TEST_SCRIPT=''
export  START_KONG='false'

if [ "${1}" == "--test-scenario" ]
then
    shift

    export  SINGLE_TEST_SCENARIO="${1}"

    if [ -z "${SINGLE_TEST_SCENARIO}" ]
    then
	    echo -e "${RED}[error] option --test-scenario requires an argument with test scenario name${NOCOLOR}"
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
export  TEST_SCRIPT=testScenario.ksh
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
	    return 1
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
	        return 1
        fi
    else
        if [ "$( cat ${ERROR} )" == "${EXPECTED_RESULT}" ]
        then
            echo -e "[run-test] ${GREEN}   --- PASS${NOCOLOR}"
        else
	        echo -e "${RED}[error] unexpected error:${LIGHTGRAY} '$( cat ${ERROR} )' should be '${EXPECTED_RESULT}'${NOCOLOR}"
	        return 1
        fi
    fi

    return 0
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

    export  SUCCESSFUL_TESTS=0
    export  FAILED_TESTS=0

    for TEST_SCENARIO_FILE in ${TEST_DIRECTORY}/testScenario*.json
    do
        export TEST_SCENARIO=$( cat ${TEST_SCENARIO_FILE} | jq --raw-output '.["scenario"]' )
        export DESCRIPTION=$( cat ${TEST_SCENARIO_FILE} | jq --raw-output '.["description"]' )
        export TARGET_OPTIONS=$( cat ${TEST_SCENARIO_FILE} | jq --raw-output '.["option"]' )
        export EXPECTED_EXIT_STATUS=$( cat ${TEST_SCENARIO_FILE} | jq --raw-output '.["expected-result"] | .["status"]' )
        export EXPECTED_EXIT_OUTPUT=$( cat ${TEST_SCENARIO_FILE} | jq --raw-output '.["expected-result"] | .["output"]' )
        export EXPECTED_EXIT_FORMAT=$( cat ${TEST_SCENARIO_FILE} | jq --raw-output '.["expected-result"] | .["format"]' )

        export PRE_TEST_SCRIPT=$( cat ${TEST_SCENARIO_FILE} | jq --raw-output '.["pre-test-script"]' )
        export POST_TEST_SCRIPT=$( cat ${TEST_SCENARIO_FILE} | jq --raw-output '.["post-test-script"]' )

        if [ -z "${SINGLE_TEST_SCENARIO}" -o "${TEST_SCENARIO}" == "${SINGLE_TEST_SCENARIO}" ]
        then
            cat <<TEST_SCRIPT > ${TEST_SCRIPT}
#!  /usr/bin/ksh

export TEST_SCENARIO='${TEST_SCENARIO}'
export DESCRIPTION='${DESCRIPTION}'

export TARGET_OPTIONS="${TARGET_OPTIONS}"
export EXPECTED_EXIT_STATUS=${EXPECTED_EXIT_STATUS}
export EXPECTED_RESULT='${EXPECTED_EXIT_OUTPUT}'
export EXPECTED_RESULT_TYPE='${EXPECTED_EXIT_FORMAT}'
TEST_SCRIPT

            if [ "${PRE_TEST_SCRIPT}" != "null" -a ! -z "${PRE_TEST_SCRIPT}" ]
            then
                echo ${PRE_TEST_SCRIPT} >> ${TEST_SCRIPT}
            fi

            cat <<TEST_SCRIPT >> ${TEST_SCRIPT}

performFunctionalTestScenario

TEST_SCRIPT

            if [ "${POST_TEST_SCRIPT}" != "null" -a ! -z "${POST_TEST_SCRIPT}" ]
            then
                echo ${POST_TEST_SCRIPT} >> ${TEST_SCRIPT}
            fi

            chmod u+x ${TEST_SCRIPT}
            . ./${TEST_SCRIPT}
            EXIT_STATUS=$?

            if [ ${EXIT_STATUS} -eq 0 ]
            then
                SUCCESSFUL_TESTS=$(( ${SUCCESSFUL_TESTS} + 1 ))
            else
                FAILED_TESTS=$(( ${FAILED_TESTS} + 1 ))
            fi
        fi
    done

    #   stop Kong afterwards
    if [ "${START_KONG}" == 'true' ]
    then
        . cmd/stopKong.sh
    fi

    #   final results
    echo -e "[run-test] ${SUCCESSFUL_TESTS} successful tests"

    if [ ${FAILED_TESTS} -ne 0 ]
    then
        echo -e "[run-test] ${RED}${FAILED_TESTS} failed tests"
        return 1
    fi

    return 0
}

export  TARGET=functional-test
export  PROJECT_TARGET=kconf

functionalTestTarget
EXIT_STATUS=$?

rm -f ${TEST_SCRIPT} ${OUTPUT} ${ERROR}

if [ ${EXIT_STATUS} -ne 0 ]
then
    exit 1
fi
