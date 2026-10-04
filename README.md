# botid

![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)
![Go 1.25](https://img.shields.io/badge/Go-1.25-00ADD8.svg)

The cl0q crawler's bot identity: the User-Agent strings its crawlers send and the
robots.txt user-agent names that address them (`cl0q-crawler`, `cl0q-discovery`,
`cl0q-favicon`, `cl0q-node`, `cl0q`). Shared by the crawlers and the robots.txt parser
so the name a site owner blocks and the name we send always match.

```go
ua, _ := botid.UserAgent("my-node")   // Mozilla/5.0 (compatible; cl0q-node/1.0; operator=my-node; +https://cl0q.com/bot)
botid.IsOurs("cl0q-crawler")           // does this robots.txt user-agent name us?
```

## Quick start

```bash
go test ./...
go build ./...
```

## How it works

One file (`botid.go`), stdlib only. The bot names and User-Agent formats live here
— and only here — so the name a site owner sees in their logs is always the name
they can block in robots.txt. The crawlers take the identity from this package;
it is not configurable per deployment.

## Configuration

None — this is a library with no knobs.

## Related repositories

- [robotstxt](../robotstxt) — the robots.txt parser that matches these names.
- [crawler-node](../crawler-node) — the crawler that sends these User-Agents (vendored as `botid/`).

Tests: `go test ./...`. MIT licensed.
