package main

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func TestMCPServer(t *testing.T) {
	ctx := context.Background()

	server := mcp.NewServer(
		&mcp.Implementation{
			Name:    "aict",
			Version: "1.0.0",
		},
		nil,
	)

	for name, spec := range toolSpecs {
		schemaJSON, err := json.Marshal(spec.InputSchema)
		if err != nil {
			t.Logf("warning: failed to marshal schema for %s: %v", name, err)
			continue
		}

		var schemaMap map[string]interface{}
		if err := json.Unmarshal(schemaJSON, &schemaMap); err != nil {
			t.Logf("warning: failed to unmarshal schema for %s: %v", name, err)
			continue
		}

		server.AddTool(&mcp.Tool{
			Name:        spec.Name,
			Description: spec.Description,
			InputSchema: schemaMap,
		}, toolHandler)
	}

	t1, t2 := mcp.NewInMemoryTransports()
	serverSession, err := server.Connect(ctx, t1, nil)
	if err != nil {
		t.Fatalf("failed to connect server: %v", err)
	}
	defer serverSession.Close()

	client := mcp.NewClient(&mcp.Implementation{Name: "test-client", Version: "1.0.0"}, nil)
	clientSession, err := client.Connect(ctx, t2, nil)
	if err != nil {
		t.Fatalf("failed to connect client: %v", err)
	}
	defer clientSession.Close()

	tools, err := clientSession.ListTools(ctx, nil)
	if err != nil {
		t.Fatalf("failed to list tools: %v", err)
	}

	if len(tools.Tools) != 7 {
		t.Errorf("expected 7 tools, got %d", len(tools.Tools))
	}

	toolNames := make(map[string]bool)
	for _, tool := range tools.Tools {
		toolNames[tool.Name] = true
	}

	expectedTools := []string{"ls", "grep", "cat", "find", "stat", "wc", "diff"}
	for _, name := range expectedTools {
		if !toolNames[name] {
			t.Errorf("expected tool %s not found", name)
		}
	}

	t.Logf("All %d tools registered successfully\n", len(tools.Tools))
}

