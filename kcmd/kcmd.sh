#!/bin/bash

kcmddir=$KTHOME/kcmd
argc=$#

function usage_help(){
    echo "Usage: kcmd [-e|-h]"
    echo "Options:"
    echo "-e, --edit    edit kcmd file"
    echo "-h, --help    print this help message"
}

while [[ $# -gt 0 ]]
do
    option="$1"

    case $option in
        -e|--edit)
            echo "Edit $kcmddir"
            code $kcmddir
            shift # past argument
            break
            ;;
        -h|--help)
            usage_help
            shift # past argument
            break
            ;;
        -l|--list)
            grep ">>>>>>>>>>.*" $kcmddir/kcmd.txt
            shift # past argument
            break
            ;;
        -s|--search)
            sed -n "/>>>>>>>>>> $2/,/<<<<<<<<<< $2/p" $kcmddir/kcmd.txt
            shift # past argument
            shift # past value
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
    cat $kcmddir/kcmd.txt
fi


