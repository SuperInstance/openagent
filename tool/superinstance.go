// Copyright 2026 The SuperInstance Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with in writing, software
// distributed under the License on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tool

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/ThinkInAIXYZ/go-mcp/protocol"
	"github.com/the-open-agent/openagent/superinstance"
	"github.com/the-open-agent/openagent/tool/builtin_tool"
)

// SuperInstanceTool provides MCP tools for querying the SuperInstance ecosystem.
type SuperInstanceTool struct{}

func (t *SuperInstanceTool) BuiltinTools() []BuiltinTool {
	return nil
}

// NewSuperInstanceQueryBuiltin creates a builtin tool for querying the ecosystem.
func NewSuperInstanceQueryBuiltin() builtin_tool.BuiltinTool {
	return &superInstanceQueryBuiltin{resolver: superinstance.NewContextResolver()}
}

// NewSuperInstanceRepoStatusBuiltin creates a builtin tool for checking repo status.
func NewSuperInstanceRepoStatusBuiltin() builtin_tool.BuiltinTool {
	return &superInstanceRepoStatusBuiltin{}
}

// NewSuperInstanceFleetBuiltin creates a builtin tool for fleet status.
func NewSuperInstanceFleetBuiltin() builtin_tool.BuiltinTool {
	return &superInstanceFleetBuiltin{}
}

// --- Query Tool ---

type superInstanceQueryBuiltin struct {
	resolver *superinstance.ContextResolver
}

func (t *superInstanceQueryBuiltin) GetName() string {
	return "superinstance_query"
}

func (t *superInstanceQueryBuiltin) GetDescription() string {
	return "Query the SuperInstance ecosystem knowledge base. Use this to resolve questions about repos, agents, theory, fleet, constraints, flux, or any SuperInstance concept."
}

func (t *superInstanceQueryBuiltin) GetInputSchema() interface{} {
	return map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"query": map[string]interface{}{
				"type":        "string",
				"description": "Natural language query about the SuperInstance ecosystem (e.g., 'constraint repos', 'fleet status', 'dials theory').",
			},
		},
		"required": []string{"query"},
	}
}

func (t *superInstanceQueryBuiltin) Execute(_ context.Context, arguments map[string]interface{}) (*protocol.CallToolResult, error) {
	query, _ := arguments["query"].(string)
	if query == "" {
		return errorResult("missing required parameter: query"), nil
	}

	result := t.resolver.Resolve(query)

	var b strings.Builder
	b.WriteString(fmt.Sprintf("# SuperInstance Query: %s\n\n", query))

	if len(result.Repos) > 0 {
		b.WriteString("## Related Repos\n")
		for _, r := range result.Repos {
			b.WriteString(fmt.Sprintf("- **%s** (%s): %s\n", r.Name, r.Language, r.Description))
		}
		b.WriteString("\n")
	}

	if len(result.Agents) > 0 {
		b.WriteString("## Related Agents\n")
		for _, a := range result.Agents {
			b.WriteString(fmt.Sprintf("- **%s** — %s [%s] Gen %d\n", a.Name, a.Role, a.Phase, a.Generation))
		}
		b.WriteString("\n")
	}

	if result.TheoryContext != "" {
		b.WriteString(fmt.Sprintf("## Theory Context\n%s\n", result.TheoryContext))
	}

	if len(result.Suggestions) > 0 {
		b.WriteString(fmt.Sprintf("\n## Suggestions\n%s\n", strings.Join(result.Suggestions, " · ")))
	}

	return textResult(b.String()), nil
}

// --- Repo Status Tool ---

type superInstanceRepoStatusBuiltin struct{}

func (t *superInstanceRepoStatusBuiltin) GetName() string {
	return "superinstance_repo_status"
}

func (t *superInstanceRepoStatusBuiltin) GetDescription() string {
	return "List all SuperInstance repositories with their metadata, or filter by language/tag."
}

func (t *superInstanceRepoStatusBuiltin) GetInputSchema() interface{} {
	return map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"filter": map[string]interface{}{
				"type":        "string",
				"description": "Optional filter: a language name (e.g., 'Rust', 'Python') or tag (e.g., 'music', 'core').",
			},
			"format": map[string]interface{}{
				"type":        "string",
				"description": "Output format: 'list' (default) or 'json'.",
			},
		},
	}
}

