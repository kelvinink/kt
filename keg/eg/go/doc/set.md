# Eg: Create
```go
----
myset := make(map[string]bool)
```

# Eg: Insert
```go
----
myset["foo"] = true
```

# Eg: Delete
```go
----
delete(myset, "foo")
```

# Eg: Membership
```go
----
exists := myset["foo"]
```

# Eg: Size
```go
----
size := len(myset)
```

# Eg: Iterate
```go
----
for k := range {
    fmt.Println(k)
}
```