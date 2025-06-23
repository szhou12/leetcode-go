# [2434. Using a Robot to Print the Lexicographically Smallest String](https://leetcode.com/problems/using-a-robot-to-print-the-lexicographically-smallest-string/description/)

## Solution idea
### Stack - Next Smaller Element
1. We only need to figure out the timing that we need to push into an element into the stack and the timing when we need to pop out an element and append it to the result.
2. stack is `t` described in the problem.

Time complexity = $O(n)$

Space complexity = $O(n)$


## Resource
1. [【每日一题】LeetCode 2434. Using a Robot to Print the Lexicographically Smallest String](https://www.youtube.com/watch?v=zVHCBTPIj9Q&ab_channel=HuifengGuan)
2. [CPP / JAVA / PYTHON ||🫅Greedy ||👍100% beats ||😊easy to understand](https://leetcode.com/problems/using-a-robot-to-print-the-lexicograshically-smallest-string/solutions/6815333/cpp-java-python-greedy-100-beats-easy-to-understand/)
    - See implementation `robotWithString_lc_soln()`