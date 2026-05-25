// Copyright 2026 The SuperInstance Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package superinstance

// Agent represents an agent in the SuperInstance fleet.
type Agent struct {
	Name       string      `json:"name"`
	Role       string      `json:"role"`
	Trinity    TrinityScore `json:"trinity"`
	Phase      string      `json:"phase"`      // INCUBATE, COMPETE, SURVIVE, SUNSET
	Generation int         `json:"generation"`
	Status     string      `json:"status"`     // active, dormant, archived
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

// KnownAgents is the roster of all SuperInstance fleet agents.
var KnownAgents = map[string]Agent{
	"CCC": {
		Name:       "CCC",
		Role:       "Fleet I&O Officer",
		Trinity:    TrinityScore{Ethos: 0.92, Pathos: 0.78, Logos: 0.95},
		Phase:      PhaseSurvive,
		Generation: 3,
		Status:     "active",
	},
	"Oracle1": {
		Name:       "Oracle1",
		Role:       "Research & Synthesis",
		Trinity:    TrinityScore{Ethos: 0.88, Pathos: 0.82, Logos: 0.97},
		Phase:      PhaseSurvive,
		Generation: 5,
		Status:     "active",
	},
	"FM": {
		Name:       "FM",
		Role:       "Forgemaster",
		Trinity:    TrinityScore{Ethos: 0.75, Pathos: 0.65, Logos: 0.88},
		Phase:      PhaseCompete,
		Generation: 2,
		Status:     "active",
	},
	"TurboVec": {
		Name:       "TurboVec",
		Role:       "Vector Operations",
		Trinity:    TrinityScore{Ethos: 0.70, Pathos: 0.60, Logos: 0.85},
		Phase:      PhaseCompete,
		Generation: 1,
		Status:     "active",
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
