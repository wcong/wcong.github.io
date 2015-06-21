---
layout: post
title: java class 文件存储结构 
tags: java class 
category: java
---

### java class 文件存储结构 -- openjdk7

#### 基础

  openjdk的源码可以在[openjdk官网](http://www.planetjdk.org/)下载到。

  java class 的编译器javac是用java写的，可以在源码的langtools/src/share/class 下查看java 源代码

  com.sun.tools.javac包是主要的javac的实现逻辑 

  com.sun.tools.javac.jvm.ClassWriter的方法 writeClass(ClassSymbol c) 是写入class文件内容的主要逻辑

#### class 文件的结构

|序号|含义|字节数|
|----|----|------|
|1|JAVA_MAGIC,固定值,表示java文件|4|
|2|minorVersion,最小java版本|1|
|3|majorVersion,主要版本号|1|
|4|pool,类文件中用到的所有项目(类，类名,常量等),给后面的索引提供字典|变长|
|5|writeInnerClasses,内部类，递归|变长|
|6|类flag|1|
|7|类poll的索引|1|
|8|父类poll的索引，没有的话是0|1|
|9|interface,数量|1|
|10|每个interface在pool的索引|1*interface.length()|
|11|field数量|1|
|12|field描述|变长|
|13|method数量|1|
|14|method描述|变长|
|15|writeFlagAttrs|变长|
|16|writeJavaAnnotations|变长|
|17|writeEnclosingMethodAttribute|变长|

#### java flag
  com.sun.tools.javac.code.Flags 定义了 java 的 flag 包括:Access flags and other modifiers for Java classes and members

  使用 int 和long 的每一bit(0或1)代表是否有相应的flag

  比如 java 中支持泛型和重载的 SYNTHETIC 和 BRIDGE 标签

``` java
public static final int SYNTHETIC    = 1<<12;
public static final long BRIDGE          = 1L<<31;
```

#### write pool
  write  Write constant pool to pool buffer
  把常量池中的常量都写入 poolbuf 中 

#### write method

1. wrire flag
2. write code
3. write exception
4. write AnnotationDefault
5. writeMemberAttrs
6. writeParameterAttrs

##### write method code
  写入方法代码逻辑

1. max_stack
2. max_locals
3. current pointer
4. catchInfo
5. lineInfo
6. CharacterRangeTable
7. varBufferSize
8. LocalVariableTypeTable
9. stackMapBufferSize
