# SuperInstance Fork & Integration Strategy

## Overview

This document outlines what we adopt from II-Commons-Skills and LLVM, what we study deeper, and concrete next steps for integration into the SuperInstance ecosystem.

---

## 1. What We Adopt from II-Commons-Skills

### 1.1 SKILL.md Format

**Adopt: Yes, directly.**

The SKILL.md frontmatter + routing rules pattern is clean and agent-friendly:

```yaml
---
name: research
description: One-line description for agent routing.
---
```

**What we use it for:**
- All SuperInstance skills should use this format
- The `description` field doubles as the agent's routing hint
- Routing rules (prefer deterministic, freshness checks, corpus selection) are expressed as natural-language guidelines

**Implemented in:** `skills/research/SKILL.md`

### 1.2 Agent YAML Frontmatter

**Adopt: Yes.**

The `name` + `description` frontmatter pattern is minimal and effective. We extend it with SuperInstance-specific fields:

```yaml
---
name: research
description: Use II-Commons for arXiv/PubMed/policy search.
ecosystem: superinstance
domains: [music, agent, compiler, audio, constraint]
---
```

### 1.3 CLI Structure

**Adopt: Yes, with wrapping.**

II-Commons uses a simple `command subcommand "query" [filters]` pattern:

```
ii-commons search arxiv "topic" --max-results 10
ii-commons cutoff
ii-commons meta "arXiv:ID"
```

We wrap this with domain-specific shortcuts in `scripts/research.js`:

```
node scripts/research.js music "harmonic tension"
node scripts/research.js agent "multi-agent coordination"
```

**Why wrap instead of use directly:** Domain shortcuts add arXiv category filters automatically, reducing cognitive load for the agent.

### 1.4 Error Handling Pattern

**Adopt: Yes.**

II-Commons writes JSON to stdout, errors as JSON on stderr. This is machine-readable and agent-friendly. Our wrapper preserves this pattern.

### 1.5 Auth Strategy

**Adopt: Yes.**

Three-tier auth: env var → config file → anonymous. Clean and flexible.

---

## 2. What We Study Deeper from LLVM/MLIR

### 2.1 MLIR Dialect Definition (Priority: HIGH)

**Study:** How to define custom MLIR dialects with ops, types, and attributes.

**Why:** A constraint dialect is the most impactful LLVM integration for SuperInstance. It would let us:
- Compile constraint specifications to optimized machine code
- Verify constraint satisfaction at the IR level
- Compose constraint analysis with standard compiler optimizations

**Next steps:**
- Work through MLIR tutorial: https://mlir.llvm.org/docs/Tutorials/UnderstandingTheIRStructure/
- Study existing dialects: `affine`, `linalg`, `tosa` as reference implementations
- Prototype `constraint` dialect with 5-10 core ops

### 2.2 LLVM ORC JIT v2 (Priority: HIGH)

**Study:** Runtime compilation using ORC JIT for live music performance.

**Why:** Live music requires recompiling constraint parameters in real-time. JIT compilation with <10ms latency would enable:
- On-the-fly dial position changes
- Live tradition switching
- Real-time constraint reweighting

**Next steps:**
- Build minimal JIT example with ORC v2
- Benchmark compilation + execution latency
- Test with constraint-audio real-time paths

### 2.3 LLVM Pass Infrastructure (Priority: MEDIUM)

**Study:** New pass manager for custom optimization passes.

**Why:** Constraint-guided optimization can improve audio processing performance by:
- Eliminating unnecessary constraint checks at compile time
- Unrolling loops with known constraint bounds
- Hoisting invariant constraint computations

**Next steps:**
- Implement a minimal analysis pass that reads function attributes
- Test with constraint-audio Rust crate (via `RUSTFLAGS`)

### 2.4 LLVM BOLT (Priority: LOW)

**Study:** Post-link binary optimization.

**Why:** Fleet binaries deployed to many instances can benefit from profile-guided layout optimization. BOLT can improve instruction cache behavior for hot audio processing paths.

**Next steps:**
- Collect performance profiles from production fleet
- Apply BOLT optimization to release binaries
- Measure improvement in audio processing latency

### 2.5 SIMD Vectorization (Priority: MEDIUM)

**Study:** LLVM's auto-vectorization and explicit SIMD for audio buffers.

**Why:** Real-time audio processing is embarrassingly parallel on fixed-size buffers. LLVM can auto-vectorize, but explicit guidance improves results.

**Next steps:**
- Benchmark auto-vectorized vs hand-written SIMD in constraint-audio
- Profile with `llvm-mca` to identify bottlenecks
- Add `#[target_feature]` annotations to hot paths

---

## 3. Concrete Next Steps

### Immediate (This Week)

1. **✅ DONE**: Create `skills/research/SKILL.md` with II-Commons integration
2. **✅ DONE**: Create `skills/research/scripts/research.js` wrapper
3. **✅ DONE**: Write LLVM-INTEGRATION.md analysis
4. **✅ DONE**: Write this FORK-STRATEGY.md
5. **TODO**: Install `@intelligentinternet/ii-commons` in dev environment
6. **TODO**: Test research wrapper against live arXiv queries

### Short-term (Next 2 Weeks)

7. **Build LLVM toolchain** — Set up clang/LLVM build environment for fleet C/Rust code
8. **PGO for constraint-audio** — Profile-guided optimization of audio hot paths
9. **Prototype MLIR dialect** — Minimal constraint dialect with 3-5 ops
10. **ORC JIT spike** — Prove <10ms compilation for simple constraint functions

### Medium-term (Next Month)

11. **Constraint metadata pipeline** — Go → LLVM attribute emission from constraint specs
12. **Audio SIMD optimization** — Vectorize critical paths in constraint-audio
13. **Custom LLVM pass** — Constraint-aware optimization pass for audio functions
14. **Integration testing** — End-to-end: constraint spec → compiled → running in fleet

### Long-term (Next Quarter)

15. **Full constraint dialect** — All ops from constraint theory formalized in MLIR
16. **JIT live performance** — Real-time recompilation during live music sets
17. **BOLT optimization** — Post-link optimization for production fleet binaries
18. **WASM target** — Compile constraints to WebAssembly for browser-based instances

---

## 4. Repository Structure

```
SuperInstance/openagent/
├── skills/
│   └── research/
│       ├── SKILL.md              # ✅ Research skill definition
│       └── scripts/
│           └── research.js       # ✅ Domain-specific wrapper
├── docs/
│   ├── LLVM-INTEGRATION.md       # ✅ LLVM analysis
│   └── FORK-STRATEGY.md          # ✅ This document
├── superinstance/
│   ├── constraint-dialect/       # 📋 Future: MLIR dialect
│   ├── constraint-passes/        # 📋 Future: LLVM passes
│   └── jit-runtime/              # 📋 Future: ORC JIT integration
└── ...
```

---

## 5. Dependencies

| Dependency | Version | Purpose | Required For |
|-----------|---------|---------|-------------|
| `@intelligentinternet/ii-commons` | latest | Research CLI | Skills (now) |
| LLVM | 18+ | Compiler infrastructure | Phases 2-4 |
| MLIR | (bundled with LLVM) | Dialect definition | Phase 3 |
| Rust nightly | latest | `#[llvm_fn_attrs]` | Phase 2+ |
| `llvm-profdata` | (bundled with LLVM) | PGO data merging | Phase 1+ |
| `llvm-bolt` | (bundled with LLVM 17+) | Post-link optimization | Phase 4 |

---

*Strategy document for the SuperInstance ecosystem — practical integration, not theoretical purity.*
