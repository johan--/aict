package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	_ "github.com/synseqack/aict/tools/cat"
	_ "github.com/synseqack/aict/tools/diff"
	_ "github.com/synseqack/aict/tools/find"
	_ "github.com/synseqack/aict/tools/grep"
	_ "github.com/synseqack/aict/tools/ls"
	_ "github.com/synseqack/aict/tools/stat"
	_ "github.com/synseqack/aict/tools/wc"
)

type ToolSpec struct {
	Name        string
	Description string
	InputSchema map[string]interface{}
	Handler     func(args map[string]interface{}) ([]string, error)
}

type param struct {
	key  string
	flag string
}

type stringParam struct {
	key  string
	flag string
}

type intParam struct {
	key  string
	flag string
}

var toolSpecs = map[string]ToolSpec{
	"ls": {
		Name:        "ls",
		Description: "List directory contents with file metadata including permissions, size, and modification time",
		InputSchema: map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"path":      map[string]interface{}{"type": "string", "description": "Directory or file path to list"},
				"all":       map[string]interface{}{"type": "boolean", "description": "Show hidden files (starting with .)"},
				"almostAll": map[string]interface{}{"type": "boolean", "description": "Show almost all (exclude . and ..)"},
				"sortTime":  map[string]interface{}{"type": "boolean", "description": "Sort by modification time, newest first"},
				"reverse":   map[string]interface{}{"type": "boolean", "description": "Reverse sort order"},
				"recursive": map[string]interface{}{"type": "boolean", "description": "List subdirectories recursively"},
				"pretty":    map[string]interface{}{"type": "boolean", "description": "Pretty print XML output"},
				"help":      map[string]interface{}{"type": "boolean", "description": "Show help"},
			},
		},
		Handler: lsHandler,
	},
	"grep": {
		Name:        "grep",
		Description: "Search for patterns in files with line numbers and context",
		InputSchema: map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"pattern":          map[string]interface{}{"type": "string", "description": "Search pattern (regex or literal)"},
				"path":             map[string]interface{}{"type": "string", "description": "File or directory to search (default: current directory)"},
				"recursive":        map[string]interface{}{"type": "boolean", "description": "Search recursively"},
				"lineNumbers":      map[string]interface{}{"type": "boolean", "description": "Show line numbers"},
				"filesWithMatches": map[string]interface{}{"type": "boolean", "description": "Show only matching file names"},
				"caseInsensitive":  map[string]interface{}{"type": "boolean", "description": "Case insensitive search"},
				"wordMatch":        map[string]interface{}{"type": "boolean", "description": "Match whole words only"},
				"context":          map[string]interface{}{"type": "integer", "description": "Number of context lines to show"},
				"countOnly":        map[string]interface{}{"type": "boolean", "description": "Count matches only"},
				"include":          map[string]interface{}{"type": "string", "description": "Include files matching pattern (e.g., *.go)"},
				"help":             map[string]interface{}{"type": "boolean", "description": "Show help"},
			},
			"required": []string{"pattern"},
		},
		Handler: grepHandler,
	},
	"cat": {
		Name:        "cat",
		Description: "Read and output file contents with metadata",
		InputSchema: map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"path":        map[string]interface{}{"type": "string", "description": "File path to read"},
				"lineNumbers": map[string]interface{}{"type": "boolean", "description": "Show line numbers"},
				"help":        map[string]interface{}{"type": "boolean", "description": "Show help"},
			},
			"required": []string{"path"},
		},
		Handler: catHandler,
	},
	"find": {
		Name:        "find",
		Description: "Find files by name, type, or modification time",
		InputSchema: map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"path":     map[string]interface{}{"type": "string", "description": "Search root directory (default: current directory)"},
				"name":     map[string]interface{}{"type": "string", "description": "File name pattern (supports * and ?)"},
				"type":     map[string]interface{}{"type": "string", "description": "File type: f (regular), d (directory), l (symlink)"},
				"mtime":    map[string]interface{}{"type": "integer", "description": "Modified within N days"},
				"maxDepth": map[string]interface{}{"type": "integer", "description": "Maximum directory depth"},
				"invert":   map[string]interface{}{"type": "boolean", "description": "Invert match conditions"},
				"help":     map[string]interface{}{"type": "boolean", "description": "Show help"},
			},
		},
		Handler: findHandler,
	},
	"stat": {
		Name:        "stat",
		Description: "Display detailed file metadata including timestamps, permissions, and ownership",
		InputSchema: map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"path": map[string]interface{}{"type": "string", "description": "File path to stat"},
				"help": map[string]interface{}{"type": "boolean", "description": "Show help"},
			},
			"required": []string{"path"},
		},
		Handler: statHandler,
	},
	"wc": {
		Name:        "wc",
		Description: "Count lines, words, and bytes in files",
		InputSchema: map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"path":  map[string]interface{}{"type": "string", "description": "File path to count"},
				"bytes": map[string]interface{}{"type": "boolean", "description": "Count bytes"},
				"words": map[string]interface{}{"type": "boolean", "description": "Count words"},
				"lines": map[string]interface{}{"type": "boolean", "description": "Count lines"},
				"help":  map[string]interface{}{"type": "boolean", "description": "Show help"},
			},
			"required": []string{"path"},
		},
		Handler: wcHandler,
	},
	"diff": {
		Name:        "diff",
		Description: "Compare two files or directories and show differences",
		InputSchema: map[string]interface{}{
			"type": "object",
			"properties": map[string]interface{}{
				"path1": map[string]interface{}{"type": "string", "description": "First file or directory"},
				"path2": map[string]interface{}{"type": "string", "description": "Second file or directory"},
				"brief": map[string]interface{}{"type": "boolean", "description": "Output only whether files differ"},
				"help":  map[string]interface{}{"type": "boolean", "description": "Show help"},
			},
			"required": []string{"path1", "path2"},
		},
		Handler: diffHandler,
	},
}

