# AGENTS.md — Krewire Agent

Operational guide and behavioral rules for autonomous AI agents working inside `krewire/agent`.

## Architecture & Boundaries

- Pure Multi-Module Monorepo with constituent modules in `core/`, `conformance/`, `version/`.
- Zero root `go.mod`; workspace orchestrated through native `go.work`.
- Strictly pure-Go standard library first (`CGO_ENABLED=0`).
- No CLI binaries are built here; all developer CLI entrypoints live in `github.com/krewire/krew`.

## Quality Gates

Before completing changes:
```bash
gofmt -l .
go vet ./...
go test ./...
CGO_ENABLED=0 go build ./...
```
