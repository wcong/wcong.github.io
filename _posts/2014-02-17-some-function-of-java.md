---
layout: page
title: Some Function Of Java 
description: "Some Function Of Java( array_chunk,implode,and string partition )"
modified: 2014-02-17
tags: [java]
---

#### partition string

```java
// special char set
public static Set<String> specialCharSet = new HashSet<String>();

static{
    // special char string which to be splited into set 
    String regEx = "[`~!@#$%^&*()+=|{}':;',\\<>/?~！@#￥%……&*（）——+|{}【】‘；：”“’。，、？]-";
    for( int i =0;i< regEx.length(); i ++ ){
        specialCharSet.add(regEx.substring(i,i+1));
    }
}

/**
 * 规范字符串
 * 1 特殊字符 转化为空格
 * 2 字符转换为小写
 * 3 英文与汉字隔离 数字与汉字隔离
 * @param title
 * @return
 */
public static String tplString(String title){
    StringBuilder sb = new StringBuilder(title);
    for( int i =0 ; i < sb.length();i++ ){
        if( specialCharSet.contains(sb.substring(i, i+1)) ){
            sb.setCharAt(i, ' ');
            continue;
        }
        char temp = sb.charAt(i);
        if( temp >='A' && temp <='Z'){
            temp = (char)(((int)temp) + 32);
        }
        sb.setCharAt(i, temp);
        if( (temp >='a' && temp <='z') || (temp >='0' && temp<='9') ){
            if( i > 0 ){
                char lastChar = sb.charAt(i-1);
                if(lastChar>=19968 && lastChar <=171941){
                    sb.insert(i, ' ');
                    i+=1;
                }
            }
            if( i < (sb.length()-1) ){
                char nextChat = sb.charAt(i+1);
                if( nextChat>=19968 && nextChat <=171941){
                    sb.insert(i+1, ' ');
                    i+=1;
                }
            }
        }
    }
    return sb.toString();
}
```

#### array_chunk

```java
/**
 * 把list分组
 * @param <T>
 * @param list
 * @param num
 * @return
 */
public static <T> List<List<T>> chunkListByNum(List<T> list,Integer num){
    List<List<T>> chunkedList = new LinkedList<List<T>>();
    if( num==0 || list.isEmpty()){
        return chunkedList;
    }
    Integer chunkNum = list.size() / num;
    if (( list.size() % num )!= 0 ) {
        chunkNum += 1;
    }
    Integer startIndex = 0;
    for (Integer i = 0; i < chunkNum - 1; i++) {
        chunkedList.add(list.subList(startIndex, startIndex + num));
        startIndex += num;
    }
    chunkedList.add(list.subList(startIndex, list.size()));
    return chunkedList;
}
```
#### implode
```java
/**
 * 把list里面的内容以 join 连接
 * 如：0^^1^^2^^3^^4^^5^^6^^7^^8^^9^^10
 * @param <T>
 * @param list
 * @param join
 * @return
 */
public static <T> String implode(List<T> list, String join){
    StringBuilder sb = new StringBuilder();
    if(list.isEmpty()){
        return sb.toString();
    }
    for( T t: list){
        sb.append(t);
        sb.append(join);
    }
    sb.delete(sb.length()-join.length(), sb.length());
    return sb.toString();
}
```
