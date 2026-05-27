// Copyright 2026 The SuperInstance Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package builtin_tool

import (
	"context"
	"testing"

	"github.com/ThinkInAIXYZ/go-mcp/protocol"
)

// mockTool implements BuiltinTool for testing the registry
type mockTool struct {
	name        string
	description string
	schema      interface{}
	executeErr  error
}

func (m *mockTool) GetName() string                { return m.name }
func (m *mockTool) GetDescription() string          { return m.description }
func (m *mockTool) GetInputSchema() interface{}     { return m.schema }
func (m *mockTool) Execute(_ context.Context, args map[string]interface{}) (*protocol.CallToolResult, error) {
	if m.executeErr != nil {
		return nil, m.executeErr
	}
	return &protocol.CallToolResult{
		IsError: false,
		Content: []protocol.Content{
			&protocol.TextContent{Type: "text", Text: "mock result"},
		},
	}, nil
}

func TestNewToolRegistry(t *testing.T) {
	registry := NewToolRegistry()
	if registry == nil {
		t.Fatal("NewToolRegistry returned nil")
	}
	if len(registry.GetAllTools()) != 0 {
		t.Error("new registry should have no tools")
	}
}

func TestRegisterAndGetTool(t *testing.T) {
	registry := NewToolRegistry()
	tool := &mockTool{name: "test_tool", description: "A test tool", schema: map[string]interface{}{}}

	registry.RegisterTool(tool)

	got, exists := registry.GetTool("test_tool")
	if !exists {
		t.Error("tool should exist after registration")
	}
	if got.GetName() != "test_tool" {
		t.Errorf("got tool name %q, want %q", got.GetName(), "test_tool")
	}
}

func TestGetToolNotFound(t *testing.T) {
	registry := NewToolRegistry()
	_, exists := registry.GetTool("nonexistent")
	if exists {
		t.Error("non-existent tool should not be found")
	}
}

func TestGetAllTools(t *testing.T) {
	registry := NewToolRegistry()
	registry.RegisterTool(&mockTool{name: "tool1"})
	registry.RegisterTool(&mockTool{name: "tool2"})
	registry.RegisterTool(&mockTool{name: "tool3"})

	tools := registry.GetAllTools()
	if len(tools) != 3 {
		t.Errorf("expected 3 tools, got %d", len(tools))
	}
}

func TestGetToolsAsProtocolTools(t *testing.T) {
	registry := NewToolRegistry()
	registry.RegisterTool(&mockTool{
		name:        "test_tool",
		description: "A test tool",
		schema: map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"input": map[string]string{"type": "string"},
			},
		},
	})

	protoTools := registry.GetToolsAsProtocolTools()
	if len(protoTools) != 1 {
		t.Fatalf("expected 1 protocol tool, got %d", len(protoTools))
	}
	if protoTools[0].Name != "test_tool" {
		t.Errorf("protocol tool name = %q, want %q", protoTools[0].Name, "test_tool")
	}
	if protoTools[0].Description != "A test tool" {
		t.Errorf("protocol tool description = %q, want %q", protoTools[0].Description, "A test tool")
	}
}

// Note: ExecuteTool is tested indirectly through the tool-specific tests
// since the registry.ExecuteTool signature returns the protocol types.

func TestRegisterToolOverwrite(t *testing.T) {
	registry := NewToolRegistry()
	registry.RegisterTool(&mockTool{name: "tool", description: "v1"})
	registry.RegisterTool(&mockTool{name: "tool", description: "v2"})

	got, _ := registry.GetTool("tool")
	if got.GetDescription() != "v2" {
		t.Errorf("overwritten tool description = %q, want %q", got.GetDescription(), "v2")
	}
}

func TestGetToolsAsProtocolToolsEmpty(t *testing.T) {
	registry := NewToolRegistry()
	protoTools := registry.GetToolsAsProtocolTools()
	if len(protoTools) != 0 {
		t.Errorf("empty registry should return empty protocol tools, got %d", len(protoTools))
	}
}
