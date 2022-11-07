package leetcode

// DFS
// 题目要求只要能从room 0出发到达所有其他rooms就行
// 注意：graph可能有环，所以用visited排除走回到已经看过的node
func canVisitAllRooms(rooms [][]int) bool {
	visited := make(map[int]bool)

	dfs(rooms, 0, &visited)

	// combo = append(combo, 0)
	// return dfs(rooms, 0, combo, &visited)

	// check if all nodes visited
	for i := 0; i < len(rooms); i++ {
		if _, ok := visited[i]; !ok {
			return false
		}
	}
	return true
}

func dfs(rooms [][]int, index int, visited *map[int]bool) {
	(*visited)[index] = true

	for _, nextRoom := range rooms[index] {
		if !(*visited)[nextRoom] {
			dfs(rooms, nextRoom, visited)
		}
	}
}

// 下面的写法是找 [room 0, room 1, ... room n-1] 可以走遍所有rooms的路径, 不符题意
// func dfs(rooms [][]int, index int, combo []int, visited *map[int]bool) bool {
//     // base case
//     if len(combo) == len(rooms) {
//         return true
//     }

//     (*visited)[index] = true
//     for _, nextRoom := range rooms[index] {
//         if !(*visited)[nextRoom] {
//             combo = append(combo, nextRoom)
//             if dfs(rooms, nextRoom, combo, visited) {
//                 return true
//             }
//             combo = combo[:len(combo) - 1]
//         }

//     }
//     return false
// }
