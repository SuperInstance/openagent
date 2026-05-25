// Copyright 2026 The SuperInstance Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with in writing, software
// distributed under the License on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package superinstance

import (
	"fmt"
	"strings"
)

// ContextResolver maps user queries to SuperInstance ecosystem knowledge.
// This enables the agent to natively understand references to "the system",
// "constraint", "fleet", "dials", etc.
type ContextResolver struct {
	ecosystem *Ecosystem
}

// NewContextResolver creates a resolver with the default ecosystem.
func NewContextResolver() *ContextResolver {
	return &ContextResolver{
		ecosystem: FullEcosystem(),
	}
}

// Query represents a resolved user query with ecosystem context.
type Query struct {
	Intent      string   `json:"intent"`       // What the user is asking about
	Repos       []Repo   `json:"repos"`        // Relevant repos
	Agents      []Agent  `json:"agents"`       // Relevant agents
	Protocols   []Protocol `json:"protocols"` // Relevant protocols
	TheoryContext string  `json:"theory_context"` // Related theoretical framework
	Suggestions []string `json:"suggestions"`  // Follow-up suggestions
}

// Resolve maps a natural language query to SuperInstance context.
func (c *ContextResolver) Resolve(query string) *Query {
	q := strings.ToLower(query)
	result := &Query{Intent: query}

	// Detect intent categories
	switch {
	case containsAny(q, "system", "ecosystem", "superinstance", "all repos", "overview"):
		c.resolveEcosystem(result)
	case containsAny(q, "constraint", "music theory", "synth", "counterpoint"):
		c.resolveConstraint(result)
	case containsAny(q, "fleet", "monitoring", "ccc", "operations", "health"):
		c.resolveFleet(result)
	case containsAny(q, "dials", "dial", "dimensions", "traditions", "framework"):
		c.resolveDials(result)
	case containsAny(q, "flux", "tensor", "midi", "genome", "hyperbolic"):
		c.resolveFlux(result)
	case containsAny(q, "forge", "forgemaster", "build", "deploy"):
		c.resolveForge(result)
	case containsAny(q, "agent", "trinity", "ethos", "pathos", "logos", "lifecycle", "phase"):
		c.resolveAgents(result)
	case containsAny(q, "conservation", "tension", "theorem", "hypothesis"):
		c.resolveConservation(result)
	case containsAny(q, "innovation", "cycle", "discovery", "codification", "ubiquity", "boredom", "rebellion"):
		c.resolveInnovation(result)
	case containsAny(q, "creative", "engine", "writing", "art"):
		c.resolveCreative(result)
	default:
		result.Suggestions = []string{
			"Try asking about: fleet, constraints, dials, flux, forge, agents, theory",
		}
	}

	return result
}

func (c *ContextResolver) resolveEcosystem(r *Query) {
	r.TheoryContext = fmt.Sprintf("SuperInstance ecosystem: %d repos, %d agents, %d protocols",
		len(c.ecosystem.Repos), len(c.ecosystem.Agents), len(c.ecosystem.Protocols))
	for _, repo := range c.ecosystem.Repos {
		r.Repos = append(r.Repos, repo)
	}
	for _, agent := range c.ecosystem.Agents {
		r.Agents = append(r.Agents, agent)
	}
	r.Suggestions = []string{"Ask about a specific repo", "Check fleet status", "Explore theory"}
}

func (c *ContextResolver) resolveConstraint(r *Query) {
	r.TheoryContext = Theory.Conservation.Description
	for _, name := range []string{"constraint-toolkit", "constraint-synth", "constraint-audio", "constraint-mux", "counterpoint-engine", "constraint-theory-core"} {
		if repo, ok := c.ecosystem.Repos[name]; ok {
			r.Repos = append(r.Repos, repo)
		}
	}
	r.Protocols = append(r.Protocols, c.ecosystem.Protocols["constraint-protocol"])
	r.Suggestions = []string{"Run constraint analysis", "Check counterpoint rules", "Explore constraint theory"}
}

func (c *ContextResolver) resolveFleet(r *Query) {
	r.TheoryContext = fmt.Sprintf("Fleet: %d active agents in phase %s",
		len(c.ecosystem.FleetState.ActiveAgents), c.ecosystem.FleetState.Phase)
	for _, name := range []string{"ccc-os", "sunset-ecosystem", "cocapn-health"} {
		if repo, ok := c.ecosystem.Repos[name]; ok {
			r.Repos = append(r.Repos, repo)
		}
	}
	r.Protocols = append(r.Protocols, c.ecosystem.Protocols["fleet-heartbeat"])
	for _, agent := range c.ecosystem.Agents {
		r.Agents = append(r.Agents, agent)
	}
	r.Suggestions = []string{"Check fleet health", "View agent trinity scores", "See dependency graph"}
}

