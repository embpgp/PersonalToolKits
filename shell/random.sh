#!/bin/bash
#
#
#Filename:random.sh
#
#Description:generate random number by shell
#
#Arthor:rutk1t0r
#
#Date:2016.12.31
#
#GPL
#
#Reference:Tinylab.org 《shellbook 》

#the env var RANDOM can generate 0~32767
#awk's rand func can generate 0~1

#example:if you use dash on ubuntu, it may not be work,so use bash
for i in $(seq 1 10); do echo $i $(($RANDOM/8192+3)) $(($RANDOM/10+3000)); done

echo "" | awk '{srand(); printf("%f", rand());}'
echo ""
