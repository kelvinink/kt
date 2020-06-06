# Retry
```go
----
for retry := 0; retry < 3; retry++ {
    err := trydo()
    if err != nil{
        logs.Info()
        time.Sleep(time.Second*3)
    } else {
        break
    }
}

----
for retry := 0; retry < 3; retry++ {
    err := trydo()
    if err == nil{
        break
    }

    logs.Info()
    time.Sleep(time.Second*3)
}
```