# ls — Directory Listing

List directory contents with language detection, MIME types, and structured metadata.

## Usage

```bash
aict ls [flags] [directory...]
```

## Flags

| Flag | Description |
|------|-------------|
| `-l` | Long format (default in XML mode) |
| `-a` | Include hidden files |
| `-A` | Include hidden files except `.` and `..` |
| `-h` | Human-readable sizes |
| `-t` | Sort by modification time |
| `-r` | Reverse sort order |
| `-R` | Recursive listing |

## XML Output

```xml
<ls timestamp="1234567890" total_entries="3" path="src/" absolute="/project/src">
  <file name="main.go" path="src/main.go" absolute="/project/src/main.go"
        size_bytes="2048" size_human="2.0 KiB"
        modified="1234567890" modified_ago_s="3600"
        permissions="rw-r--r--" mode="0644"
        owner="user" group="group"
        mime="text/x-go" language="go"
        binary="false" executable="false"/>
  <directory name="internal" path="src/internal" absolute="/project/src/internal"
             modified="1234567890" modified_ago_s="3600"
             permissions="rwxr-xr-x" mode="0755"/>
  <symlink name="link" path="src/link" absolute="/project/src/link"
           target="main.go" broken="false"/>
</ls>
```

## Plain Text Output

```bash
$ aict ls src/ --plain
main.go
utils.go
internal/
```
