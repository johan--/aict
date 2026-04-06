package main

import (
	"encoding/json"
	"fmt"

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
	meta := tool.AllMeta()

	for name, m := range meta {
		fmt.Printf("=== %s ===\n", name)
		jsonBytes, err := json.MarshalIndent(m.InputSchema, "", "  ")
		if err != nil {
			fmt.Printf("Error marshaling: %v\n", err)
		} else {
			fmt.Println(string(jsonBytes))
		}
		fmt.Println()
	}
}
