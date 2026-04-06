# cat — File Read

Read file contents with encoding detection, language identification, and line counting.

## Usage

```bash
aict cat [flags] [file...]
```

## Flags

| Flag | Description |
|------|-------------|
| `-n` | Show line numbers |

## XML Output

```xml
<cat timestamp="1234567890" files="1" total_bytes="2048" total_lines="50">
  <file path="main.go" absolute="/project/main.go"
        size_bytes="2048" size_human="2.0 KiB"
        lines="50" encoding="utf-8" language="go"
        mime="text/x-go" binary="false"
        modified="1234567890" modified_ago_s="3600">
    <content><![CDATA[package main...]]></content>
  </file>
</cat>
```

## Notes

- Binary files omit `<content>` and set `binary="true"`
- Multi-file concatenation is supported
- Encoding detection: UTF-8, UTF-8-BOM, binary
