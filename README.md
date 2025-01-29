# Go WaitGroup and Channel Deadlock

This repository demonstrates a potential deadlock scenario in Go when using a `sync.WaitGroup` with a channel.  The example shows how improper handling of goroutines and channel closure can lead to unexpected behavior.

## Bug Description
The provided code creates 10 goroutines, each sending a value to a channel. A separate goroutine waits for all sending goroutines to finish using `wg.Wait()` before closing the channel.  However, if the channel's buffer is not large enough to accommodate all the values sent by the goroutines before `wg.Wait()` completes, the sending goroutines will block, causing a deadlock because the channel cannot be closed and the WaitGroup counter is not decremented.