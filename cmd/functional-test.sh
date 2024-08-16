#!  /usr/bin/ksh

#   color constants
export  RED='\033[0;31m'
export  GREEN='\033[0;32m'
export  LIGHTGRAY='\033[0;37m'
export  NOCOLOR='\033[0m'

#   set environment
export  VERBOSE='true'
export OUTPUT=testScenario.out
export ERROR=testScenario.err

echo -e "[functional-test] ${TARGET}: ${GREEN}running functional tests on project ${PROJECT_TARGET}${NOCOLOR}"

for TEST_SCRIPT in test/runTestScenario*
do
    ksh -c ${TEST_SCRIPT}
done

rm -f ${OUTPUT} ${ERROR}
