package main

import (
	"fmt"
	"net/url"
)

func constructRelativePath() {
	Url := url.URL{}
	Url.Path = "user/profile"
	parms := url.Values{}
	parms.Add("name", "kelvin")
	parms.Add("age", "25")

	Url.RawQuery = parms.Encode()
	fmt.Printf(Url.String())
}

func constructUrl() {
	u := &url.URL{
		Scheme:   "https",
		User:     url.UserPassword("me", "pass"),
		Host:     "example.com",
		Path:     "foo/bar",
		RawQuery: "x=1&y=2",
		Fragment: "anchor",
	}
	fmt.Println(u.String())
	u.Opaque = "||"
	fmt.Println(u.String())
}

func parseAndSetUrl() {
	Url, err := url.Parse("http://bing.com/search?q=dotnet")
	if err != nil {
		fmt.Println("parse url error")
		return
	}
	Url.Scheme = "https"
	Url.Host = "google.com"
	q := Url.Query()
	q.Set("q", "golang")
	Url.RawQuery = q.Encode()
	fmt.Println(Url)
}

func escapeUrlLink() {
	link := "http://images.com/cat.png"
	fmt.Println("http://mywebpage.com?image=" + url.QueryEscape(link))
}

func main() {
	//constructRelativePath()
	//constructUrl()
	//parseAndSetUrl()
	escapeUrlLink()
}

/* start
---------------------------------------------

Doc: URL Definition
--------------------
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

References
----------
[1] Golang Url Package Documentation: https://golang.org/pkg/net/url
[2] Golang Url Package Examples: https://golang.org/src/net/url/example_test.go

---------------------------------------------
end */
