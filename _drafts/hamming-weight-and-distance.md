---
layout: post
title: 汉明距离和汉明重量
tags: go 汉明距离 汉明重量
description: 汉明距离和汉明重量 
---
### 前提
*   比较字符串
*   字符串等长
### 汉明距离
两个等长字符串之间的汉明距离是两个字符串对应位置的不同字符的个数,它就是将一个字符串变换成另外一个字符串所需要替换的字符个数
例子:
*   1011101与1101001之间的汉明距离是3
*   abcd与abce的汉明距离是1
### 汉明重量
汉明重量是一串符号中非零符号的个数。因此它等同于同样长度的全零符号串的汉明距离。在最为常见的数据位符号串中，它是 1 的个数
### 例子
#### Number of 1 Bits 
    Write a function that takes an unsigned integer and returns the number of '1' bits it has (also known as the Hamming weight).

    For example, the 32-bit integer '11' has binary representation 00000000000000000000000000001011, so the function should return 3.
#### code

``` python
class Solution:
    # @param n, an integer
    # @return an integer
    def hammingWeight(self, n):
        result =0
        while n >0:
            if n%2 == 1:
                result +=1
            n= n/2
        return result
```
