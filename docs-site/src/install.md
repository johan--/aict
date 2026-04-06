# Installation

## Option 1: Go Install (Recommended)

```bash
# Install both binaries
go install github.com/synseqack/aict@latest
go install github.com/synseqack/aict/cmd/mcp@latest
```

This places `aict` and `aict-mcp` in your `$GOPATH/bin` directory.

## Option 2: Build from Source

```bash
git clone https://github.com/synseqack/aict
cd aict
go build -o aict .
go build -o aict-mcp ./cmd/mcp
```

## Option 3: Docker

```bash
docker build -t aict .
docker run --rm -v "$(pwd)":/work -w /work aict ls .
```

## Verify Installation

```bash
aict --help
aict ls .
```

You should see a list of available commands and an XML directory listing.

## Shell Completion

```bash
# Bash
source completions/aict.bash

# Zsh
source completions/aict.zsh
```

Add the appropriate line to your `.bashrc` or `.zshrc` for persistent completion.
