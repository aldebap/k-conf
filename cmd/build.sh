#!  /usr/bin/ksh

#   color constants
export  RED='\033[0;31m'
export  GREEN='\033[0;32m'
export  LIGHTGRAY='\033[0;37m'
export  NOCOLOR='\033[0m'

#   set environment
export  BINARY_DIRECTORY='./bin'
export  VERBOSE='true'

#   function to execute the "build" target action
function buildTarget {

    echo -e "[build] ${TARGET}: ${GREEN}build target ${PROJECT_TARGET}${NOCOLOR}"

    if [ ! -d ${BINARY_DIRECTORY} ]
    then
        mkdir ${BINARY_DIRECTORY}
    fi

    if [ ! -z "${PROJECT_TARGET}" ]
    then
        go build -o ${BINARY_DIRECTORY}/${PROJECT_TARGET} ${BUILD_PARAMS}
    fi
}

TARGET=build
PROJECT_TARGET=kconf
BUILD_PARAMS=

buildTarget
