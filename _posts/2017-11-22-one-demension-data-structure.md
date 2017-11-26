---
layout: post
title: One Dimension Data Structure 
categories: [interview, book]
description: algorithm about one dimension data structure
keywords: List, Stack, Queue, Hashtable, String
---
### At First
those are questions about one dimension structure such as arrat and list and transformation from that.

### Questions
1. Is Unique: Implement an algorithm to determine if a string has all unique characters. What if you cannot use additional data structures?
analysis: 
	1. string is a char array
	2. english characters are just 26, 52 including capital  
complexity
	1. time:O(n)
	2. space:O(1)
```
public boolean isUnique(String string){
	if( string == null || string.length == 0 ){
		return true;
	}
	int[] charTable = new int[52];
	for( char soloChar:string.toCharArray() ){
		int index = (int)soloChar - (int)'a';
		charTable[index]+=1;
		if( charTable[index] > 1 ){
			return false;
		}
	}
	return true;
}
```
2. Check Permutation: Given two strings,write a method to decide if one is a permutation of the other.
analysis:
	1.just the same 
3. URLify: Write a method to replace all spaces in a string with '%20  You may assume that the string has suf cient space at the end to hold the additional characters,and that you are given the "true" length of the string. (Note: If implementing in Java,please use a character array so that you can perform this operation in place.)
4. Palindrome Permutation: Given a string, write a function to check if it is a permutation of a palindrome. A palindrome is a word or phrase that is the same forwards and backwards. A permutation is a rearrangement of letters. The palindrome does not need to be limited to just dictionary words.
5. One Away: There are three types of edits that can be performed on strings: insert a character, remove a character, or replace a character. Given two strings, write a function to check if they are one edit (or zero edits) away.
	analysis: 1.recursive 
```
public boolean isOneAway(String a,String b){
	int distance = a.length()-b.length();
	distance = distance<0?-distance:distance;		
	if( distance>1 ){
		return false;	
	}
}
private boolean isOneAwayRecursive(String a,String b,int aIndex,int bIndex){
	for( ;aIndex<a.length();aIndex++ ){
		if( aIndex>=b.length() ){
			return false;	
		}
		if( a.charAt(aIndex) != b.charAt(aIndex) ){
			return false;
		}
	}
	return aIndex == b.length();
}
```
6. String Compression: Implement a method to perform basic string compression using the counts of repeated characters. For example, the string aabcccccaaa would become a2blc5a3. If the "compressed" string would not become smaller than the original string, your method should return the original string. You can assume the string has only uppercase and lowercase letters (a - z).
7. Rotate Matrix: Given an image represented by an NxN matrix, where each pixel in the image is 4 bytes, write a method to rotate the image by 90 degrees. Can you do this in place?
8. Zero Matrix: Write an algorithm such that if an element in an MxN matrix is 0, its entire row and column are set to 0.
9. String Rotation:Assumeyou have a method isSubstringwhich checks if one word is a substring of another. Given two strings, sl and s2, write code to check if s2 is a rotation of sl using only one call to isSubstring (e.g.,"waterbottle" is a rotation of"erbottlewat").
10. Write code to remove duplicates from an unsorted linked list. How would you solve this problem if a temporary buffer is not allowed?
11. Return Kth to Last: Implement an algorithm to find the kth to last element of a singly linked list.
12. Delete Middle Node: Implement an algorithm to delete a node in the middle (i.e., any node but the first and last node, not necessarily the exact middle) of a singly linked list, given only access to that node.
13. Partition: Write code to partition a linked list around a value x, such that all nodes less than x come before all nodes greater than or equal to x. If x is contained within the list, the values of x only need to be after the elements less than x (see below). The partition element x can appear anywhere in the "right partition"; it does not need to appear between the left and right partitions.
14. Sum Lists: You have two numbers represented by a linked list, where each node contains a single digit.The digits are stored in reverse order, such that the 1's digit is at the head of the list. Write a function that adds the two numbers and returns the sum as a linked list.
15. Palindrome: Implement a function to check if a linked list is a palindrome.
16. Intersection: Given two (singly) linked lists, determine if the two lists intersect. Return the intersecting node. Note that the intersection is defined based on reference, not value.That is, if the kth node of the first linked list is the exact same node (by reference) as the jth node of the second linked list, then they are intersecting.
17. Loop Detection: Given a circular linked list, implement an algorithm that returns the node at the
beginning of the loop.
18. Three in One: Describe how you could use a single array to implement three stacks.
19. Stack Min: How would you design a stack which, in addition to push and pop, has a function min which returns the minimum element? Push, pop and min should all operate in 0(1) time.
20. Stack of Plates: Imagine a (literal) stack of plates. If the stack gets too high, it might topple.Therefore, in real life, we would likely start a new stack when the previous stack exceeds some threshold. Implement a data structure SetOfStacks that mimics this. SetOfStacks should be composed of several stacks and should create a new stack once the previous one exceeds capacity. SetOfStacks.push() and SetOfStacks.pop() should behave identically to a single stack (that is, pop() should return the same values as it would if there were just a single stack).
21. Queue via Stacks: Implement a MyQueue class which implements a queue using two stacks.
22. Sort Stack: Write a program to sort a stack such that the smallest items are on the top. You can use an additional temporary stack, but you may not copy the elements into any other data structure (such as an array). The stack supports the following operations: push, pop, peek, and isEmpty.
23. Animal Shelter: An animal shelter, which holds only dogs and cats, operates on a strictly"first in, first out" basis. People must adopt either the"oldest" (based on arrival time) of all animals at the shelter, or they can select whether they would prefer a dog or a cat (and will receive the oldest animal of that type). They cannot select which specific animal they would like. Create the data structures to maintain this system and implement operations such as enqueue, dequeueAny, dequeueDog, and dequeueCat. You may use the built-in Linked list data structure.
