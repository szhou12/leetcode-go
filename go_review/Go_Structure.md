# [Structure](https://github.com/szhou12/leetcode-go/blob/main/go_review/README.md)

## Contents
* [Files Under Same Package (Redeclare error)](#files-under-same-package)

## Files Under Same Package
[Go: "instance" redeclared in this block](https://stackoverflow.com/questions/34344172/go-instance-redeclared-in-this-block)
* TLDR: 
    * 同一个package下的所有文件内容都想象成同在一整张白纸上，分隔的files只是为了developer方便管理。在compiler看来都是在同一张白纸上，所以 variables naming 不能重复。
    * "Files have no real meaning for go, unlike in java, python and many others, they are just for you to organize your code as you see fit."
    * "In go variables are visible `package` wide. Hence the compiler complains about having two global variables with the same name."