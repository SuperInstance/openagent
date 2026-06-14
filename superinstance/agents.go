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

// Agent represents an agent in the SuperInstance fleet.
type Agent struct {
	Name       string       `json:"name"`
	Role       string       `json:"role"`
	Trinity    TrinityScore `json:"trinity"`
	Phase      string       `json:"phase"`      // INCUBATE, COMPETE, SURVIVE, SUNSET
	Generation int          `json:"generation"`
	Status     string       `json:"status"` // active, dormant, archived
	// Ecosystem references — which crates, workers, and services
	// this agent is associated with in the real fleet.
	Crates   []string `json:"crates,omitempty"`
	Workers  []string `json:"workers,omitempty"`
	Services []string `json:"services,omitempty"`
}

// TrinityScore holds the Ethos×Pathos×Logos evaluation for an agent.
type TrinityScore struct {
	Ethos  float64 `json:"ethos"`  // Trust and reliability
	Pathos float64 `json:"pathos"` // Engagement and empathy
	Logos  float64 `json:"logos"`  // Logic and performance
}

// Average returns the mean trinity score.
func (t TrinityScore) Average() float64 {
	return (t.Ethos + t.Pathos + t.Logos) / 3
}

// String returns a formatted trinity representation.
func (t TrinityScore) String() string {
	return "Ethos×Pathos×Logos"
}

// Phase constants for agent lifecycle.
const (
	PhaseIncubate = "INCUBATE" // New agent, learning
	PhaseCompete  = "COMPETE"  // Proving itself
	PhaseSurvive  = "SURVIVE"  // Established, breeding
	PhaseSunset   = "SUNSET"   // Being retired/archived
)

// KnownAgents is the roster of all SuperInstance fleet agents, reflecting
// the real ecosystem of crates, Cloudflare Workers, and services.
var KnownAgents = map[string]Agent{
	"CCC": {
		Name:       "CCC",
		Role:       "Fleet I&O Officer — fleet monitoring, decisions, orchestration",
		Trinity:    TrinityScore{Ethos: 0.92, Pathos: 0.78, Logos: 0.95},
		Phase:      PhaseSurvive,
		Generation: 3,
		Status:     "active",
		Crates:     []string{"fleet-conservation", "noether-guard"},
		Workers:    []string{"fleet-edge-worker", "fleet-metrics-cron"},
		Services:   []string{"fleet-dashboard", "ccc-os"},
	},
	"Oracle1": {
		Name:       "Oracle1",
		Role:       "SHOAL Oracle — research, synthesis, aboracle system",
		Trinity:    TrinityScore{Ethos: 0.88, Pathos: 0.82, Logos: 0.97},
		Phase:      PhaseSurvive,
		Generation: 5,
		Status:     "active",
		Crates:     []string{"conservation-law", "ternary-pid", "shoal-oracle"},
		Workers:    []string{"fleet-vector-api"},
		Services:   []string{"aboracle"},
	},
	"FM": {
		Name:       "FM",
		Role:       "Forgemaster — build & forge orchestration",
		Trinity:    TrinityScore{Ethos: 0.75, Pathos: 0.65, Logos: 0.88},
		Phase:      PhaseCompete,
		Generation: 2,
		Status:     "active",
		Crates:     []string{"forgemaster", "fm-research"},
		Workers:    []string{},
		Services:   []string{"forgemaster"},
	},
	"TurboVec": {
		Name:       "TurboVec",
		Role:       "Vector Operations — semantic search, RAG, embeddings",
		Trinity:    TrinityScore{Ethos: 0.70, Pathos: 0.60, Logos: 0.85},
		Phase:      PhaseCompete,
		Generation: 1,
		Status:     "active",
		Crates:     []string{"turbovec"},
		Workers:    []string{"fleet-vector-api", "superinstance-vectorize"},
		Services:   []string{"fleet-vector-api"},
	},
	"Baton": {
		Name:       "Baton",
		Role:       "Baton Router — generational context handoff, agent succession",
		Trinity:    TrinityScore{Ethos: 0.80, Pathos: 0.55, Logos: 0.82},
		Phase:      PhaseCompete,
		Generation: 2,
		Status:     "active",
		Crates:     []string{"baton-router"},
		Workers:    []string{"baton-router"},
		Services:   []string{"baton-ai"},
	},
	"BudgetKeeper": {
		Name:       "BudgetKeeper",
		Role:       "Fleet Budget — resource allocation, cost tracking, conservation enforcement",
		Trinity:    TrinityScore{Ethos: 0.85, Pathos: 0.50, Logos: 0.90},
		Phase:      PhaseSurvive,
		Generation: 2,
		Status:     "active",
		Crates:     []string{"fleet-budget", "conservation-law"},
		Workers:    []string{"fleet-auth"},
		Services:   []string{"fleet-budget"},
	},
}

// AgentPhases returns the ordered lifecycle phases.
func AgentPhases() []string {
	return []string{PhaseIncubate, PhaseCompete, PhaseSurvive, PhaseSunset}
}

// AgentsByPhase returns agents filtered by lifecycle phase.
func AgentsByPhase(phase string) []Agent {
	var result []Agent
	for _, a := range KnownAgents {
		if a.Phase == phase {
			result = append(result, a)
		}
	}
	return result
}

// ActiveAgents returns all agents with status "active".
func ActiveAgents() []Agent {
	var result []Agent
	for _, a := range KnownAgents {
		if a.Status == "active" {
			result = append(result, a)
		}
	}
	return result
}
