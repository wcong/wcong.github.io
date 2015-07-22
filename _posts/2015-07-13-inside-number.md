---
layout: post
title: java数字存储
tags: number bit byte
category: java
---
### java数字存储
java有三大基本类型:char,boolean,number类型。number的包含两类:一类是类int型，另一类是类float型。

java中为了表示的2的n次方，用p来表示，比如：ox1.1p21f 代表1.1乘以2的21次方。

### 类int型
类int型主要是，byte，short，int，long,表示的方法类型，主要是bit length不一样，所以表示的范围也不一样。

#### little-endian 或 big-endian

对于大于两个byte的number型都会存在这个问题。java采用的是big-endian。大小位的区分是以byte为单位的。
比如int型的8：

* little-endian的表示是小位排在前面，16进制的表示是：08000000
* big-endian的表示是大位排在前面，16进制的表示是：00000008

#### int的表示

* int是sign的首位标记正负，0代表正数，1代表负数，后面的31为代表数的值。
* 0属于正数的范围,所以负数是从-1（oxf0000000）开始，-(((2^31)-1)+1)结束（即oxffffffff）,正数从0（ox00000000）开始，((2^31)-1)结束，即(0x7fffffff)。

### 类float型
这种类型的数的bits被分为了三个部分，sign，exponent，significand

float 是1个sign位，8个exponent位，23个significand位。
double是1个sign位，11个exponent位，52个signficand位。

以float为例，float的表示方法是，(sign(+或-))((1或0).(2^significand))x(2^exponent-127)

![float](/static/images/float.png)

#### point
* 有正负0
* 分为标准float和非标准
* exponent 也可能为负数，所以实际值要减去一个中间数。

#### sign位

* 0：代表正数
* 1：代表负数

#### exponent位 

代表significand需要乘的2的指数部分，因为要表示负数，所以要减去一个中间数。
比如 float 的exponent是8位，中间数是 (2^(8-1))-1,
另外，0x00跟0xff都是特殊值，有不一样的规则

* 0x00时，浮点数的指数等于1-127(不是0-127)，有效数字M不再加上第一位的1，而是还原为0.xxxxxx的小数。这样做是为了表示±0，以及接近于0的很小的数字
* 0xff时，这时，如果significand 全为0，表示±无穷大（正负取决sign）；如果有效数字M不全为0，表示这个数不是一个数（NaN）。

#### significand

代表实际的数值，但都是表示小数部分，需要乘以2^exponent来表示真实地数值。

* 标准float中，需要加1，(比如二进制(10000000000000000000000)代表0.10000000000000000000000)在计算实际值需要加上1，就是1.10000000000000000000000.
* 非标准的不需要加1

#### float 的二进制输出

Integer.toBinaryString() 会省略首位的0

``` java
System.out.println(Integer.toBinaryString(Float.floatToRawIntBits(1.11f)));
```
