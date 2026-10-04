# botid

![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)
![Go 1.25](https://img.shields.io/badge/Go-1.25-00ADD8.svg)

The cl0q crawler's bot identity: the User-Agent strings its crawlers send and the
robots.txt user-agent names that address them (`cl0q-crawler`, `cl0q-discovery`,
`cl0q-favicon`, `cl0q-node`, `cl0q`). Shared by the crawlers and the robots.txt parser
so the name a site owner blocks and the name we send always match.

## Installation

```bash
go get cl0q/botid
```

Requires Go 1.25+.

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

`botid.MustUserAgent()` reads the `CL0Q_NODE_ID` environment variable (the operator id
assigned at enrollment) and exits if it is unset: a node must never crawl anonymously.
Library users who manage the id themselves should prefer `botid.UserAgent(id)`.

## Usage

```go
import "cl0q/botid"

// Generate a User-Agent for your crawler node
ua, _ := botid.UserAgent("my-node")
// Mozilla/5.0 (compatible; cl0q-node/1.0; operator=my-node; +https://cl0q.com/bot)

// Check if a robots.txt user-agent name refers to us
botid.IsOurs("cl0q-crawler")  // true
botid.IsOurs("googlebot")     // false
```

### Environment-based configuration

`botid.MustUserAgent()` reads the `CL0Q_NODE_ID` environment variable (the operator id
assigned at enrollment) and exits if it is unset: a node must never crawl anonymously.

```go
// Exits if CL0Q_NODE_ID is not set
ua := botid.MustUserAgent()
```

Library users who manage the id themselves should prefer `botid.UserAgent(id)`.

## API Reference

### `UserAgent(nodeID string) (string, error)`

Generates a User-Agent string for the given node ID. Returns an error if the node ID
is empty.

### `MustUserAgent() string`

Like `UserAgent` but reads the node ID from `CL0Q_NODE_ID` and exits on failure.
For use in crawler binaries, not libraries.

### `IsOurs(name string) bool`

Reports whether the given robots.txt user-agent name addresses a cl0q crawler.
Matches `cl0q-crawler`, `cl0q-discovery`, `cl0q-favicon`, `cl0q-node`, and `cl0q`
(case-insensitive).

## Configuration

None — this is a library with no knobs.

## Testing

```bash
go test ./...
go test -race ./...
```

## Contributing

Pull requests welcome. Please:
- Run `gofmt` on all changes
- Add tests for new functionality
- Keep the User-Agent format stable (sites may parse it)

## Related repositories

- [robotstxt](../robotstxt) — the robots.txt parser that matches these names.
- [crawler-node](../crawler-node) — the crawler that sends these User-Agents (vendored as `botid/`).

## License

MIT. See `LICENSE`.
