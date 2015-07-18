---
layout: post
title: java design pattern
description: java design pattern
keywords: java
---

### Java设计模式

  没有完整的学习过设计模式，正好码代码也有一段时间了，是时间学习，整理一下了。

  Java 的设计模式主要可以分为三种

  1. 创造型模式 
  2. 结构型模式
  3. 行为型模式

### 创造型模式

  创造型模式是关于，对象的创建。

  创建型模式主要有4种类型

1. 工厂模式:对象的创建有另一个类管理
	* 普通工厂模式，就是建立一个工厂类，对实现了同一接口的一些类进行实例的创建
	* 抽象工厂模式（即工厂也有多个，工厂也需要继承接口），而且不再使用字符串来判断实例化的对象，而是在方法中直接实例化。
2. 单例模式:一个类只有唯一一个对象
3. 建造者模式（builder）:针对对象的属性来说，为对象的多个属性创建实例  
4. 原型模式:将一个对象作为原型，对其进行复制、克隆，产生一个和原对象类似的新对象

### 结构模式
   结构模式主要是修改类的结构：主要分为7个模式，4种类型
#### 结构模式的类型
1. 在原有类的基础上修改，继承类，或继承接口：适配器模式，装饰器模式，代理模式
2. 不管原来的类，直接持有并使用实例：外观模式，bridge模式
3. 类自己组合，形成特定的数据结构：组合模式
4. 对象池：享元模式

#### 所有结构模式

1. 适配器模式：增加方法
	1. 增加类的方法：创建一个新类，继承原有的类，实现新的接口
	2. 增加对象的方法：可以创建一个类，持有原类的一个实例，在着这个类的方法中，调用实例的方法就行。
2. 装饰模式：修改原有方法的逻辑：新建一个类，持有原类的实例，继承原类的接口，实现接口的方法，在方法中添加自己的逻辑，并调用原方法
3. 代理模式：修改原有方法的逻辑：跟装饰模式的区别是，类的实例是自己创建，装饰模式的原类的实例是传入的。
4. 外观模式：不需要管原方法，持有原类的实例，按照需要的逻辑动态使用实例的方法。
5. bridge模式：跟外观模式的区别是，原类的类型会变，所以持有的是实例的接口。所以会有一组类实现相同的接口，使用bridge模式是，选一个合适的类实例化，并转化为接口传入bridge
6. 组合模式：处理类似树形结构，原类的属性中包含原类。
7. 享元模式（flyweight）：即对象池。

### 行为模式

   行为模式主要描述类和类之间的关系，主要有11个模式。主要分为3个类型
#### 行为模式类型

1. 类之间有相同点，可以替换的：策略模式，模板模式
2. 两个类之间有逻辑关系：观察者模式，迭代子模式，责任链模式，命令模式，访问者模式，中介者模式，解释器模式
3. 类的状态：备忘录模式，状态模式

#### 所有行为模式

1. 策略模式：设计一个接口，多个实现类实现该接口：主要适用于不同算法的无缝切换
2. 模板模式：跟策略模式的区别是这些算法有统一的入口和可以重用的模块你，所以用抽象类替换接口。
3. 观察者模式：类似订阅的方式，把对象A委托给对象B，对象B会在自己发生变化时，相应的调用对象A的方法。
4. 责任链模式：就是多个类的观察者模式，会有委托链路。
5. 命令模式：类A想调用类B的方法，把调用的逻辑写在类C里面。
6. 访问者模式：类A想调用类B的方法，就让类B开了个借口，接受A的实例，并触发A的调用方法。调用方法抽象成一个接口，就会有更好的扩展性。
7. 中介者模式：类A跟B不再相互传递自己了，用一个类C，都持有A和B的实例，让类C实现调用逻辑。
8. 解释器模式：类A要根据类B的内容实现逻辑，类A用了模板模式.
9. 迭代子模式：迭代器。
10. 备忘录模式：类A想要保存自己的一些属性，新建类B结构中设计这些属性。类A持有类B，用于保存数据
11. 状态模式：对象A持有对象b，同一个A的方法，b的属性不同，方法的逻辑也不同。