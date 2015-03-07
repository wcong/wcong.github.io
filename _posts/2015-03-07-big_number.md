---
layout: page
title: 大整数
tags: go big-number
category: go
---
### 大整数
大整数是区别于程序内置的证书来说的
程序内置的证书受计算机的限制(64位或32位)
对于某些加密算法的无法满足需求,所以就有了大整数的需求
#### 实现原理
内置 []int 保存10进制各个位数
#### 代码实现
##### struct
posotion 用来表示正负,
number 是一个[]byte 用来存放  

```go
type BigNumber struct {
    positive bool
    number   []byte
}
```

##### 加,减,乘,幂
代码实现在[bignumber](https://github.com/wcong/big-number)

加减法是先计算正负值,根据正负值计算[]byte加减结果
参见一下方法

```go
func (this *BigNumber) SliceSubtract(number []byte) {}
func (this *BigNumber) SliceAdd(number []byte) {}
```

乘法是先计算正负,再计算乘数每一个和被乘数相乘,得到结果集再相加
参见方法

```go
func (this *BigNumber) SliceMultiply(number []byte) {}
```

除法是比较复杂的,同样是先计算正负,在根据基本的除法计算,得到[]byte的除法计算值,由于计算过程中,存在借位和首位为0的情况,所以做了很多处理
参见方法

```go
func (this *BigNumber) SliceDivide(number []byte) {}
```

