# du / df — Disk Usage

Analyze directory sizes and filesystem usage.

## du — Directory Usage

```bash
aict du [flags] [path...]
```

### Flags

| Flag | Description |
|------|-------------|
| `-s` | Summary only |
| `-h` | Human-readable sizes |
| `-a` | Show all files, not just directories |
| `--max-depth N` | Maximum depth |

### XML Output

```xml
<du timestamp="1234567890" total_size_bytes="1048576" total_size_human="1.0 MiB">
  <entry path="src/" absolute="/project/src"
         size_bytes="524288" size_human="512.0 KiB" depth="0"/>
</du>
```

## df — Filesystem Usage

```bash
aict df [flags]
```

### Flags

| Flag | Description |
|------|-------------|
| `-h` | Human-readable sizes |

### XML Output

```xml
<df timestamp="1234567890">
  <filesystem device="/dev/sda1" mount="/" type="ext4"
              size_bytes="107374182400" size_human="100.0 GiB"
              used_bytes="53687091200" used_human="50.0 GiB"
              avail_bytes="53687091200" avail_human="50.0 GiB"
              use_pct="50" inodes_total="6553600" inodes_used="3276800"/>
</df>
```

### Platform Support

`df` uses `syscall.Statfs` and is Linux/macOS only. Not available on Windows.
