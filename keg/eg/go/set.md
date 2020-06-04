# Create
```go
----
myset := make(map[string]bool)
```

# Insert
```go
----
myset["foo"] = true
```

# Delete
```go
----
delete(myset, "foo")
```

# Membership
```go
----
exists := myset["foo"]
```

# Size
```go
----
size := len(myset)
```

# Iterate
```go
----
for k := range {
    fmt.Println(k)
}
```