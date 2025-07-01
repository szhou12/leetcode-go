# [981. Time Based Key-Value Store](https://leetcode.com/problems/time-based-key-value-store/)

Time complexity of `set()` = $O(M*L)$ where $M$ is the number of calls to `set()` and $L$ is it takes O(L) time to hash the string.

Time complexity of `get()` = $O(N*L*logM)$ where $N$ is the number of calls to `get()`.