func buildBoolArgs(params []param, args map[string]interface{}) []string {
	var result []string
	for _, p := range params {
		if toBool(args[p.key]) && p.flag != "" {
			result = append(result, p.flag)
		}
	}
	return result
}

func buildStringArgs(params []stringParam, args map[string]interface{}) []string {
	var result []string
	for _, p := range params {
		if v, ok := args[p.key].(string); ok && v != "" && p.flag != "" {
			result = append(result, p.flag, v)
		}
	}
	return result
}

func buildIntArgs(params []intParam, args map[string]interface{}) []string {
	var result []string
	for _, p := range params {
		if v, ok := args[p.key].(float64); ok && p.flag != "" {
			result = append(result, p.flag, fmt.Sprintf("%d", int(v)))
		}
	}
	return result
}

func getString(args map[string]interface{}, key, defaultVal string) string {
	if v, ok := args[key].(string); ok && v != "" {
		return v
	}
	return defaultVal
}

func lsHandler(args map[string]interface{}) ([]string, error) {
	result := buildBoolArgs([]param{
		{key: "all", flag: "-a"},
		{key: "almostAll", flag: "-A"},
		{key: "sortTime", flag: "-t"},
		{key: "reverse", flag: "-r"},
		{key: "recursive", flag: "-R"},
		{key: "pretty", flag: "--pretty"},
		{key: "help", flag: "-h"},
	}, args)

	if path := getString(args, "path", ""); path != "" {
		result = append(result, path)
	}

	return result, nil
}

func grepHandler(args map[string]interface{}) ([]string, error) {
	result := buildBoolArgs([]param{
		{key: "recursive", flag: "-r"},
		{key: "lineNumbers", flag: "-n"},
		{key: "filesWithMatches", flag: "-l"},
		{key: "caseInsensitive", flag: "-i"},
		{key: "wordMatch", flag: "-w"},
		{key: "countOnly", flag: "-c"},
		{key: "help", flag: "-h"},
	}, args)

	if pattern := getString(args, "pattern", ""); pattern != "" {
		result = append(result, pattern)
	}

	result = append(result, getString(args, "path", "."))

	result = append(result, buildIntArgs([]intParam{
		{key: "context", flag: "-C"},
	}, args)...)

	result = append(result, buildStringArgs([]stringParam{
		{key: "include", flag: "--include"},
	}, args)...)

	return result, nil
}

func catHandler(args map[string]interface{}) ([]string, error) {
	result := buildBoolArgs([]param{
		{key: "lineNumbers", flag: "-n"},
		{key: "help", flag: "-h"},
	}, args)

	if path := getString(args, "path", ""); path != "" {
		result = append(result, path)
	}

	return result, nil
}

