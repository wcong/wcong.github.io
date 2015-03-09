---
layout: post
title: replication of elasticsearch
tags: elasticsearch replication
---

### replication of elasticsearch
#### what would you do if make a back up of shard
there are some method must have
*     read all the replication
*     accept the request and add one 
*     check shard and replication are the same
#### in the code
##### incremental
**AsyncShardOperationAction** in **TransportShardReplicationOperationAction**
**performReplicas* in ***performOnPrimary** in **doStart**
 