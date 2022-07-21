# [13. Roman to Integer](https://leetcode.com/problems/roman-to-integer/)

## Solution idea

### My Solution - Two Pointers/Sliding Window of size 2
Case 1: If current pointed two roman letters can form a number, then add this number to result AND increment both `left` and `right` pointers by 2.

Case 2: O/w, if not, then add only the number that `left` points to AND increment both `left` and `right` pointers by 1.

Time complexity = $O(n)$ where $n$ is the length of input string.

### Another Solution
Look at each single roman letter once at a time from back to front.

If the current letter's value < its next letter's value, they two can form a number. Otherwise, they can't.

Time complexity = $O(n)$ where $n$ is the length of input string.