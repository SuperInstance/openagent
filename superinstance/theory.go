// Copyright 2026 The SuperInstance Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with in writing, software
// distributed under the License on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package superinstance

// ConservationLaw represents the I_vert + I_horiz conservation hypothesis.
// Originally posited as a theorem, later demoted to hypothesis after
// cross-tradition analysis showed only weak correlation (~0.436).
type ConservationLaw struct {
	Correlation   float64 `json:"correlation"`    // ~0.436 (weak)
	CV            float64 `json:"cv"`             // ~14.4% coefficient of variation
	Status        string  `json:"status"`         // "hypothesis" (demoted from theorem)
	MeantoneRatio float64 `json:"meantone_ratio"` // 1.003 — meantone is the attractor
	Description   string  `json:"description"`
}

// DialsFramework represents the "Dials Not Laws" theory — musical traditions
// are points in a continuous dial space, not rule-bound categories.
type DialsFramework struct {
	Dimensions     []string `json:"dimensions"`
	Traditions     []string `json:"traditions"`
	Unexplored     float64  `json:"unexplored"`      // 82% of dial space unexplored
	VKHCorrelation float64  `json:"vkh_correlation"` // -0.935 vertical/horizontal inverse
	Description    string   `json:"description"`
}

// InnovationCycle represents the Discovery→Codification→Ubiquity→Boredom→Rebellion
// cycle observed across musical traditions.
type InnovationCycle struct {
	Phases       []string `json:"phases"`
	Description  string   `json:"description"`
	CurrentPhase string   `json:"current_phase"` // Where AI-generated music sits
}

// Theory holds all theoretical frameworks in the SuperInstance ecosystem.
var Theory = struct {
	Conservation ConservationLaw
	Dials        DialsFramework
	Innovation   InnovationCycle
}{
	Conservation: ConservationLaw{
		Correlation:   0.436,
		CV:            14.4,
		Status:        "hypothesis",
		MeantoneRatio: 1.003,
		Description:   "Hypothesis that vertical (harmonic) and horizontal (melodic) tension are conserved. Weak correlation found cross-traditionally. Meantone tuning (ratio ~1.003) appears as an attractor.",
	},
	Dials: DialsFramework{
		Dimensions: []string{
			"harmonic_tension",
			"rhythmic_complexity",
			"spectral_density",
			"interval_diversity",
			"temporal_symmetry",
			"register_span",
			"articulation_variance",
		},
		Traditions: []string{
			"Jazz", "Classical", "Gamelan", "Gagaku", "Hindustani",
			"African Polyrhythm", "EDM", "Blues", "Hip-hop", "Latin",
		},
		Unexplored:     0.82,
		VKHCorrelation: -0.935,
		Description:    "Musical traditions are positions in a continuous multi-dimensional dial space, not discrete categories. 82% of this space remains unexplored. Vertical and horizontal tension show strong inverse correlation (-0.935).",
	},
	Innovation: InnovationCycle{
		Phases: []string{
			"Discovery",
			"Codification",
			"Ubiquity",
			"Boredom",
			"Rebellion",
		},
		Description:  "Observed cycle where new musical techniques go from discovery through codification to ubiquity, then provoke boredom-driven rebellion. AI-generated music currently sits in the Codification phase.",
		CurrentPhase: "Codification",
	},
}

// TrinityDimensions returns the three dimensions of the Trinity scoring system.
func TrinityDimensions() []string {
	return []string{"Ethos", "Pathos", "Logos"}
}

// PhaseTransitions returns the valid lifecycle phase transitions.
func PhaseTransitions() map[string][]string {
	return map[string][]string{
		PhaseIncubate: {PhaseCompete, PhaseSunset},
		PhaseCompete:  {PhaseSurvive, PhaseSunset},
		PhaseSurvive:  {PhaseSunset},
		PhaseSunset:   {},
	}
}
