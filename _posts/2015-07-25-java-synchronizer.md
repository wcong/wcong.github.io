---
layout: post
title: java synchronizer
tags: java synchronizer 
category: java
---

# java线程模型和同步机制

## 简单线程模型

最简单的线程模型是，每个线程都独立的完成各自的工作，互不干涉，线程的切换由系统的线程调度器自行决定。

* yield() 可以主动让出执行权
* sleep() 会让当前线程block，当block完成后，线程调度器会重新分配线程

## 协作线程模型

当需要多个线程合作做一件事情的时候，事情会比较复杂。

### 简单排队模型

最常见的是排队模型，每个线程排队执行某个操作。

比如注册用户时需要保证手机号唯一,把数据库简化为list的的代码如下

``` java
import java.util.List;
import java.util.LinkedList;
public class SyncTest implements Runnable{

        private static List<String> mobileList = new LinkedList<String>();

        private String mobile;

        public SyncTest(String mobile){
                this.mobile = mobile;
        }

        public void run(){
                addMobile();
        }

        public synchronized void addMobile(){
                if(! mobileList.contains(mobile) ){
                        mobileList.add(mobile);
                }
        }

        public static void main(String[] args){
                SyncTest register1 = new SyncTest("18667026331");
                SyncTest register2 = new SyncTest("18667026331");
                Thread threa1 = new Thread(register1);
                Thread threa2 = new Thread(register2);
                threa1.start();
                threa2.start();
                System.out.println(mobileList);
        }
}
```

### 合作模型

除了基本的线程调度器外，还有线程锁内的调度器。线程锁内的调度器优先于线程的调度器，所以sleep()和yield()都不会影响锁的持有者。
当线程运行到同一个synchronized块时，锁内的调度器就开始接管线程的调度了。

* wait()会告诉调度器，停止当前线程，并一直等待，直到随机的notify()选中它，或者nofityAll()通知到，才会有资格继续运行。
* nifity()会告诉调度器，可以随机唤醒一个wait()的线程
* notifyAll()会告诉调度器，唤醒所有的线程

一个典型的例子是线程合作打印数字，大致要求是，开启两个线程，从0开始每个线程有序的交替的打印三个数字，到30得时候都停止。
大致的结果如下:
线程1：1
线程1：2
线程1：3
线程2：4
线程2：5
线程2：6
.......


``` java
import java.lang.Thread;
import java.lang.Runnable;
public class ThreadTest implements Runnable{

        public static int num =1;

        public static int thread =1;

        private Integer threadNum=thread++;

        public static Object obj = new Object();

        public void run(){
                try{
                synchronized(obj){

                        while(num <=30){
                                for( int i=0;i<3;i++ ){
                                        System.out.println(threadNum+":"+num++);
                                }
                                obj.notify();
                                obj.wait();
                        }
                        obj.notifyAll();
                }
                }catch(InterruptedException e){
                }

        }

        public static void main(String[] argv){
                ThreadTest task1 = new ThreadTest();
                ThreadTest task2 = new ThreadTest();
                Thread thread1= new Thread(task1);
                thread1.start();
                Thread thread2 = new Thread(task2);
                thread2.start();
        }

}
```

### 更加复杂的合作模型

从上面的例子可以看出，线程之间只会有一个等待条件，有时候会有多个等待调价，比如上面的例子扩展一下，开启3个线程，每个打印3个数字，到30结束,因为nofity()会随机选出一个线程，所以3个线程就不能按顺序启动。

这时候就会用到 Lock,lock可以产生多个condition，每个condition都会生成一个所内的等待队列，等待被唤醒。
对应的condition会有 await(),signal()和signalAll()方法

``` java
import java.lang.Thread;
import java.lang.Runnable;
import java.util.concurrent.locks.ReentrantLock;
import  java.util.concurrent.locks.Condition;
public class LockTest implements Runnable{

	public static int num =1;

	public static int thread =1;

	private Integer threadNum=thread++;

	private static ReentrantLock lock = new ReentrantLock();

	public static Condition condition1 = lock.newCondition();

	public static Condition condition2 = lock.newCondition();

	public static Condition condition3 = lock.newCondition();

	private Condition myCondition;

	private Condition nextCondition;

	public LockTest( Condition myCondition,Condition nextCondition ){
		this.myCondition = myCondition;
		this.nextCondition = nextCondition;
	}

	public void run(){
		lock.lock();
		try{
			while(num <=30){
				for( int i=0;i<3;i++ ){
					System.out.println(threadNum+":"+num++);
				}
				nextCondition.signal();
				myCondition.await();
			}
			nextCondition.signal();
		}catch(InterruptedException e){
		}finally{
			lock.unlock();
		}
	}
	
	public static void main(String[] argv){
		LockTest task1 = new LockTest(condition1,condition2);
		LockTest task2 = new LockTest(condition2,condition3);
		LockTest task3 = new LockTest(condition3,condition1);
		Thread thread1= new Thread(task1);
		thread1.start();
		Thread thread2 = new Thread(task2);
		thread2.start();
		Thread thread3 = new Thread(task3);
		thread3.start();
	}

}
```
