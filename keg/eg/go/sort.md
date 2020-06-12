# Eg: Sort Strings
```go
----
var strlist []string
sort.Strings(strlist)
```

# Eg: Sort any Slice
```go
----
sort.Slice(myList, func(i, j int) bool { return myList[i] < myList[j] })
```

# References
[1] Golang Sort Package Documentation: https://golang.org/pkg/sort/
