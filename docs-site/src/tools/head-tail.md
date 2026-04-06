# head / tail — Partial File Read

Read the first or last N lines/bytes of a file with enrichment metadata.

## Usage

```bash
aict head [flags] [file...]
aict tail [flags] [file...]
```

## Flags

| Flag | Description |
|------|-------------|
| `-n N` | Number of lines (default: 10) |
| `-c N` | Number of bytes |

## XML Output

```xml
<head timestamp="1234567890" lines_requested="10" lines_returned="10"
        file_total_lines="500" bytes_returned="256" file_total_bytes="12800"
        truncated="false">
  <content><![CDATA[line 1...]]></content>
</head>
```

## Notes

- `tail` supports `-f` for follow mode (not available in XML stream mode)
- Language and MIME enrichment applied to returned content
