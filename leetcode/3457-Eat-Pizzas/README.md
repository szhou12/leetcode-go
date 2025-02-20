# [3457. Eat Pizzas!](https://leetcode.com/problems/eat-pizzas/description/)

## Solution idea
### Greedy + Math
We have 4n pizzas, so we can definitely form 4 pizzas (even randomly selected) to eat every day. On odd days, we pick the heaviest pizzas, so we can preserve m heaviest pizzas for total of m odd days. On even days, we pick the second heaviest pizzas, so 2-pizza as a group, we pick the ligher pizza from the group per even day. Note that we cannot reuse the heavier pizza.

Time complexity = $O(n\log n)$

## Resource
[Java || C++ || Python || One pass 🚀|| Two pass || Picture Explanation ||](https://leetcode.com/problems/eat-pizzas/solutions/6427675/java-c-python-one-pass-two-pass-picture-explanation/)