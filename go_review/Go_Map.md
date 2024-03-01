# [Map](https://github.com/szhou12/leetcode-go/blob/main/go_review/README.md)

## Contents
* [Contains a Key](#contains-a-key)
* [Delete a Key](#delete-a-key-from-map)
* [Return Value From Non-Existent Key](#return-value-from-non-existent-key)
* [Slice vs Array. Which Can Be Used as Key?](#slice-vs-array-which-can-be-used-as-key)


## Contains a Key
[How to check if a map contains a key in Go?](https://stackoverflow.com/questions/2050391/how-to-check-if-a-map-contains-a-key-in-go)
```go
if val, ok := myMap["foo"]; ok {
    //If the key exists, do something here
}
```

## Delete a Key From Map
[How to delete a key from a map in Golang?](https://www.tutorialspoint.com/how-to-delete-a-key-from-a-map-in-golang)
```go
delete(map, key)
```

## Return Value From Non-Existent Key
[Access map element (that does not exist)](https://code-maven.com/slides/golang/access-map-element)

[Effective Go - Maps section](https://go.dev/doc/effective_go#maps)
* 个人总结：
    * Go中，如果访问一个`map`中不存在的key，将会以‘0假空’原则返回value，而不会报错。如果不提前进行`containsKey()`的检查，这个特性可能会引起bug。需要特别注意！
    * An attempt to fetch a map value with a key that is not present in the map will return the zero value for the type of the entries in the map.
```go
father := make(map[int]int)
fmt.Println(father) // map[]
fmt.Println(father[12])  // 0

father := make(map[int]bool)
fmt.Println(father) // map[]
fmt.Println(father[12]) // false

father := make(map[int][]int)
fmt.Println(father) // map[]
fmt.Println(father[12]) // []
```

## Slice vs Array. Which Can Be Used as Key?
[Slice as a key in map](https://stackoverflow.com/questions/20297503/slice-as-a-key-in-map)
* Slice CANNOT be map key!
* Array CAN be map key!
```go
s := make(map[[2]int]bool) // [2]int is array. Can be map key.

t := make(map[[]int]bool) // []int is slice. CANNOT be map key.
```