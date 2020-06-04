# Create
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

# Insert
```go
----
mymap["age"] = 25
```

# Delete
```go
----
delete(mymap, "age")
```

# Membership
```go
----
exists := mymap["age"]
```

# Size
```go
----
size := len(mymap)
```

# Iterate
```go
----
for k, v := range mymap{
    fmt.Println(k, v)
}
```