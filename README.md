# Phonebook


This is my first project on the DevOps roadmap.
I follow these goals on it:

- Linux operation system and Its overall structure 
- basic networking concepts(UDP, TCP, IP, DNS, ARP, Router, gateway) and working with network tools on Linux( netstat, arp, route, dig, tcpdump, ...) 
- Firewall concepts and iptable command webservers and their various usages DMZ concept
- database concepts and SQL 
- HTTP protocol
- TLS handshake concepts, symmetric and asymmetric encryption, ca and certificate backup and recovery concepts
- s3
- familiarity with systemd and service management in Linux 

This project contains four main modules: Edge web Server, Rest-API server, Database Server, and GateWay.

# Edge web Server
- A webServer on a cloud server.
- This server has two network interfaces. One with public IP and another one with private IP .
- Webserver serves two domains(for example, hachi.ir, and monamoghadampanah.ir)   on HTTPS .
- First one  must serve a WordPress website .
- Second one serves a static page, and also if user requests to monamoghadampanah.ir/phonebook, user requests must send to Rest-Api server.
- Rest-API requests must support GET, POST, DELETE, PUT methods. Requests forwarded to Rest-API must be HTTP.
- This server is only available on ports 22, 443, and 80 using IPTABLES. 
- Local Ip must be a static IP address on os. 
- Use ubuntu 20.04 distribution 

# Rest-API server:

- On this server, an application as a service serves on 18080 port .
- The application receives requests forwarded from the webserver and does actions on the database (The restful standard must be implemented.)
- 18080 port must only be available for webserver private IP.
- This server only has private IP and access to the internet only available from gateway server.
- SSH to this server only available from gateway server and ssh key .Local IP must be static.
- Use ubuntu 20.04 distribution.
- Rest application must start automatically after server reboot.


# GateWay:
- As Mysql and Rest-API server only has local IP, this server act as an internet gateway for them.
- All backups available on /op/db-backup must send to ArvanClous storage.
- All backups on ArvanClous storage older than five days must be removed.
- Use ubuntu 20.04 distribution.
- Local IP must be static.
