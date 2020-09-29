# Eg: Create
```go
----
mymap := make(map[string]int)

----
mymap := map[string]int{
    "age": 18,
    "weight": 55,
    "height": 170
}

----
var mymap map[string]int
```

# Eg: Insert
```go
----
mymap["age"] = 25
```

# Eg: Delete
```go
----
delete(mymap, "age")
```

# Eg: Membership
```go
----
exists := mymap["age"]
```

# Eg: Size
```go
----
size := len(mymap)
```

# Eg: Iterate
```go
----
for k, v := range mymap{
    fmt.Println(k, v)
}
```