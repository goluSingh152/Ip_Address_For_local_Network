# Ip_Address_For_local_Network
We can get all connected device's IP address of a local network.<br>
<b>Python</b><br>
1. I followed the number of steps:<br>
Based on user Requirement we generate each IP address.<br>
Pass this Ip address to sendPingRequest function and using subprocess and after getting its result we store each IP based on its live or not.<br>

<b>GOLANG</b><br>
I followed the same algorithm  for GOLANG which I used in Python but for the wide range of IP address, in that case, we need output very fast, For that scenario I used GOLANG.<br>
Here i used <b><i>goroutine</i></b> and <b><i>channel</i></b> to call multiple Ping function at a time and store into channel and that channel pass to any other method to check or validate that which device and which device is closed.
<br>
<p>This script is very useful when we have a number of device inside our organization, in that case, we can easily monitor them, or somehow we change our network, in that case, IP changed (we do not mention static IP address) so in that case go each device and check their IP address is very tough, but suppose we know our domain and IP range, in that case, we can easily monitor our all devices easily.
