# openagent

**Personal AI assistant in Go** — a lightweight, extensible agent runtime for running AI assistants locally.

## What This Gives You

- **Go-native** — fast startup, low memory, single binary deployment
- **Extensible** — plugin architecture for adding capabilities
- **Local-first** — runs on your machine, no cloud required
- **Multi-model** — connect to any LLM provider

## Quick Start

```bash
go build -o openagent ./cmd/openagent
./openagent
```

## How It Fits

Go implementation of the agent runtime pattern used across the SuperInstance ecosystem. Complements the TypeScript/Node.js `claw` gateway with a lighter-weight Go alternative.

## License

MIT
