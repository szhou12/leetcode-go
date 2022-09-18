# [409. Longest Palindrome](https://leetcode.com/problems/longest-palindrome/)

## Solution idea

Loop through every character, use a `hashset` to check if we can find its pair.

If so, add $2$ to length of palindrome.

At the end, check if we can add one more character at the center of padlindrome.

Time complexity = $O(n)$