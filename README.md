# botid

The cl0q crawler's bot identity: the User-Agent strings its crawlers send and the
robots.txt user-agent names that address them (`cl0q-crawler`, `cl0q-discovery`,
`cl0q-favicon`, `cl0q-node`, `cl0q`). Shared by the crawlers and the robots.txt parser
so the name a site owner blocks and the name we send always match.

```go
ua, _ := botid.UserAgent("my-node")   // Mozilla/5.0 (compatible; cl0q-node/1.0; operator=my-node; +https://cl0q.com/bot)
botid.IsOurs("cl0q-crawler")           // does this robots.txt user-agent name us?
```

Tests: `go test ./...`
