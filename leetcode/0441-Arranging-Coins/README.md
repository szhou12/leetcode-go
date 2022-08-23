# [441. Arranging Coins](https://leetcode.com/problems/arranging-coins/)

## Solution idea

### 模拟
每次增加1，并看是否有剩余

### 等差数列求和公式

$S = \frac{n(n+1)}{2} \Rightarrow n = \lfloor \sqrt{2S+1/4} - 1/2 \rfloor$