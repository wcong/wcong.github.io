---
layout: post
title: java 内存管理
tags: java 内存 gc
category: java
---

### java 内存管理

  针对openjdk7

  学习**深入理解java虚拟机**，这里写下自己的理解，主要包括以下三个方面

  * 内存布局
  * 内存gc算法
  * 内存gc收集器

### 1 内存布局
  java的内存主要分为以下几项：

#### 1.1 heap

  最大的应该heap(堆):大部分对象的实体都放在heap里面。也是gc的主要工作对象。根据gc算法的不同，具体布局也不同。

  目前的主要算法是分代算法，分为young generation和old generation。

  young generation还分为eden，from survivor,to survivor.一般比例是8:1:1。一般新生对象会分配在eden中，发生gc是会把幸存的对象放到，某一个survivor,当sorvivor满了，回把对象在放到,old generation。 

#### 1.2 程序计数器
  线程私有的，是当前进程所执行的字节码的行号指示器。

#### 1.3 java虚拟机栈
  线程私有，每个方法在执行的时候会创建虚拟机栈帧，用于存储局部变量表，操作数栈，动态链接，方法出口。每一个方法从调用到执行完成，对应着，栈帧在虚拟机栈中入栈出栈的过程。

#### 1.4 本地方法栈
  线程私有，使用的本地方法，如io,thread.
  
#### 1.5 方法区
  线程共享，保存类信息，常量，静态变量等信息。

### 2 gc算法

 gc的主要作用是标记，处理(清除，拷贝，整理)

#### 2.1 标记 
  1. 引用计数法。是最基本的标记算法，很难解决对象之间循环引用的问题。
  2. 可达性分析。通过gc roots 为起始点，向下搜索，没有引用链的，证明对象不可用。

	java中的gc roots包括下面几个:
	* 虚拟机栈中的引用对象
	* 方法区中类静态属性引用的对象
	* 方法七中常量引用的对象
	* 本地方法栈引用的对象

#### 2.2 处理的类型
  标记完成后回对对象处理，处理的方式不同，分为不同的算法:目前分为清除，整理，拷贝三中方法。

  不同的处理算法涉及到对象内存的分配：

1. 指针碰撞法，是指在内存空间中对象之间紧密排列，分配新对象说只需要拿到末尾指针，往后分配对应大小的空间就行了。对应的是整理和拷贝方法，都是把幸存的对象整齐排列，重新计算末尾指针位置

2. free list:单独维护一个列表那些地址没有分配，对应的是清除方法。gc时，只是在free list中标记释放的内存，下次分配对象时可以重新使用。缺点是，产生内存碎片，但是回收速度快，不需要整理内存。

  java的hotspot是使用分代处理算法，young generation，和old generation 使用不同的回收算法。


##### 2.2.1 标记-清除(mark-sweep)算法
* 标记所有需要回收的对象，在标记完成后同意回收被标记的对象，
* 最主要的问题是会产生内存碎片

##### 2.2.2 标记-拷贝(mark-copy)算法
* 事先把内存分为相等的两块内存（A,B），使用其中一块(A)，gc完成时把幸存的对象整齐的拷贝到另一块(B)，在继续使用（B）分配新对象，下次gc后在使用（A），以此类推。
* 缺点是浪费内存

##### 2.2.3 标记-整理(mark-compact)算法
* gc完成后，把幸存的对象重新整理。
* 缺点是整理耗时。

##### 2.2.4 分代整理
* 综合上面几个算法，不同的generation使用不同的算法
* 比如，young genration 每次gc都有大量对象死去，所以采用mark-copy算法，
* 一般，新对象分配在eden区，gc时，把幸存的对象拷贝到survivor区，survivo也有两个区，每次gc切换survivor。
* old generation 对象存活率较高，使用mark-sweep或，mark-compact

### 3 gc收集器

* 采用分代的收集器，每个代的收集器都不一样。

* g1收集器不采用分代，统一一个算法

#### 3.1 分代收集器
  主要是分为串行（serial），并行（parallel），并发（concurrent）三个大的类别
##### 3.1.1 young genetation

* Serial(mark-copy):stop the world,然后单线程收集,可以配合CMS
* ParNew(mark-copy):stop the world,然后多线程收集,可以配合CMS
* Parallel Scavenge(mark-copy):stop the world,有限制的多线程收集

##### 3.1.2 old generation

* Concurrent Mark Sweep(as title):部分 stop the world,多线程收集
* Serial Old(mark-compact): stop the world,单线程收集，CMS备选方案
* Parallel Old(mark-compact):stop the world,多线程收集   
  
#### 3.2 G1
  将java heap 划分为多个 region。
  整体上看 是 mark-compact
  局部上看(region之间)上看 是 mark-copy
