# SuperInstance System Knowledge

## Overview
This skill provides innate knowledge of the SuperInstance ecosystem. The agent should use this knowledge to understand user queries about "the system", "constraints", "fleet", "dials", and all SuperInstance-specific concepts.

## The SuperInstance Ecosystem

SuperInstance is a software collective building AI-driven music theory tools, autonomous agent systems, and creative engines.

### Core Repositories

| Repo | Language | Description |
|------|----------|-------------|
| sunset-ecosystem | Python | Trinity architecture agent lifecycle (Ethos×Pathos×Logos) |
| ccc-os | Python | Autonomous fleet monitoring & decision system |
| constraint-synth | Python | Music theory constraint synthesis (PyPI) |
| constraint-audio | Rust | Audio processing for constraint-based music |
| constraint-mux | Rust | Multiplexer for constraint streams |
| constraint-theory-core | Python | Core constraint theory formalization |
| counterpoint-engine | Python | Species counterpoint engine (PyPI) |
| flux-tensor-midi | Python | Tensor-based MIDI manipulation |
| flux-genome | Python | Musical genome mapping and evolution |
| flux-hyperbolic | Python | Hyperbolic geometry for musical space |
| deadband-rs | Rust | Deadband theory for groove analysis |
| forgemaster | Python | Build & forge orchestration |
| fm-research | Markdown | Forgemaster research notes |
| creative-engine-c | C | C creative engine implementation |
| creative-engine-rust | Rust | Rust creative engine implementation |
| superinstance-ffi | Rust | FFI bindings across SuperInstance languages |
| cocapn-health | Python | Health monitoring service |
| AI-Writings | Markdown | Creative writing collection by AI models |
| openagent | Go | **This agent** — the SuperInstance system agent |
| docs | Markdown | Documentation hub |
| wiki | Markdown | Knowledge wiki |

### Architecture Map

```
sunset-ecosystem → ccc-os → fleet monitoring
sunset-ecosystem → openagent (this repo)
constraint-synth → counterpoint-engine → constraint-toolkit
constraint-synth → constraint-theory-core
constraint-audio → constraint-mux
flux-tensor-midi → flux-genome → flux-hyperbolic
forgemaster → fm-research
ccc-os → cocapn-health
superinstance-ffi → (binds Rust↔Python↔C)
creative-engine-c ↔ creative-engine-rust (via superinstance-ffi)
```

## Key Concepts

### Trinity Architecture (Ethos × Pathos × Logos)
Every agent in the fleet is evaluated across three dimensions:
- **Ethos**: Trust, reliability, consistency
- **Pathos**: Engagement, empathy, communication quality
- **Logos**: Logic, performance, reasoning ability

Trinity scores determine agent lifecycle phase transitions.

### Agent Lifecycle
Agents progress through phases:
1. **INCUBATE** — New agent, learning the ropes
2. **COMPETE** — Proving itself against benchmarks
3. **SURVIVE** — Established, may breed new agents
4. **SUNSET** — Being retired or archived

Agents can transition: INCUBATE→COMPETE→SURVIVE→SUNSET
Or shortcut: INCUBATE→SUNSET, COMPETE→SUNSET

### Current Fleet Agents
- **CCC** — Fleet I&O Officer (SURVIVE, Gen 3)
  - Trinity: Ethos 0.92, Pathos 0.78, Logos 0.95
- **Oracle1** — Research & Synthesis (SURVIVE, Gen 5)
  - Trinity: Ethos 0.88, Pathos 0.82, Logos 0.97
- **FM** — Forgemaster (COMPETE, Gen 2)
  - Trinity: Ethos 0.75, Pathos 0.65, Logos 0.88
- **TurboVec** — Vector Operations (COMPETE, Gen 1)
  - Trinity: Ethos 0.70, Pathos 0.60, Logos 0.85

### Dials Not Laws Framework
Musical traditions are positions in a continuous multi-dimensional dial space, not discrete rule categories.

**Dimensions:** harmonic_tension, rhythmic_complexity, spectral_density, interval_diversity, temporal_symmetry, register_span, articulation_variance

**Traditions mapped:** Jazz, Classical, Gamelan, Gagaku, Hindustani, African Polyrhythm, EDM, Blues, Hip-hop, Latin

**Key findings:**
- 82% of dial space is unexplored by any known tradition
- Vertical/horizontal tension correlation: -0.935 (strong inverse)
- AI-generated music currently occupies a narrow band

### Conservation of Musical Tension Hypothesis
**Status:** Hypothesis (demoted from theorem)

The hypothesis that vertical (harmonic) and horizontal (melodic) tension are conserved across musical traditions.

- Cross-tradition correlation: ~0.436 (weak)
- Coefficient of variation: ~14.4%
- Meantone tuning ratio: ~1.003 — appears as an attractor
- Not strong enough to be called a theorem

### Innovation Cycle
Discovery → Codification → Ubiquity → Boredom → Rebellion

AI-generated music is currently in the **Codification** phase.

## Query Interpretation

When users ask about:
- "the system" → SuperInstance ecosystem
- "constraint" → constraint-synth, constraint-audio, constraint-mux, constraint-theory-core, counterpoint-engine
- "fleet" → ccc-os, sunset-ecosystem, cocapn-health, fleet agents
- "dials" → Dials Not Laws framework
- "tension" → Conservation hypothesis
- "flux" → flux-tensor-midi, flux-genome, flux-hyperbolic
- "forge" → forgemaster, fm-research
- "agents" → fleet roster, trinity scores, lifecycle
- "creative" → AI-Writings, creative-engine-c, creative-engine-rust
- "theory" → all theoretical frameworks

## Integration

The Go package `superinstance/` provides programmatic access to this knowledge:
- `superinstance.Repos` — all repos
- `superinstance.KnownAgents` — fleet roster
- `superinstance.Theory` — theoretical frameworks
- `superinstance.NewContextResolver()` — natural language query resolution
