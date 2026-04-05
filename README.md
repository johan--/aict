# aict

Your command line, but built for AI.

## The Problem

Every time an AI agent runs `ls`, `grep`, or `cat`, it wastes tokens parsing human-readable output. The agent has to guess which column is the filename, which is the size, which is the date. This guesswork costs money and introduces errors.

aict gives you the same tools you already know, but the output is structured. No parsing. No regex. Just data.

## What You Get

```
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
go install github.com/synseqack/aict@latest
```

Or build from source:

```bash
git clone https://github.com/synseqack/aict
cd aict
go build -o aict
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
- `cat` - File contents with encoding detection
- `head` / `tail` - First or last lines of a file
- `file` - File type detection via magic bytes
- `stat` - File metadata with all timestamps
- `wc` - Line, word, and byte counts

**Directory & Search**
- `ls` - Directory listings with language and MIME type detection
- `find` - Filesystem search by name, type, or modification time
- `grep` - Pattern search with context lines and recursive support
- `diff` - File and directory comparison

**Path Utilities**
- `realpath` - Resolve absolute paths
- `basename` - Extract filename from path
- `dirname` - Extract directory from path
- `pwd` - Print working directory

**Text Processing**
- `sort` - Sort lines with options
- `uniq` - Remove or count duplicate lines
- `cut` - Extract columns from delimited text
- `tr` - Translate or delete characters

**System & Environment**
- `env` - Environment variables with secret redaction
- `system` - OS, runtime, and user information
- `ps` - Running process list
- `df` - Disk space usage
- `du` - Directory size analysis
- `checksums` - MD5, SHA1, and SHA256 hashes

## MCP Server

aict can run as an MCP server so AI assistants like ChatGPT can call these tools directly.

```bash
go build -o aict-mcp ./cmd/mcp
```

Configure your AI client to use `aict-mcp` as a command-line MCP server. Each tool becomes a callable function with typed arguments and structured JSON output.

## Why This Exists

We built AI coding agents that needed to read files, search codebases, and compare directories. Standard CLI tools are designed for humans. Every parsing attempt was brittle.

This gives you the same capabilities, but the output is unambiguous. The agent does not guess. It reads.

## Design Choices

- Single binary, no dependencies beyond Go standard library
- Every tool works in XML, JSON, or plain text
- All timestamps are Unix epoch integers
- All sizes are in bytes with human-readable companions
- Errors are structured data, never stderr
- Paths are always absolute

## Something Missing?

This is an open project. If you need a tool added or have a feature request, open an issue.

## License

MIT
