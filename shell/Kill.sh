#!/bin/sh
#
#
#Filename:Kill.sh
#
#Description:fast kill process by names
#
#Arthor:rutk1t0r
#
#Date:2016.10.25
#
#GPL
#

Usage()
{
	echo "Usage:"
	echo "Please input like $0 [PROCESS_NAME]"
	echo "The [PROCESS_NAME] arg don't a need full name,but it must be unique for your needs"
}

#the number of args must be greater than 1
if [ $# -lt 1 ] ; then
	Usage $@
	exit 1
fi

#directly kill all process by pid
for process_name in $@ 
do
	kill -9 `ps -aux | grep $process_name | awk '{print $2}'`
done
