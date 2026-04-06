# Benchmarks

Measured on a 3.2 GHz Linux x86-64 host against GNU coreutils.

## Results

| Tool | GNU | aict | Ratio | Status |
|------|-----|------|-------|--------|
| `ls` (1,000 files) | ~2ms | ~15ms | 7x | PASS |
| `grep` (100k lines) | ~1ms | ~100ms | 100x | SLOW |
| `find` (deep tree) | ~2ms | ~9ms | 5x | PASS |
| `cat` (100k lines) | ~1ms | ~23ms | 17x | SLOW |
| `diff` (1,000 lines) | ~1ms | ~10ms | 10x | PASS |

## Why is aict slower?

aict provides significantly more functionality than GNU coreutils:

1. **Language Detection** — Each file is analyzed for programming language
2. **MIME Detection** — Magic bytes analysis for file type
3. **Structured Output** — XML/JSON structures instead of plain text
4. **Enriched Metadata** — Timestamps, permissions, ownership, etc.

## Optimization Tips

- Use `--plain` to skip enrichment when you only need content
- Use `--include` filters in `grep` to reduce files scanned
- Use `--maxdepth` in `find` to limit directory traversal
- For maximum speed, use GNU coreutils directly when structured output isn't needed

The trade-off is intentional: more tokens spent on parsing vs. more semantic information returned.

## Run Benchmarks Yourself

```bash
go build -o aict .
go run ./benchmarks/bench.go
```
