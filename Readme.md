##### GitHub OAuth HTTP handler


See <http://godoc.org/github.com/heroku/herokugoauth> for documentation.

###### Example

```go
h := &githubauth.Handler{
	RequireDomain:   "heroku.com",
	// e.g. faba0c08be7474a785b272c4f4154c998c0943b51e662637be11b1a0ecda43b3
	Keys:         os.Getenv("KEY")),
	ClientID:     os.Getenv("OAUTH_CLIENT_ID"),
	ClientSecret: os.Getenv("OAUTH_CLIENT_SECRET"),
}
http.ListenAndServe(":8080", h)
```
