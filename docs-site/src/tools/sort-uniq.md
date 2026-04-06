# sort / uniq — Sorting and Deduplication

## sort

Sort lines with various options.

```bash
aict sort [flags] [file...]
```

### Flags

| Flag | Description |
|------|-------------|
| `-n` | Numeric sort |
| `-r` | Reverse order |
| `-k N` | Sort by field N |
| `-t <delim>` | Field delimiter |

### XML Output

```xml
<sort timestamp="1234567890" lines_in="100" lines_out="100"
      key="1" order="ascending">
  <content><![CDATA[sorted lines...]]></content>
</sort>
```

## uniq

Remove or count duplicate lines.

```bash
aict uniq [flags] [file...]
```

### Flags

| Flag | Description |
|------|-------------|
| `-c` | Prefix lines with count |
| `-d` | Only print duplicates |
| `-u` | Only print unique lines |

### XML Output

```xml
<uniq timestamp="1234567890" lines_in="100" lines_out="50"
      duplicates_removed="50" counted="true">
  <entry count="5"><![CDATA[duplicate line]]></entry>
</uniq>
```
