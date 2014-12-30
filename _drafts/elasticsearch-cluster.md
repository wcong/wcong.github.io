---
layout: post
title: elasticsearch cluster
tags:[elasicsearch cluster]
---


### elasticsearch cluster
#### socket dead loop
``` java
MulticastChannel
``` 
``` java
Plain
```
``` java 
MulticastSocket
```

### think
#### what a cluster would do
*.   discovery node by configure
*.   distribute request
*.   accept request and deal with it
*.   managenent node : backup 
#### so what will you do when you discovery a node
1.   start in *ZenDiscovery*  where *innerJoinCluster* is the main function
     *.	   handleJoinRequest
2.   start in *InternalClusterService*
