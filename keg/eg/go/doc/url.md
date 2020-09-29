# Eg: Construct Relative Path
```go
----
Url := url.URL{}
Url.Path = "user/profile"
parms := url.Values{}
parms.Add("name", "kelvin")
parms.Add("age", "25")

Url.RawQuery = parms.Encode()
fmt.Printf(Url.String())
// Output:
// user/profile?age=25&name=kelvin
```

# Eg: Construct Url
```go
----
u := &url.URL{
    Scheme:   "https",
    User:     url.UserPassword("me", "pass"),
    Host:     "example.com",
    Path:     "foo/bar",
    RawQuery: "x=1&y=2",
    Fragment: "anchor",
}
fmt.Println(u.String())
u.Opaque = "opaque"
fmt.Println(u.String())
// Output:
// https://me:pass@example.com/foo/bar?x=1&y=2#anchor
// https:opaque?x=1&y=2#anchor
```

# Eg: Parse and Set Url
```go
----
Url, err := url.Parse("http://bing.com/search?q=dotnet")
if err != nil {
    logs.Error("parse url error")
}
Url.Scheme = "https"
Url.Host = "google.com"
q := Url.Query()
q.Set("q", "golang")
Url.RawQuery = q.Encode()
fmt.Println(Url)
// Output: 
// https://google.com/search?q=golang
```

# Eg: Escape Url Link
```go
----
link := "http://images.com/cat.png"
fmt.Println("http://mywebpage.com?image=" + url.QueryEscape(link))
// Output:
// http://mywebpage.com?image=http%3A%2F%2Fimages.com%2Fcat.png
```

---------------------------------------------------------------------------------------
---------------------------------------------------------------------------------------
# Doc: URL Definition
```go
----
type URL struct {
    Scheme     string
    Opaque     string    // encoded opaque data
    User       *Userinfo // username and password information
    Host       string    // host or host:port
    Path       string    // path (relative paths may omit leading slash)
    RawPath    string    // encoded path hint (see EscapedPath method); added in Go 1.5
    ForceQuery bool      // append a query ('?') even if RawQuery is empty; added in Go 1.7
    RawQuery   string    // encoded query values, without '?'
    Fragment   string    // fragment for references, without '#'
}
````

# References
[1] Golang Url Package Documentation: https://golang.org/pkg/net/url
[2] Golang Url Package Examples: https://golang.org/src/net/url/example_test.go