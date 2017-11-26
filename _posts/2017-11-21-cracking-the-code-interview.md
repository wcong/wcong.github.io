---
layout: post
title: Cracking The Code Interview 
categories: [interview, book]
description: start to prepare the code interview 
keywords: interview, book
---
### At Beginning
Recently I start to prepare for the foreign interview. So i start to read the book <Cracking The Code Interview>. Use this blog to record what I learned.

### Basic
#### Big O Big omega Big theta
* O (big 0): In academia, big O describes an upper bound on the time. An algorithm that prints all the values in an array could be described as O(N), but it could also be described as O(N2), O(N3), or 0(2N) (or many other big O times). The algorithm is at least as fast as each of these; therefore they are upper bounds on the runtime. This is similar to a less-than-or-equal-to relationship. Likewise, a simple algorithm to print the values in an array is O(N) as well as O(N3 ) or any runtime bigger than O(N).
* Ω (big omega): In academia, Ω is the equivalent concept but for lower bound. Printing the values in an array is Ω(N) as well as Ω(log N) and 0(1). After all,you know that it won't be faster than those runtimes.
* θ (big theta): In academia, e means both O and Ω. That is, an algorithm is θ(N) if it is both O(N) and Ω(N). θ gives a tight bound on runtime.
#### Best Case, Worst Case, Expected Case
Quick sort picks a random element as a "pivot" and then swaps values in the array such that the elements less than pivot appear before elements greater than pivot
* Best Case: If all elements are equal, then quick sort will, on average, just traverse through the array once. This is O(N). (This actually depends slightly on the implementation of quick sort. There are implementations, though, that will run very quickly on a sorted array.)
* Worst Case: What if we get really unlucky and the pivot is repeatedly the biggest element in the array? (Actually, this can easily happen. If the pivot is chosen to be the first element in the subarray and the array is sorted in reverse order, we'll have this situation.) In this case, our recursion doesn't divide the array in half and recurse on each half. It just shrinks the subarray by one element. This will degenerate to anO(N2) runtime.
* Expected Case: Usually, though, these wonderful or terrible situations won't happen. Sure, sometimes the pivot will be very low or very high, but it won't happen over and over again. We can expect a runtime ofO(N log N).
#### Space Complexity
Space complexity is a parallel concept to time complexity. If we need to create an array of size n, this will require 0(n) space. If we need a two-dimensional array of size nxn, this will requireO(n2) space
#### Others
* Drop the Constants
* DroptheNon-DominantTerms
* AmortizedTime
* Log N Runtimes
* Recursive Runtimes
