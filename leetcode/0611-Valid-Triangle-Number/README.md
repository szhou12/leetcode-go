# [611. Valid Triangle Number](https://leetcode.com/problems/valid-triangle-number/)

## Solution Idea:

First, I've tried out DFS + pruning but it would ETL.

Then, I've tried out brute force (O(n)) and it still ETL.

The current solution is greedy which sorts the input array first and 

then for each A[i], find all triangles that can be formed with A[i] being one edge.

Use Induction to prove this greedy solution returns the optimal solution.

Proof:

Let $f(a_1, \cdots, a_i) = $ the number of possible traingles found by the Algo. 

Let $f(a_1, \cdots, a_i)* =$ all possible triangles.

WTS $f(a_1, \cdots, a_i) \geq f(a_1, \cdots, a_i)*$

IH: the Algo correctly returns all possible traingles for $a_1, \cdots, a_i$. ie. $f(a_1, \cdots, a_{i-1}) = f(a_1, \cdots, a_{i-1})*$

$f(a_1, \cdots, a_i) = f(a_1, \cdots, a_{i-1}) + $ all triangles with one edge $= a_i$.

Since the Algo sorts the input array in increasing order, all triangles with one edge $= a_i$ is equivalent to 
all triangles with the largest edge $= a_i$.

The Algo checks all other two edges that are less than $a_i$ that can form a triangle.

Hence, the Algo correctly counts the number of all possible triangles with one edge $= a_i$ since otherwise there must be at least one other edge $>a_i$ which is not possible to find in $a_1, \cdots, a_i$.

ie. $f(a_1, \cdots, a_i) = f(a_1, \cdots, a_{i-1}) + $ all triangles with one edge $= a_i \geq f(a_1, \cdots, a_i)*$

Time complexity: $O(n^2)$