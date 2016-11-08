#!/usr/bin/env python
# -*- coding: UTF-8 -*-
#
#file name:SendP2Email.py
#
#description:Verity the three switch by command and send the IP to email
#
#author:rutk1t0r
#
#Date:2016.08.29
#
#GPL
#

import smtplib
from email.mime.text import MIMEText
from email.header import Header
import fcntl
import struct 

# 作用：自动构造数据包认证校园网账号密码
# 描述：经过抓包可知在完成TCP的三次握手后浏览器加载了认证界面，
#		但是最核心的是获取到了账户和密码，直接构造http的post包到
#		认证交换机上的801端口，而且账户和密码都是明文打包在http数据包中
#  		因此不安全，但是为了方便，还是走一波吧，此python脚本可以开机启动
#		以达到自动认证校园网的功能,有兴趣的童鞋还可以实现注销的功能
#
# 导入 socket 模块
import socket
import commands
#网卡名称
iface_name = commands.getstatusoutput("ifconfig | awk '{print $1}' | head -n 1")[1]  #默认取了第一个元组内容

def Verify_account(user_account,user_pass):
	s = socket.socket()         # 创建 socket 对象
	host = "192.168.7.220"      #核心认证三层交换机IP地址 
	port = 801                  # 设置端口
	account = user_account      #校园网账号
	password = user_pass        #校园网账号的密码
	#构造http数据包头部字段和数据，几乎直接抄就行
	Post_header = "POST http://%s:%d/eportal/?c=ACSetting&a=Login&wlanuserip=null&wlanacip=null&wlanacname=null&port=&iTermType=1&session=null HTTP/1.1\r\n"%(host,port)
	Accept_header = "Accept: text/html, application/xhtml+xml, */*\r\n"
	Referer_header = "Referer: http://%s/a70.htm\r\n"%(host)
	Language_header = "Accept-Language: zh-CN\r\n"
	User_Agent_header = "User-Agent: Mozilla/5.0 (Windows NT 6.1; WOW64; Trident/7.0; rv:11.0) like Gecko\r\n"
	Content_Type_header = "Content-Type: application/x-www-form-urlencoded\r\n"
	Accept_Encoding_header = "Accept-Encoding: gzip, deflate\r\n"
	Host_header = "Host: %s:%d\r\n"%(host,port)
	#content_length，按照顺序此处为长度字段
	Connection_header = "Connection: Keep-Alive\r\n"
	Pragma_header = "Pragma: no-cache\r\n\r\n"
	Payload = "DDDDD=%s&upass=%s&R1=0&R2=&R6=0&para=00&0MKKey=123456"%(account,password)
	Content_Length_header = "Content-Length: %d\r\n"%(len(Payload))
	try:
		s.connect((host, port))   #先握个手，然后再发数据，并且等待回送，否则主机可就直接四次byebye啦
		data=Post_header + Accept_header + Referer_header + Language_header + User_Agent_header + Content_Type_header \
	 	+ Accept_Encoding_header + Host_header + Content_Length_header + Connection_header \
	 	+ Pragma_header + Payload
		s.send(data)
		print s.recv(1024)        #顺便打印一下吧
	except:
		print "[-]error..........."

	s.close()                    #正常的四次挥手~~不过此后貌似还是会有数据包的交互，但是已经认证成功了就没关系了


#其实用shell命令更快，hostname -I就可以
def get_local_ip(ifname): 
	s = socket.socket(socket.AF_INET, socket.SOCK_DGRAM) 
	inet = fcntl.ioctl(s.fileno(), 0x8915, struct.pack('256s', ifname[:15])) 
	ret = socket.inet_ntoa(inet[20:24]) 
	return ret

#验证网络是否通,抄网上的代码．．．
def is_internet(host="8.8.8.8", port=53, timeout=3):
	"""
	Host: 8.8.8.8 (google-public-dns-a.google.com)
	OpenPort: 53/tcp
	Service: domain (DNS/TCP)
	"""
	try:
		socket.setdefaulttimeout(timeout)
		socket.socket(socket.AF_INET, socket.SOCK_STREAM).connect((host, port))
		print "网络已经通畅"
		return True
	except Exception as ex:
		print ex.message
		return False

def main():
	ip_add =  get_local_ip(iface_name)  #网卡名称,一般PC的第一个都是以太网网卡
	# 第三方 SMTP 服务
	mail_host="smtp.163.com"  #设置服务器,此处用的163
	mail_user="username"    #用户名,example,或者手机号码,例如邮箱是xxxx@163.com,此处填写xxxx
	mail_pass="password"   #第三方服务口令 ,example，需要开通第三方授权,自己设置,最好不要随便泄露因为某些客户端直接需要授权码可以读取你的邮件


	sender = 'sender_account'  #发送者邮箱号,例如xxxx@163.com
	receivers = ['receiver_account']  # 接收邮件，可设置为你的QQ邮箱或者其他邮箱，其实都是自己。。。。。发现好像只能发给自己

	message = MIMEText('Python 邮件发送测试...以太网的IP地址：'+ip_add, 'plain', 'utf-8')
	message['From'] = "send_user"
	message['To'] =  "receive_user"

	subject = 'Python SMTP 邮件测试'
	message['Subject'] = Header(subject, 'utf-8')
	
	if not is_internet():
		Verify_account('account','passwd')  #!!!这里填写校园网账号密码的地方例如：2013********,********
	try:
		smtpObj = smtplib.SMTP() 
		print '[+]smtp对象建立成功!'
		smtpObj.connect(mail_host, 25)    # 25 为 SMTP 端口号
		print '[+]连接服务器端口成功'
		smtpObj.login(mail_user,mail_pass)  
		print '[+]登陆成功'
		smtpObj.sendmail(sender, receivers, message.as_string())
		print "[+]邮件发送成功"
	except smtplib.SMTPException:
		print "[-]Error: 无法发送邮件"
	
	smtpObj.close()

if __name__ == '__main__':
	main()

