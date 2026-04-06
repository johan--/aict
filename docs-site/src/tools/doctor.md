# doctor — Self-Diagnostic

Check that `aict` is properly installed and functioning on your system.

## Usage

```bash
aict doctor
```

## XML Output

```xml
<doctor timestamp="1234567890" status="ok">
  <check name="binary" status="pass" message="aict found in PATH"/>
  <check name="platform" status="pass" message="linux/amd64"/>
  <check name="tools" status="pass" message="22 tools registered"/>
  <check name="mcp" status="pass" message="aict-mcp available"/>
</doctor>
```

## Status Values

| Status | Meaning |
|--------|---------|
| `pass` | Check succeeded |
| `warn` | Non-critical issue |
| `fail` | Critical issue found |

Run this command if `aict` isn't behaving as expected. It verifies the binary location, platform compatibility, tool registration, and MCP server availability.
