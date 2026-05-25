// Copyright 2026 The SuperInstance Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package superinstance provides innate knowledge of the SuperInstance
// software ecosystem, its agents, theoretical frameworks, and fleet state.
package superinstance

// Organization is the GitHub organization for all SuperInstance repos.
const Organization = "SuperInstance"

// Ecosystem represents the full SuperInstance software stack.
type Ecosystem struct {
	Repos      map[string]Repo
	Agents     map[string]Agent
	Protocols  map[string]Protocol
	FleetState FleetState
}

// Repo describes a repository in the SuperInstance ecosystem.
type Repo struct {
	Name          string   `json:"name"`
	Description   string   `json:"description"`
	Language      string   `json:"language"`
	DefaultBranch string   `json:"default_branch"`
	DependsOn     []string `json:"depends_on"`
	Tags          []string `json:"tags"`
	URL           string   `json:"url"`
}

// Protocol describes a communication protocol used in the ecosystem.
type Protocol struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	UsedBy      []string `json:"used_by"`
}

// FleetState represents the current state of the SuperInstance agent fleet.
type FleetState struct {
	TotalAgents    int      `json:"total_agents"`
	ActiveAgents   []string `json:"active_agents"`
	Phase          string   `json:"phase"`
	LastGeneration int      `json:"last_generation"`
}

