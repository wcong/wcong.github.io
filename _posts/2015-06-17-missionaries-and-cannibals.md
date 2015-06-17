---
layout: post
title: 食人魔过河 
tags: java algorithm
category: java
---

### 问题描述

食人魔过河是经典的的状态空间问题。描述如下:

	有N个传教士和N个野人要过河，现在有一条船只能承载K个人（包括野人），K<N，在任何时刻，如果有野人和传教士在一起，必须要求传教士的人数多于或等于野人的人数。 设M为传教士的人数，C为野人的人数。

这里面有一些隐藏的条件：

1. 船是有来有回的，所以要考虑取的人数，和回来的人数。
2. 在岸坐边，岸右边，船上，必须要满足,(M==0 && (M+C) >0) || (M>=C)

### 状态空间

因为船的来回状态，所以用程序逻辑来推演结果是很复杂的。

所以要程序遍历所有的条件，所有的组合，在后在推断可能的解。

### 程序逻辑

1. 生成所有符合条件的状态,比如N=3时，岸左边(M=3,C=1),船上(M=0,C=2),对应的船到岸时右边(M=0,C=2)
2. 推演所有状态的关联关系，比如共有A,B,C,D四个状态,推演出B的上一步可能是A，C，下一个是D。以此类推
3. 从根状态出发，得出完整的路径，比如N=3时，根状态是岸左边(M=3,C=3)，船上(M=0,C=0)

### code

```java
import java.util.*;

/**
 * Created by wcong on 15/5/17.
 */
public class MissionariesAndCannibals {

    private static class Status {

        List<Status> next = new LinkedList<Status>();

        Integer missionaries;

        Integer cannibals;

        boolean isLeft;

        Status(boolean isLight) {
            this.isLeft = isLight;
        }

        public void num(Integer missionaries, Integer cannibals) {
            this.missionaries = missionaries;
            this.cannibals = cannibals;
        }

        public boolean isEnd(int n) {
            return !isLeft && missionaries == n && cannibals == n;
        }

        @Override
        public String toString() {
            return isLeft + ":" + missionaries + ":" + cannibals;
        }
    }


    Integer n;
    Integer k;
    List<Status> leftList = new LinkedList<Status>();
    List<Status> rightList = new LinkedList<Status>();

    public MissionariesAndCannibals(Integer n, Integer k) {
        this.n = n;
        this.k = k;
    }

    public void link(Status status, Set<Status> set) {
        if (status.isEnd(n)) {
            return;
        }
        List<Status> list = status.isLeft ? rightList : leftList;
        set.add(status);
        for (Status nextStatus : list) {
            if (set.contains(nextStatus)) {
                continue;
            }
            int goCannibals = status.cannibals + nextStatus.cannibals - n;
            int goMissionaries = status.missionaries + nextStatus.missionaries - n;
            int goNum = goCannibals + goMissionaries;
            int leftCannibals = status.cannibals - goCannibals;
            int leftMissionaries = status.missionaries - goMissionaries;
            if (goCannibals >= 0 &&
                    goMissionaries >= 0 &&
                    goNum > 0 &&
                    goNum <= k &&
                    !(leftMissionaries > 0 && leftMissionaries < leftCannibals) &&
                    !(goMissionaries > 0 && goMissionaries < goCannibals)) {
                if (!status.next.contains(nextStatus)) {
                    status.next.add(nextStatus);
                }
                link(nextStatus, new HashSet<Status>(set));
            }
        }
    }

    public void makeStatusList() {
        for (int i = 0; i <= n; i++) {
            for (int j = 0; j <= n; j++) {
                if (i > 0 && j > i) {
                    break;
                }
                Status leftStatus = new Status(true);
                leftList.add(leftStatus);
                leftStatus.num(i, j);
                Status rightStatus = new Status(false);
                rightList.add(rightStatus);
                rightStatus.num(i, j);
            }
        }
    }

    public boolean clean(Status status, Set<Status> lastSet) {
        if (status.isEnd(n)) {
            return false;
        }
        if (status.next.isEmpty()) {
            return true;
        }
        boolean isClean = true;
        lastSet.add(status);
        Iterator<Status> statusIterator = status.next.iterator();
        while (statusIterator.hasNext()) {
            Status nextStatus = statusIterator.next();
            if (lastSet.contains(nextStatus)) {
                continue;
            }
            Set<Status> nextSet = new HashSet<Status>(lastSet);
            if (!clean(nextStatus, nextSet)) {
                isClean = false;
            } else {
                statusIterator.remove();
            }
        }
        return isClean;
    }

    public void calculate() {
        makeStatusList();
        Status firstStatus = leftList.get(leftList.size() - 1);
        link(firstStatus, new HashSet<Status>());
        clean(firstStatus, new HashSet<Status>());
    }

    public void printOne() throws InterruptedException {
        Status status = leftList.get(leftList.size() - 1);
        do {
            System.out.println(status);
            Thread.sleep(1000);
            status = status.next.get(0);
        } while (!status.isEnd(n));
    }

    public static void main(String[] args) throws InterruptedException {
        int n = Integer.parseInt(args[0]);
        int k = Integer.parseInt(args[1]);
        MissionariesAndCannibals test = new MissionariesAndCannibals(n, k);
        test.calculate();
        test.printOne();
    }

}

```
