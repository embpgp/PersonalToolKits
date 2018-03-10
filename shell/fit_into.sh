#!/bin/sh
#目的:
#两个文本，按行遍历第一个文本。取行里的第一个单词。
#然后在循环体里，遍历第二个文本，查找从第一个文本中得到单词的行，再查找行中的channel

#$1 is file1  #$2 is file2

if [ -z $1 ];then
	orders_file="orders.txt"
else
	orders_file=$1
fi

if [ -z $2 ];then
	properties_file="properties.txt"
else
	properties_file=$2
fi

orders_first_col=`awk '{print $1}' $orders_file`

for i in $orders_first_col 
do
	properties_val=`grep $i $properties_file | awk -F '=' '{print $3}'`
	orders_row=`grep $i $orders_file`
	sed -i  "s/$orders_row/$orders_row $properties_val/g" $orders_file
done
