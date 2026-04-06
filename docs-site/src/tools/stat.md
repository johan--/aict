# stat — File Metadata

Get detailed file metadata with all timestamps, permissions, ownership, and enrichment.

## Usage

```bash
aict stat [flags] [file...]
```

## Flags

| Flag | Description |
|------|-------------|
| `-L` | Follow symlinks |

## XML Output

```xml
<stat timestamp="1234567890">
  <file path="main.go" absolute="/project/main.go"
        inode="123456" links="1" device="259,0"
        permissions="rw-r--r--" mode_octal="0644"
        uid="1000" gid="1000" owner="user" group="group"
        size_bytes="2048" size_human="2.0 KiB"
        atime="1234567890" atime_ago_s="3600"
        mtime="1234567890" mtime_ago_s="3600"
        ctime="1234567890" ctime_ago_s="3600"
        birth="1234567890" birth_ago_s="7200"
        language="go" mime="text/x-go"/>
</stat>
```
