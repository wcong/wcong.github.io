---
layout: post
title: why set integer max value for size perform that way
tags: elasticsearch size 
---

### elasticsearch about size
we have meet a situation of large gc,where we found elasticsearch allocate the value of size literature,if you set the max value of integer ,it will ask memory to split that memcache
#### about the code
about the chain
``` java
RestSearchAction
```
``` java
NodeClient
```
``` java
TransportAction
```
``` java
TransportSearchAction
```
``` java
ShardSearchTransportRequest
```
acture send the request **SearchServiceTransportAction**
**SearchService** do the search in lucene
the reasult of search **QuerySearchResult**
#### where is transport handle
``` java
TransportRequestHandler
```
and **ActionNames** bind all the handle