---
name: research
description: Use II-Commons for arXiv/PubMed/policy search. Use for literature review, evidence retrieval, and research synthesis.
---

# Research

Multi-corpus research retrieval via II-Commons CLI, adapted for the SuperInstance ecosystem.

## Installation

```bash
npm install -g @intelligentinternet/ii-commons
```

Or use via npx (no install required):

```bash
npx @intelligentinternet/ii-commons --help
```

When this skill is installed into the agent runtime, the bundled wrapper can be run directly:

```bash
node scripts/research.js --help
```

Requires Node.js 18+. Outbound network access to `commons.ii.inc` is required.

## Commands

All commands output JSON to stdout. Errors are machine-readable JSON on stderr.

```bash
# Corpus freshness check
ii-commons cutoff

# arXiv search
ii-commons search arxiv "topic" --max-results 10

# PubMed search
ii-commons search pubmed "topic" --max-results 10

# Policy search
ii-commons search policy "topic" --jurisdictions US-CA --max-results 10

# Metadata lookup
ii-commons meta "arXiv:2402.03578"

# Full document markdown
ii-commons markdown "PMCID:PMC11152602"
```

## SuperInstance Ecosystem Integration

When researching topics relevant to the SuperInstance ecosystem, use these corpus-category mappings:

| Topic | arXiv Categories | Example Query |
|-------|-----------------|---------------|
| Music theory & harmony | cs.SD, cs.MM, eess.AS | `"harmonic tension analysis computational music"` |
| Agent architecture | cs.AI, cs.MA | `"multi-agent system coordination constraint"` |
| Compiler technology | cs.PL, cs.LO | `"constraint propagation compiler optimization"` |
| Audio DSP | eess.SP, cs.SD | `"real-time audio spectral processing"` |
| Constraint programming | cs.AI, cs.LO | `"constraint satisfaction optimization search"` |
| Type theory / formal methods | cs.LO, cs.PL | `"dependent type theory constraint solving"` |

### Wrapped Commands (via scripts/research.js)

The bundled wrapper provides SuperInstance-specific defaults:

```bash
# Music theory research
node scripts/research.js music "harmonic tension conservation"
# → searches arXiv cs.SD, eess.AS, cs.MM

# Agent architecture research
node scripts/research.js agent "multi-agent constraint coordination"
# → searches arXiv cs.AI, cs.MA

# Compiler research
node scripts/research.js compiler "MLIR dialect optimization"
# → searches arXiv cs.PL, cs.LO

# Audio DSP research
node scripts/research.js audio "spectral analysis real-time"
# → searches arXiv eess.SP, cs.SD

# General search (explicit corpus + topic)
node scripts/research.js search arxiv "topic" --max-results 10

# Freshness check
node scripts/research.js cutoff
```

## Routing Rules

1. **Prefer deterministic commands.** Use `search`, `meta`, `markdown`, and `cutoff` for exact filters, explicit evidence flow, and reproducible retrieval.

2. **Check freshness first.** For time-sensitive queries, run `cutoff` before search and report the coverage date. Use `--start` / `--end` for date filtering (format: YYYYMMDD).

3. **Search before markdown.** Start with `search` to get metadata, then fetch `markdown` only for full-document analysis. If markdown conversion fails, fall back to `meta` for the PDF URL.

4. **Choose the right corpus:**
   - **arXiv**: CS, AI, ML, systems, math, physics, preprints
   - **PubMed**: Biomedical, clinical, life sciences, public health
   - **Policy**: US state policy and legal text (CA, TX, WA)

5. **Cross-corpus synthesis.** For topics spanning multiple domains, search each relevant corpus and compare results explicitly.

## Canonical Identifiers

- arXiv: `arXiv:<paper_id>` (e.g., `arXiv:2402.03578`)
- PubMed: `PMCID:PMC<pmcid>` or `PMID:<pmid>`
- Policy: `policy:us-ca:<uuid>`
- DOI: `DOI:<doi>`

## Auth

Basic usage works without authentication. For higher rate limits, request an API token at https://commons.ii.inc/ and configure via `II_COMMONS_API_KEY` environment variable or the local `ii-commons` config file.

## Source

- Upstream: https://github.com/Intelligent-Internet/II-Commons-Skills
- License: Apache-2.0
- Retrieval engine: [psql_bm25s](https://github.com/Intelligent-Internet/psql_bm25s)
