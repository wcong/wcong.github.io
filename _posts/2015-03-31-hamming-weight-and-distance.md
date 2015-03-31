---
layout: post
title: 汉明距离和汉明重量
tags: go 汉明距离 汉明重量 信息论
description: 汉明距离和汉明重量 
---
### 汉明距离
两个等长字符串之间的汉明距离是两个字符串对应位置的不同字符的个数,它就是将一个字符串变换成另外一个字符串所需要替换的字符个数
例子:
*   1011101与1101001之间的汉明距离是3
*   abcd与abce的汉明距离是1
### 汉明重量
汉明重量是一串符号中非零符号的个数。因此它等同于同样长度的全零符号串的汉明距离。在最为常见的数据位符号串中，它是 1 的个数
#### 实现
    X与X-1相与得到的最低位永远是0，
    减1 操作将最右边的符号从0变到1，从1变到0，与操作将会移除最右端的1。
    如果最初X有N个1，那么经过N次这样的迭代运算，X将减到0 

``` go

package main

import (
	"fmt"
)

// calculate hamming weight
// base on one fact x&x-1 make the left 1 move to right postion
func HammingWeight(input uint64) int {
	count := 0
	for ; input > 0; count++ {
		input &= input - 1
	}
	return count
}

func main() {
	fmt.Println(HammingWeight(uint64(11111)))
}
``
```
