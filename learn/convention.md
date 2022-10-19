goroutines

```
func r () {
    ...doSomething
}

func main () {
    go r()
}
```


note:
    when creating a goroutine, know how it will end, avoid memory leaks
    check for race conditions at compile time

    Synchronization
        use sync.WaitGroup to wait for groups of goroutines to complete
        use sync.Mutex and sync.RWMutex to protect data access
    Parallelism
        by default, go will use CPU threads = available cores
        change with runtime.GOMAXPROCS()
        more threads can increase performance, but too many can slow it down