# Eg: Marshal
```go
----
bytes, err := json.Marshal(object)
if err != nil {
    logs.Info("Json marshal error")
}
```

# Eg: Unmarshal
```go
----
err := json.Unmarshal(bytes, &model)
if err != nil {
    logs.Info("Json unmarshal error")
}
```