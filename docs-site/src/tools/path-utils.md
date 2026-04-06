# Path Utilities

## realpath

Resolve a path to its absolute, canonical form.

```bash
aict realpath [path...]
```

```xml
<realpath timestamp="1234567890" path="src/../main.go"
          absolute="/project/main.go" exists="true" type="file"/>
```

## basename

Extract the filename and optionally the stem/extension from a path.

```bash
aict basename [path...]
```

```xml
<basename timestamp="1234567890" path="/project/main.go"
          name="main.go" stem="main" extension=".go"/>
```

## dirname

Extract the parent directory from a path.

```bash
aict dirname [path...]
```

```xml
<dirname timestamp="1234567890" path="/project/main.go"
         directory="/project"/>
```

## pwd

Print the current working directory with home-relative path.

```bash
aict pwd
```

```xml
<pwd timestamp="1234567890" path="/project/src"
     home="/home/user" relative_to_home="~/project/src"/>
```
