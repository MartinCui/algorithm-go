# algorithm-go

Personal project to practice, implement and verify different algorithms in Golang.

## usage

```shell script
make test folder=${{folderName}}
make bench folder=${{folderName}}
```

# structvspointer

People say in Golang points are not necessary and Struct should be passed whenever possible.
I wrote this benchmark to show how significant this decision will affect Golang performance when passing parameters by 
struct instead of by pointer.
So my understanding:

1. performance wise, when struct has more than 10 fields, "passing by struct" performance drops severely. It slows down 10+ times when struct has 20 fields  
2. correctness wise, they say using struct can avoid nil pointer error. But I say it's better crash instead of "wrong and move on" in most of the software cases.
