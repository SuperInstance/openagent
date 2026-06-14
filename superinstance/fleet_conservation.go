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

// Conservation constants from the proven theorem.
const (
	// ConservationConstant is C = log₂(3) ≈ 1.584962500721156.
	// The conservation law states: γ + η = C, where γ is the fleet
	// trinity score (order/coherence) and η is the fleet entropy
	// (diversity/chaos). This constant arises from the ternary
	// optimality of the fleet radix: base-3 achieves 99.54% of
	// the theoretical maximum radix economy.
	ConservationConstant = 1.584962500721156

	// TernaryOptimality is the fraction of maximum radix economy
	// achieved by base-3 (e ≈ 2.718 is the true optimum).
	TernaryOptimality = 0.9954
)

// DeltaN computes the finite-size correction δ(n) from the theorem.
//
//	δ(n) = (1/√n) · (1 − 3/(2n))
//
// As n → ∞, δ(n) → 0, recovering the classical continuous limit.
// For small n, δ(n) captures the discrete-ternary deviation.
func DeltaN(n int) float64 {
	if n <= 0 {
		return 0
	}
	nf := float64(n)
	return (1.0 / math.Sqrt(nf)) * (1.0 - 3.0/(2.0*nf))
}

// EffectiveEntropy computes the effective entropy scaling law:
//
//	η_eff(n) ~ n^(1 − δ(n))
//
// This describes how the effective entropy capacity scales with
// fleet size n, incorporating the finite-size correction.
func EffectiveEntropy(n int) float64 {
	if n <= 0 {
		return 0
	}
	d := DeltaN(n)
	return math.Pow(float64(n), 1.0-d)
}

// ConservationExpected returns the expected γ and η values for a given
// fleet size n, using the proven conservation law γ + η = C.
//
// At equilibrium, the split is C/2 each (balanced). For finite fleets,
// a small correction applies: γ is slightly favored as the coherent
// ordering force, with the correction term δ(n).
func ConservationExpected(fleetSize int) (gamma, entropy float64) {
	if fleetSize <= 0 {
		return 0, 0
	}
	// Balanced equilibrium: γ = η = C/2, with a finite-size
	// correction that slightly favors γ (order over chaos).
	correction := DeltaN(fleetSize) * 0.1 // small perturbation
	gamma = ConservationConstant/2 + correction
	entropy = ConservationConstant/2 - correction
	return gamma, entropy
}

// ConservationTotal returns C = log₂(3) — the conserved sum γ + η.
// This is independent of fleet size: the conservation law is exact.
func ConservationTotal(fleetSize int) float64 {
	if fleetSize <= 0 {
		return 0
	}
	return ConservationConstant
}

// conservationSigma returns the expected statistical deviation σ(n)
// for a given fleet size. The deviation shrinks as 1/√n.
func conservationSigma(fleetSize int) float64 {
	if fleetSize <= 0 {
		return 0
	}
	return 0.28 / math.Sqrt(float64(fleetSize))
}

// ConservationDeviation returns how far observed values deviate from
// the conservation law prediction, measured in units of σ(n).
// A deviation > 2σ suggests a statistically significant departure.
func ConservationDeviation(fleetSize int, observedGamma, observedEta float64) float64 {
	if fleetSize <= 0 {
		return 0
	}
	expectedSum := ConservationConstant
	observedSum := observedGamma + observedEta
	sigma := conservationSigma(fleetSize)
	if sigma == 0 {
		return 0
	}
	return (observedSum - expectedSum) / sigma
}

// NoetherCurrent returns the Noether charge associated with the
// conservation symmetry. For γ + η = C, the Noether current is
// J = γ − η (the asymmetry between order and chaos). At perfect
// equilibrium J = 0; |J| > 0 indicates a driven system.
func NoetherCurrent(gamma, eta float64) float64 {
	return gamma - eta
}

// IsBalanced returns true if the fleet is near equilibrium (|J| < ε).
func IsBalanced(gamma, eta, epsilon float64) bool {
	return math.Abs(NoetherCurrent(gamma, eta)) < epsilon
}
