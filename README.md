# aict

[![CI](https://github.com/synseqack/aict/actions/workflows/ci.yml/badge.svg)](https://github.com/synseqack/aict/actions)
[![Go Version](https://img.shields.io/github/go-mod/go-version/synseqack/aict)](https://github.com/synseqack/aict)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Latest Release](https://img.shields.io/github/v/release/synseqack/aict)](https://github.com/synseqack/aict/releases)

A CLI tool that outputs XML/JSON, built for AI agents to consume directly.

**Unix coreutils rebuilt for AI agents — structured XML/JSON output, zero parsing, zero ambiguity.**

[![asciicast](https://asciinema.org/a/placeholder.svg)](https://asciinema.org/a/placeholder)

> 📹 *Demo: `aict ls`, `aict grep`, `aict cat` in AI mode — replace this placeholder once a recording is published.*

## The Problem

When an AI agent runs `ls`, `grep`, or `cat`, it gets human-readable plaintext. The agent must parse column positions, guess which field is the filename, and handle inconsistent formats. This parsing is brittle and breaks easily.

## What You Get

```xml
$ aict ls src/
<ls timestamp="1234567890" total_entries="3">
  <file name="main.go" path="src/main.go" absolute="/project/src/main.go"
        size_bytes="2048" size_human="2.0 KiB" modified="1234567890" modified_ago_s="3600"
        language="go" mime="text/x-go" binary="false"/>
  <file name="utils.go" path="src/utils.go" absolute="/project/src/utils.go"
        size_bytes="1024" size_human="1.0 KiB" modified="1234567890" modified_ago_s="3600"
        language="go" mime="text/x-go" binary="false"/>
  <directory name="internal" path="src/internal"/>
</ls>
```

Every field is labeled. Every path is absolute. Every timestamp is a Unix epoch integer. The agent knows exactly what it is looking at.

## Install

```bash
# Install both binaries (recommended)
go install github.com/synseqack/aict@latest
go install github.com/synseqack/aict/cmd/mcp@latest
```

Or build from source:

```bash
git clone https://github.com/synseqack/aict
cd aict
go build -o aict .
go build -o aict-mcp ./cmd/mcp
```

## Usage

```bash
# AI mode (XML output)
AICT_XML=1 aict ls src/

# or with flag
aict ls src/ --xml

# Plain text when you need it
aict ls src/ --plain

# JSON for programmatic use
aict ls src/ --json
```

## Available Tools

**File Inspection**

| Tool | Description | MCP? |
|------|-------------|------|
| `cat` | File contents with encoding detection | ✅ |
| `head` / `tail` | First or last lines of a file | ✅ |
| `file` | File type detection via magic bytes | ✅ |
| `stat` | File metadata with all timestamps | ✅ |
| `wc` | Line, word, and byte counts | ✅ |

**Directory & Search**

| Tool | Description | MCP? |
|------|-------------|------|
| `ls` | Directory listings with language and MIME type detection | ✅ |
| `find` | Filesystem search by name, type, or modification time | ✅ |
| `grep` | Pattern search with context lines and recursive support | ✅ |
| `diff` | File and directory comparison | ✅ |

**Path Utilities**

| Tool | Description | MCP? |
|------|-------------|------|
| `realpath` | Resolve absolute paths | ✅ |
| `basename` | Extract filename from path | ✅ |
| `dirname` | Extract directory from path | ✅ |
| `pwd` | Print working directory | ✅ |

**Text Processing**

| Tool | Description | MCP? |
|------|-------------|------|
| `sort` | Sort lines with options | ✅ |
| `uniq` | Remove or count duplicate lines | ✅ |
| `cut` | Extract columns from delimited text | ✅ |
| `tr` | Translate or delete characters | ✅ |

**System & Environment**

| Tool | Description | MCP? |
|------|-------------|------|
| `env` | Environment variables with secret redaction | ✅ |
| `system` | OS, runtime, and user information | ✅ |
| `ps` | Running process list | ✅ |
| `df` | Disk space usage | ✅ |
| `du` | Directory size analysis | ✅ |
| `checksums` | MD5, SHA1, and SHA256 hashes | ✅ |

## MCP Server

`aict` ships a standalone MCP server binary so AI assistants can call every tool directly — no shell spawning, no output parsing.

```bash
# Build the MCP server
go build -o aict-mcp ./cmd/mcp

# Or install both binaries at once
go install github.com/synseqack/aict@latest
go install github.com/synseqack/aict/cmd/mcp@latest
```

### Claude (claude.ai / Claude Desktop)

Add to your Claude MCP config (`~/.config/claude/claude_desktop_config.json` or via Settings → MCP):

```json
{
  "mcpServers": {
    "aict": {
      "command": "aict-mcp",
      "args": []
    }
  }
}
```

### Cursor

Add to `.cursor/mcp.json` in your project root (or global `~/.cursor/mcp.json`):

```json
{
  "mcpServers": {
    "aict": {
      "command": "aict-mcp",
      "args": []
    }
  }
}
```

Once connected, every `aict` tool (`ls`, `grep`, `cat`, `find`, `stat`, `diff`, `git`, `ps`, `checksums`, etc.) becomes a typed, callable function. The model receives structured JSON — not raw terminal output.

## FAQ

### vs ripgrep (`rg`)?

`ripgrep` is faster and better for interactive search in terminals. `aict grep` is slower but returns structured XML/JSON with line numbers, byte offsets, language metadata, and context — all in one response. For AI agents that need to reason about search results, `aict grep` eliminates a parsing layer entirely. (Phase 4 will optionally delegate to `rg --json` when available.)

### vs eza / lsd?

`eza` and `lsd` are beautiful terminal replacements for humans. `aict ls` is for machines — the output is XML with absolute paths, MIME types, language tags, binary flags, and epoch timestamps. There is no colour code to strip, no column alignment to guess.

### Why XML and not JSON by default?

XML is the default because:
- Attributes carry metadata without nesting — a `<file size_bytes="2048" language="go"/>` is 40 chars; the JSON equivalent is 60+ with mandatory quotes and colons.
- AI context windows are token-limited; denser encoding means more results per call.
- Structured errors (`<error code="2" msg="no such file"/>`) compose naturally into the parent element.

Pass `--json` any time you want JSON. The schema is identical.

### Does it work on Windows?

Partially. `ls`, `cat`, `stat`, `wc`, `find`, `diff`, `grep`, `head`, `tail`, `sort`, `uniq`, `cut`, `tr`, `checksums`, `realpath`, `basename`, `dirname`, `pwd`, `env` all work. `ps` and `df` are Linux/macOS only (they read `/proc` and use `syscall.Statfs`).

### Can I use it without `AICT_XML=1`?

Yes. Pass `--xml`, `--json`, or `--plain` per invocation. The env var is a convenience for shells configured for AI pipelines.

## Benchmarks

Measured on a 3.2 GHz Linux x86-64 host against GNU coreutils. Overhead comes from language detection, MIME sniffing, and structured output — all intentional.

| Tool | GNU | aict | Ratio | Notes |
|------|-----|------|-------|-------|
| `ls` (1 000 files) | ~2 ms | ~15 ms | 7× | ✅ <10× target |
| `grep` (100 k lines) | ~1 ms | ~100 ms | 100× | ⚠️ enrichment cost |
| `find` (deep tree) | ~2 ms | ~9 ms | 5× | ✅ <10× target |
| `cat` (100 k lines) | ~1 ms | ~23 ms | 17× | ⚠️ enrichment cost |
| `diff` (1 000 lines) | ~1 ms | ~10 ms | 10× | ✅ <10× target |

`grep` and `cat` are slow because every matched file is MIME-typed and language-detected. Use `--plain` to skip enrichment when you only need content.

Run benchmarks yourself:

```bash
go build -o aict .
go run ./benchmarks/bench.go
```

## Cross-Platform

| Platform | Supported Tools | Notes |
|----------|----------------|-------|
| Linux | All tools | Full support |
| macOS | All except `ps` (partial) | `ps` uses `/proc` which is Linux-only |
| Windows | Core tools (see FAQ) | `ps`, `df` unavailable; path separators handled automatically |

## FAQ

**Why not just pipe to jq?**

You can: `aict ls . --json | jq '.total_entries'`

But jq doesn't help with `ls`, `cat`, `find`, or `stat` - those don't output JSON by default. aict gives you structured output natively for every tool.

**Why XML instead of JSON by default?**

XML is more readable for debugging and supports attributes alongside content. Both are supported: use `--json` if you prefer.

**What about eza/ripgrep?**

eza is a prettier `ls`. ripgrep is a faster `grep`. Both still output human-readable text. aict is designed for machine consumption first.

## Design Choices

- Single binary, no dependencies beyond Go standard library
- Every tool works in XML, JSON, or plain text
- All timestamps are Unix epoch integers
- All sizes are in bytes with human-readable companions
- Errors are structured data, never stderr
- Paths are always absolute

## Built by AI, for AI

We built AI coding agents that needed to read files, search codebases, and compare directories. Standard CLI tools are designed for humans. Every parsing attempt was brittle.

This gives you the same capabilities, but the output is unambiguous. The agent does not guess. It reads.

## Something Missing?

This is an open project. If you need a tool added or have a feature request, open an issue.

## License

MIT
