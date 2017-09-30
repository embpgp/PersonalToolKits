#!/bin/sh
#add this file to /etc/cron.d or other dir that it can be runned by cron

resolv_conf=/etc/resolv.conf
dns_entry="nameserver 114.114.114.114"
grep "${dns_entry}" ${resolv_conf} >> /dev/null
if [ $? -ne 0 ]; then
	echo "${dns_entry}" >> ${resolv_conf}
fi
