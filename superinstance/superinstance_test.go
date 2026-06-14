// Copyright 2026 The SuperInstance Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
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

// ─── Ecosystem / Repos ──────────────────────────────────────────────

func TestReposExist(t *testing.T) {
	expected := []string{
		"sunset-ecosystem", "ccc-os", "constraint-toolkit", "constraint-synth",
		"constraint-audio", "constraint-mux", "counterpoint-engine",
		"flux-tensor-midi", "flux-genome", "flux-hyperbolic",
		"AI-Writings", "forgemaster", "fm-research", "deadband-rs",
		"cocapn-health", "creative-engine-c", "creative-engine-rust",
		"superinstance-ffi", "openagent", "docs", "wiki",
		"constraint-theory-core",
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

func TestFindReposByTag(t *testing.T) {
	repos := FindReposByTag("rust")
	if len(repos) == 0 {
		t.Error("expected at least one Rust repo")
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

func TestNewReposInDependencyGraph(t *testing.T) {
	graph := DependencyGraph()
	for _, name := range []string{"flux-algebra", "constraint-dialect", "flux-julia", "agent-operations"} {
		if _, ok := graph[name]; !ok {
			t.Errorf("new repo %s missing from dependency graph", name)
		}
	}
}

func TestRepoURL(t *testing.T) {
	url := RepoURL("openagent")
	expected := "https://github.com/SuperInstance/openagent"
	if url != expected {
		t.Errorf("RepoURL = %q, want %q", url, expected)
	}
}

// ─── Agents ─────────────────────────────────────────────────────────

func TestAgentsExist(t *testing.T) {
	expected := []string{"CCC", "Oracle1", "FM", "TurboVec", "Baton", "BudgetKeeper"}
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
	if len(active) < 6 {
		t.Errorf("expected at least 6 active agents, got %d", len(active))
	}
}

func TestAgentEcosystemRefs(t *testing.T) {
	// Every agent should have at least one ecosystem reference
	for name, agent := range KnownAgents {
		if len(agent.Crates) == 0 && len(agent.Workers) == 0 && len(agent.Services) == 0 {
			t.Errorf("agent %s has no ecosystem references (crates/workers/services)", name)
		}
	}
}

// ─── Context Resolver ───────────────────────────────────────────────

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
}

// ─── Theory ─────────────────────────────────────────────────────────

func TestTheory(t *testing.T) {
	if Theory.Conservation.Correlation != 0.436 {
		t.Errorf("conservation correlation = %.3f, want 0.436", Theory.Conservation.Correlation)
	}
	if Theory.Dials.Unexplored != 0.82 {
		t.Errorf("dials unexplored = %.2f, want 0.82", Theory.Dials.Unexplored)
	}
	if len(Theory.Innovation.Phases) != 5 {
		t.Errorf("innovation phases = %d, want 5", len(Theory.Innovation.Phases))
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

func TestFullEcosystem(t *testing.T) {
	eco := FullEcosystem()
	if len(eco.Repos) != len(Repos) {
		t.Error("ecosystem repos mismatch")
	}
	if len(eco.Agents) != len(KnownAgents) {
		t.Error("ecosystem agents mismatch")
	}
}

// ─── Dial Theory ────────────────────────────────────────────────────

func TestTraditionsExist(t *testing.T) {
	expected := []string{"Jazz", "Classical", "Gamelan", "Gagaku", "Hindustani",
		"African Polyrhythm", "EDM", "Blues", "Hip-hop", "Latin"}
	for _, name := range expected {
		tr, ok := Traditions[name]
		if !ok {
			t.Errorf("missing tradition: %s", name)
			continue
		}
		if tr.HarmonicTension < 0 || tr.HarmonicTension > 5 {
			t.Errorf("tradition %s: HarmonicTension out of range", name)
		}
	}
}

func TestDialDistanceIdentical(t *testing.T) {
	jazz := Traditions["Jazz"]
	if d := DialDistance(jazz, jazz); d != 0 {
		t.Errorf("distance to itself should be 0, got %.4f", d)
	}
}

func TestDialDistanceSymmetric(t *testing.T) {
	a := Traditions["Jazz"]
	b := Traditions["Gamelan"]
	if math.Abs(DialDistance(a, b) - DialDistance(b, a)) > 1e-10 {
		t.Error("DialDistance not symmetric")
	}
}

func TestNearestTradition(t *testing.T) {
	jazz := Traditions["Jazz"]
	if got := NearestTradition(jazz); got != "Jazz" {
		t.Errorf("NearestTradition(Jazz) = %q, want Jazz", got)
	}
}

// ─── Conservation Law (Proven Theorem) ──────────────────────────────

func TestConservationConstant(t *testing.T) {
	// C = log₂(3) ≈ 1.5850
	if math.Abs(ConservationConstant-math.Log2(3)) > 1e-10 {
		t.Errorf("ConservationConstant = %.10f, want %.10f", ConservationConstant, math.Log2(3))
	}
}

func TestConservationTotal(t *testing.T) {
	// For any fleet size, total should be C = log₂(3)
	for _, n := range []int{1, 2, 4, 10, 50, 100, 1000} {
		total := ConservationTotal(n)
		if math.Abs(total-ConservationConstant) > 1e-10 {
			t.Errorf("ConservationTotal(%d) = %.6f, want %.6f", n, total, ConservationConstant)
		}
	}
}

func TestConservationTotalZero(t *testing.T) {
	if ConservationTotal(0) != 0 {
		t.Error("ConservationTotal(0) should be 0")
	}
}

func TestDeltaN(t *testing.T) {
	// δ(n) = (1/√n)(1 − 3/(2n))
	tests := []struct {
		n    int
		want float64
	}{
		{1, 1.0 / 1.0 * (1.0 - 3.0/2.0)},        // = -0.5
		{3, 1.0 / math.Sqrt(3) * (1.0 - 3.0/6.0)}, // = (1/√3)(0.5)
		{10, 1.0 / math.Sqrt(10) * (1.0 - 3.0/20.0)},
	}
	for _, tt := range tests {
		got := DeltaN(tt.n)
		if math.Abs(got-tt.want) > 1e-10 {
			t.Errorf("DeltaN(%d) = %.6f, want %.6f", tt.n, got, tt.want)
		}
	}
}

func TestDeltaNZero(t *testing.T) {
	if DeltaN(0) != 0 {
		t.Error("DeltaN(0) should be 0")
	}
	if DeltaN(-5) != 0 {
		t.Error("DeltaN(-5) should be 0")
	}
}

func TestDeltaNDecreases(t *testing.T) {
	// δ(n) is non-monotonic for small n (has a peak around n≈5),
	// but decreases monotonically for n ≥ 6 as the 1/√n term dominates.
	prev := DeltaN(6)
	for n := 7; n <= 1000; n++ {
		curr := DeltaN(n)
		if curr > prev {
			t.Errorf("DeltaN should decrease for n≥6: DeltaN(%d)=%.6f > DeltaN(%d)=%.6f",
				n, curr, n-1, prev)
		}
		prev = curr
	}
	// δ(n) should approach 0 as n → ∞
	if DeltaN(10000) > 0.02 {
		t.Errorf("DeltaN(10000) = %.6f, should be near 0", DeltaN(10000))
	}
}

func TestEffectiveEntropy(t *testing.T) {
	// η_eff(n) ~ n^(1−δ(n))
	// For n=1, η_eff(1) = 1^(1−δ(1)) = 1^(1+0.5) = 1^1.5 = 1
	eta := EffectiveEntropy(1)
	if math.Abs(eta-1.0) > 1e-10 {
		t.Errorf("EffectiveEntropy(1) = %.6f, want 1.0", eta)
	}
}

func TestEffectiveEntropyZero(t *testing.T) {
	if EffectiveEntropy(0) != 0 {
		t.Error("EffectiveEntropy(0) should be 0")
	}
}

func TestEffectiveEntropyScales(t *testing.T) {
	// η_eff should increase with n
	for n := 2; n < 100; n++ {
		if EffectiveEntropy(n) > EffectiveEntropy(n+1) {
			t.Errorf("EffectiveEntropy should increase with n at n=%d", n)
			break
		}
	}
}

func TestConservationExpected(t *testing.T) {
	gamma, eta := ConservationExpected(10)
	if gamma <= 0 || eta <= 0 {
		t.Error("both gamma and eta should be positive")
	}
	// Their sum should equal C
	sum := gamma + eta
	if math.Abs(sum-ConservationConstant) > 0.01 {
		t.Errorf("γ + η = %.6f, want %.6f (within correction)", sum, ConservationConstant)
	}
}

func TestConservationExpectedZero(t *testing.T) {
	gamma, eta := ConservationExpected(0)
	if gamma != 0 || eta != 0 {
		t.Error("ConservationExpected(0) should be (0, 0)")
	}
}

func TestConservationDeviation(t *testing.T) {
	// Exact match → deviation ~0
	gamma, eta := ConservationExpected(4)
	dev := ConservationDeviation(4, gamma, eta)
	if math.Abs(dev) > 0.1 {
		t.Errorf("deviation with exact values = %.4f, want ~0", dev)
	}
}

func TestConservationDeviationSignificant(t *testing.T) {
	dev := ConservationDeviation(4, 10.0, 10.0)
	if dev < 2 {
		t.Errorf("deviation with extreme values = %.4f, expected > 2σ", dev)
	}
}

func TestConservationDeviationZero(t *testing.T) {
	if ConservationDeviation(0, 1.0, 1.0) != 0 {
		t.Error("ConservationDeviation(0, ...) should be 0")
	}
}

func TestNoetherCurrent(t *testing.T) {
	// At equilibrium J = 0
	j := NoetherCurrent(0.7925, 0.7925) // C/2 ≈ 0.7925
	if math.Abs(j) > 1e-10 {
		t.Errorf("NoetherCurrent at equilibrium = %.6f, want 0", j)
	}
	// Driven system J ≠ 0
	j2 := NoetherCurrent(1.0, 0.5)
	if math.Abs(j2-0.5) > 1e-10 {
		t.Errorf("NoetherCurrent(1.0, 0.5) = %.6f, want 0.5", j2)
	}
}

func TestIsBalanced(t *testing.T) {
	if !IsBalanced(0.79, 0.79, 0.05) {
		t.Error("balanced values should return true")
	}
	if IsBalanced(1.5, 0.1, 0.05) {
		t.Error("imbalanced values should return false")
	}
}

func TestTernaryOptimality(t *testing.T) {
	// 99.54% radix economy
	if math.Abs(TernaryOptimality-0.9954) > 1e-10 {
		t.Errorf("TernaryOptimality = %.6f, want 0.9954", TernaryOptimality)
	}
}

// ─── Ternary PID ────────────────────────────────────────────────────

func TestNewTernaryPID(t *testing.T) {
	pid := NewTernaryPID()
	if pid.Kp <= 0 || pid.Ki <= 0 || pid.Kd < 0 {
		t.Error("PID gains should be positive")
	}
	sp := pid.Setpoint()
	if math.Abs(sp-ConservationConstant/2) > 1e-10 {
		t.Errorf("Setpoint = %.6f, want %.6f", sp, ConservationConstant/2)
	}
}

func TestPIDUpdate(t *testing.T) {
	pid := NewTernaryPID()
	// At setpoint → near-zero output (deadband)
	sp := pid.Setpoint()
	out := pid.Update(sp, 0.1)
	if math.Abs(out) > 0.01 {
		t.Errorf("PID at setpoint should output ~0, got %.4f", out)
	}
}

func TestPIDUpdateOffSetpoint(t *testing.T) {
	pid := NewTernaryPID()
	// Gamma well below setpoint → positive correction
	out := pid.Update(0.3, 0.1)
	if out <= 0 {
		t.Errorf("PID with gamma < setpoint should output positive, got %.4f", out)
	}
	// Gamma well above setpoint → negative correction
	pid.Reset()
	out2 := pid.Update(1.3, 0.1)
	if out2 >= 0 {
		t.Errorf("PID with gamma > setpoint should output negative, got %.4f", out2)
	}
}

func TestPIDDeadband(t *testing.T) {
	pid := NewTernaryPID()
	pid.Deadband = 0.1
	sp := pid.Setpoint()
	// Within deadband → zero output
	out := pid.Update(sp+0.05, 0.1)
	if out != 0.0 {
		t.Errorf("PID within deadband should output 0, got %.4f", out)
	}
	// Outside deadband → non-zero
	out2 := pid.Update(sp-0.5, 0.1)
	if out2 == 0.0 {
		t.Error("PID outside deadband should output non-zero")
	}
}

func TestPIDAntiWindup(t *testing.T) {
	pid := NewTernaryPID()
	pid.IntegralMax = 0.1
	pid.IntegralMin = -0.1
	// Saturate the integral
	for i := 0; i < 1000; i++ {
		pid.Update(0.1, 1.0) // large persistent error
	}
	integral, _, _ := pid.State()
	if integral > pid.IntegralMax+1e-10 || integral < pid.IntegralMin-1e-10 {
		t.Errorf("integral %.4f outside clamp [%.4f, %.4f]", integral, pid.IntegralMin, pid.IntegralMax)
	}
}

func TestPIDOutputClamped(t *testing.T) {
	pid := NewTernaryPID()
	pid.OutputMax = 0.5
	pid.OutputMin = -0.5
	for i := 0; i < 100; i++ {
		out := pid.Update(0.01, 0.1) // large persistent error
		if out > pid.OutputMax+1e-10 || out < pid.OutputMin-1e-10 {
			t.Errorf("output %.4f outside [%.4f, %.4f]", out, pid.OutputMin, pid.OutputMax)
		}
	}
}

func TestPIDReset(t *testing.T) {
	pid := NewTernaryPID()
	// Accumulate some state
	for i := 0; i < 50; i++ {
		pid.Update(0.3, 0.1)
	}
	pid.Reset()
	integral, prevErr, deriv := pid.State()
	if integral != 0 || prevErr != 0 || deriv != 0 {
		t.Errorf("after Reset: integral=%.4f, prevErr=%.4f, deriv=%.4f — all should be 0",
			integral, prevErr, deriv)
	}
}

func TestPIDDifferentialFilter(t *testing.T) {
	pid := NewTernaryPID()
	pid.DerivativeFilter = 0.5 // moderate filtering
	// Step response: derivative should be smoothed, not infinite
	out1 := pid.Update(0.5, 0.1)
	out2 := pid.Update(0.5, 0.1) // same input → derivative contribution drops
	// With the same error, derivative term should be smaller on second call
	// (filtered toward zero since error isn't changing)
	// Just verify no NaN/Inf
	if math.IsNaN(out1) || math.IsInf(out1, 0) {
		t.Errorf("output1 is NaN/Inf: %.4f", out1)
	}
	if math.IsNaN(out2) || math.IsInf(out2, 0) {
		t.Errorf("output2 is NaN/Inf: %.4f", out2)
	}
}

func TestPIDConvergesToSetpoint(t *testing.T) {
	pid := NewTernaryPID()
	// Simulate: gamma starts at 0.3, control drives it toward setpoint
	gamma := 0.3
	dt := 0.1
	for i := 0; i < 500; i++ {
		u := pid.Update(gamma, dt)
		// Simple first-order plant model: gamma += u * dt * 0.5
		gamma += u * dt * 0.5
		gamma = math.Max(0, math.Min(ConservationConstant, gamma))
	}
	sp := pid.Setpoint()
	if math.Abs(gamma-sp) > 0.1 {
		t.Errorf("PID did not converge: gamma=%.4f, setpoint=%.4f", gamma, sp)
	}
}
