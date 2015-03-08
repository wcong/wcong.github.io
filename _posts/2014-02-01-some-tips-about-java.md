---
layout: post
title: Some Tips About Java
description: "some little but import thing about java"
modified: 2014-02-01
tags: [java]

---
## here are some tips about java

### 1 == and  equals between Long,Integer,Byte

####  1.1 Byte between 0 and 128
== could work,but equals do not work,and I do not known the detail of it

#### test code


    public class ByteEqure {
        public static void main(String[] args){
            Byte a =10;
            if( a.equals(10) ){
                System.out.println("not print");
            }else{
                System.out.println("print");
            }
            if( a == 10 ){
                System.out.println("print");
            }else{
                System.out.println("not print");
            }
        }
    }
### 1.2 Long ,Integer between 0 and 128
both == and equals could work

#### text code

    public class TestLong {
        public static void main(String[] args) {
            // TODO Auto-generated method stub
            Long a= 10l;
            Long b = 10l;
            System.out.println(a==b);
        }
    }
### 2 java split 
this is a example of split string by all the blank;

    String[] strs = str.split("\\s+”); 

it is the same function of follow code

    String[] strs = str.split(“ +”);  

you can split the other char by relpace chars before '+';
