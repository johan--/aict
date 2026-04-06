# diff — File Comparison

Compare files and directories with Myers diff algorithm, structured hunks, and line numbers.

## Usage

```bash
aict diff [flags] <old> <new>
```

## Flags

| Flag | Description |
|------|-------------|
| `-u` | Unified diff format |
| `--label <name>` | Custom labels |
| `-r` | Recursive directory comparison |
| `--ignore-all-space` | Ignore whitespace changes |
| `-q` | Brief output |

## XML Output

```xml
<diff timestamp="1234567890" added_lines="5" removed_lines="3"
      changed_hunks="2" identical="false">
  <hunk old_start="10" old_count="5" new_start="10" new_count="7">
    <context line="10">func example() {</context>
    <removed line="11">    old_line();</removed>
    <added line="11">    new_line();</added>
    <added line="12">    another_new();</added>
    <context line="12">    return nil</context>
  </hunk>
</diff>
```

## Identical Files

```xml
<diff timestamp="1234567890" added_lines="0" removed_lines="0"
      changed_hunks="0" identical="true"/>
```
