---
layout: page 
title: 素数与加密算法
tags: 素数 go 加密算法
category: math
---

### menu
*	素数
*	关于素数的一些程序
*	RSA 加密算法

#### 素数
	数学上定义:素数，又称质数，是指在一个大于1的自然数中，除了1 和此整数自身之外，无法被其他自然数整除的数。
*	素数有个很奇妙的特性就是，所有的自然数都可以由n个素数的乘积得到
*	素数在自然数中不规则出现，没有一个固定的函数可以得到所有素数

#### 关于素数的一些程序
主要有两个关于素数的算法
*	判断一个数是否是素数
*	选出n以内的素数
##### 判断一个数是否是素数
最简答的方法就是根据定义来，从2 到n 除到 n-1 看能否被整除。
当饭根据乘法的交换律，只用除到m(m*m=n)就行了。
```go
package main

import (
	"fmt"
)

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isPrime(2))
}
```

##### 选出n以内的素数
最简单的方法就是，从1 到n每个都通过 isPrime 函数判断一下
因为所有的自然数都是有素数组成的，所以可以用筛选法，比如判断到2是素数的时候，所以n以内可以被2整除的就可以排除了

```go

func selectPrime(n int) []int {
	resultArray := make([]byte, n)
	for i := 0; i < n; i++ {
		resultArray[i] = 1
	}
	for i := 0; i < n; i++ {
		if resultArray[i] == 0 {
			continue
		}
		num := i + 1
		prime := isPrime(num)
		if !prime {
			resultArray[i] = 0
			continue
		}
		for j := i + 1; j < n; j++ {
			if (j+1)%num == 0 {
				resultArray[j] = 0
			}
		}
	}
	primeSlice := make([]int, 0)
	for i := 0; i < n; i++ {
		if resultArray[i] == 1 {
			primeSlice = append(primeSlice, i+1)
		}
	}
	return primeSlice
}
```

#### RSA cryptosystem
RSA 是一种流行的非对称的加密技术，即公开秘钥加密算法。

RSA 算法基础数RSA定理

	若P和Q是两个相异质数，另有正整数R和M,其中M的值与(P-1)(Q-1)的值互质，并使得(PM)mod(P-1)(Q-1)=1.有正整数A,且A<PQ,设:
	C=A的R次方 mod PQ
	B=C的M次方 mod PQ
	则有:A=B
	
所以 RSA 的程序逻辑是这样的:

*	随意选择两个大的素数P和Q,P!=Q
*	将P,Q两素数想乘得到一个数N,即N=PQ
*	将P,Q分别减1,再相乘,得到一个数T,即T=(P-1)(Q-1)
*	选择一个整数E,作为一个密钥,使E和T互质,且E必须小于T
*	根据公式DEmodT=1,计算出D的值,作为另一个密钥
*	通过以上步骤计算得出N,E,D3个数据,其中(N,E)作为公钥,(N,D)作为私钥,可以互换

程序稍后实现
	
