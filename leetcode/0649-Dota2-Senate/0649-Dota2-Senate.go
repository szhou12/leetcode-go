package leetcode

// Two queues 模拟
func predictPartyVictory(senate string) string {
	n := len(senate)
	senateArr := []byte(senate)

	rQueue := make([]int, 0) // 存入 R 的index
	dQueue := make([]int, 0) // 存入 D 的index
	for idx, senate := range senateArr {
		if senate == 'R' {
			rQueue = append(rQueue, idx)
		} else {
			dQueue = append(dQueue, idx)
		}
	}

	for len(rQueue) > 0 && len(dQueue) > 0 {
		rIdx := rQueue[0]
		rQueue = rQueue[1:]
		dIdx := dQueue[0]
		dQueue = dQueue[1:]

		// index小的存活, append到后面
		if rIdx < dIdx {
			rQueue = append(rQueue, rIdx+n)
		} else {
			dQueue = append(dQueue, dIdx+n)
		}

	}

	if len(rQueue) > 0 {
		return "Radiant"
	} else {
		return "Dire"
	}
}

// Greedy
func predictPartyVictory_greedy(senate string) string {
	// R = true表示本轮循环结束后，依然有R存活。D同理
	R, D := true, true
	// 排在当前senate之前对手议员的个数
	flag := 0

	senateArr := []byte(senate)
	for R && D { //一旦R或者D为false，就结束循环，说明本轮结束后只剩下R或者D了
		R = false
		D = false
		for i := 0; i < len(senateArr); i++ {
			if senateArr[i] == 'R' {
				if flag < 0 { // 敌方D议员数 >= 1, R被消灭, R此时为false
					senateArr[i] = 0
				} else {
					R = true // 如果没被消灭，本轮循环结束有R存活
				}
				flag++
			}
			if senateArr[i] == 'D' {
				if flag > 0 { // 敌方R议员数 >= 1, R被消灭, D此时为false
					senateArr[i] = 0
				} else {
					D = true // 如果没被消灭，本轮循环结束有R存活
				}
				flag--
			}
		}

	}

	// 循环结束之后，R和D只能有一个为true
	if R {
		return "Radiant"
	} else {
		return "Dire"
	}

}
