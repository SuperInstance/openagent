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
)

// FleetConservation models the fleet conservation law:
//
//	γ + H = 1.283 - 0.159·log(V) ± σ(V)
//
// where γ is the average fleet trinity score (gamma), H is the fleet
// entropy (Helmholtz term), and V is the fleet size. σ(V) shrinks as
// fleet size grows, reflecting tighter conservation at scale.

// conservationSigma returns the expected deviation σ(V) for a given
// fleet size. The deviation shrinks approximately as 1/√V.
func conservationSigma(fleetSize int) float64 {
	if fleetSize <= 0 {
		return 0
	}
	return 0.28 / math.Sqrt(float64(fleetSize))
}

// ConservationExpected returns the expected γ and H values for a given
// fleet size according to the conservation law.
func ConservationExpected(fleetSize int) (gamma, helmholtz float64) {
	if fleetSize <= 0 {
		return 0, 0
	}
	// The conservation law constrains the sum, but the split between
	// γ and H is observed to favour γ at larger fleet sizes.
	total := 1.283 - 0.159*math.Log(float64(fleetSize))
	// Empirical split: γ takes ~62% of the total, H takes ~38%
	gamma = total * 0.62
	helmholtz = total * 0.38
	return gamma, helmholtz
}

// ConservationDeviation returns how far observed values deviate from
// the conservation law prediction, measured in units of σ(V).
// A deviation > 2 suggests a statistically significant departure.
func ConservationDeviation(fleetSize int, observedGamma, observedH float64) float64 {
	expectedG, expectedH := ConservationExpected(fleetSize)
	sigma := conservationSigma(fleetSize)
	if sigma == 0 {
		return 0
	}
	expectedSum := expectedG + expectedH
	observedSum := observedGamma + observedH
	return (observedSum - expectedSum) / sigma
}

// ConservationTotal returns the expected γ + H total for a fleet size.
func ConservationTotal(fleetSize int) float64 {
	if fleetSize <= 0 {
		return 0
	}
	return 1.283 - 0.159*math.Log(float64(fleetSize))
}
