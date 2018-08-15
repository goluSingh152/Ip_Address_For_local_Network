import os
import subprocess
import socket

#local ip address of current machine when connected through internet
def getIpAddress():
	s = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
	s.connect(("8.8.8.8", 80))
	return s.getsockname()[0]


#send ping to given address and send yes or not based on result
def sendPingRequest(ip):
	response = subprocess.Popen("ping -b -c1 -W1 " + ip, shell=True, stdout=subprocess.PIPE).stdout.read()
	retValue = '_'.join(response.decode("utf-8").split(',')[-2].split(" "))
	print(retValue)
	if retValue == '_100%_packet_loss':
		return False
	else: 
		return True
#generate ip address form given ip class
def ip():
	upIpList = []
	first = 192
	second = 168
	third  = 43
	fourth = 0
	for i in range (43, 255):
		for j in range(110,255):
			lst = [str(first),str(second),str(i),str(j)]
			ip = '.'.join(lst)
			print(ip)
			if sendPingRequest(ip):
				print("Match", ip)
				upIpList.append(ip)
				if len(upIpList) == 2:
					return upIpList
	return upIpList	


def main():
	ipList = ip()
	print("Final Ip List :", ipList)
	

if __name__ == "__main__":
	main()
			
