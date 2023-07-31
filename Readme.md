
# YAGI 

YAGI is a command-line tool written in Go (Golang) that offers various network-related functionalities and utilities. It aims to simplify and streamline network troubleshooting, analysis, and management tasks, making it easier for users to interact with their network infrastructure from the command line.


## Features

 1) Traceroute: YAGI allows users to perform traceroutes to a specified destination. Traceroute helps diagnose the path taken by data packets across the network, helping to identify possible bottlenecks or connectivity issues.

2) Reverse DNS Lookup: The tool enables users to perform reverse DNS lookups for given IP addresses. This feature helps identify domain names associated with specific IP addresses.

3) Bandwidth Testing: YAGI provides the ability to conduct bandwidth testing, helping users measure network performance and identify potential bandwidth constraints.

4) DHCP Management: Users can manage Dynamic Host Configuration Protocol (DHCP) leases with YAGI. This includes options to release, renew, or flush DHCP leases for better control over network IP assignments.

5) Network Topology: YAGI can display the network topology, offering an overview of how network devices are interconnected.

6) Network Information: Users can obtain essential network information, such as IP addresses, subnet masks, gateway addresses, and more.

7) Port Scanning: YAGI supports port scanning on specified destinations. Port scanning helps identify open ports on network devices, aiding in security assessments.

















## Getting Started

Installation: To use YAGI, you need to install Go on your system and then compile and run the YAGI source code.

Command Syntax: The general syntax for using YAGI is command [destination]. Some commands may require additional parameters or flags.

Help: Typing help or -h within the YAGI prompt displays a manual page that lists all available commands and their usage instructions.
## Usage/Examples

The YAGI tool is used via the command line interface (CLI). Users interact with the tool by entering commands and, in some cases, providing a destination IP address or domain name as input for specific functionalities.


Perform a traceroute:
```Javasript 
tr google.com
>>Tracing route to google.com [142.250.183.78]
over a maximum of 30 hops:

  1     6 ms     6 ms     5 ms  0.0.0.0
  2     6 ms     6 ms     5 ms  telmedia-0.0.0.0.broadband [0.0.0.0]
  3     7 ms     7 ms     6 ms  corporate-0.0.0.0.airtel [0.0.0.0]
  4     9 ms     9 ms     9 ms  0.0.0.0
  5    11 ms    11 ms    10 ms  0.0.0.0
  6    13 ms    13 ms    13 ms  0.0.0.0
  7     9 ms    10 ms     9 ms  0.0.0.0
  8    11 ms    11 ms    11 ms  bam-h177.00.net [0.0.0.0]

Trace complete.

```
IP addresses are removed for privacy reasons 

Reverse DNS lookup: 
```Javascript 
rDNS 8.8.8.8
>>Domain names associated with IP address 8.8.8.8:dns.google.
```

Perform bandwidth testing:
```Javascript
bwidth 
>>Downloaded 21371 bytes in 915.6744ms
  Estimated Bandwidth: 0.18 Mbps
```

DHCP management:
```Javascript 
DHCPManage -release 
>>DHCP lease release in progress...
  DHCP lease released successfully for IP address 0.0.0.0

DHCPManage -renew
>>DHCP lease renewed successfully.
 New lease details:
   IP Address: 0.0.0.0
   Subnet Mask: 255.255.255.255
   Default Gateway: 127.0.0.1
   Lease Duration: 86400 seconds
   DNS Servers: 8.8.8.8, 8.8.4.4

DHCPManage -flush
>>Flushing DHCP lease in progress...
 DHCP lease flushed successfully.

 The IP address 0.0.0.0 has been released and is now available for lease.
 ```

Perform a port scan:
```Javascript 
portscan 8.8.8.8
>>Open ports for 8.8.8.8 (8.8.8.8)
 53      dns
 110     pop3
 443     https
```

Print network information::
```Javascript
netinfo
>>Network information
  Interface: eth0
  MAC Address: 12:34:56:78:90:AB
  IP Address: 192.168.1.100
  Subnet Mask: 255.255.255.0
  Default Gateway: 192.168.1.1
  DNS Servers: 8.8.8.8, 8.8.4.4

Interface: wlan0
  MAC Address: AA:BB:CC:DD:EE:FF
  IP Address: 10.0.0.50
  Subnet Mask: 255.255.255.0
  Default Gateway: 10.0.0.1
  DNS Servers: 10.0.0.1
