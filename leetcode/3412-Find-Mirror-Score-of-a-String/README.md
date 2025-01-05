# [3412. Find Mirror Score of a String](https://leetcode.com/problems/find-mirror-score-of-a-string/description/)

## Solution idea
### HashMap + Simulation
1. HashMap: key = "mirror", value = list of indices seen so far.
2. Simulation:
    1. iterate over the string.
    2. at the current letter, if its "mirror" exists in the HashMap ("mirror" exists AND the length of list > 0): use the LAST index of the list to update the result. Then, remove the last index from the list.
    3. otherwise, add the current letter to the HashMap as a valid "mirror" candidate for later use (create a new list if the key does not exist yet).

Time complexity = $O(n)$

## Resource
[[C++, Java] Using map of list - Explained](https://leetcode.com/problems/find-mirror-score-of-a-string/solutions/6232184/c-java-using-map-of-list-explained/)
