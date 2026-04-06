# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Cross-platform support documentation
- CONTRIBUTING.md guide
- GitHub Actions CI workflow
- Docker build configuration

### Changed
- Improved test coverage for edge cases

## [1.0.0] - 2026-04-06

### Added
- **Phase 0**: Foundation
  - Go module and directory structure
  - Internal packages (xml, detect, path, format, meta)
  - `ls` tool with full XML output

- **Phase 1**: Core Reads
  - `cat` - File read with encoding detection
  - `grep` - Recursive regex search
  - `find` - Filesystem search
  - `stat` - File metadata
  - `wc` - Line/word/char/byte counting
  - `diff` - Myers diff algorithm

- **Phase 2**: Contextual Enrichment
  - `file` - Type detection
  - `head`/`tail` - Partial file read
  - `du`/`df` - Disk usage
  - `realpath`/`basename`/`dirname` - Path utilities
  - `pwd` - Working directory
  - `sort`/`uniq` - Sorting and deduplication
  - `cut`/`tr` - Text processing
  - `env` - Environment with secret redaction
  - `system` - Combined system info
  - `ps` - Process listing
  - `checksums` - Hash computation
  - MCP server (`cmd/mcp`)

### Features
- XML output (default)
- JSON output (`--json`)
- Plain text output (`--plain`)
- `AICT_XML=1` environment variable
- Structured error elements
- Language detection
- MIME type detection

[Unreleased]: https://github.com/synseqack/aict/compare/v1.0.0...HEAD
[1.0.0]: https://github.com/synseqack/aict/releases/tag/v1.0.0