func findHandler(args map[string]interface{}) ([]string, error) {
	result := buildBoolArgs([]param{
		{key: "invert", flag: "!"},
		{key: "help", flag: "-h"},
	}, args)

	result = append(result, getString(args, "path", "."))

	result = append(result, buildStringArgs([]stringParam{
		{key: "name", flag: "-name"},
		{key: "type", flag: "-type"},
	}, args)...)

	result = append(result, buildIntArgs([]intParam{
		{key: "mtime", flag: "-mtime"},
		{key: "maxDepth", flag: "-maxdepth"},
	}, args)...)

	return result, nil
}

func statHandler(args map[string]interface{}) ([]string, error) {
	result := buildBoolArgs([]param{
		{key: "help", flag: "-h"},
	}, args)

	if path := getString(args, "path", ""); path != "" {
		result = append(result, path)
	}

	return result, nil
}

func wcHandler(args map[string]interface{}) ([]string, error) {
	result := buildBoolArgs([]param{
		{key: "bytes", flag: "-c"},
		{key: "words", flag: "-w"},
		{key: "lines", flag: "-l"},
		{key: "help", flag: "-h"},
	}, args)

	if path := getString(args, "path", ""); path != "" {
		result = append(result, path)
	}

	return result, nil
}

func diffHandler(args map[string]interface{}) ([]string, error) {
	result := buildBoolArgs([]param{
		{key: "brief", flag: "--brief"},
		{key: "help", flag: "-h"},
	}, args)

	if path1 := getString(args, "path1", ""); path1 != "" {
		result = append(result, path1)
	}
	if path2 := getString(args, "path2", ""); path2 != "" {
		result = append(result, path2)
	}

	return result, nil
}

func toBool(v interface{}) bool {
	if b, ok := v.(bool); ok {
		return b
	}
	if s, ok := v.(string); ok {
		return s == "true" || s == "1"
	}
	return false
}

func findAICTBinary() string {
	if binaryPath := os.Getenv("AICT_BINARY"); binaryPath != "" {
		return binaryPath
	}

	execPath, err := os.Executable()
	if err == nil {
		dir := filepath.Dir(execPath)
		candidate := filepath.Join(dir, "aict")
		if _, err := os.Stat(candidate); err == nil {
			return candidate
		}
	}

	dir, err := os.Getwd()
	if err == nil {
		candidate := filepath.Join(dir, "aict")
		if _, err := os.Stat(candidate); err == nil {
			return candidate
		}
	}

	return "aict"
}

func runAICT(args []string) (string, error) {
	binaryPath := findAICTBinary()

	aictArgs := append([]string{"--json"}, args...)

	cmd := exec.Command(binaryPath, aictArgs...)
	output, err := cmd.Output()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			return string(exitErr.Stderr), fmt.Errorf("aict error: %s", strings.TrimSpace(string(exitErr.Stderr)))
		}
		return "", fmt.Errorf("failed to run aict: %w", err)
	}

	return string(output), nil
}

func toolHandler(ctx context.Context, req *mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	toolName := req.Params.Name

	spec, ok := toolSpecs[toolName]
	if !ok {
		return errorResult(fmt.Sprintf("unknown tool: %s", toolName)), nil
	}

	argsMap := parseArgs(req.Params.Arguments)

	if toBool(argsMap["help"]) {
		return errorResult(fmt.Sprintf("usage: aict %s [options]", toolName)), nil
	}

	aictArgs, err := spec.Handler(argsMap)
	if err != nil {
		return errorResult(fmt.Sprintf("error building args: %v", err)), nil
	}

	output, err := runAICT(aictArgs)
	if err != nil {
		return errorResult(err.Error()), nil
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: output},
		},
	}, nil
}

func errorResult(msg string) *mcp.CallToolResult {
	return &mcp.CallToolResult{
		IsError: true,
		Content: []mcp.Content{
			&mcp.TextContent{Text: msg},
		},
	}
}

func parseArgs(args any) map[string]interface{} {
	argsMap := make(map[string]interface{})
	if args == nil {
		return argsMap
	}
	data, err := json.Marshal(args)
	if err == nil {
		_ = json.Unmarshal(data, &argsMap)
	}
	return argsMap
}

func main() {
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
			log.Printf("warning: failed to marshal schema for %s: %v", name, err)
			continue
		}

		var schemaMap map[string]interface{}
		if err := json.Unmarshal(schemaJSON, &schemaMap); err != nil {
			log.Printf("warning: failed to unmarshal schema for %s: %v", name, err)
			continue
		}

		server.AddTool(&mcp.Tool{
			Name:        spec.Name,
			Description: spec.Description,
			InputSchema: schemaMap,
		}, toolHandler)
	}

	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Printf("MCP server error: %v", err)
	}
}