// Repos contains all known SuperInstance repositories.
var Repos = map[string]Repo{
	"sunset-ecosystem": {
		Name:          "sunset-ecosystem",
		Description:   "Trinity architecture agent lifecycle (Ethos×Pathos×Logos)",
		Language:      "Python",
		DefaultBranch: "main",
		DependsOn:     []string{},
		Tags:          []string{"core", "orchestration", "evolution"},
		URL:           "https://github.com/SuperInstance/sunset-ecosystem",
	},
	"ccc-os": {
		Name:          "ccc-os",
		Description:   "Autonomous fleet monitoring & decision system",
		Language:      "Python",
		DefaultBranch: "main",
		DependsOn:     []string{"sunset-ecosystem"},
		Tags:          []string{"monitoring", "fleet", "operations"},
		URL:           "https://github.com/SuperInstance/ccc-os",
	},
	"constraint-toolkit": {
		Name:          "constraint-toolkit",
		Description:   "Workspace-native constraint theory toolkit (not a standalone repo)",
		Language:      "Python",
		DefaultBranch: "main",
		DependsOn:     []string{"constraint-synth"},
		Tags:          []string{"music", "theory", "toolkit", "workspace"},
		URL:           "https://github.com/SuperInstance/constraint-toolkit",
	},
	"constraint-synth": {
		Name:          "constraint-synth",
		Description:   "Music theory constraint synthesis (PyPI package)",
		Language:      "Python",
		DefaultBranch: "main",
		DependsOn:     []string{},
		Tags:          []string{"music", "theory", "synthesis"},
		URL:           "https://github.com/SuperInstance/constraint-synth",
	},
	"constraint-audio": {
		Name:          "constraint-audio",
		Description:   "Rust audio processing for constraint-based music",
		Language:      "Rust",
		DefaultBranch: "main",
		DependsOn:     []string{},
		Tags:          []string{"audio", "rust", "music"},
		URL:           "https://github.com/SuperInstance/constraint-audio",
	},
	"constraint-mux": {
		Name:          "constraint-mux",
		Description:   "Rust multiplexer for constraint streams",
		Language:      "Rust",
		DefaultBranch: "main",
		DependsOn:     []string{"constraint-audio"},
		Tags:          []string{"mux", "rust", "streaming"},
		URL:           "https://github.com/SuperInstance/constraint-mux",
	},
	"counterpoint-engine": {
		Name:          "counterpoint-engine",
		Description:   "Species counterpoint engine (PyPI)",
		Language:      "Python",
		DefaultBranch: "main",
		DependsOn:     []string{"constraint-synth"},
		Tags:          []string{"music", "counterpoint", "theory"},
		URL:           "https://github.com/SuperInstance/counterpoint-engine",
	},
	"flux-tensor-midi": {
		Name:          "flux-tensor-midi",
		Description:   "Tensor-based MIDI manipulation",
		Language:      "Python",
		DefaultBranch: "main",
		DependsOn:     []string{},
		Tags:          []string{"midi", "tensor", "flux"},
		URL:           "https://github.com/SuperInstance/flux-tensor-midi",
	},
	"flux-genome": {
		Name:          "flux-genome",
		Description:   "Musical genome mapping and evolution",
		Language:      "Python",
		DefaultBranch: "main",
		DependsOn:     []string{"flux-tensor-midi"},
		Tags:          []string{"genome", "evolution", "flux"},
		URL:           "https://github.com/SuperInstance/flux-genome",
	},
	"flux-hyperbolic": {
		Name:          "flux-hyperbolic",
		Description:   "Hyperbolic geometry for musical space exploration",
		Language:      "Python",
		DefaultBranch: "main",
		DependsOn:     []string{"flux-genome"},
		Tags:          []string{"hyperbolic", "geometry", "flux"},
		URL:           "https://github.com/SuperInstance/flux-hyperbolic",
	},
	"AI-Writings": {
		Name:          "AI-Writings",
		Description:   "Creative writing collection by AI models",
		Language:      "Markdown",
		DefaultBranch: "main",
		DependsOn:     []string{},
		Tags:          []string{"creative", "writing", "art"},
		URL:           "https://github.com/SuperInstance/AI-Writings",
	},
	"forgemaster": {
		Name:          "forgemaster",
		Description:   "Build & forge orchestration system",
		Language:      "Python",
		DefaultBranch: "main",
		DependsOn:     []string{},
		Tags:          []string{"build", "forge", "orchestration"},
		URL:           "https://github.com/SuperInstance/forgemaster",
	},
	"fm-research": {
		Name:          "fm-research",
		Description:   "Forgemaster research notes and experiments",
		Language:      "Markdown",
		DefaultBranch: "main",
		DependsOn:     []string{"forgemaster"},
		Tags:          []string{"research", "notes"},
		URL:           "https://github.com/SuperInstance/fm-research",
	},
	"deadband-rs": {
		Name:          "deadband-rs",
		Description:   "Rust deadband theory for groove analysis",
		Language:      "Rust",
		DefaultBranch: "main",
		DependsOn:     []string{},
		Tags:          []string{"rust", "deadband", "groove", "rhythm"},
		URL:           "https://github.com/SuperInstance/deadband-rs",
	},
	"cocapn-health": {
		Name:          "cocapn-health",
		Description:   "Health monitoring service for the fleet",
		Language:      "Python",
		DefaultBranch: "main",
		DependsOn:     []string{"ccc-os"},
		Tags:          []string{"health", "monitoring"},
		URL:           "https://github.com/SuperInstance/cocapn-health",
	},
	"creative-engine-c": {
		Name:          "creative-engine-c",
		Description:   "C implementation of the creative engine",
		Language:      "C",
		DefaultBranch: "main",
		DependsOn:     []string{},
		Tags:          []string{"creative", "engine", "c"},
		URL:           "https://github.com/SuperInstance/creative-engine-c",
	},
	"creative-engine-rust": {
		Name:          "creative-engine-rust",
		Description:   "Rust implementation of the creative engine",
		Language:      "Rust",
		DefaultBranch: "main",
		DependsOn:     []string{},
		Tags:          []string{"creative", "engine", "rust"},
		URL:           "https://github.com/SuperInstance/creative-engine-rust",
	},
	"superinstance-ffi": {
		Name:          "superinstance-ffi",
		Description:   "FFI bindings across SuperInstance languages (Rust↔Python↔C)",
		Language:      "Rust",
		DefaultBranch: "main",
		DependsOn:     []string{},
		Tags:          []string{"ffi", "bindings", "interop"},
		URL:           "https://github.com/SuperInstance/superinstance-ffi",
	},
	"openagent": {
		Name:          "openagent",
		Description:   "The SuperInstance system agent (this repo)",
		Language:      "Go",
		DefaultBranch: "master",
		DependsOn:     []string{"ccc-os", "sunset-ecosystem"},
		Tags:          []string{"agent", "core", "go"},
		URL:           "https://github.com/SuperInstance/openagent",
	},
	"docs": {
		Name:          "docs",
		Description:   "SuperInstance documentation hub",
		Language:      "Markdown",
		DefaultBranch: "main",
		DependsOn:     []string{},
		Tags:          []string{"docs", "documentation"},
		URL:           "https://github.com/SuperInstance/docs",
	},
	"wiki": {
		Name:          "wiki",
		Description:   "SuperInstance knowledge wiki",
		Language:      "Markdown",
		DefaultBranch: "main",
		DependsOn:     []string{},
		Tags:          []string{"wiki", "knowledge"},
		URL:           "https://github.com/SuperInstance/wiki",
	},
	"constraint-theory-core": {
		Name:          "constraint-theory-core",
		Description:   "Core constraint theory formalization",
		Language:      "Python",
		DefaultBranch: "main",
		DependsOn:     []string{},
		Tags:          []string{"theory", "constraints", "formal"},
		URL:           "https://github.com/SuperInstance/constraint-theory-core",
	},
	"flux-algebra": {
		Name:          "flux-algebra",
		Description:   "Algebraic structures for musical flux operations",
		Language:      "Python",
		DefaultBranch: "main",
		DependsOn:     []string{"flux-tensor-midi"},
		Tags:          []string{"flux", "algebra", "math"},
		URL:           "https://github.com/SuperInstance/flux-algebra",
	},
	"constraint-dialect": {
		Name:          "constraint-dialect",
		Description:   "Dialectical constraint resolution for competing musical rules",
		Language:      "Python",
		DefaultBranch: "main",
		DependsOn:     []string{"constraint-synth", "constraint-theory-core"},
		Tags:          []string{"constraints", "dialectic", "resolution"},
		URL:           "https://github.com/SuperInstance/constraint-dialect",
	},
	"flux-julia": {
		Name:          "flux-julia",
		Description:   "Julia implementation of flux tensor operations for high-performance music math",
		Language:      "Julia",
		DefaultBranch: "main",
		DependsOn:     []string{},
		Tags:          []string{"flux", "julia", "performance"},
		URL:           "https://github.com/SuperInstance/flux-julia",
	},
	"agent-operations": {
		Name:          "agent-operations",
		Description:   "Operational tooling for fleet agent management and deployment",
		Language:      "Go",
		DefaultBranch: "main",
		DependsOn:     []string{"openagent", "ccc-os"},
		Tags:          []string{"operations", "fleet", "deployment"},
		URL:           "https://github.com/SuperInstance/agent-operations",
	},
	"constraint-viz": {
		Name:          "constraint-viz",
		Description:   "Visualization toolkit for constraint spaces and dial positions",
		Language:      "Python",
		DefaultBranch: "main",
		DependsOn:     []string{"constraint-synth"},
		Tags:          []string{"visualization", "constraints", "dials"},
		URL:           "https://github.com/SuperInstance/constraint-viz",
	},
	"flux-rust": {
		Name:          "flux-rust",
		Description:   "Rust implementation of flux tensor core for real-time audio pipelines",
		Language:      "Rust",
		DefaultBranch: "main",
		DependsOn:     []string{},
		Tags:          []string{"flux", "rust", "performance"},
		URL:           "https://github.com/SuperInstance/flux-rust",
	},
	"tonal-archaeology": {
		Name:          "tonal-archaeology",
		Description:   "Historical tonal system analysis and reconstruction toolkit",
		Language:      "Python",
		DefaultBranch: "main",
		DependsOn:     []string{"constraint-synth", "deadband-rs"},
		Tags:          []string{"tonal", "history", "analysis"},
		URL:           "https://github.com/SuperInstance/tonal-archaeology",
	},
	"fleet-conservation": {
		Name:          "fleet-conservation",
		Description:   "Fleet conservation law models and statistical validation",
		Language:      "Python",
		DefaultBranch: "main",
		DependsOn:     []string{"constraint-theory-core"},
		Tags:          []string{"fleet", "conservation", "statistics"},
		URL:           "https://github.com/SuperInstance/fleet-conservation",
	},
	"dial-theory": {
		Name:          "dial-theory",
		Description:   "Formal 'Dials Not Laws' theory implementation and tradition mapping",
		Language:      "Python",
		DefaultBranch: "main",
		DependsOn:     []string{"constraint-theory-core"},
		Tags:          []string{"dials", "theory", "traditions"},
		URL:           "https://github.com/SuperInstance/dial-theory",
	},
	"harmonic-atlas": {
		Name:          "harmonic-atlas",
		Description:   "Worldwide harmonic practice atlas mapped to dial positions",
		Language:      "Python",
		DefaultBranch: "main",
		DependsOn:     []string{"dial-theory", "tonal-archaeology"},
		Tags:          []string{"harmonic", "atlas", "world-music"},
		URL:           "https://github.com/SuperInstance/harmonic-atlas",
	},
	"spectral-taxonomy": {
		Name:          "spectral-taxonomy",
		Description:   "Taxonomy of spectral density patterns across traditions",
		Language:      "Python",
		DefaultBranch: "main",
		DependsOn:     []string{"constraint-audio"},
		Tags:          []string{"spectral", "taxonomy", "analysis"},
		URL:           "https://github.com/SuperInstance/spectral-taxonomy",
	},
	"breeding-grounds": {
		Name:          "breeding-grounds",
		Description:   "Agent breeding experiment infrastructure and genealogy tracking",
		Language:      "Python",
		DefaultBranch: "main",
		DependsOn:     []string{"sunset-ecosystem"},
		Tags:          []string{"breeding", "agents", "evolution"},
		URL:           "https://github.com/SuperInstance/breeding-grounds",
	},
	"polyglot-bridge": {
		Name:          "polyglot-bridge",
		Description:   "Polyglot RPC bridge connecting Python/Rust/Julia/Go services",
		Language:      "Rust",
		DefaultBranch: "main",
		DependsOn:     []string{"superinstance-ffi"},
		Tags:          []string{"rpc", "bridge", "interop"},
		URL:           "https://github.com/SuperInstance/polyglot-bridge",
	},
}

