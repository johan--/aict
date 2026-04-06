# wc — Line/Word/Char/Byte Count

Count lines, words, characters, and bytes with per-file and total statistics.

## Usage

```bash
aict wc [flags] [file...]
```

## Flags

| Flag | Description |
|------|-------------|
| `-l` | Lines only |
| `-w` | Words only |
| `-c` | Bytes only |
| `-m` | Characters only |

## XML Output

```xml
<wc timestamp="1234567890" files="3">
  <file path="main.go" absolute="/project/main.go"
        lines="50" words="200" chars="1800" bytes="2048" language="go"/>
  <file path="utils.go" absolute="/project/utils.go"
        lines="30" words="100" chars="900" bytes="1024" language="go"/>
  <total lines="80" words="300" chars="2700" bytes="3072"/>
</wc>
```
