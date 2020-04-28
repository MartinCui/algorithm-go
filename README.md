# algorithm-go

Personal project to practice, implement and verify different algorithms in Golang.

## usage

```shell script
make test folder=${{folderName}}
make bench folder=${{folderName}}
```

## data structures

1. Linked List
1. Trees, Tries & Graphs
1. Stacks & Queues
1. Heaps
1. Vectors/ArrayLists
1. Hash Tables (very important)

## algorithms

1. Breadth-First Search
1. Depth-First Search
1. Binary Search
1. Merge Sort
1. Quick Sort

## concepts

1. Bit Manipulation
1. Memory (Stack vs Heap)
1. Recursion
1. Dynamic Programming
1. Big O Time & Space

# structvspointer

People say in Golang points are not necessary and Struct should be passed whenever possible.
I wrote this benchmark to show how significant this decision will affect Golang performance when passing parameters by 
struct instead of by pointer.
So my understanding:

1. performance wise, when struct has more than 10 fields, "passing by struct" performance drops severely. It slows down 10+ times when struct has 20 fields  
2. correctness wise, they say using struct can avoid nil pointer error. But I say it's better crash instead of "wrong and move on" in most of the software cases.