func (c *ContextResolver) resolveDials(r *Query) {
	r.TheoryContext = Theory.Dials.Description
	r.Suggestions = []string{"Explore unexplored dial space", "Compare traditions", "Map tradition positions"}
}

func (c *ContextResolver) resolveFlux(r *Query) {
	r.TheoryContext = "Flux subsystem: tensor-based MIDI manipulation, genome mapping, hyperbolic exploration"
	for _, name := range []string{"flux-tensor-midi", "flux-genome", "flux-hyperbolic"} {
		if repo, ok := c.ecosystem.Repos[name]; ok {
			r.Repos = append(r.Repos, repo)
		}
	}
	r.Protocols = append(r.Protocols, c.ecosystem.Protocols["flux-protocol"])
	r.Suggestions = []string{"Analyze MIDI tensor", "Map genome", "Explore hyperbolic space"}
}

func (c *ContextResolver) resolveForge(r *Query) {
	r.TheoryContext = "Forgemaster: build and forge orchestration system"
	for _, name := range []string{"forgemaster", "fm-research"} {
		if repo, ok := c.ecosystem.Repos[name]; ok {
			r.Repos = append(r.Repos, repo)
		}
	}
	r.Protocols = append(r.Protocols, c.ecosystem.Protocols["forge-protocol"])
	r.Suggestions = []string{"Check build status", "View forge history"}
}

func (c *ContextResolver) resolveAgents(r *Query) {
	r.TheoryContext = "Trinity Architecture: Ethos (trust) × Pathos (engagement) × Logos (logic)"
	for _, agent := range c.ecosystem.Agents {
		r.Agents = append(r.Agents, agent)
	}
	r.Repos = append(r.Repos, c.ecosystem.Repos["sunset-ecosystem"])
	r.Suggestions = []string{"View trinity scores", "Check lifecycle phases", "See breeding history"}
}

func (c *ContextResolver) resolveConservation(r *Query) {
	r.TheoryContext = fmt.Sprintf("Conservation hypothesis: I_vert + I_horiz conserved? Correlation=%.3f, CV=%.1f%%, Status=%s",
		Theory.Conservation.Correlation, Theory.Conservation.CV, Theory.Conservation.Status)
	r.Suggestions = []string{"Run conservation analysis", "Cross-tradition comparison", "Check meantone attractor"}
}

func (c *ContextResolver) resolveInnovation(r *Query) {
	r.TheoryContext = fmt.Sprintf("Innovation cycle: %s. Current phase for AI music: %s",
		strings.Join(Theory.Innovation.Phases, " → "), Theory.Innovation.CurrentPhase)
	r.Suggestions = []string{"Map tradition positions in cycle", "Predict next rebellion"}
}

func (c *ContextResolver) resolveCreative(r *Query) {
	r.TheoryContext = "Creative subsystem: writing, engines, art generation"
	for _, name := range []string{"AI-Writings", "creative-engine-c", "creative-engine-rust"} {
		if repo, ok := c.ecosystem.Repos[name]; ok {
			r.Repos = append(r.Repos, repo)
		}
	}
	r.Suggestions = []string{"Browse AI writings", "Generate creative output", "Compare engine implementations"}
}

// Summary generates a human-readable summary of the entire ecosystem.
func (c *ContextResolver) Summary() string {
	var b strings.Builder
	b.WriteString("# SuperInstance Ecosystem\n\n")
	b.WriteString(fmt.Sprintf("## Repositories (%d)\n", len(c.ecosystem.Repos)))
	for name, repo := range c.ecosystem.Repos {
		b.WriteString(fmt.Sprintf("- **%s** (%s): %s\n", name, repo.Language, repo.Description))
	}
	b.WriteString(fmt.Sprintf("\n## Active Agents (%d)\n", len(c.ecosystem.Agents)))
	for _, agent := range c.ecosystem.Agents {
		avg := agent.Trinity.Average()
		b.WriteString(fmt.Sprintf("- **%s** [%s] Gen %d — Trinity: %.2f\n", agent.Name, agent.Role, agent.Generation, avg))
	}
	b.WriteString("\n## Key Theories\n")
	b.WriteString(fmt.Sprintf("- Conservation: %s (r=%.3f)\n", Theory.Conservation.Status, Theory.Conservation.Correlation))
	b.WriteString(fmt.Sprintf("- Dials: %d dimensions, %.0f%% unexplored\n", len(Theory.Dials.Dimensions), Theory.Dials.Unexplored*100))
	b.WriteString(fmt.Sprintf("- Innovation: %s (currently: %s)\n", strings.Join(Theory.Innovation.Phases, "→"), Theory.Innovation.CurrentPhase))
	return b.String()
}

// RepoURL returns the full GitHub URL for a repo name.
func RepoURL(name string) string {
	return fmt.Sprintf("https://github.com/%s/%s", Organization, name)
}

func containsAny(s string, terms ...string) bool {
	for _, t := range terms {
		if strings.Contains(s, t) {
			return true
		}
	}
	return false
}
