# find — Filesystem Search

Search the filesystem by name, type, modification time, and more.

## Usage

```bash
aict find [path...] [conditions]
```

## Conditions

| Condition | Description |
|-----------|-------------|
| `-name <pattern>` | Match filename (glob) |
| `-type <f|d|l>` | Match file type |
| `-mtime <N>` | Modified N days ago |
| `-size <N>` | File size |
| `-maxdepth <N>` | Maximum directory depth |
| `-not` | Negate next condition |
| `-o` | OR operator |

## XML Output

```xml
<find timestamp="1234567890" total_results="42" search_root="/project">
  <condition type="name" value="*.go"/>
  <condition type="maxdepth" value="3"/>
  <result path="main.go" absolute="/project/main.go" type="file"
          size_bytes="2048" modified="1234567890" modified_ago_s="3600"
          language="go" mime="text/x-go" depth="0"/>
</find>
```
