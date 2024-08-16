#!  /usr/bin/ksh

#   test scenatio #01
export TEST_SCENARIO='01'
export DESCRIPTION='version option'
export PROJECT_TARGET=kconf
export TARGET_OPTIONS='-version'
export EXPECTED_EXIT_STATUS=0
export EXPECTED_RESULT='kconf 0.1'

echo -e "[automated-test] ${PROJECT_TARGET}: ${GREEN}running test scenarion #${TEST_SCENARIO}: ${DESCRIPTION}${NOCOLOR}"

bin/${PROJECT_TARGET} ${TARGET_OPTIONS} > ${OUTPUT} 2> ${ERROR}

if [ $? -ne ${EXPECTED_EXIT_STATUS} ]
then
	echo -e "${RED}[error] unexpected exit status:${LIGHTGRAY} $?: expected: ${EXPECTED_EXIT_STATUS}${NOCOLOR}"
	exit 1
fi

if [ $? -eq 0 ]
then
    if [ "$( cat ${OUTPUT} )" != "${EXPECTED_RESULT}" ]
    then
	    echo -e "${RED}[error] unexpected result:${LIGHTGRAY} '$( cat ${OUTPUT} )' should be '${EXPECTED_RESULT}'${NOCOLOR}"
	    exit 1
    else
        echo -e "[automated-test] ${GREEN}   --- PASS${NOCOLOR}"
    fi
else
    if [ "$( cat ${ERROR} )" != "${EXPECTED_RESULT}" ]
    then
	    echo -e "${RED}[error] unexpected error:${LIGHTGRAY} '$( cat ${ERROR} )' should be '${EXPECTED_RESULT}'${NOCOLOR}"
	    exit 1
    else
        echo -e "[automated-test] ${GREEN}   --- PASS${NOCOLOR}"
    fi
fi

exit 0
