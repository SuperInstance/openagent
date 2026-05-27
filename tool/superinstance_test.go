// Copyright 2026 The SuperInstance Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tool

import (
	"context"
	"strings"
	"testing"

	"github.com/ThinkInAIXYZ/go-mcp/protocol"
	"github.com/the-open-agent/openagent/superinstance"
)

func TestNewSuperInstanceQueryBuiltin(t *testing.T) {
	tool := NewSuperInstanceQueryBuiltin()
	if tool == nil {
		t.Fatal("NewSuperInstanceQueryBuiltin returned nil")
	}
	if tool.GetName() != "superinstance_query" {
		t.Errorf("GetName() = %q, want %q", tool.GetName(), "superinstance_query")
	}
	if tool.GetDescription() == "" {
		t.Error("GetDescription() returned empty string")
	}
	if tool.GetInputSchema() == nil {
		t.Error("GetInputSchema() returned nil")
	}
}

func TestNewSuperInstanceRepoStatusBuiltin(t *testing.T) {
	tool := NewSuperInstanceRepoStatusBuiltin()
	if tool == nil {
		t.Fatal("NewSuperInstanceRepoStatusBuiltin returned nil")
	}
	if tool.GetName() != "superinstance_repo_status" {
		t.Errorf("GetName() = %q, want %q", tool.GetName(), "superinstance_repo_status")
	}
	if tool.GetDescription() == "" {
		t.Error("GetDescription() returned empty string")
	}
}

func TestNewSuperInstanceFleetBuiltin(t *testing.T) {
	tool := NewSuperInstanceFleetBuiltin()
	if tool == nil {
		t.Fatal("NewSuperInstanceFleetBuiltin returned nil")
	}
	if tool.GetName() != "superinstance_fleet" {
		t.Errorf("GetName() = %q, want %q", tool.GetName(), "superinstance_fleet")
	}
	if tool.GetDescription() == "" {
		t.Error("GetDescription() returned empty string")
	}
}

// --- Query Tool Execute tests ---

