# [630. Course Schedule III](https://leetcode.com/problems/course-schedule-iii/)

## Solution idea:

Psuedo-code:

```
Earliest(courses):
    Sort courses by finish time and re-number the courses so that f1 <= f2 <= ... <= fn
    Init S = {} to be the selected courses
    for i := 1 to n:
        if course[i] is compatible with S {
            add course[i] to S
        }
    return S.length()
```

NOTE (IMPORTANT!!!)

We need to check TWO cases whether `course[i] is compatible with S`:

Case 1: if adding `course[i]` doesn't pass the deadline, directly add `course[i]`

(假设我们当前已经修了Ｎ门．那么接下来又到了下一个deadline，也就是deadline放宽了，但我们又新添了一门课变成N+1门．如果能赶在新deadline之前搞定这门新课，我们自然就是能上就上（那样就是N+1门))

Case 2: if adding `course[i]` pass the deadline, we need to swap `course[i]` with a course in `S` that has the longest duration.

(如果不能呢？我们自然怪罪当前N+1课程列表里最长的那门，我们只要把那门最长的踢掉就一定满足deadline的要求．(为什么？因为我们之前保证了Ｎ门可以满足上一个deadline，那么现在的Ｎ门一定也可以满足当前的deadline.))