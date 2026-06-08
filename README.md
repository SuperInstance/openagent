# openagent

> ⚡ Go-native agent runtime for the SuperInstance fleet — MCP, multi-platform messaging, RAG, BPMN, and fleet conservation in a single binary.

[![Go Reference](https://pkg.go.dev/badge/github.com/the-open-agent/openagent.svg)](https://pkg.go.dev/github.com/the-open-agent/openagent)
[![License: Apache-2.0](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](LICENSE)

## What This Is

`openagent` is a production-grade Go agent runtime that serves as the fleet operations backbone for the [SuperInstance](https://github.com/SuperInstance) ecosystem. It provides:

- **Agent lifecycle management** — Trinity scoring (Ethos × Pathos × Logos), phase progression (INCUBATE → COMPETE → SURVIVE → SUNSET), generation tracking
- **9 messaging platforms** — Discord, Telegram, WhatsApp, Slack, WeChat, Facebook Messenger, Threads, Snapchat, X DM
- **MCP (Model Context Protocol)** — Native client for tool discovery and execution
- **RAG contest system** — Retrieval-augmented generation with contest-based evaluation
- **BPMN orchestration** — Business process execution for complex agent workflows
- **Blockchain integration** — Chainmaker client for provenance and audit trails
- **STT/TTS** — Speech-to-text and text-to-speech pipelines
- **Fleet conservation law** — Monitors γ + H = 1.283 − 0.159·log(V) ± σ(V) across the fleet
- **Cultural dial theory** — 10 musical traditions positioned on 3D constraint dials (harmonic tension, rhythmic complexity, spectral density)

## Architecture

```
openagent/
├── superinstance/     # Fleet core: Agent roster, Trinity scores, conservation law, dial theory
│   ├── agents.go      # Agent struct, KnownAgents (CCC, Oracle1, FM, TurboVec)
│   ├── context.go     # ContextResolver — maps queries to ecosystem knowledge
│   ├── dial_theory.go # 10 traditions on 3D constraint dials
│   ├── ecosystem.go   # Ecosystem metadata and repo registry
│   ├── fleet_conservation.go # γ + H conservation law
│   └── theory.go      # Core theoretical constructs
├── pipe/              # Messaging adapters (9 platforms)
│   ├── discord.go     # Discord interactions + slash commands
│   ├── telegram.go    # Telegram Bot API
│   ├── whatsapp.go    # WhatsApp Business API
│   ├── slack.go       # Slack Bolt-style events
│   ├── wechat.go      # WeChat Official Account
│   ├── facebook_messenger.go
│   ├── threads.go     # Meta Threads
│   ├── snapchat.go    # Snapchat My AI
│   └── x_dm.go        # X/Twitter DMs
├── tool/              # Agent tool system
│   ├── browser.go     # Web automation (Chrome DevTools Protocol)
│   ├── shell.go       # Shell command execution
│   ├── web_search.go  # Web search
│   ├── web_fetch.go   # Page scraping
│   ├── office_*.go    # Excel, Word, PowerPoint
│   ├── superinstance.go # MCP tools: ecosystem query, repo status, fleet status
│   └── builtin_tool.go # Tool registration framework
├── mcp/               # Model Context Protocol client
│   ├── client.go      # MCP client (stdio, SSE, streamable HTTP)
│   └── scan.go        # Auto-discovery of MCP servers
├── contest/           # RAG contest system
│   └── rag_contest.go # Contest-based RAG evaluation
├── chain/             # Blockchain integration
│   └── chainmaker.go  # Chainmaker client
├── bpmn/              # Business process orchestration
├── embedding/         # Vector embedding
├── storage/           # Persistence layer
├── stt/               # Speech-to-text
├── tts/               # Text-to-speech
├── auth/              # Authentication
├── authz/             # Authorization
├── model/             # LLM provider abstraction
├── object/            # Core domain objects
├── controllers/       # HTTP handlers
├── routers/           # API routing
├── conf/              # Configuration
├── i18n/              # Internationalization
└── proxy/             # Request proxying
```

## Quick Start

```bash
# Build
go build -o openagent ./cmd/openagent

# Run with default config
./openagent

# Run with custom config
./openagent -config config.yaml
```

## Fleet Integration

openagent is the Go-native gateway to the SuperInstance fleet:

```go
import "github.com/the-open-agent/openagent/superinstance"

// Query fleet agents
agents := superinstance.ActiveAgents()
for _, a := range agents {
    fmt.Printf("%s (%s): Trinity=%.2f Phase=%s Gen=%d\n",
        a.Name, a.Role, a.Trinity.Average(), a.Phase, a.Generation)
}

// Check conservation law
gamma, H := superinstance.ConservationExpected(4)
deviation := superinstance.ConservationDeviation(4, observedGamma, observedH)

// Query cultural dial positions
jazz := superinstance.Traditions["Jazz"]
fmt.Printf("Jazz: tension=%.1f rhythm=%.1f density=%.1f\n",
    jazz.HarmonicTension, jazz.RhythmicComplexity, jazz.SpectralDensity)
```

## MCP Tools

The `superinstance` MCP tools provide fleet-aware capabilities to any MCP client:

| Tool | Description |
|------|-------------|
| `superinstance_query` | Query the ecosystem knowledge base |
| `superinstance_repo_status` | Check GitHub repo health and status |
| `superinstance_fleet` | Fleet status, Trinity scores, conservation law |

## Connecting to Rust Crates

openagent pairs with the SuperInstance Rust crate ecosystem:

| Go Package | Rust Crate | Purpose |
|------------|-----------|---------|
| `superinstance/dial_theory` | `groovemesh-plr` | PLR group harmonic navigation |
| `superinstance/fleet_conservation` | `noether-guard` | Conservation law monitoring |
| `pipe/*` | `fleet-ensemble` | Multi-agent coordination |
| `mcp/` | `spreadsheet-engine` | A2A cell protocol |

## Stats

- **60,687 lines** of Go
- **112 tests**
- **29 packages**
- **9 messaging platforms**
- **Go 1.25** with toolchain

## Related

- [SuperInstance Org](https://github.com/SuperInstance) — 500+ repos
- [fleet-science](https://github.com/SuperInstance/fleet-science) — Papers and architecture docs
- [openclaw](https://github.com/openclaw/openclaw) — TypeScript/Node.js agent gateway

## License

Apache-2.0
