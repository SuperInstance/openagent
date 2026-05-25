// Copyright 2026 The SuperInstance Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package superinstance

import (
	"math"
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
		// New repos added this session
		"flux-algebra", "constraint-dialect", "flux-julia", "agent-operations",
		"constraint-viz", "flux-rust", "tonal-archaeology", "fleet-conservation",
		"dial-theory", "harmonic-atlas", "spectral-taxonomy",
		"breeding-grounds", "polyglot-bridge",
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

// --- Dial Theory tests ---

func TestTraditionsExist(t *testing.T) {
	expected := []string{"Jazz", "Classical", "Gamelan", "Gagaku", "Hindustani",
		"African Polyrhythm", "EDM", "Blues", "Hip-hop", "Latin"}
	for _, name := range expected {
		tr, ok := Traditions[name]
		if !ok {
			t.Errorf("missing tradition: %s", name)
			continue
		}
		if tr.Name != name {
			t.Errorf("tradition %s: Name mismatch", name)
		}
		if tr.HarmonicTension < 0 || tr.HarmonicTension > 5 {
			t.Errorf("tradition %s: HarmonicTension %.2f out of range [0,5]", name, tr.HarmonicTension)
		}
		if tr.RhythmicComplexity < 0 || tr.RhythmicComplexity > 5 {
			t.Errorf("tradition %s: RhythmicComplexity %.2f out of range [0,5]", name, tr.RhythmicComplexity)
		}
		if tr.SpectralDensity < 0 || tr.SpectralDensity > 5 {
			t.Errorf("tradition %s: SpectralDensity %.2f out of range [0,5]", name, tr.SpectralDensity)
		}
	}
}

func TestDialDistanceIdentical(t *testing.T) {
	jazz := Traditions["Jazz"]
	d := DialDistance(jazz, jazz)
	if d != 0 {
		t.Errorf("distance from a tradition to itself should be 0, got %.4f", d)
	}
}

func TestDialDistanceSymmetric(t *testing.T) {
	a := Traditions["Jazz"]
	b := Traditions["Gamelan"]
	ab := DialDistance(a, b)
	ba := DialDistance(b, a)
	if math.Abs(ab-ba) > 1e-10 {
		t.Errorf("DialDistance not symmetric: %.4f vs %.4f", ab, ba)
	}
}

func TestDialDistancePositive(t *testing.T) {
	for aname, a := range Traditions {
		for bname, b := range Traditions {
			if aname >= bname {
				continue
			}
			d := DialDistance(a, b)
			if d < 0 {
				t.Errorf("distance between %s and %s is negative: %.4f", aname, bname, d)
			}
		}
	}
}

func TestNearestTradition(t *testing.T) {
	// A position identical to Jazz should resolve to Jazz.
	jazz := Traditions["Jazz"]
	if got := NearestTradition(jazz); got != "Jazz" {
		t.Errorf("NearestTradition(Jazz position) = %q, want Jazz", got)
	}
	// Slightly perturbed position near Classical should still be Classical.
	classical := Traditions["Classical"]
	perturbed := Tradition{
		Name:               "test",
		HarmonicTension:    classical.HarmonicTension + 0.05,
		RhythmicComplexity: classical.RhythmicComplexity + 0.05,
		SpectralDensity:    classical.SpectralDensity + 0.05,
	}
	if got := NearestTradition(perturbed); got != "Classical" {
		t.Errorf("NearestTradition(near Classical) = %q, want Classical", got)
	}
}

func TestNearestTraditionReturnsKnown(t *testing.T) {
	pos := Tradition{Name: "unknown", HarmonicTension: 2.5, RhythmicComplexity: 3.0, SpectralDensity: 2.5}
	result := NearestTradition(pos)
	if _, ok := Traditions[result]; !ok {
		t.Errorf("NearestTradition returned unknown tradition %q", result)
	}
}

// --- Fleet Conservation tests ---

func TestConservationExpected(t *testing.T) {
	gamma, h := ConservationExpected(4)
	if gamma <= 0 || h <= 0 {
		t.Errorf("ConservationExpected(4) = (%.4f, %.4f), both should be positive", gamma, h)
	}
}

func TestConservationExpectedZero(t *testing.T) {
	gamma, h := ConservationExpected(0)
	if gamma != 0 || h != 0 {
		t.Errorf("ConservationExpected(0) = (%.4f, %.4f), want (0, 0)", gamma, h)
	}
}

func TestConservationTotal(t *testing.T) {
	// γ + H should equal 1.283 - 0.159 * ln(4) for fleet size 4
	total := ConservationTotal(4)
	expected := 1.283 - 0.159*math.Log(4)
	if math.Abs(total-expected) > 1e-10 {
		t.Errorf("ConservationTotal(4) = %.6f, want %.6f", total, expected)
	}
}

func TestConservationTotalDecreasesWithSize(t *testing.T) {
	t1 := ConservationTotal(2)
	t2 := ConservationTotal(10)
	t3 := ConservationTotal(50)
	if t1 <= t2 || t2 <= t3 {
		t.Errorf("conservation total should decrease with fleet size: t(2)=%.4f, t(10)=%.4f, t(50)=%.4f", t1, t2, t3)
	}
}

func TestConservationDeviation(t *testing.T) {
	// Exactly matching expected values should give deviation ~0
	gamma, h := ConservationExpected(4)
	dev := ConservationDeviation(4, gamma, h)
	if math.Abs(dev) > 0.01 {
		t.Errorf("ConservationDeviation with exact values = %.4f, want ~0", dev)
	}
}

func TestConservationDeviationSignificant(t *testing.T) {
	// Wildly off values should give large deviation
	dev := ConservationDeviation(4, 10.0, 10.0)
	if dev < 2 {
		t.Errorf("ConservationDeviation with extreme values = %.4f, expected > 2", dev)
	}
}

func TestConservationDeviationZero(t *testing.T) {
	dev := ConservationDeviation(0, 1.0, 1.0)
	if dev != 0 {
		t.Errorf("ConservationDeviation(0, ...) = %.4f, want 0", dev)
	}
}

func TestNewReposInDependencyGraph(t *testing.T) {
	graph := DependencyGraph()
	// Verify new repos are in the graph
	for _, name := range []string{"flux-algebra", "constraint-dialect", "flux-julia", "agent-operations"} {
		if _, ok := graph[name]; !ok {
			t.Errorf("new repo %s missing from dependency graph", name)
		}
	}
	// flux-algebra depends on flux-tensor-midi
	deps := graph["flux-algebra"]
	found := false
	for _, d := range deps {
		if d == "flux-tensor-midi" {
			found = true
			break
		}
	}
	if !found {
		t.Error("flux-algebra should depend on flux-tensor-midi")
	}
}
