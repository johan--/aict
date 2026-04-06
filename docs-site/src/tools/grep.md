# grep — Pattern Search

Search files for regex patterns with context lines, byte offsets, and language metadata.

## Usage

```bash
aict grep [flags] <pattern> [path...]
```

## Flags

| Flag | Description |
|------|-------------|
| `-r` | Recursive directory search |
| `-n` | Show line numbers |
| `-l` | List matching files only |
| `-i` | Case-insensitive search |
| `-w` | Match whole words only |
| `-A N` | Show N lines after match |
| `-B N` | Show N lines before match |
| `-C N` | Show N lines of context |
| `-c` | Count matches per file |
| `-v` | Invert match |
| `-E` | Extended regex |
| `-F` | Fixed string (literal) |
| `--include` | Include files matching glob |
| `--exclude-dir` | Exclude directories matching glob |

## XML Output

```xml
<grep timestamp="1234567890" pattern="func" recursive="true"
      case_sensitive="true" match_type="regex"
      searched_files="100" matched_files="3" total_matches="12"
      search_root="/project">
  <file path="main.go" absolute="/project/main.go" matches="5" language="go">
    <match line="10" col="1" offset_bytes="200">
      <before>package main</before>
      <text>func main() {</text>
      <after>    fmt.Println("hello")</after>
    </match>
  </file>
</grep>
```

## Empty Results

```xml
<grep timestamp="1234567890" pattern="neverexists" recursive="false"
      case_sensitive="true" match_type="regex"
      searched_files="0" matched_files="0" total_matches="0"
      search_root=".">
</grep>
```
