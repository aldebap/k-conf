#!  /usr/bin/ksh

#   color constants
export  RED='\033[0;31m'
export  GREEN='\033[0;32m'
export  LIGHTGRAY='\033[0;37m'
export  NOCOLOR='\033[0m'

#   set environment
export  VERBOSE='true'

#   function to execute the "release" target action
function releaseTarget {

    echo -e "[build] ${TARGET}: ${GREEN}build release target ${LIGHTGRAY}${PROJECT_TARGET}${NOCOLOR}"

    if [ ! -z "${PROJECT_TARGET}" ]
    then
        git tag -a "v${RELEASE_TAG}" -m "Release v${RELEASE_TAG}"
        git push origin "v${RELEASE_TAG}"

        goreleaser --clean
    fi
}

TARGET=release
PROJECT_TARGET=kconf
RELEASE_TAG=v0.3.0

releaseTarget
