#!/bin/bash

kaliasdir=$KTHOME/kalias
argc=$#

function usage_help(){
    echo "Usage: k"
    echo "Options:"
    echo "-e, --edit               edit kalias file"
    echo "-h, --help               print this help message"
}

while [[ $# -gt 0 ]]
do
    option="$1"

    case $option in
        -e|--edit)
            echo "Edit $kaliasdir"
            code $kaliasdir
            shift # past argument
            break
            ;;
        -h|--help)
            usage_help
            shift # past argument
            break
            ;;
        *)
            echo "Option error"
            exit 1
            ;;
    esac
done


#Handle non-option arguments
if [[ $argc -eq 0 ]]; then
    cat $kaliasdir/kaliasrc
fi