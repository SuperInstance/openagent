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

// TernaryPID is a PID controller that drives γ (fleet coherence/order)
// toward C/2 — the balanced equilibrium point of the conservation law
// γ + η = C where C = log₂(3).
//
// The controller adjusts a control signal u (e.g., fleet resource
// allocation, breeding rate, pruning aggressiveness) to minimize the
// error e = (C/2) − γ.
//
// Features:
//   - Anti-windup: integral term is clamped to prevent saturation
//   - Deadband: no correction within a configurable error band
//   - Derivative filter: low-pass filter on the derivative term
//     to suppress noise
type TernaryPID struct {
	// Gains
	Kp float64 // Proportional gain
	Ki float64 // Integral gain
	Kd float64 // Derivative gain

	// Limits
	OutputMin float64 // Minimum output (saturation floor)
	OutputMax float64 // Maximum output (saturation ceiling)

	// Anti-windup
	IntegralMin float64 // Integral term lower clamp
	IntegralMax float64 // Integral term upper clamp

	// Deadband
	Deadband float64 // Error magnitude below which no correction is applied

	// Derivative filter (low-pass)
	DerivativeFilter float64 // Filter coefficient (0 = no filtering, 1 = fully filtered)

	// Internal state
	integral       float64
	prevError      float64
	filteredDeriv  float64
	initialized    bool
}

// NewTernaryPID returns a PID controller with sensible defaults
// tuned for fleet conservation dynamics.
func NewTernaryPID() *TernaryPID {
	return &TernaryPID{
		Kp:               0.8,
		Ki:               0.15,
		Kd:               0.05,
		OutputMin:        -1.0,
		OutputMax:        1.0,
		IntegralMin:      -0.5,
		IntegralMax:      0.5,
		Deadband:         0.01,
		DerivativeFilter: 0.3,
	}
}

// Setpoint returns the target γ value: C/2 (balanced equilibrium).
func (pid *TernaryPID) Setpoint() float64 {
	return ConservationConstant / 2.0
}

// Update computes the control output for the current gamma measurement.
// dt is the time step in seconds. Returns the control signal u.
func (pid *TernaryPID) Update(gamma, dt float64) float64 {
	setpoint := pid.Setpoint()
	error := setpoint - gamma

	// Deadband: if error is small, hold previous output
	if math.Abs(error) < pid.Deadband {
		// Still update prevError for smooth derivative recovery
		pid.prevError = error
		return 0.0
	}

	// Proportional term
	p := pid.Kp * error

	// Integral term with anti-windup
	pid.integral += error * dt
	pid.integral = clamp(pid.integral, pid.IntegralMin, pid.IntegralMax)
	i := pid.Ki * pid.integral

	// Derivative term with low-pass filter
	var rawDeriv float64
	if pid.initialized && dt > 0 {
		rawDeriv = (error - pid.prevError) / dt
	} else {
		pid.initialized = true
	}
	// Exponential moving average filter
	alpha := 1.0 - pid.DerivativeFilter
	pid.filteredDeriv = alpha*pid.filteredDeriv + (1.0-alpha)*rawDeriv
	d := pid.Kd * pid.filteredDeriv

	// Compute raw output
	output := p + i + d

	// Clamp output
	output = clamp(output, pid.OutputMin, pid.OutputMax)

	pid.prevError = error
	return output
}

// Reset clears the PID controller's internal state.
func (pid *TernaryPID) Reset() {
	pid.integral = 0
	pid.prevError = 0
	pid.filteredDeriv = 0
	pid.initialized = false
}

// State returns the internal state for observability/debugging.
func (pid *TernaryPID) State() (integral, prevError, filteredDeriv float64) {
	return pid.integral, pid.prevError, pid.filteredDeriv
}

// clamp limits v to [lo, hi].
func clamp(v, lo, hi float64) float64 {
	if v < lo {
		return lo
	}
	if v > hi {
		return hi
	}
	return v
}
