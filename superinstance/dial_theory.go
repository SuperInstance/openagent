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
)

// Tradition represents a musical tradition positioned on constraint dials.
// Each dimension is scored 0–5 where 0 is minimal and 5 is maximal.
type Tradition struct {
	Name               string  `json:"name"`
	HarmonicTension    float64 `json:"harmonic_tension"`     // 0-5
	RhythmicComplexity float64 `json:"rhythmic_complexity"`  // 0-5
	SpectralDensity    float64 `json:"spectral_density"`     // 0-5
}

// Traditions maps known musical traditions to their dial positions.
// Positions derived from ethnomusicological analysis and constraint
// theory modelling. Values are approximate and represent centroid
// positions within each tradition's natural variance.
var Traditions = map[string]Tradition{
	"Jazz": {
		Name:               "Jazz",
		HarmonicTension:    3.8,
		RhythmicComplexity: 3.5,
		SpectralDensity:    2.8,
	},
	"Classical": {
		Name:               "Classical",
		HarmonicTension:    3.2,
		RhythmicComplexity: 2.0,
		SpectralDensity:    2.5,
	},
	"Gamelan": {
		Name:               "Gamelan",
		HarmonicTension:    1.5,
		RhythmicComplexity: 4.2,
		SpectralDensity:    3.5,
	},
	"Gagaku": {
		Name:               "Gagaku",
		HarmonicTension:    1.0,
		RhythmicComplexity: 1.5,
		SpectralDensity:    1.8,
	},
	"Hindustani": {
		Name:               "Hindustani",
		HarmonicTension:    2.2,
		RhythmicComplexity: 4.0,
		SpectralDensity:    2.0,
	},
	"African Polyrhythm": {
		Name:               "African Polyrhythm",
		HarmonicTension:    1.8,
		RhythmicComplexity: 4.8,
		SpectralDensity:    3.0,
	},
	"EDM": {
		Name:               "EDM",
		HarmonicTension:    2.5,
		RhythmicComplexity: 2.8,
		SpectralDensity:    4.5,
	},
	"Blues": {
		Name:               "Blues",
		HarmonicTension:    3.0,
		RhythmicComplexity: 2.2,
		SpectralDensity:    1.5,
	},
	"Hip-hop": {
		Name:               "Hip-hop",
		HarmonicTension:    2.0,
		RhythmicComplexity: 4.0,
		SpectralDensity:    3.2,
	},
	"Latin": {
		Name:               "Latin",
		HarmonicTension:    2.8,
		RhythmicComplexity: 4.3,
		SpectralDensity:    2.5,
	},
}

// DialDistance computes the Euclidean distance between two traditions
// in the 3-dimensional dial space (HarmonicTension, RhythmicComplexity,
// SpectralDensity).
func DialDistance(a, b Tradition) float64 {
	dh := a.HarmonicTension - b.HarmonicTension
	dr := a.RhythmicComplexity - b.RhythmicComplexity
	ds := a.SpectralDensity - b.SpectralDensity
	return math.Sqrt(dh*dh + dr*dr + ds*ds)
}

// NearestTradition finds the closest known tradition to a given dial
// position. Returns the tradition name and the distance.
func NearestTradition(pos Tradition) string {
	var closest string
	minDist := math.MaxFloat64
	for _, t := range Traditions {
		d := DialDistance(pos, t)
		if d < minDist {
			minDist = d
			closest = t.Name
		}
	}
	return closest
}