func TestLsHandler(t *testing.T) {
	tests := []struct {
		name     string
		args     map[string]interface{}
		expected []string
	}{
		{
			name:     "empty args",
			args:     map[string]interface{}{},
			expected: []string{},
		},
		{
			name:     "with path",
			args:     map[string]interface{}{"path": "/tmp"},
			expected: []string{"/tmp"},
		},
		{
			name:     "all flags",
			args:     map[string]interface{}{"all": true, "sortTime": true, "reverse": true},
			expected: []string{"-a", "-t", "-r"},
		},
		{
			name:     "path with flags",
			args:     map[string]interface{}{"path": ".", "all": true, "recursive": true},
			expected: []string{"-a", "-R", "."},
		},
		{
			name:     "almost all",
			args:     map[string]interface{}{"almostAll": true},
			expected: []string{"-A"},
		},
		{
			name:     "pretty flag",
			args:     map[string]interface{}{"pretty": true},
			expected: []string{"--pretty"},
		},
		{
			name:     "help flag",
			args:     map[string]interface{}{"help": true},
			expected: []string{"-h"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := lsHandler(tt.args)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if len(result) != len(tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
			for i, v := range tt.expected {
				if i >= len(result) || result[i] != v {
					t.Errorf("expected %v, got %v", tt.expected, result)
				}
			}
		})
	}
}

func TestGrepHandler(t *testing.T) {
	tests := []struct {
		name     string
		args     map[string]interface{}
		expected []string
	}{
		{
			name:     "pattern only",
			args:     map[string]interface{}{"pattern": "func"},
			expected: []string{"func", "."},
		},
		{
			name:     "pattern with path",
			args:     map[string]interface{}{"pattern": "func", "path": "/src"},
			expected: []string{"func", "/src"},
		},
		{
			name:     "pattern with flags",
			args:     map[string]interface{}{"pattern": "test", "recursive": true, "caseInsensitive": true},
			expected: []string{"-r", "-i", "test", "."},
		},
		{
			name:     "with context",
			args:     map[string]interface{}{"pattern": "foo", "context": 3.0},
			expected: []string{"foo", ".", "-C", "3"},
		},
		{
			name:     "with include",
			args:     map[string]interface{}{"pattern": "bar", "include": "*.go"},
			expected: []string{"bar", ".", "--include", "*.go"},
		},
		{
			name:     "help flag",
			args:     map[string]interface{}{"pattern": "x", "help": true},
			expected: []string{"-h", "x", "."},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := grepHandler(tt.args)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if len(result) != len(tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
			for i, v := range tt.expected {
				if i >= len(result) || result[i] != v {
					t.Errorf("expected %v, got %v", tt.expected, result)
				}
			}
		})
	}
}

func TestCatHandler(t *testing.T) {
	tests := []struct {
		name     string
		args     map[string]interface{}
		expected []string
	}{
		{
			name:     "path only",
			args:     map[string]interface{}{"path": "/etc/hosts"},
			expected: []string{"/etc/hosts"},
		},
		{
			name:     "path with line numbers",
			args:     map[string]interface{}{"path": "file.txt", "lineNumbers": true},
			expected: []string{"-n", "file.txt"},
		},
		{
			name:     "help flag",
			args:     map[string]interface{}{"path": "x", "help": true},
			expected: []string{"-h", "x"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := catHandler(tt.args)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if len(result) != len(tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
			for i, v := range tt.expected {
				if i >= len(result) || result[i] != v {
					t.Errorf("expected %v, got %v", tt.expected, result)
				}
			}
		})
	}
}

func TestFindHandler(t *testing.T) {
	tests := []struct {
		name     string
		args     map[string]interface{}
		expected []string
	}{
		{
			name:     "empty args",
			args:     map[string]interface{}{},
			expected: []string{"."},
		},
		{
			name:     "with name",
			args:     map[string]interface{}{"name": "*.go"},
			expected: []string{".", "-name", "*.go"},
		},
		{
			name:     "with type",
			args:     map[string]interface{}{"type": "f"},
			expected: []string{".", "-type", "f"},
		},
		{
			name:     "with mtime",
			args:     map[string]interface{}{"mtime": 7.0},
			expected: []string{".", "-mtime", "7"},
		},
		{
			name:     "with maxDepth",
			args:     map[string]interface{}{"maxDepth": 3.0},
			expected: []string{".", "-maxdepth", "3"},
		},
		{
			name:     "with path",
			args:     map[string]interface{}{"path": "/src", "name": "*.txt"},
			expected: []string{"/src", "-name", "*.txt"},
		},
		{
			name:     "help flag",
			args:     map[string]interface{}{"help": true},
			expected: []string{"-h", "."},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := findHandler(tt.args)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if len(result) != len(tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
			for i, v := range tt.expected {
				if i >= len(result) || result[i] != v {
					t.Errorf("expected %v, got %v", tt.expected, result)
				}
			}
		})
	}
}

func TestStatHandler(t *testing.T) {
	tests := []struct {
		name     string
		args     map[string]interface{}
		expected []string
	}{
		{
			name:     "path only",
			args:     map[string]interface{}{"path": "/etc/hosts"},
			expected: []string{"/etc/hosts"},
		},
		{
			name:     "help flag",
			args:     map[string]interface{}{"path": "x", "help": true},
			expected: []string{"-h", "x"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := statHandler(tt.args)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if len(result) != len(tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
			for i, v := range tt.expected {
				if i >= len(result) || result[i] != v {
					t.Errorf("expected %v, got %v", tt.expected, result)
				}
			}
		})
	}
}

func TestWcHandler(t *testing.T) {
	tests := []struct {
		name     string
		args     map[string]interface{}
		expected []string
	}{
		{
			name:     "path only",
			args:     map[string]interface{}{"path": "file.txt"},
			expected: []string{"file.txt"},
		},
		{
			name:     "path with lines",
			args:     map[string]interface{}{"path": "file.txt", "lines": true},
			expected: []string{"-l", "file.txt"},
		},
		{
			name:     "all count flags",
			args:     map[string]interface{}{"path": "f", "bytes": true, "words": true, "lines": true},
			expected: []string{"-c", "-w", "-l", "f"},
		},
		{
			name:     "help flag",
			args:     map[string]interface{}{"path": "x", "help": true},
			expected: []string{"-h", "x"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := wcHandler(tt.args)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if len(result) != len(tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
			for i, v := range tt.expected {
				if i >= len(result) || result[i] != v {
					t.Errorf("expected %v, got %v", tt.expected, result)
				}
			}
		})
	}
}

func TestDiffHandler(t *testing.T) {
	tests := []struct {
		name     string
		args     map[string]interface{}
		expected []string
	}{
		{
			name:     "two paths",
			args:     map[string]interface{}{"path1": "a.txt", "path2": "b.txt"},
			expected: []string{"a.txt", "b.txt"},
		},
		{
			name:     "brief flag",
			args:     map[string]interface{}{"path1": "a", "path2": "b", "brief": true},
			expected: []string{"--brief", "a", "b"},
		},
		{
			name:     "help flag",
			args:     map[string]interface{}{"path1": "a", "path2": "b", "help": true},
			expected: []string{"-h", "a", "b"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := diffHandler(tt.args)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if len(result) != len(tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
			for i, v := range tt.expected {
				if i >= len(result) || result[i] != v {
					t.Errorf("expected %v, got %v", tt.expected, result)
				}
			}
		})
	}
}

func TestToBool(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected bool
	}{
		{"bool true", true, true},
		{"bool false", false, false},
		{"string true", "true", true},
		{"string 1", "1", true},
		{"string false", "false", false},
		{"string 0", "0", false},
		{"string empty", "", false},
		{"nil", nil, false},
		{"int", 1, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := toBool(tt.input)
			if result != tt.expected {
				t.Errorf("toBool(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestParseArgs(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected map[string]interface{}
	}{
		{
			name:     "nil",
			input:    nil,
			expected: map[string]interface{}{},
		},
		{
			name:     "map",
			input:    map[string]interface{}{"path": "/tmp"},
			expected: map[string]interface{}{"path": "/tmp"},
		},
		{
			name:     "json raw",
			input:    json.RawMessage(`{"path": "/tmp"}`),
			expected: map[string]interface{}{"path": "/tmp"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parseArgs(tt.input)
			if len(result) != len(tt.expected) {
				t.Errorf("parseArgs(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestToolHandler(t *testing.T) {
	binaryPath := findAICTBinary()
	if _, err := os.Stat(binaryPath); err != nil {
		t.Skipf("aict binary not found at %s, skipping integration test", binaryPath)
	}

	ctx := context.Background()

	server := mcp.NewServer(
		&mcp.Implementation{Name: "aict", Version: "1.0.0"},
		nil,
	)

	for _, spec := range toolSpecs {
		schemaJSON, _ := json.Marshal(spec.InputSchema)
		var schemaMap map[string]interface{}
		json.Unmarshal(schemaJSON, &schemaMap)

		server.AddTool(&mcp.Tool{
			Name:        spec.Name,
			Description: spec.Description,
			InputSchema: schemaMap,
		}, toolHandler)
	}

	t1, t2 := mcp.NewInMemoryTransports()
	serverSession, _ := server.Connect(ctx, t1, nil)
	defer serverSession.Close()

	client := mcp.NewClient(&mcp.Implementation{Name: "test", Version: "1.0.0"}, nil)
	clientSession, _ := client.Connect(ctx, t2, nil)
	defer clientSession.Close()

	result, err := clientSession.CallTool(ctx, &mcp.CallToolParams{
		Name:      "ls",
		Arguments: map[string]interface{}{"path": "."},
	})
	if err != nil {
		t.Fatalf("CallTool failed: %v", err)
	}

	if result.IsError {
		t.Errorf("tool returned error: %v", result.Content)
	}

	if len(result.Content) == 0 {
		t.Errorf("expected content, got empty")
	}
}

func TestLsIntegration(t *testing.T) {
	binaryPath := findAICTBinary()
	if _, err := os.Stat(binaryPath); err != nil {
		t.Skipf("aict binary not found at %s, skipping integration test", binaryPath)
	}

	dir := t.TempDir()
	testFile := filepath.Join(dir, "test.txt")
	if err := os.WriteFile(testFile, []byte("hello"), 0644); err != nil {
		t.Fatalf("failed to create test file: %v", err)
	}

	result, err := runAICT([]string{"ls", dir})
	if err != nil {
		t.Fatalf("runAICT failed: %v", err)
	}

	var output map[string]interface{}
	if err := json.Unmarshal([]byte(result), &output); err != nil {
		t.Fatalf("invalid JSON output: %v", err)
	}

	if output["ls"] == nil {
		t.Errorf("expected ls key in output, got %v", output)
	}
}

func TestFindAICTBinary(t *testing.T) {
	binary := findAICTBinary()

	if _, err := os.Stat(binary); err != nil {
		t.Skipf("aict not found at %s", binary)
	}

	t.Logf("found aict at: %s", binary)
}
