# [1432. Max Difference You Can Get From Changing an Integer](https://leetcode.com/problems/max-difference-you-can-get-from-changing-an-integer/description/)

## Solution idea
### Brute Force
1. x try out all options (0-9), y try out all options (0-9)

Time complexity = $O(10^2 * \log_{10} num)$ where `num` is the input number. We use a double loop to enumerate all possible replacement methods, which takes $O(10^2)$ time. For each replacement method, we convert num to a string and perform the replacement operation. The time required for this is proportional to the number of digits in num, which is $O(\log_{10}(num))$.

Space complexity = $O(\log_{10} num)$. The algorithm stores string representations of the input number, which have length proportional to the number of digits.


