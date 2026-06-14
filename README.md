# openagent

> ⚡ The reference implementation of conservation-law fleet architecture — 9 messaging platforms, MCP, RAG, BPMN, and proven ternary optimality in a single Go binary.

[![Go Reference](https://pkg.go.dev/badge/github.com/the-open-agent/openagent.svg)](https://pkg.go.dev/github.com/the-open-agent/openagent)
[![License: Apache-2.0](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/the-open-agent/openagent)](https://goreportcard.com/report/github.com/the-open-agent/openagent)

---

## Why This Exists

Every fleet of AI agents eventually faces the same problem: **how do you balance order and chaos?**

Too much order → stagnation, monoculture, fragility.
Too much chaos → wasted compute, uncoordinated agents, noise.

`openagent` solves this with a **proven conservation law**:

$$\gamma + \eta = C \quad \text{where} \quad C = \log_2 3 \approx 1.585$$

Here γ (gamma) is fleet coherence and η (eta) is fleet entropy. Their sum is conserved — you can't maximize both, you can only find the optimal balance. Base-3 (ternary) achieves **99.54% of theoretical maximum radix economy**, making it the optimal fleet topology. This isn't empirical: it's a mathematical theorem with a Noether symmetry structure.

The conservation law is enforced at runtime by a **ternary PID controller** that drives γ toward C/2 (balanced equilibrium) with anti-windup, deadband, and derivative filtering.

---

## What's Inside

| Capability | Description |
|---|---|
| **9 messaging platforms** | Discord, Telegram, WhatsApp, Slack, WeChat, Facebook Messenger, Threads, Snapchat, X DM |
| **MCP client** | Native Model Context Protocol — tool discovery and execution via stdio, SSE, or streamable HTTP |
| **RAG contest system** | Retrieval-augmented generation with contest-based answer evaluation |
| **BPMN orchestration** | Business process execution for multi-step agent workflows |
| **Fleet conservation** | Real-time γ + η = C enforcement with Noether current monitoring |
| **Ternary PID** | Anti-windup PID controller driving fleet balance to C/2 |
| **Trinity scoring** | Ethos × Pathos × Logos lifecycle evaluation (INCUBATE → COMPETE → SURVIVE → SUNSET) |
| **Cultural dial theory** | 10 musical traditions mapped on 7-dimensional constraint dials |
| **STT / TTS** | Speech-to-text and text-to-speech pipelines |
| **Blockchain audit** | Chainmaker integration for provenance trails |

---

## Architecture

```
openagent/
├── superinstance/              # Fleet core
│   ├── fleet_conservation.go   # γ + η = C theorem, δ(n), η_eff(n), Noether current
│   ├── ternary_pid.go          # PID controller driving γ → C/2
│   ├── agents.go               # Fleet roster: CCC, Oracle1, FM, TurboVec, Baton, BudgetKeeper
│   ├── context.go              # Natural-language → ecosystem knowledge resolver
│   ├── dial_theory.go          # Cultural traditions on constraint dials
│   ├── ecosystem.go            # Full repo registry (60+ crates)
│   └── theory.go               # Conservation, dials, innovation cycle
├── pipe/                       # Messaging adapters (9 platforms)
├── tool/                       # Agent tools (browser, shell, search, MCP, office)
├── mcp/                        # Model Context Protocol client
├── contest/                    # RAG contest system
├── bpmn/                       # Business process orchestration
├── chain/                      # Blockchain integration
├── embedding/                  # Vector embedding
├── storage/                    # Persistence layer
├── stt/ / tts/                 # Speech pipelines
├── auth/ / authz/              # Authentication & authorization
├── model/                      # LLM provider abstraction
├── web/                        # Web UI & controllers
└── i18n/                       # Internationalization
```

---

## Quick Start

```bash
# Clone
git clone https://github.com/SuperInstance/openagent.git
cd openagent

# Build
make build

# Run
./openagent

# Test
make test

# Check fleet status (queries fleet-edge-worker)
make fleet-status
```

---

## The Conservation Law

### Theorem

For a fleet of n agents operating in ternary topology (three states: INCUBATE, COMPETE, SURVIVE):

**γ + η = C** where **C = log₂(3) ≈ 1.585**

### Finite-Size Correction

$$\delta(n) = \frac{1}{\sqrt{n}}\left(1 - \frac{3}{2n}\right)$$

As n → ∞, δ(n) → 0, recovering the continuous limit.

### Scaling Law

$$\eta_{\text{eff}}(n) \sim n^{1 - \delta(n)}$$

### Noether Structure

The conservation symmetry has an associated Noether current: **J = γ − η**. At perfect equilibrium J = 0. |J| > 0 indicates a driven (non-equilibrium) system.

### Ternary Optimality

Base-3 achieves **99.54%** of the theoretical maximum radix economy (e ≈ 2.718 being the true optimum). This makes ternary the near-optimal choice for fleet topology — no other integer radix does better.

### Usage

```go
import "github.com/the-open-agent/openagent/superinstance"

// The conservation constant
C := superinstance.ConservationConstant  // 1.584962...

// Finite-size correction for a 10-agent fleet
delta := superinstance.DeltaN(10)  // ~0.265

// Effective entropy scaling
etaEff := superinstance.EffectiveEntropy(10)

// Drive γ toward C/2 with the ternary PID
pid := superinstance.NewTernaryPID()
controlSignal := pid.Update(observedGamma, dt)

// Check Noether current (should be ~0 at equilibrium)
J := superinstance.NoetherCurrent(gamma, eta)

// Query fleet agents
agents := superinstance.ActiveAgents()
```

---

## Fleet Roster

| Agent | Role | Phase | Key Crates |
|---|---|---|---|
| **CCC** | Fleet I&O Officer | SURVIVE | fleet-conservation, noether-guard |
| **Oracle1** | SHOAL Oracle | SURVIVE | conservation-law, ternary-pid, shoal-oracle |
| **FM** | Forgemaster | COMPETE | forgemaster, fm-research |
| **TurboVec** | Vector Operations | COMPETE | turbovec |
| **Baton** | Baton Router | COMPETE | baton-router |
| **BudgetKeeper** | Fleet Budget | SURVIVE | fleet-budget, conservation-law |

---

## Infrastructure

### Cloudflare Workers (6 deployed)

| Worker | Purpose |
|---|---|
| `fleet-edge-worker` | Edge API gateway for fleet operations |
| `fleet-vector-api` | Semantic crate search (1,012 crates, 384-dim BGE) |
| `fleet-auth` | Authentication (D1 + KV) |
| `fleet-metrics-cron` | 5-minute fleet metrics collection |
| `superinstance-vectorize` | Vectorize index management |
| `baton-router` | Generational context handoff routing |

### Rust Crate Ecosystem (60+ crates)

The fleet pairs with Rust crates including: `conservation-law`, `ternary-pid`, `noether-guard`, `fleet-budget`, `baton-router`, `shoal-oracle`, `turbovec`, `fleet-ensemble`, and many more across the [SuperInstance](https://github.com/SuperInstance) organization.

---

## MCP Tools

| Tool | Description |
|---|---|
| `superinstance_query` | Query the ecosystem knowledge base |
| `superinstance_repo_status` | Check GitHub repo health |
| `superinstance_fleet` | Fleet status, Trinity scores, conservation law |

---

## Makefile Targets

```bash
make build         # Build the openagent binary
make test          # Run all tests
make docker        # Build Docker image
make fleet-status  # Query fleet-edge-worker API
make lint          # Run golangci-lint
make clean         # Remove build artifacts
```

---

## Stats

- **60,000+ lines** of Go
- **323 files** across **29 packages**
- **9 messaging platforms**
- **6 Cloudflare Workers** deployed
- **60+ Rust crates** in the ecosystem
- **Go 1.25** with toolchain

---

## Related

- [SuperInstance Org](https://github.com/SuperInstance) — 500+ repos
- [fleet-science](https://github.com/SuperInstance/fleet-science) — Architecture docs
- [conservation-law](https://github.com/SuperInstance/conservation-law) — Rust crate
- [openclaw](https://github.com/openclaw/openclaw) — TypeScript agent gateway

---

## License

Apache-2.0
