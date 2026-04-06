# cut / tr — Text Processing

## cut

Extract columns from delimited text.

```bash
aict cut [flags] [file...]
```

### Flags

| Flag | Description |
|------|-------------|
| `-d <delim>` | Field delimiter (default: tab) |
| `-f <fields>` | Fields to extract (comma-separated) |

### XML Output

```xml
<cut timestamp="1234567890" delimiter="," fields="1,3" lines_processed="100">
  <content><![CDATA[extracted columns...]]></content>
</cut>
```

## tr

Translate or delete characters.

```bash
aict tr [flags] <set1> [set2]
```

### Flags

| Flag | Description |
|------|-------------|
| `-d` | Delete characters in set1 |
| `-s` | Squeeze repeated characters |

### XML Output

```xml
<tr timestamp="1234567890" lines_processed="100" bytes_processed="5000">
  <content><![CDATA[transformed text...]]></content>
</tr>
```
