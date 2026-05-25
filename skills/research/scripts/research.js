#!/usr/bin/env node
//
// research.js — SuperInstance research wrapper around II-Commons CLI
//
// Provides domain-specific shortcuts for ecosystem-relevant research topics
// with sensible defaults for arXiv category filtering.
//
// Usage:
//   node scripts/research.js cutoff
//   node scripts/research.js search arxiv "topic" --max-results 10
//   node scripts/research.js music "harmonic tension"
//   node scripts/research.js agent "multi-agent coordination"
//   node scripts/research.js compiler "MLIR dialect"
//   node scripts/research.js audio "spectral analysis"
//   node scripts/research.js constraint "constraint satisfaction"
//

"use strict";

const { execFile } = require("node:child_process");
const path = require("node:path");

// ── Domain presets ──────────────────────────────────────────────────────────
// Maps shorthand domain names to arXiv category lists and default max-results.
const DOMAINS = {
  music: {
    categories: ["cs.SD", "cs.MM", "eess.AS"],
    corpus: "arxiv",
    maxResults: 10,
    description: "Music theory, harmony, sound computing",
  },
  agent: {
    categories: ["cs.AI", "cs.MA"],
    corpus: "arxiv",
    maxResults: 10,
    description: "Agent architecture, multi-agent systems",
  },
  compiler: {
    categories: ["cs.PL", "cs.LO"],
    corpus: "arxiv",
    maxResults: 10,
    description: "Compilers, programming languages, MLIR",
  },
  audio: {
    categories: ["eess.SP", "cs.SD"],
    corpus: "arxiv",
    maxResults: 10,
    description: "Audio DSP, signal processing",
  },
  constraint: {
    categories: ["cs.AI", "cs.LO"],
    corpus: "arxiv",
    maxResults: 10,
    description: "Constraint programming, optimization",
  },
};

// ── Resolve the ii-commons binary ──────────────────────────────────────────
function resolveBinary() {
  // Try local node_modules first, then global, then npx fallback
  const localBin = path.join(
    __dirname,
    "..",
    "node_modules",
    ".bin",
    "ii-commons"
  );
  try {
    require("fs").accessSync(localBin);
    return localBin;
  } catch {
    // Fall through
  }

  // Check if globally installed
  const { execSync } = require("node:child_process");
  try {
    const globalPath = execSync("which ii-commons 2>/dev/null", {
      encoding: "utf8",
    }).trim();
    if (globalPath) return globalPath;
  } catch {
    // Fall through
  }

  // Fall back to npx invocation
  return null;
}

// ── Build and execute an ii-commons command ─────────────────────────────────
function run(args) {
  const binary = resolveBinary();
  const binArgs = binary ? [binary] : [];
  const cmd = binary || "npx";
  const finalArgs = binary ? args : ["@intelligentinternet/ii-commons", ...args];

  const child = execFile(cmd, finalArgs, {
    maxBuffer: 10 * 1024 * 1024, // 10 MB for large results
    timeout: 120_000,
    env: { ...process.env },
  });

  child.stdout.on("data", (d) => process.stdout.write(d));
  child.stderr.on("data", (d) => process.stderr.write(d));
  child.on("exit", (code) => process.exit(code ?? 0));
}

// ── Domain shortcut handler ─────────────────────────────────────────────────
function domainSearch(domain, query, extraArgs = []) {
  const preset = DOMAINS[domain];
  if (!preset) {
    process.stderr.write(
      JSON.stringify({
        error: `Unknown domain "${domain}". Available: ${Object.keys(DOMAINS).join(", ")}`,
      }) + "\n"
    );
    process.exit(1);
  }

  const args = [
    "search",
    preset.corpus,
    query,
    "--categories",
    preset.categories.join(","),
    "--max-results",
    String(preset.maxResults),
    ...extraArgs,
  ];

  run(args);
}

// ── CLI entry point ─────────────────────────────────────────────────────────
function main() {
  const argv = process.argv.slice(2);
  if (argv.length === 0) {
    console.log(
      [
        "SuperInstance Research Wrapper — II-Commons CLI",
        "",
        "Usage:",
        "  node scripts/research.js <command> [args...]",
        "",
        "Domain shortcuts (search arXiv with preset categories):",
        ...Object.entries(DOMAINS).map(
          ([name, p]) => `  ${name} "query"        — ${p.description}`
        ),
        "",
        "Direct commands (passed through to ii-commons):",
        "  cutoff                          — corpus freshness dates",
        '  search <corpus> "topic" [opts]  — direct search',
        '  meta "arXiv:ID"                 — metadata lookup',
        '  markdown "PMCID:ID"             — full document',
        "",
        "Extra options forwarded: --max-results, --start, --end, etc.",
      ].join("\n")
    );
    process.exit(0);
  }

  const command = argv[0];

  // Domain shortcuts
  if (DOMAINS[command]) {
    if (argv.length < 2) {
      process.stderr.write(
        JSON.stringify({ error: `Usage: research.js ${command} "query"` }) +
          "\n"
      );
      process.exit(1);
    }
    domainSearch(command, argv[1], argv.slice(2));
    return;
  }

  // Pass-through to ii-commons
  run(argv);
}

main();