func TestQueryToolExecuteMissingParam(t *testing.T) {
	tool := NewSuperInstanceQueryBuiltin()
	result, err := tool.Execute(context.Background(), map[string]interface{}{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("expected error result for missing query parameter")
	}
}

func TestQueryToolExecuteEmptyQuery(t *testing.T) {
	tool := NewSuperInstanceQueryBuiltin()
	result, err := tool.Execute(context.Background(), map[string]interface{}{
		"query": "",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !result.IsError {
		t.Error("expected error result for empty query")
	}
}

func TestQueryToolExecuteConstraintQuery(t *testing.T) {
	tool := NewSuperInstanceQueryBuiltin()
	result, err := tool.Execute(context.Background(), map[string]interface{}{
		"query": "constraint repos",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.IsError {
		t.Error("unexpected error for valid query")
	}
	text := resultContentText(result)
	if !strings.Contains(text, "constraint") {
		t.Error("result should mention constraint")
	}
}

func TestQueryToolExecuteFleetQuery(t *testing.T) {
	tool := NewSuperInstanceQueryBuiltin()
	result, err := tool.Execute(context.Background(), map[string]interface{}{
		"query": "fleet status",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.IsError {
		t.Error("unexpected error for fleet query")
	}
	text := resultContentText(result)
	if !strings.Contains(text, "CCC") {
		t.Error("fleet query should mention CCC agent")
	}
}

func TestQueryToolExecuteDialsQuery(t *testing.T) {
	tool := NewSuperInstanceQueryBuiltin()
	result, err := tool.Execute(context.Background(), map[string]interface{}{
		"query": "tell me about the dials theory",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.IsError {
		t.Error("unexpected error for dials query")
	}
}

func TestQueryToolExecuteFluxQuery(t *testing.T) {
	tool := NewSuperInstanceQueryBuiltin()
	result, err := tool.Execute(context.Background(), map[string]interface{}{
		"query": "flux tensor midi",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.IsError {
		t.Error("unexpected error for flux query")
	}
	text := resultContentText(result)
	if !strings.Contains(text, "flux-tensor-midi") {
		t.Error("flux query should mention flux-tensor-midi")
	}
}

func TestQueryToolExecuteForgeQuery(t *testing.T) {
	tool := NewSuperInstanceQueryBuiltin()
	result, err := tool.Execute(context.Background(), map[string]interface{}{
		"query": "forgemaster build",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.IsError {
		t.Error("unexpected error for forge query")
	}
	text := resultContentText(result)
	if !strings.Contains(text, "forgemaster") {
		t.Error("forge query should mention forgemaster")
	}
}

func TestQueryToolExecuteAgentQuery(t *testing.T) {
	tool := NewSuperInstanceQueryBuiltin()
	result, err := tool.Execute(context.Background(), map[string]interface{}{
		"query": "agent trinity scores",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.IsError {
		t.Error("unexpected error for agent query")
	}
	text := resultContentText(result)
	if !strings.Contains(text, "Trinity") {
		t.Error("agent query should mention Trinity")
	}
}

func TestQueryToolExecuteConservationQuery(t *testing.T) {
	tool := NewSuperInstanceQueryBuiltin()
	result, err := tool.Execute(context.Background(), map[string]interface{}{
		"query": "conservation hypothesis",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.IsError {
		t.Error("unexpected error for conservation query")
	}
	text := resultContentText(result)
	if !strings.Contains(text, "0.436") {
		t.Error("conservation query should reference correlation")
	}
}

func TestQueryToolExecuteInnovationQuery(t *testing.T) {
	tool := NewSuperInstanceQueryBuiltin()
	result, err := tool.Execute(context.Background(), map[string]interface{}{
		"query": "innovation cycle",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.IsError {
		t.Error("unexpected error for innovation query")
	}
	text := resultContentText(result)
	if !strings.Contains(text, "Codification") {
		t.Error("innovation query should mention current phase")
	}
}

func TestQueryToolExecuteCreativeQuery(t *testing.T) {
	tool := NewSuperInstanceQueryBuiltin()
	result, err := tool.Execute(context.Background(), map[string]interface{}{
		"query": "creative engine",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.IsError {
		t.Error("unexpected error for creative query")
	}
	text := resultContentText(result)
	if !strings.Contains(text, "creative") {
		t.Error("creative query should mention creative")
	}
}

func TestQueryToolExecuteEcosystemQuery(t *testing.T) {
	tool := NewSuperInstanceQueryBuiltin()
	result, err := tool.Execute(context.Background(), map[string]interface{}{
		"query": "show me the superinstance ecosystem",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.IsError {
		t.Error("unexpected error for ecosystem query")
	}
	text := resultContentText(result)
	if !strings.Contains(text, "Related Repos") {
		t.Error("ecosystem query should list repos")
	}
}

func TestQueryToolExecuteUnknownQuery(t *testing.T) {
	tool := NewSuperInstanceQueryBuiltin()
	result, err := tool.Execute(context.Background(), map[string]interface{}{
		"query": "xyzzy foobar baz quux",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// Unknown queries should still return results (with suggestions)
	if result.IsError {
		t.Error("unknown query should not return an error")
	}
}

// --- Repo Status Tool Execute tests ---

func TestRepoStatusToolAllRepos(t *testing.T) {
	tool := NewSuperInstanceRepoStatusBuiltin()
	result, err := tool.Execute(context.Background(), map[string]interface{}{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.IsError {
		t.Error("unexpected error")
	}
	text := resultContentText(result)
	if !strings.Contains(text, "openagent") {
		t.Error("repo status should include openagent")
	}
}

func TestRepoStatusToolFilterByLanguage(t *testing.T) {
	tool := NewSuperInstanceRepoStatusBuiltin()
	result, err := tool.Execute(context.Background(), map[string]interface{}{
		"filter": "Rust",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.IsError {
		t.Error("unexpected error")
	}
	text := resultContentText(result)
	if !strings.Contains(text, "constraint-audio") {
		t.Error("Rust filter should include constraint-audio")
	}
}

func TestRepoStatusToolFilterByTag(t *testing.T) {
	tool := NewSuperInstanceRepoStatusBuiltin()
	result, err := tool.Execute(context.Background(), map[string]interface{}{
		"filter": "music",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.IsError {
		t.Error("unexpected error")
	}
	text := resultContentText(result)
	if !strings.Contains(text, "constraint-synth") {
		t.Error("music tag filter should include constraint-synth")
	}
}

func TestRepoStatusToolJSONFormat(t *testing.T) {
	tool := NewSuperInstanceRepoStatusBuiltin()
	result, err := tool.Execute(context.Background(), map[string]interface{}{
		"format": "json",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.IsError {
		t.Error("unexpected error")
	}
	text := resultContentText(result)
	if !strings.Contains(text, `"name"`) {
		t.Error("JSON format should contain name field")
	}
}

func TestRepoStatusToolFilterJSON(t *testing.T) {
	tool := NewSuperInstanceRepoStatusBuiltin()
	result, err := tool.Execute(context.Background(), map[string]interface{}{
		"filter": "Go",
		"format": "json",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.IsError {
		t.Error("unexpected error")
	}
	text := resultContentText(result)
	if !strings.Contains(text, "openagent") {
		t.Error("Go filter should include openagent")
	}
}

// --- Fleet Tool Execute tests ---

func TestFleetToolSummary(t *testing.T) {
	tool := NewSuperInstanceFleetBuiltin()
	result, err := tool.Execute(context.Background(), map[string]interface{}{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.IsError {
		t.Error("unexpected error")
	}
	text := resultContentText(result)
	if !strings.Contains(text, "CCC") {
		t.Error("fleet summary should mention CCC")
	}
	if !strings.Contains(text, "Phase") {
		t.Error("fleet summary should mention phase")
	}
}

func TestFleetToolSummaryExplicit(t *testing.T) {
	tool := NewSuperInstanceFleetBuiltin()
	result, err := tool.Execute(context.Background(), map[string]interface{}{
		"detail": "summary",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.IsError {
		t.Error("unexpected error")
	}
	text := resultContentText(result)
	if !strings.Contains(text, "|") {
		t.Error("summary format should contain table rows")
	}
}

func TestFleetToolFull(t *testing.T) {
	tool := NewSuperInstanceFleetBuiltin()
	result, err := tool.Execute(context.Background(), map[string]interface{}{
		"detail": "full",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.IsError {
		t.Error("unexpected error")
	}
	text := resultContentText(result)
	if !strings.Contains(text, "Ethos") {
		t.Error("full detail should include Ethos")
	}
	if !strings.Contains(text, "Conservation") {
		t.Error("full detail should include Conservation theory")
	}
	if !strings.Contains(text, "Dials") {
		t.Error("full detail should include Dials theory")
	}
}

// --- SuperInstanceTool tests ---

func TestSuperInstanceToolBuiltinTools(t *testing.T) {
	var st SuperInstanceTool
	if st.BuiltinTools() != nil {
		t.Error("SuperInstanceTool.BuiltinTools() should return nil")
	}
}

// --- Helper functions tests ---

func TestHelperTextResult(t *testing.T) {
	result := textResult("hello world")
	if result.IsError {
		t.Error("textResult should not be error")
	}
	text := resultContentText(result)
	if text != "hello world" {
		t.Errorf("textResult text = %q, want %q", text, "hello world")
	}
}

func TestHelperErrorResult(t *testing.T) {
	result := errorResult("something broke")
	if !result.IsError {
		t.Error("errorResult should be error")
	}
	text := resultContentText(result)
	if text != "something broke" {
		t.Errorf("errorResult text = %q, want %q", text, "something broke")
	}
}

// --- Integration: verify all tools work with the registry ---

func TestAllSuperInstanceToolsHaveRequiredFields(t *testing.T) {
	tools := []struct {
		name string
		tool interface {
			GetName() string
			GetDescription() string
			GetInputSchema() interface{}
		}
	}{
		{"query", NewSuperInstanceQueryBuiltin()},
		{"repo_status", NewSuperInstanceRepoStatusBuiltin()},
		{"fleet", NewSuperInstanceFleetBuiltin()},
	}

	for _, tt := range tools {
		t.Run(tt.name, func(t *testing.T) {
			if tt.tool.GetName() == "" {
				t.Error("tool name is empty")
			}
			if tt.tool.GetDescription() == "" {
				t.Error("tool description is empty")
			}
			schema := tt.tool.GetInputSchema()
			if schema == nil {
				t.Error("tool input schema is nil")
			}
		})
	}
}

func TestQueryToolReturnsRelevantRepos(t *testing.T) {
	// Verify that queries for specific repos actually return those repos
	tool := NewSuperInstanceQueryBuiltin()

	tests := []struct {
		query     string
		wantRepo  string
	}{
		{"constraint synth", "constraint-synth"},
		{"fleet monitoring ccc", "ccc-os"},
		{"forgemaster", "forgemaster"},
		{"creative engine rust", "creative-engine-rust"},
	}

	for _, tt := range tests {
		t.Run(tt.query, func(t *testing.T) {
			result, err := tool.Execute(context.Background(), map[string]interface{}{
				"query": tt.query,
			})
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if result.IsError {
				t.Fatalf("unexpected error result")
			}
			text := resultContentText(result)
			if !strings.Contains(text, tt.wantRepo) {
				t.Errorf("query %q: result does not contain %q\nGot: %s", tt.query, tt.wantRepo, text[:min(200, len(text))])
			}
		})
	}
}

func TestRepoStatusToolFilterByNonexistentLanguage(t *testing.T) {
	tool := NewSuperInstanceRepoStatusBuiltin()
	result, err := tool.Execute(context.Background(), map[string]interface{}{
		"filter": "COBOL",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// Should not error, just return empty or all repos
	if result.IsError {
		t.Error("filtering by non-existent language should not error")
	}
}

func TestFleetFullDetailTheories(t *testing.T) {
	tool := NewSuperInstanceFleetBuiltin()
	result, err := tool.Execute(context.Background(), map[string]interface{}{
		"detail": "full",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	text := resultContentText(result)

	// Verify conservation theory details
	if !strings.Contains(text, "hypothesis") {
		t.Error("full fleet detail should show conservation status")
	}

	// Verify innovation cycle
	if !strings.Contains(text, "Discovery") {
		t.Error("full fleet detail should show innovation phases")
	}

	// Verify agent details
	if !strings.Contains(text, "Oracle1") {
		t.Error("full fleet detail should mention Oracle1")
	}
}

// --- Ecosystem data consistency tests via tools ---

func TestQueryToolEcosystemReposMatchSuperinstance(t *testing.T) {
	tool := NewSuperInstanceQueryBuiltin()
	result, err := tool.Execute(context.Background(), map[string]interface{}{
		"query": "superinstance overview",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	text := resultContentText(result)

	// Check that key repos appear
	keyRepos := []string{"openagent", "sunset-ecosystem", "ccc-os", "forgemaster"}
	for _, repo := range keyRepos {
		if !strings.Contains(text, repo) {
			t.Errorf("ecosystem overview should mention %s", repo)
		}
	}
}

func TestRepoStatusAllReposCount(t *testing.T) {
	tool := NewSuperInstanceRepoStatusBuiltin()
	result, err := tool.Execute(context.Background(), map[string]interface{}{
		"format": "json",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	text := resultContentText(result)
	// Count occurrences of "name" field to estimate repo count
	count := strings.Count(text, `"name"`)
	expected := len(superinstance.Repos)
	if count != expected {
		t.Errorf("repo count: got %d name fields, expected %d repos", count, expected)
	}
}

// --- resultContentText helper ---

func resultContentText(result *protocol.CallToolResult) string {
	for _, c := range result.Content {
		if tc, ok := c.(*protocol.TextContent); ok {
			return tc.Text
		}
	}
	return ""
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