func (t *superInstanceRepoStatusBuiltin) Execute(_ context.Context, arguments map[string]interface{}) (*protocol.CallToolResult, error) {
	filter, _ := arguments["filter"].(string)
	format, _ := arguments["format"].(string)

	var repos []superinstance.Repo
	if filter != "" {
		repos = superinstance.FindReposByTag(filter)
		if len(repos) == 0 {
			repos = superinstance.FindReposByLanguage(filter)
		}
	} else {
		for _, r := range superinstance.Repos {
			repos = append(repos, r)
		}
	}

	// Sort by name
	sort.Slice(repos, func(i, j int) bool {
		return repos[i].Name < repos[j].Name
	})

	if format == "json" {
		data, err := json.MarshalIndent(repos, "", "  ")
		if err != nil {
			return errorResult(fmt.Sprintf("JSON marshal error: %v", err)), nil
		}
		return textResult(string(data)), nil
	}

	var b strings.Builder
	b.WriteString(fmt.Sprintf("# SuperInstance Repositories (%d)\n\n", len(repos)))
	for _, r := range repos {
		b.WriteString(fmt.Sprintf("- **%s** (%s): %s\n", r.Name, r.Language, r.Description))
	}

	return textResult(b.String()), nil
}

// --- Fleet Status Tool ---

type superInstanceFleetBuiltin struct{}

func (t *superInstanceFleetBuiltin) GetName() string {
	return "superinstance_fleet"
}

func (t *superInstanceFleetBuiltin) GetDescription() string {
	return "Show the SuperInstance agent fleet: roster, trinity scores, lifecycle phases, and fleet state."
}

func (t *superInstanceFleetBuiltin) GetInputSchema() interface{} {
	return map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"detail": map[string]interface{}{
				"type":        "string",
				"description": "Detail level: 'summary' (default) or 'full' (includes all trinity scores and dependencies).",
			},
		},
	}
}

func (t *superInstanceFleetBuiltin) Execute(_ context.Context, arguments map[string]interface{}) (*protocol.CallToolResult, error) {
	detail, _ := arguments["detail"].(string)
	if detail == "" {
		detail = "summary"
	}

	fleet := superinstance.DefaultFleetState()

	var b strings.Builder
	b.WriteString("# SuperInstance Fleet Status\n\n")
	b.WriteString(fmt.Sprintf("**Phase:** %s | **Total Agents:** %d | **Last Generation:** %d\n\n",
		fleet.Phase, fleet.TotalAgents, fleet.LastGeneration))

	agents := superinstance.ActiveAgents()
	sort.Slice(agents, func(i, j int) bool {
		return agents[i].Trinity.Average() > agents[j].Trinity.Average()
	})

	if detail == "full" {
		for _, a := range agents {
			b.WriteString(fmt.Sprintf("## %s — %s\n", a.Name, a.Role))
			b.WriteString(fmt.Sprintf("- Phase: %s (Gen %d)\n", a.Phase, a.Generation))
			b.WriteString(fmt.Sprintf("- Trinity: Ethos=%.2f | Pathos=%.2f | Logos=%.2f | Avg=%.2f\n",
				a.Trinity.Ethos, a.Trinity.Pathos, a.Trinity.Logos, a.Trinity.Average()))
			b.WriteString(fmt.Sprintf("- Status: %s\n\n", a.Status))
		}

		b.WriteString("## Theories\n")
		b.WriteString(fmt.Sprintf("- **Conservation**: %s (r=%.3f, CV=%.1f%%)\n",
			superinstance.Theory.Conservation.Status,
			superinstance.Theory.Conservation.Correlation,
			superinstance.Theory.Conservation.CV))
		b.WriteString(fmt.Sprintf("- **Dials**: %d dimensions, %.0f%% unexplored, VKH r=%.3f\n",
			len(superinstance.Theory.Dials.Dimensions),
			superinstance.Theory.Dials.Unexplored*100,
			superinstance.Theory.Dials.VKHCorrelation))
		b.WriteString(fmt.Sprintf("- **Innovation**: %s (phase: %s)\n",
			strings.Join(superinstance.Theory.Innovation.Phases, " → "),
			superinstance.Theory.Innovation.CurrentPhase))
	} else {
		b.WriteString("| Agent | Role | Phase | Gen | Trinity Avg |\n")
		b.WriteString("|-------|------|-------|-----|-------------|\n")
		for _, a := range agents {
			b.WriteString(fmt.Sprintf("| %s | %s | %s | %d | %.2f |\n",
				a.Name, a.Role, a.Phase, a.Generation, a.Trinity.Average()))
		}
	}

	return textResult(b.String()), nil
}

// --- Helpers ---

func textResult(text string) *protocol.CallToolResult {
	return &protocol.CallToolResult{
		IsError: false,
		Content: []protocol.Content{
			&protocol.TextContent{Type: "text", Text: text},
		},
	}
}

func errorResult(text string) *protocol.CallToolResult {
	return &protocol.CallToolResult{
		IsError: true,
		Content: []protocol.Content{
			&protocol.TextContent{Type: "text", Text: text},
		},
	}
}
