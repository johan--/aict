# env — Environment Variables

List environment variables with automatic secret redaction and PATH parsing.

## Usage

```bash
aict env [flags]
```

## XML Output

```xml
<env timestamp="1234567890" total="42" secrets_redacted="3">
  <variable name="HOME" value="/home/user" type="path"/>
  <variable name="PATH" type="path_list">
    <path_entry path="/usr/bin" exists="true"/>
    <path_entry path="/usr/local/bin" exists="true"/>
  </variable>
  <variable name="API_KEY" present="true" redacted="true" type="secret"/>
</env>
```

## Secret Detection

Variables matching these patterns are automatically redacted:

- Names containing `KEY`, `SECRET`, `TOKEN`, `PASSWORD`, `DSN`
- URLs containing `://user:pass@`

Redacted variables show `present="true" redacted="true"` instead of their value.
