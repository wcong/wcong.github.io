---
layout: post
title: Build App with Flex
description: "read and write flex in pic operate"
modified: 2014-02-05
tags: [flex,flash,actionscript,as]
---
### flex/flash开发应用,

#### 技术背景
1.  flex 4.6
2.  flash builder 4.6
3.  anction script 3

#### 开发细节

##### 代码布局

``` actionscript 
<?xml version="1.0" encoding="utf-8"?>
<s:Application xmlns:fx="http://ns.adobe.com/mxml/2009"
               xmlns:s="library://ns.adobe.com/flex/spark"
               xmlns:mx="library://ns.adobe.com/flex/mx"
               minWidth="955" minHeight="600" creationComplete="init()">
    <fx:Declarations>
        <!-- 将非可视元素（例如服务、值对象）放在此处 -->
    </fx:Declarations>
    <fx:Script>
        <![CDATA[ 
            <!-- 代码 -->
        ]]>
    </fx:Script>
    <!-- 组件 -->
</s:Application>
```
##### 初始化应用
s:Application 添加 creationComplete 的处理函数

``` actionscript 
<s:Application  .... 
    creationComplete="init()">
```

代码部分

``` actionscript 
public function init():void{
    // 这里写入初始化代码
}
```

##### json 解析

``` actionscript 
var result:Object = JSON.parse(String(event.result)); 
```

##### 拖拽实现(drag and drop) 

1.  绑定  
组件:

``` actionscript
<s:Image source ="@Embed('img/image_text.png')" width="100%" height="25%" name = "image_text" mouseDown="modalDragHandle(event);"/>
```

代码:

``` actionscript
// modal drag start
public function modalDragHandle(e:MouseEvent):void{
    var dragInitiator:Image=Image(e.currentTarget);
    var ds:DragSource = new DragSource();
    ds.addData(dragInitiator, "modal");               
    DragManager.doDrag(dragInitiator, ds, e);
}
```

2.  accept and drop 
组件:

``` actionscript
<s:VGroup id="phone_img_box" width="100%" height="100%"
                         dragDrop="modalDropHandler(event);" dragEnter="modalEnterHandler(event);"
                        horizontalAlign="center" verticalAlign="middle">
</s:VGroup>
```

代码:

``` actionscript 
// modal drag end
public function modalEnterHandler(event:DragEvent):void{
    if (event.dragSource.hasFormat("modal")){   
        //phone_img_box.addElement(Image(event.dragInitiator));
        DragManager.acceptDragDrop(phone_img_box);
    }
}
public function modalDropHandler(event:DragEvent):void{
    var image:Image = Image(event.dragInitiator);
    var name:String = image.name;
    var type:int = Module.judgeModalType(name);
    var border:Module = Module(makeModal(type));
    var index:int = calculateElementIndex(event);
    phone_img_box.addElementAt(border,index + 1);
    border.validateNow();
    desc_img_box.addElementAt(border.makeMiniViewBox(),index + 1);
}
```

##### 文字转图片

``` actionscript
public function drawText(data:BitmapData,text:String,rect:Rectangle):BitmapData{
    var textField:TextField = new TextField();
    textField.antiAliasType = AntiAliasType.ADVANCED;
    var textFormate:TextFormat = new TextFormat();
    textField.defaultTextFormat = textFormate;
    textField.height = rect.height;
    textField.width = rect.width;
    textField.text = text;
    var mat:Matrix = new Matrix();
    mat.translate(rect.x, rect.y);
    data.draw(textField,mat);
    return data;
}
```

##### 压缩图片

``` actionscript 
public static function makeCompressImageData(data:BitmapData,width:int,height:int):BitmapData{
    var matrix:Matrix = new Matrix();
    var widthScale:Number =  width / data.width ;
    var heightScale:Number = height /data.height ;
    matrix.scale(widthScale,heightScale);
    var newBitMap:BitmapData = new BitmapData(width,height,true,0x000000);
    newBitMap.draw(data,matrix,null,null,null,true);
    return newBitMap;
}
```

