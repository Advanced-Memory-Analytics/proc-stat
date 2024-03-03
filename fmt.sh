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

    # cd into directory of fmt.sh
    cd "${THIS_DIR}"

    gofmt -s -w .
}

[[ "${BASH_SOURCE[0]}" == "${0}" ]] && main "$@"