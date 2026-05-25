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
	"strings"
	"testing"
)

func TestReposExist(t *testing.T) {
	expected := []string{
		"sunset-ecosystem", "ccc-os", "constraint-toolkit", "constraint-synth",
		"constraint-audio", "constraint-mux", "counterpoint-engine",
		"flux-tensor-midi", "flux-genome", "flux-hyperbolic",
		"AI-Writings", "forgemaster", "fm-research", "deadband-rs",
		"cocapn-health", "creative-engine-c", "creative-engine-rust",
		"superinstance-ffi", "openagent", "docs", "wiki",
		"constraint-theory-core",
	}
	for _, name := range expected {
		repo, ok := Repos[name]
		if !ok {
			t.Errorf("missing repo: %s", name)
			continue
		}
		if repo.Name != name {
			t.Errorf("repo %s: Name mismatch", name)
		}
		if repo.Description == "" {
			t.Errorf("repo %s: empty Description", name)
		}
		if repo.Language == "" {
			t.Errorf("repo %s: empty Language", name)
		}
	}
}

func TestAgentsExist(t *testing.T) {
	expected := []string{"CCC", "Oracle1", "FM", "TurboVec"}
	for _, name := range expected {
		agent, ok := KnownAgents[name]
		if !ok {
			t.Errorf("missing agent: %s", name)
			continue
		}
		if agent.Phase != PhaseIncubate && agent.Phase != PhaseCompete &&
			agent.Phase != PhaseSurvive && agent.Phase != PhaseSunset {
			t.Errorf("agent %s: invalid phase %q", name, agent.Phase)
		}
		avg := agent.Trinity.Average()
		if avg < 0 || avg > 1 {
			t.Errorf("agent %s: trinity average %.2f out of range", name, avg)
		}
	}
}

func TestTrinityAverage(t *testing.T) {
	tr := TrinityScore{Ethos: 1.0, Pathos: 0.0, Logos: 0.5}
	if avg := tr.Average(); avg != 0.5 {
		t.Errorf("expected average 0.5, got %.2f", avg)
	}
}

func TestFindReposByTag(t *testing.T) {
	repos := FindReposByTag("rust")
	if len(repos) == 0 {
		t.Error("expected at least one Rust repo")
	}
	for _, r := range repos {
		if r.Language != "Rust" {
			t.Errorf("FindReposByTag(\"rust\") returned non-Rust repo: %s (%s)", r.Name, r.Language)
		}
	}
}

func TestFindReposByLanguage(t *testing.T) {
	repos := FindReposByLanguage("Go")
	found := false
	for _, r := range repos {
		if r.Name == "openagent" {
			found = true
			break
		}
	}
	if !found {
		t.Error("expected to find openagent in Go repos")
	}
}

func TestDependencyGraph(t *testing.T) {
	graph := DependencyGraph()
	if len(graph) == 0 {
		t.Error("dependency graph is empty")
	}
	// ccc-os depends on sunset-ecosystem
	deps, ok := graph["ccc-os"]
	if !ok {
		t.Error("ccc-os missing from dependency graph")
	}
	found := false
	for _, d := range deps {
		if d == "sunset-ecosystem" {
			found = true
			break
		}
	}
	if !found {
		t.Error("ccc-os should depend on sunset-ecosystem")
	}
}

func TestContextResolver(t *testing.T) {
	cr := NewContextResolver()

	tests := []struct {
		query    string
		wantRepo string
	}{
		{"constraint synth", "constraint-synth"},
		{"fleet status", "ccc-os"},
		{"flux tensor", "flux-tensor-midi"},
		{"forgemaster build", "forgemaster"},
		{"the system overview", "sunset-ecosystem"},
	}

	for _, tt := range tests {
		result := cr.Resolve(tt.query)
		found := false
		for _, r := range result.Repos {
			if r.Name == tt.wantRepo {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Resolve(%q): expected repo %s not found", tt.query, tt.wantRepo)
		}
	}
}

func TestContextResolverSummary(t *testing.T) {
	cr := NewContextResolver()
	summary := cr.Summary()
	if !strings.Contains(summary, "SuperInstance") {
		t.Error("summary missing SuperInstance")
	}
	if !strings.Contains(summary, "CCC") {
		t.Error("summary missing CCC agent")
	}
	if !strings.Contains(summary, "Conservation") {
		t.Error("summary missing Conservation theory")
	}
}

func TestTheory(t *testing.T) {
	if Theory.Conservation.Correlation != 0.436 {
		t.Errorf("conservation correlation = %.3f, want 0.436", Theory.Conservation.Correlation)
	}
	if Theory.Conservation.Status != "hypothesis" {
		t.Errorf("conservation status = %q, want \"hypothesis\"", Theory.Conservation.Status)
	}
	if Theory.Dials.Unexplored != 0.82 {
		t.Errorf("dials unexplored = %.2f, want 0.82", Theory.Dials.Unexplored)
	}
	if Theory.Dials.VKHCorrelation != -0.935 {
		t.Errorf("VKH correlation = %.3f, want -0.935", Theory.Dials.VKHCorrelation)
	}
	if len(Theory.Innovation.Phases) != 5 {
		t.Errorf("innovation phases = %d, want 5", len(Theory.Innovation.Phases))
	}
	if Theory.Innovation.CurrentPhase != "Codification" {
		t.Errorf("innovation current phase = %q, want \"Codification\"", Theory.Innovation.CurrentPhase)
	}
}

func TestPhaseTransitions(t *testing.T) {
	transitions := PhaseTransitions()
	if len(transitions[PhaseIncubate]) != 2 {
		t.Errorf("INCUBATE should have 2 transitions, got %d", len(transitions[PhaseIncubate]))
	}
	if len(transitions[PhaseSunset]) != 0 {
		t.Errorf("SUNSET should have 0 transitions, got %d", len(transitions[PhaseSunset]))
	}
}

func TestAgentsByPhase(t *testing.T) {
	compete := AgentsByPhase(PhaseCompete)
	if len(compete) == 0 {
		t.Error("expected at least one COMPETE agent")
	}
	survive := AgentsByPhase(PhaseSurvive)
	if len(survive) == 0 {
		t.Error("expected at least one SURVIVE agent")
	}
}

func TestActiveAgents(t *testing.T) {
	active := ActiveAgents()
	if len(active) == 0 {
		t.Error("expected active agents")
	}
}

func TestFullEcosystem(t *testing.T) {
	eco := FullEcosystem()
	if len(eco.Repos) != len(Repos) {
		t.Error("ecosystem repos mismatch")
	}
	if len(eco.Agents) != len(KnownAgents) {
		t.Error("ecosystem agents mismatch")
	}
}

func TestRepoURL(t *testing.T) {
	url := RepoURL("openagent")
	expected := "https://github.com/SuperInstance/openagent"
	if url != expected {
		t.Errorf("RepoURL = %q, want %q", url, expected)
	}
}
