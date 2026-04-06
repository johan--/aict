# system — System Information

Combined OS, runtime, and user information.

## Usage

```bash
aict system
```

## XML Output

```xml
<system timestamp="1234567890">
  <user uid="1000" gid="1000" username="user" home="/home/user" shell="/bin/bash">
    <group name="users" gid="100"/>
    <group name="sudo" gid="27"/>
  </user>
  <os name="linux" arch="amd64" hostname="myhost" kernel="5.15.0" distro="Ubuntu 24.04"/>
  <runtime go_version="go1.25" num_cpu="8" go_max_procs="8"/>
</system>
```

## Platform Support

- Linux: Full support including distro detection from `/etc/os-release`
- macOS: Partial — distro detection unavailable
- Windows: Partial — group lookup and distro detection unavailable
