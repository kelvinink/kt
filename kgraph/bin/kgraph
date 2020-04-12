#!/bin/bash

kgraphdir=$KTHOME/kgraph
argc=$#

function usage_help(){
    echo "Usage: kgraph [-e|-h]"
    echo "Options:"
    echo "-e, --edit                  edit kgraph file"
    echo "-h, --help                  print this help message"
    echo "-l, --list                  list all kgraph images"
    echo "-s, --search <img name>     search kgraph image by name"
    echo "-o, --open                  open image dir"
}

while [[ $# -gt 0 ]]
do
    option="$1"

    case $option in
        -e|--edit)
            echo "Edit $kgraphdir"
            code $kgraphdir
            shift # past argument
            break
            ;;
        -h|--help)
            usage_help
            shift # past argument
            break
            ;;
        -l|--list)
            ls $kgraphdir/img | cat
            shift # past argument
            break
            ;;
        -s|--search)
            ls  $kgraphdir/img | grep --color=auto $2
            shift # past argument
            shift # past value
            break
            ;;
        -o|--open)
            open $kgraphdir/img
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
    ls $kgraphdir/img | cat
fi


