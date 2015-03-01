---
layout: post
title: squid proxy server
tags: squid proxy server crawler
---
### squid
[Optimising Web Delivery](http://www.squid-cache.org/) 

### install 
ubuntu
``` bash
apt-get install squid3
```
### configure
file path */etc/squid3/squid.conf *
#### http proxy
find *http_port* set the port you want
#### allow ip 
find *http_access* 
write 
```
acl my src my_ip
http_access allow my
```
before 
```
http_access deny all
```
### reload configure
