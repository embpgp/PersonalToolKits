#!/bin/bash
#
#Filename:calc.sh
#
#Description:simple calc for conversion system
#
#Author:rutk1t0r
#
#Date:2017.3.19
#
#GPL
#

Usage()
{
	echo "Usgae:"
	echo "$0 -i ibase -o obase number"
	echo "the base should be range 2 to 16 and above 10 must be capital"
	echo "example:$0 -i 16 -o 10 FF"
}

if [ $# -ne 5 ]; then
	Usage $@
	exit 1
fi

while getopts "i:o:" arg
do
	case $arg in
		i)
			ibase=$OPTARG
			;;
		o)
			obase=$OPTARG
			;;
		?)
			;;
	esac

done

shift $(($OPTIND - 1))

echo "obase=$obase;ibase=$ibase;$1" | bc -l
