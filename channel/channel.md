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
```go
close(ch)
close(ch) // panic: close of closed channel
```
close channel will return the while loop
```go
func main() {
    ch := make(chan int, 100)
    for elem := range ch {
        fmt.Println(i)
    }
}
```
We can get the channel capacity and length 
```go
func main() {
    ch := make(chan int, 100)
    ch <- 1
    fmt.Println(len(ch)) // 1
    fmt.Println(cap(ch)) // 100
}
```
len and cap is not function call, when build it will get hchan's field.
<!--stackedit_data:
eyJoaXN0b3J5IjpbOTUyNzM3MjcwLC0xNzI1MzAxNjM0LC0xND
M0NzU2NzIzXX0=
-->