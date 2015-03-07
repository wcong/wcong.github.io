---
layout: post
title: Index MySql Width Solr 
description: "index mysql width solr"
modified: 2014-02-13
tags: [java,solr,mysql,index,python]
---
#### 官方教程 

1.  官方地址:https://lucene.apache.org/solr/

2.  官方教程地址:https://lucene.apache.org/solr/4_6_1/tutorial.html

#### 具体操作

官方教程详细讲解了solr的基本使用,这里就不多说了,下面主要是具体的使用

##### 业务背景

我有4张需要做全文索引的表(keyword_hot,cat_keyword,top_20,top_50)。

##### 创建索引

这4张表没有关联,所以需要4个core,solr 原有一个core collection1 所以需要手工新建3个 core 

```bash
cp  -R collection1 keyword_hot
cp  -R collection1 top_20
cp  -R collection1 top_50
```

##### 配置索引

在每个 core 内部需要配置mysql 的数据源,开启 mysql 的数据源 需要配置3个地方

1.  在 solrconfig.xml 添加如下配置

```xml
<requestHandler name="/dataimport" class="org.apache.solr.handler.dataimport.DataImportHandler">
      <lst name="defaults">
        <str name="config">data-config.xml</str>
    </lst>
</requestHandler>
```

2.  创建并填写 data-config.xml    dataSource 的配置项batchSize 是分页读取所有数据

```xml
<dataConfig>
    <dataSource type="JdbcDataSource" driver="com.mysql.jdbc.Driver" url="jdbc:mysql://localhost/p4papp" user="root" password="*******" batchSize = "-1"/>
    <document name="keyword_hot">
        <entity name="keyword_hot" query="select id,keyword,cid from keyword_hot">
            <field column="id" name="id" />
            <field column="cid" name="cid" />
            <field name="keyword" column="keyword" />
        </entity>
    </document>
</dataConfig>
```
3.  在schema.xml 配置上面的 字段的属性 用到了中文分词的solr扩展IKA,在schema.xml 里面有许多的字段,需要详细的看文档，不要轻易删除

``` xml
    <!-- cat_keyword的字段  -->
    <field name="cid" type="int" indexed="true" stored="true" />
    <field name="keyword" type="text_ik" indexed="true" stored="true"/>
    <!-- cat_keyword的字段  -->
    <!--  配置IK分词器 -->
    <fieldType name="text_ik" class="solr.TextField">
        <analyzer class="org.wltea.analyzer.lucene.IKAnalyzer"/>
    </fieldType>
```

##### 更新索引

最后需要有一个程序每天更新索引,记录一下每个索引更新时间,sleep 一定的时间

```python
#!/usr/bin/env python
#encoding=utf8
import urllib2
import urllib
import time
import datetime
import sys
reload(sys)
sys.setdefaultencoding('utf8')
def log(log):
    date = datetime.datetime.now().strftime('%Y-%m-%d %H-%M-%S')
    print date + ':' + log
indexMap = {'keyword_hot':600,'cat_keyword':600,'top_20':60,'top_50':120}
log('开始更新solr索引')
for index in indexMap.keys():
    log('开始更新:' + index + '的索引')
    url = 'http://localhost:30003/solr/'+index+'/dataimport?command=full-import&clean=true&commit=true'
    response = urllib2.urlopen(url);
    html = response.read()
    log('更新:' + index + '的索引:sleep:' + str(indexMap[index])+ '秒')
    time.sleep(indexMap[index])
```
