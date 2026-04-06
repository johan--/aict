# ps — Process List

List running processes with detailed metadata.

## Usage

```bash
aict ps [flags]
```

## Flags

| Flag | Description |
|------|-------------|
| `aux` | BSD-style full listing |
| `-ef` | System V-style full listing |
| `-p <pid>` | Show specific PID |
| `--sort <field>` | Sort by field |

## XML Output

```xml
<ps timestamp="1234567890" total_processes="150">
  <process pid="1234" ppid="1" user="root" uid="0"
           cpu_pct="0.0" mem_pct="1.2" vsz_kb="12345" rss_kb="8192"
           state="S" state_desc="sleeping"
           started="1234567890" started_ago_s="86400"
           command="/usr/bin/python3" args="python3 server.py"
           exe="/usr/bin/python3"/>
</ps>
```

## Platform Support

- Linux: Full support via `/proc` filesystem
- macOS: Partial — uses `syscall.SysctlRaw` fallback
- Windows: Not available
