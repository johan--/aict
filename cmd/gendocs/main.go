package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/synseqack/aict/internal/tool"
	_ "github.com/synseqack/aict/tools/basename"
	_ "github.com/synseqack/aict/tools/cat"
	_ "github.com/synseqack/aict/tools/checksums"
	_ "github.com/synseqack/aict/tools/cut"
	_ "github.com/synseqack/aict/tools/df"
	_ "github.com/synseqack/aict/tools/diff"
	_ "github.com/synseqack/aict/tools/dirname"
	_ "github.com/synseqack/aict/tools/doctor"
	_ "github.com/synseqack/aict/tools/du"
	_ "github.com/synseqack/aict/tools/env"
	_ "github.com/synseqack/aict/tools/file"
	_ "github.com/synseqack/aict/tools/find"
	_ "github.com/synseqack/aict/tools/git"
	_ "github.com/synseqack/aict/tools/grep"
	_ "github.com/synseqack/aict/tools/head"
	_ "github.com/synseqack/aict/tools/ls"
	_ "github.com/synseqack/aict/tools/ps"
	_ "github.com/synseqack/aict/tools/pwd"
	_ "github.com/synseqack/aict/tools/realpath"
	_ "github.com/synseqack/aict/tools/sort"
	_ "github.com/synseqack/aict/tools/stat"
	_ "github.com/synseqack/aict/tools/system"
	_ "github.com/synseqack/aict/tools/tail"
	_ "github.com/synseqack/aict/tools/tr"
	_ "github.com/synseqack/aict/tools/uniq"
	_ "github.com/synseqack/aict/tools/wc"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: gendocs <output-file>\n")
		os.Exit(1)
	}

	outputFile := os.Args[1]

	meta := tool.AllMeta()

	var names []string
	for name := range meta {
		names = append(names, name)
	}
	sort.Strings(names)

	var sb strings.Builder

	sb.WriteString("# Tools Reference\n\n")
	sb.WriteString("Complete reference for all `aict` commands. Each tool outputs structured XML/JSON by default, with optional plain text mode for compatibility.\n\n")
	sb.WriteString("## Common Flags\n\n")
	sb.WriteString("Every tool supports these global output flags:\n\n")
	sb.WriteString("| Flag | Description |\n")
	sb.WriteString("|------|-------------|\n")
	sb.WriteString("| `--xml` | XML output (default if `AICT_XML=1`) |\n")
	sb.WriteString("| `--json` | JSON output |\n")
	sb.WriteString("| `--plain` | Plain text output |\n")
	sb.WriteString("| `--pretty` | Pretty-printed output |\n\n")

	for _, name := range names {
		m := meta[name]
		sb.WriteString(fmt.Sprintf("## %s\n\n", name))
		sb.WriteString(fmt.Sprintf("%s\n\n", m.Description))

		sb.WriteString(fmt.Sprintf("```bash\n"))
		sb.WriteString(fmt.Sprintf("aict %s [flags] [arguments...]\n", name))
		sb.WriteString(fmt.Sprintf("```\n\n"))

		if m.InputSchema != nil {
			if props, ok := m.InputSchema["properties"].(map[string]interface{}); ok && len(props) > 0 {
				sb.WriteString("### Flags\n\n")
				sb.WriteString("| Flag | Type | Description |\n")
				sb.WriteString("|------|------|-------------|\n")

				var flagNames []string
				for flagName := range props {
					flagNames = append(flagNames, flagName)
				}
				sort.Strings(flagNames)

				for _, flagName := range flagNames {
					if flagName == "xml" || flagName == "json" || flagName == "plain" || flagName == "pretty" {
						continue
					}

					prop := props[flagName].(map[string]interface{})
					flagType := prop["type"].(string)
					desc := prop["description"].(string)

					flagDisplay := fmt.Sprintf("`--%s`", flagName)
					if len(flagName) == 1 {
						flagDisplay = fmt.Sprintf("`-%s`", flagName)
					}

					sb.WriteString(fmt.Sprintf("| %s | %s | %s |\n", flagDisplay, flagType, desc))
				}
				sb.WriteString("\n")
			}
		}
	}

	if err := os.WriteFile(outputFile, []byte(sb.String()), 0644); err != nil {
		fmt.Fprintf(os.Stderr, "error writing file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Generated documentation for %d tools -> %s\n", len(names), outputFile)
}
