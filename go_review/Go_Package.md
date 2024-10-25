# [Package](https://github.com/szhou12/leetcode-go/blob/main/go_review/README.md)

## Contents
* [Package Managment](#package-management)
* [Code Scope (Redeclare error)](#code-scope)

## Package Management

### Package VS. Directory
* 必须要明确这两个概念的区别：
    * **Package (包)**
        * 逻辑上一个或多个go文件的集合，它们都由 **包的声明**`package <pacakgeName>` 所标识，它们一起组成了一个package的实现。注意：这些go文件不一定都放在同一个directory下。
        * e.g. 你可能注意到每道题solution的go文件都是以`package leetcode`开头，这是在告诉程序猿这些go文件都属于名叫 `leetcode` 这个package。如果程序猿想import一个directory下的go文件，import时如果不额外赋予名字，将使用 default alias `leetcode` (参考例子1)。
    * **Directory (目录)**
        * 也即文件夹。物理上存放一个或多个go文件的地方。这是在告诉compiler去哪里找到一个特定的package name。目录的路径是import path，以string形式表示。一个directory下是一个共享的code scope，不可以redeclare。
* By convention, 一个go文件在哪个directory下，就declare 它的 package name 为这个directory的名字。但是这并不是强制的，只是为了方便程序猿管理查找。
* **例子1**: 
    * 不遵守 naming convention 的例子: 这个project内的所有go文件package name都是`leetcode`，但是它们并不在同一个directory下。
    * 假如有一个main文件，想要import *0015-3Sum* 目录下的 `threeSum()`:
    ```go
    import (
        "leetcode-go/leetcode/0015-3Sum" // 标明import path，这里没有赋予package alias，后面调用使用默认的alias "leetcode"
    )

    func main() {
        fmt.Println(leetcode.threeSum([]int{1,2,3})) // 使用 default alias "leetcode" 来调用
    }
    ```
* **例子2**:
    * 假如这个main文件，想要同时import *0015-3Sum* 和 *0018-4Sum* 两个目录下的函数:
    * 此时，因为两个目录下的package name都是`leetcode`，所以import path时需要给定alias。
        * 2 different dir can have the same `package xxx`. But when import the 2 dir in the same go file, need to give them different aliases.
    ```go
    import (
        lc15 "leetcode-go/leetcode/0015-3Sum" // 标明package alias为 lc15
        lc18 "leetcode-go/leetcode/0018-4Sum" // 标明package alias为 lc18
    )

    func main() {
        fmt.Println(lc15.threeSum([]int{1,2,3})) // 使用 alias "lc15" 来调用
        fmt.Println(lc18.fourSum([]int{1,2,3})) // 使用 alias "lc18" 来调用
    }
    ```

### Resources
* [Relationship between a package statement and the directory of a .go file](https://stackoverflow.com/questions/43579838/relationship-between-a-package-statement-and-the-directory-of-a-go-file)
* [教程五、Go 语言包管理（Package）必知必会](https://learnku.com/go/t/27649)
* [Go基础：路径、文件名和包名的关系](https://developer.aliyun.com/article/888601)
* [【golang学习笔记】包（package）的使用](https://juejin.cn/post/7122730352023437343)


## Code Scope
[Go: "instance" redeclared in this block](https://stackoverflow.com/questions/34344172/go-instance-redeclared-in-this-block)
* TLDR: 
    * 同一个 directory 下的所有文件内容都想象成同在一整张白纸上 (i.e. 同一个scope下)，也就是说所有的function/variable declaration都是在这个 directory 下共享的，分隔的files只是为了程序猿方便管理。在compiler看来都是在同一张白纸上，所以 naming 不能重复。
    * 不同 directory 下的文件则是相互隔离的，也就是说，compiler会认为是不同的scope (好几张白纸)。这也是为什么在 directory A 定义了 `max()`，然后再在 directory B 定义 `max()` 时不会发生冲突。
* 考虑下面两种 Structures:
    * `prob1.go`, `prob2.go` 都有 `package leetcode`. 但是, Structure 2中两个 go files的code共享，Structure 1中不共享。
    * 例如: `prob1.go`中定义了 `type PQ []int`。在Structure 2中，`prob2.go`可以直接“看见”, 所以不能再在`prob2.go`重新定义 `type PQ []int`。但是在Structure 1中，`prob2.go`无法“看见”,  所以可以再在`prob2.go`重新定义 `type PQ []int`。
        * WHY? Go treats each directory as a separate package, even though you declared `package leetcode` in both files.
        * 一句话总结：代码的"可见性"取决于物理上 go files 的放置位置，直隶同一个directory `/dir` 下的go files代码互相可见 (变量名，函数名等的定义都不能重复)。分属于不同directory下的go files代码互相不可见，即使它们的第一行的package declaration相同 (e.g. `package leetcode`)。
```
<========= Structure 1 =========>
<----- Code are NOT shared between prob1.go and prob2.go ----->
leetcode/
    ├── prob1/             // prob1/ 是 prob1.go 的直隶directory
    │   └── prob1.go
    └── prob2/             // prob2/ 是 prob2.go 的直隶directory
        └── prob2.go

<========= Structure 2 =========>
<----- Code are shared between prob1.go and prob2.go ----->
leetcode/             // leetcode/ 是 prob1.go 和 prob2.go 的直隶directory
    ├── prob1.go
    └── prob2.go
```