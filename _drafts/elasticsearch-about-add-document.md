---
layout: defult
title: elasticsearch add document 
tags: elasticsearch
---
### how to add a document
**RestIndexAction** accept the rest data
**HeadersAndContextCopyClient** in **BaseRestHandler**  and abstract class **AbstractClient** 
**TransportIndexAction** *doExecute* start to do the index
**AsyncShardOperationAction** in **TransportShardReplicationOperationAction**