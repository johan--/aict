# Tools Reference

All `aict` tools are organized into five categories. Every tool supports `--xml`, `--json`, and `--plain` output modes.

## File Inspection

| Tool | Description |
|------|-------------|
| [`cat`](./tools/cat.md) | File contents with encoding detection |
| [`head` / `tail`](./tools/head-tail.md) | First or last lines of a file |
| [`file`](./tools/file.md) | File type detection via magic bytes |
| [`stat`](./tools/stat.md) | File metadata with all timestamps |
| [`wc`](./tools/wc.md) | Line, word, and byte counts |

## Directory & Search

| Tool | Description |
|------|-------------|
| [`ls`](./tools/ls.md) | Directory listings with language and MIME type detection |
| [`find`](./tools/find.md) | Filesystem search by name, type, or modification time |
| [`grep`](./tools/grep.md) | Pattern search with context lines and recursive support |
| [`diff`](./tools/diff.md) | File and directory comparison |

## Path Utilities

| Tool | Description |
|------|-------------|
| [`realpath`](./tools/path-utils.md) | Resolve absolute paths |
| [`basename`](./tools/path-utils.md) | Extract filename from path |
| [`dirname`](./tools/path-utils.md) | Extract directory from path |
| [`pwd`](./tools/path-utils.md) | Print working directory |

## Text Processing

| Tool | Description |
|------|-------------|
| [`sort`](./tools/sort-uniq.md) | Sort lines with options |
| [`uniq`](./tools/sort-uniq.md) | Remove or count duplicate lines |
| [`cut`](./tools/cut-tr.md) | Extract columns from delimited text |
| [`tr`](./tools/cut-tr.md) | Translate or delete characters |

## System & Environment

| Tool | Description |
|------|-------------|
| [`env`](./tools/env.md) | Environment variables with secret redaction |
| [`system`](./tools/system.md) | OS, runtime, and user information |
| [`ps`](./tools/ps.md) | Running process list |
| [`df`](./tools/du-df.md) | Disk space usage |
| [`du`](./tools/du-df.md) | Directory size analysis |
| [`checksums`](./tools/checksums.md) | MD5, SHA1, and SHA256 hashes |
| [`git`](./tools/git.md) | Git status, diff, log, ls-files, blame |
| [`doctor`](./tools/doctor.md) | Self-diagnostic command |
