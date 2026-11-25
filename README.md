# Krewire Agent (`github.com/krewire/agent`)

Autonomous AI agent engine, multi-provider LLM integrations, and SDLC governance runtime for the Krewire digital SDLC ecosystem.

Part of the [Krewire](https://krewire.com) digital SDLC ecosystem.

## Overview

Krewire Agent provides pure-Go agentic building blocks:
- Autonomous runtime with subagent delegation and role-based access control (`core/agent`)
- Unified LLM provider drivers: OpenAI, Anthropic, Gemini, Ollama, and mock testing (`core/providers`)
- Context window management, token budgeting, and session persistence (`core/memory`)
- SDLC quality verification and governance workflows (`core/sdlc`)
- Tool calling primitives: AST exploration, patching, execution sandboxes, and file system navigation (`core/tools`)
- Agent persona system and prompt templates (`core/persona`, `core/boost`)
- Model Context Protocol (MCP) and HTTP/SSE transport drivers (`core/transport`)

## Architecture & Invariants

- **Pure Go & CGO-Free**: Fully compilable with `CGO_ENABLED=0`.
- **Ecosystem Battery Integration**: CLI driver commands (`krew chat`, `krew agent`, `krew daemon`) are built and executed inside `github.com/krewire/krew`, consuming `github.com/krewire/agent/core`.
- **Multi-Module Monorepo**: Pure monorepo without a root `go.mod`, coordinated via `go.work`.

## Quality Gates

Before submitting changes:
```bash
gofmt -l .
go vet ./...
go test -v ./...
CGO_ENABLED=0 go build ./...
```
