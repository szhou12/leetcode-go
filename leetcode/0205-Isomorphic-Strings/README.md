# [205. Isomorphic Strings](https://leetcode.com/problems/isomorphic-strings/)

## Solution idea

Use two HashMaps:

`map dict`: map each char in `s` to the char of the same position in `t`. `map dict` is to check that if there is a repeated char in `s`, it must be mapped to the same char that was previously mapped to in `t`.

`map checkDups`: a set that contains all chars in `t` that have been scanned. `map checkDups` is to ensure "No two characters may map to the same character".

Time complexity = $O(n)$

