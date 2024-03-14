#!/bin/bash

# Acknowledgement: https://github.com/edulinq/autograder-server/blob/main/test.sh
# Get directory of this script
readonly THIS_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

function main() {
    # Make sure no arguments are passed
    if [[ $# -ne 0 ]]; then
        echo "USAGE: $0"
        exit 1
    fi

    # Exit on signal interrupt (ctrl-c)
    trap exit SIGINT

    # cd into directory of test.sh
    cd "${THIS_DIR}"

    # Track error count.
    local error_count=0

    # Run each test once, print verbose messages.
    go test -v -count=1 ./...
    if [[ ${?} -ne 0 ]] ; then
        ((error_count += 1))
    fi

    if [[ ${error_count} -gt 0 ]] ; then
        echo "Found $error_count issues."
    else
        echo "All tests passed."
    fi

    return ${error_count}
}

[[ "${BASH_SOURCE[0]}" == "${0}" ]] && main "$@"