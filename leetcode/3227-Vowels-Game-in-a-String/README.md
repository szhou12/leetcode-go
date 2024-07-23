# [3227. Vowels Game in a String](https://leetcode.com/problems/vowels-game-in-a-string/description/)

## Solution idea
### 找规律 - 打擂台
1. 如果 vowels = 0, 那么Bob铁赢。因为Alice先进行行动，首轮她就无法挑出含奇数个元音字母的substring，无法行动。注意，这里Bob赢是因为Alice无法行动赢的。
2. 如果 vowels = n:
    1. 如果 n 是偶数，那么Alice能赢。因为Alice先行动，她可以挑出含 n-1 个元音字母的substring，使Bob无法挑出偶数个元音字母。
    2. 如果 n 是奇数，那么Alice能赢。因为Alice先行动，她可以直接把含所有元音字母的substring挑出，然后到Bob的回合，因为他必须选择一个不为空的substring并且包含偶数个元音字母，显然他做不到。
* 综上所述，如果字符串中只要至少有一个元音，Alice铁赢。如果一个元音字母都没有，Bob铁赢。

Time complexity = $O(n)$

## Resource
[Find a vowel](https://leetcode.com/problems/vowels-game-in-a-string/solutions/5508944/find-a-vowel/)