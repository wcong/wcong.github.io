---
layout: post
title: go json-rpc
description: go json rpc,when valiable in struct is not always exposed
tags: [go,json,rpc]
---

#### go json rpc
go默认的[rpc](/2015/03/09/go-rpc.html)要求参数的所有成员必须是暴露的(首字母大写)

但是有些系统自带的类型某些参数不符合这个条件

比如*net/http*包内的*Request*的参数*URL*内有一个*Userinfo*内的变量不是暴露的

*net/rpc/jsonrpc*不依赖包变量是否导出,可以替换默认的*rpc*编码方式

跟默认的*rpc*使用起来有两个区别
1.  *server conn* 时使用*jsonrpc*的*ServeConn*而不是*rpc*的*ServeConn*

```go
jsonrpc.ServeConn(conn)
```
2.  *dial*时使用*jsonrpc*的*Dial*而不是*rpc*的*Dial*

```go
jsonrpc.Dial("tcp",":1100")
```
