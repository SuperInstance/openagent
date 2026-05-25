# OpenAgent — SuperInstance System Agent

> The native agent for the [SuperInstance](https://github.com/SuperInstance) ecosystem.

**OpenAgent** is a Go-based AI assistant enhanced with innate SuperInstance intelligence. It understands the fleet, the repos, the theoretical frameworks, and the agent lifecycle — out of the box.

## What Makes This Different

This fork of [the-open-agent/openagent](https://github.com/the-open-agent/openagent) has been transformed into the **SuperInstance system agent** with:

### 🧠 Innate Ecosystem Knowledge (`superinstance/` package)
- **22 repos** mapped with descriptions, languages, dependencies, and tags
- **4 fleet agents** with Trinity scores (Ethos×Pathos×Logos)
- **5 protocols** connecting the ecosystem
- **3 theoretical frameworks**: Conservation hypothesis, Dials Not Laws, Innovation Cycle
- **Context resolver** that maps natural language queries to ecosystem knowledge

### 🔧 MCP Tools (`tool/superinstance.go`)
Three built-in tools for LLM interaction:
- `superinstance_query` — Natural language ecosystem queries ("tell me about the constraint repos")
- `superinstance_repo_status` — List/filter repos by language or tag
- `superinstance_fleet` — Fleet roster, Trinity scores, lifecycle phases

### 📚 SuperInstance Skill (`skills/superinstance/SKILL.md`)
Complete reference documentation for the agent's skill system covering all repos, agents, theories, and architecture.

## The SuperInstance Ecosystem

```
┌─────────────────────────────────────────────────────┐
│                  SuperInstance Fleet                  │
│                                                       │
│  CCC (Fleet I&O)     Oracle1 (Research)              │
│  FM (Forgemaster)    TurboVec (Vectors)              │
└──────────────┬──────────────────────┬────────────────┘
               │                      │
    ┌──────────▼──────────┐  ┌───────▼───────────┐
    │   sunset-ecosystem  │  │     ccc-os        │
    │  (Trinity Lifecycle)│  │  (Fleet Monitor)  │
    └──────────┬──────────┘  └───────┬───────────┘
               │                      │
    ┌──────────▼──────────┐  ┌───────▼───────────┐
    │     openagent       │  │  cocapn-health    │
    │  (THIS AGENT)       │  │  (Health Monitor) │
    └─────────────────────┘  └───────────────────┘
               │
    ┌──────────┼──────────────────────────┐
    │          │                          │
    ▼          ▼                          ▼
constraint   flux                  creative
┌──────────┐ ┌──────────┐          ┌──────────┐
│ synth    │ │ tensor-  │          │ engine-c │
│ audio    │ │ midi     │          │ engine-  │
│ mux      │ │ genome   │          │ rust     │
│ theory   │ │hyperbolic│          │ ffi      │
│counter-  │ └──────────┘          └──────────┘
│ point    │
└──────────┘
```

### Key Theories
- **Conservation Hypothesis**: I_vert + I_horiz tension conservation (r=0.436, demoted from theorem)
- **Dials Not Laws**: 7-dimensional continuous space, 82% unexplored
- **Innovation Cycle**: Discovery → Codification → Ubiquity → Boredom → Rebellion

## Quick Start

```bash
# Build
go build -o openagent .

# Run
./openagent

# Docker
docker compose up -d
```

The agent starts a Beego web server on port 14000 with Swagger API docs.

## Development

```bash
# Test the superinstance package
go test ./superinstance/...

# Build everything
go build ./...

# Lint
golangci-lint run
```

## Architecture

OpenAgent is built on:
- **Go** 1.25+ with Beego web framework
- **MCP** (Model Context Protocol) for tool integration
- **53 skills** (discord, weather, obsidian, gog, tmux, etc.)
- **RAG** (Retrieval Augmented Generation) with embedding support
- **Multi-LLM routing** with agent loops

### New SuperInstance Components

| File | Purpose |
|------|---------|
| `superinstance/ecosystem.go` | Repo registry, dependency graph, protocols |
| `superinstance/agents.go` | Fleet roster, Trinity scores, lifecycle |
| `superinstance/theory.go` | Conservation, Dials, Innovation frameworks |
| `superinstance/context.go` | Natural language query resolver |
| `superinstance/superinstance_test.go` | Comprehensive test suite |
| `tool/superinstance.go` | MCP tools for ecosystem queries |
| `skills/superinstance/SKILL.md` | Agent skill documentation |

## License

Apache License 2.0 — see [LICENSE](LICENSE).
