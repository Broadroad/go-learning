# Channel

channel is the very important data structure in golang. It is a lock-structure, consist CSP.

## Usage
### block channel
```go
var a = make(chan int)
```
### Unblock channel
```go
var a = make(chan int, 10)
```
The most different between block and unblock channel is the capacity. Block channel is only the tool of communication and sync.
send data to channel:
```go
ch := make(chan int, 100)
ch <- 1
```
receive from channel:
```go
ch := make(chan int, 100)
ch <- 1
```
close channel
```go
close(ch)
```
re-close channel will cause panic
<!--stackedit_data:
eyJoaXN0b3J5IjpbMTM4OTU3NjU0NiwtMTcyNTMwMTYzNCwtMT
QzNDc1NjcyM119
-->