// Protocols used across the ecosystem.
var Protocols = map[string]Protocol{
	"trinity-lifecycle": {
		Name:        "trinity-lifecycle",
		Description: "Agent lifecycle protocol based on Ethos×Pathos×Logos trinity scoring",
		UsedBy:      []string{"sunset-ecosystem", "ccc-os", "openagent"},
	},
	"fleet-heartbeat": {
		Name:        "fleet-heartbeat",
		Description: "Periodic health and status reporting from fleet agents",
		UsedBy:      []string{"ccc-os", "cocapn-health"},
	},
	"constraint-protocol": {
		Name:        "constraint-protocol",
		Description: "Music theory constraint exchange format between constraint-* repos",
		UsedBy:      []string{"constraint-synth", "constraint-audio", "constraint-mux", "counterpoint-engine"},
	},
	"flux-protocol": {
		Name:        "flux-protocol",
		Description: "Tensor/gene exchange protocol for musical data",
		UsedBy:      []string{"flux-tensor-midi", "flux-genome", "flux-hyperbolic"},
	},
	"forge-protocol": {
		Name:        "forge-protocol",
		Description: "Build and deployment orchestration protocol",
		UsedBy:      []string{"forgemaster"},
	},
}

// DefaultFleetState returns the default fleet state.
func DefaultFleetState() FleetState {
	return FleetState{
		TotalAgents:    len(KnownAgents),
		ActiveAgents:   []string{"CCC", "Oracle1", "FM", "TurboVec"},
		Phase:          "COMPETE",
		LastGeneration: 5,
	}
}

// FullEcosystem returns the complete ecosystem with all default values.
func FullEcosystem() *Ecosystem {
	return &Ecosystem{
		Repos:      Repos,
		Agents:     KnownAgents,
		Protocols:  Protocols,
		FleetState: DefaultFleetState(),
	}
}

// FindReposByTag returns all repos matching a given tag.
func FindReposByTag(tag string) []Repo {
	var result []Repo
	for _, r := range Repos {
		for _, t := range r.Tags {
			if t == tag {
				result = append(result, r)
				break
			}
		}
	}
	return result
}

// FindReposByLanguage returns all repos in a given language.
func FindReposByLanguage(lang string) []Repo {
	var result []Repo
	for _, r := range Repos {
		if r.Language == lang {
			result = append(result, r)
		}
	}
	return result
}

// DependencyGraph returns the full dependency graph as adjacency list.
func DependencyGraph() map[string][]string {
	graph := make(map[string][]string)
	for name, repo := range Repos {
		graph[name] = repo.DependsOn
	}
	return graph
}
