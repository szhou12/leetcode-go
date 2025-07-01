package leetcode

/*
Example Walkthrough:
logs = [[0,5],[1,2],[0,2],[0,5],[1,3]], k = 5

map{id: {times}}
{0: {5, 2}, 1: {2, 3}}

map{unique_times: count}
{2: 1+1}

answers
 1 2 3 4 5
 0 1 2 3 4
[0,2,0,0,0]
*/

func findingUsersActiveMinutes(logs [][]int, k int) []int {
    idTimes := make(map[int]map[int]bool)

    for _, log := range logs {
        id, t := log[0], log[1]
        if _, ok := idTimes[id]; !ok {
            idTimes[id] = make(map[int]bool)
        }
        idTimes[id][t] = true
    }

    

    timesCnt := make(map[int]int)
    for _, times := range idTimes {
        timesCnt[len(times)]++
    }

    

    res := make([]int, k)
    for i := 0; i < k; i++ {
        if cnt, ok := timesCnt[i+1]; ok {
            res[i] = cnt
        }
    }

    return res
}