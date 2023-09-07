# [Package](https://github.com/szhou12/leetcode-go/blob/main/go_review/README.md)

## Contents
* [Package Managment](#package-management)
* [Files under same package (Redeclare error)](#files-under-same-package)

## Package Management

### Package VS. Directory
* 必须要明确这两个概念的区别：
    * **Package (包)**
        * 逻辑上一个或多个go文件的集合，它们都由 **包的声明**`package <pacakgeName>` 所标识，它们一起组成了一个package的实现。注意：这些go文件不一定都放在同一个directory下。
        * e.g. 你可能注意到每道题solution的go文件都是以`package leetcode`开头，这是在告诉programmer这些go文件都属于名叫 `leetcode` 这个package, which is the default alias when import。
    * **Directory (目录)**
        * 也即文件夹。物理上存放一个或多个go文件的地方。这是在告诉compiler去哪里找到一个特定的package name。目录的路径是import path，以string形式表示。一个目录下是一个共享的scope，不可以redeclare。
* By convention, 一个go文件在哪个directory下，就declare 它的 package name 为这个directory的名字。但是这并不是强制的，只是为了方便programmer管理查找。
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


## Files Under Same Package
[Go: "instance" redeclared in this block](https://stackoverflow.com/questions/34344172/go-instance-redeclared-in-this-block)
* TLDR: 
    * 同一个 directory 下的所有文件内容都想象成同在一整张白纸上 (i.e. 同一个scope下)，也就是说所有的function/variable declaration都是在这个 directory 下共享的，分隔的files只是为了developer方便管理。在compiler看来都是在同一张白纸上，所以 naming 不能重复。