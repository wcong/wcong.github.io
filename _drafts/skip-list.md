---
layout: post
title: skip list 
tags: go 数据结构 redis leveldb
category: 数据结构
---
### skip list
最近接触[redis](https://github.com/antirez/redis.git)，就看了一下redis的源码，发现redis内部存储的结构是某种链表。
又比较了同类型的[leveldb](https://github.com/google/leveldb.git)，发现和redis内部存储的结构类似。
搜索了一下发现是[skip list](http://en.wikipedia.org/wiki/Skip_list)，所以详细了解一下，并用go实现一个建议的skip list。
### 原理
