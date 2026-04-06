# git — Git Subcommands

Git operations with structured XML output.

## Usage

```bash
aict git <subcommand> [flags]
```

## Subcommands

| Subcommand | Description |
|------------|-------------|
| `status` | Working tree status |
| `diff` | Changes between commits/working tree |
| `log` | Commit history |
| `ls-files` | List tracked files |
| `blame` | Line-by-line authorship |

## git status

```xml
<git timestamp="1234567890" command="status">
  <branch name="main" ahead="0" behind="0"/>
  <staged path="main.go" absolute="/project/main.go" status="modified"/>
  <modified path="utils.go" absolute="/project/utils.go"/>
  <untracked path="new.go" absolute="/project/new.go"/>
</git>
```

## git diff

```xml
<git timestamp="1234567890" command="diff">
  <file path="main.go" absolute="/project/main.go">
    <hunk old_start="10" old_count="5" new_start="10" new_count="7">
      <added line="11">    new_line();</added>
      <removed line="11">    old_line();</removed>
    </hunk>
  </file>
</git>
```

## git log

```xml
<git timestamp="1234567890" command="log" count="10">
  <commit hash="abc123" author="user" email="user@example.com"
          date="1234567890" message="feat: add new feature"/>
</git>
```
