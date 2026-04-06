# checksums — Hash Computation

Compute MD5, SHA1, and SHA256 hashes in a single pass.

## Usage

```bash
aict checksums [flags] [file...]
```

## Flags

| Flag | Description |
|------|-------------|
| `-c` | Verify against checksum file |

## XML Output

```xml
<checksums timestamp="1234567890" files="2">
  <file path="main.go" absolute="/project/main.go"
        size_bytes="2048"
        md5="d41d8cd98f00b204e9800998ecf8427e"
        sha1="da39a3ee5e6b4b0d3255bfef95601890afd80709"
        sha256="e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"/>
</checksums>
```

## Notes

- All three hashes computed in a single pass using `io.MultiWriter`
- Verification mode (`-c`) reads a standard checksum file and validates each entry
