```go
package main

import (

        "fmt"
        "sync"
)

func main() {
        var wg sync.WaitGroup
        ch := make(chan int, 10) // Buffered channel

        for i := 0; i < 10; i++ {
                wg.Add(1)
                go func(i int) {
                        defer wg.Done()
                        ch <- i
                }(i)
        }

        go func() {
                wg.Wait()
                close(ch)
        }()

        for v := range ch {
                fmt.Println(v)
        }
}
```