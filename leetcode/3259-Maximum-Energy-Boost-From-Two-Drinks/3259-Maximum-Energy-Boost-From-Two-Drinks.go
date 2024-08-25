package leetcode

func maxEnergyBoost(energyDrinkA []int, energyDrinkB []int) int64 {
    n := len(energyDrinkA)
    dp := make([][3]int, n)

    dp[0][0] = 0
    dp[0][1] = energyDrinkA[0]
    dp[0][2] = energyDrinkB[0]

    for i := 1; i < n; i++ {
        dp[i][0] = max(max(dp[i-1][0], dp[i-1][1]), dp[i-1][2])
        dp[i][1] = max(dp[i-1][1] + energyDrinkA[i], dp[i-1][0] + energyDrinkA[i])
        dp[i][2] = max(dp[i-1][2] + energyDrinkB[i], dp[i-1][0] + energyDrinkB[i])
    }

    res:= max(dp[n-1][1], dp[n-1][2])
    return int64(res)

}
// dp[i][0] := max total boost at i-th hour when no consumption at last time consumption
// dp[i][1] := max total boost at i-th hour when the last time consumption is A
// dp[i][2] := max total boost at i-th hour when the last time consumption is B

// dp[i][0] = max(dp[i-1][1], dp[i-1][2])
// dp[i][1] = max(dp[i-1][1] + energyDrinkA[i], dp[i-1][0] + energyDrinkB[i])
// dp[i][2] = max(dp[i-1][2] + energyDrinkB[i], dp[i-1][0] + energyDrinkA[i])