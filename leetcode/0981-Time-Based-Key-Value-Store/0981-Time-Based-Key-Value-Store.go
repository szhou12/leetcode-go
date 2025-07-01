package leetcode


type Pair struct {
    value string
    timestamp int
}

type TimeMap struct {
    dict map[string][]Pair // {key1: [[t1, v1], [t2, v2]]}
}


func Constructor() TimeMap {
    dict := make(map[string][]Pair)
    return TimeMap{dict: dict}
}


func (this *TimeMap) Set(key string, value string, timestamp int)  {
    if _, ok := this.dict[key]; !ok {
        this.dict[key] = make([]Pair, 0)
    }
    this.dict[key] = append(this.dict[key], Pair{timestamp: timestamp, value:value})
}


func (this *TimeMap) Get(key string, timestamp int) string {
    if _, ok := this.dict[key]; !ok {
        return ""
    }

    if timestamp < this.dict[key][0].timestamp {
        return ""
    }

    left, right := 0, len(this.dict[key]) - 1

    for left < right {
        mid := right - (right - left) / 2
        if this.dict[key][mid].timestamp <= timestamp {
            left = mid
        } else {
            right = mid - 1
        }
    }

    if this.dict[key][left].timestamp > timestamp {
        return ""
    }
    return this.dict[key][left].value

    return ""


}


/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */