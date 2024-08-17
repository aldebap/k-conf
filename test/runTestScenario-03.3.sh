#!  /usr/bin/ksh

#   test scenatio #03.3
export TEST_SCENARIO='03.3'
export DESCRIPTION='command add service'

export TARGET_OPTIONS='-verbose add service --name=test-scenario-03.3 --url=http://localhost:8080/api/v1/test'
export EXPECTED_EXIT_STATUS=0
export EXPECTED_RESULT='^new service ID: '

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
    RESULT=$( cat ${OUTPUT} | grep "${EXPECTED_RESULT}" )

    if [ -z "${RESULT}" ]
    then
	    echo -e "${RED}[error] unexpected result:${LIGHTGRAY} '$( cat ${OUTPUT} )' should be '${EXPECTED_RESULT}'${NOCOLOR}"
	    exit 1
    else
        echo -e "[run-test] ${GREEN}   --- PASS${NOCOLOR}"
    fi
else
    if [ "$( cat ${ERROR} )" != "${EXPECTED_RESULT}" ]
    then
	    echo -e "${RED}[error] unexpected error:${LIGHTGRAY} '$( cat ${ERROR} )' should be '${EXPECTED_RESULT}'${NOCOLOR}"
	    exit 1
    else
        echo -e "[run-test] ${GREEN}   --- PASS${NOCOLOR}"
    fi
fi
