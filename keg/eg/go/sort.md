# Sort Strings
```go
----
var strlist []string
sort.Strings(strlist)
```

# Sort any Slice
```go
----
sort.Slice(myList, func(i, j int) bool { return myList[i] < myList[j] })
```

