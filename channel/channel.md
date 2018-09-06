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
If channel is closed then we can not send message to the channel, or it will panic.
```go
ch := make(chan int)
close(ch)
ch <- 1 // panic: send on closed channel
```
if close a channel, and a sender goroutine is waiting on the  blocking channel, then it will panic. 
```go
func main() {
    ch := make(chan int)
    go func() { ch <- 1 }() // panic: send on closed channel
    time.Sleep(time.Second)
    go func() { close(ch) }()
    time.Sleep(time.Second)
    x, ok := <-ch
    fmt.Println(x, ok)
}
```
close a channel will wake up all the goroutine which waiting on the channel. And then write g will find that the channel is closed, it will panic. One more thing, ch <- 1, it will block here until another g take away the message in the channel
<!--stackedit_data:
eyJoaXN0b3J5IjpbMTEzNTQyODAzNSwtMTU4NDQxODUzMywxMj
IyODcwOTcxLDk1MjczNzI3MCwtMTcyNTMwMTYzNCwtMTQzNDc1
NjcyM119
-->