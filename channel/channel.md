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
<!--stackedit_data:
eyJoaXN0b3J5IjpbLTE3MjUzMDE2MzQsLTE0MzQ3NTY3MjNdfQ
==
-->