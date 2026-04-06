# XML Schema Reference

Every `aict` tool outputs XML with a consistent structure. The same schema is used for JSON output (`--json`) with identical field names.

## Global Rules

| Rule | Detail |
|------|--------|
| Root element | Named after the tool: `<ls>`, `<grep>`, `<cat>`, etc. |
| `timestamp` attribute | Unix epoch integer on every root element |
| `path` attribute | Path as given by the user |
| `absolute` attribute | Fully resolved absolute path |
| Time values | Unix epoch integers with `_ago_s` companion attributes |
| Size values | Bytes with `_human` companion attributes |
| Booleans | `true`/`false` strings, never `1`/`0` |
| Errors | `<error code="" msg=""/>` elements, never stderr |
| Empty results | Valid XML with zero counts |
| Binary files | Never output as CDATA — omit content or use base64 |
| Language values | Lowercase canonical (`go`, `python`, `typescript`) |

---

## ls

```xml
<ls timestamp="1234567890" total_entries="3" path="src/" absolute="/project/src">
  <file name="main.go" path="src/main.go" absolute="/project/src/main.go"
        size_bytes="2048" size_human="2.0 KiB"
        modified="1234567890" modified_ago_s="3600"
        permissions="rw-r--r--" mode="0644"
        owner="user" group="group"
        mime="text/x-go" language="go"
        binary="false" executable="false"/>
  <directory name="internal" path="src/internal" absolute="/project/src/internal"
             modified="1234567890" modified_ago_s="3600"
             permissions="rwxr-xr-x" mode="0755"/>
  <symlink name="link" path="src/link" absolute="/project/src/link"
           target="main.go" broken="false"/>
</ls>
```

---

## cat

```xml
<cat timestamp="1234567890" files="1" total_bytes="2048" total_lines="50">
  <file path="main.go" absolute="/project/main.go"
        size_bytes="2048" size_human="2.0 KiB"
        lines="50" encoding="utf-8" language="go"
        mime="text/x-go" binary="false"
        modified="1234567890" modified_ago_s="3600">
    <content><![CDATA[package main...]]></content>
  </file>
</cat>
```

---

## grep

```xml
<grep timestamp="1234567890" pattern="func" recursive="true"
      case_sensitive="true" match_type="regex"
      searched_files="100" matched_files="3" total_matches="12"
      search_root="/project">
  <file path="main.go" absolute="/project/main.go" matches="5" language="go">
    <match line="10" col="1" offset_bytes="200">
      <before>package main</before>
      <text>func main() {</text>
      <after>    fmt.Println("hello")</after>
    </match>
  </file>
</grep>
```

---

## find

```xml
<find timestamp="1234567890" total_results="42" search_root="/project">
  <condition type="name" value="*.go"/>
  <condition type="maxdepth" value="3"/>
  <result path="main.go" absolute="/project/main.go" type="file"
          size_bytes="2048" modified="1234567890" modified_ago_s="3600"
          language="go" mime="text/x-go" depth="0"/>
</find>
```

---

## stat

```xml
<stat timestamp="1234567890">
  <file path="main.go" absolute="/project/main.go"
        inode="123456" links="1" device="259,0"
        permissions="rw-r--r--" mode_octal="0644"
        uid="1000" gid="1000" owner="user" group="group"
        size_bytes="2048" size_human="2.0 KiB"
        atime="1234567890" atime_ago_s="3600"
        mtime="1234567890" mtime_ago_s="3600"
        ctime="1234567890" ctime_ago_s="3600"
        birth="1234567890" birth_ago_s="7200"
        language="go" mime="text/x-go"/>
</stat>
```

---

## wc

```xml
<wc timestamp="1234567890" files="3">
  <file path="main.go" absolute="/project/main.go"
        lines="50" words="200" chars="1800" bytes="2048" language="go"/>
  <file path="utils.go" absolute="/project/utils.go"
        lines="30" words="100" chars="900" bytes="1024" language="go"/>
  <total lines="80" words="300" chars="2700" bytes="3072"/>
</wc>
```

---

## diff

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

---

## file

```xml
<file timestamp="1234567890">
  <entry path="main.go" absolute="/project/main.go"
         type="Go source" mime="text/x-go" category="text"
         language="go" charset="utf-8" executable="false"/>
</file>
```

---

## head / tail

```xml
<head timestamp="1234567890" lines_requested="10" lines_returned="10"
        file_total_lines="500" bytes_returned="256" file_total_bytes="12800"
        truncated="false">
  <content><![CDATA[line 1...]]></content>
</head>
```

---

## du

```xml
<du timestamp="1234567890" total_size_bytes="1048576" total_size_human="1.0 MiB">
  <entry path="src/" absolute="/project/src"
         size_bytes="524288" size_human="512.0 KiB" depth="0"/>
</du>
```

---

## df

```xml
<df timestamp="1234567890">
  <filesystem device="/dev/sda1" mount="/" type="ext4"
              size_bytes="107374182400" size_human="100.0 GiB"
              used_bytes="53687091200" used_human="50.0 GiB"
              avail_bytes="53687091200" avail_human="50.0 GiB"
              use_pct="50" inodes_total="6553600" inodes_used="3276800"/>
</df>
```

---

## realpath / basename / dirname

```xml
<realpath timestamp="1234567890" path="src/../main.go"
          absolute="/project/main.go" exists="true" type="file"/>

<basename timestamp="1234567890" path="/project/main.go"
          name="main.go" stem="main" extension=".go"/>

<dirname timestamp="1234567890" path="/project/main.go"
         directory="/project"/>
```

---

## pwd

```xml
<pwd timestamp="1234567890" path="/project/src"
     home="/home/user" relative_to_home="~/project/src"/>
```

---

## sort

```xml
<sort timestamp="1234567890" lines_in="100" lines_out="100"
      key="1" order="ascending">
  <content><![CDATA[sorted lines...]]></content>
</sort>
```

---

## uniq

```xml
<uniq timestamp="1234567890" lines_in="100" lines_out="50"
      duplicates_removed="50" counted="true">
  <entry count="5"><![CDATA[duplicate line]]></entry>
</uniq>
```

---

## cut

```xml
<cut timestamp="1234567890" delimiter="," fields="1,3" lines_processed="100">
  <content><![CDATA[extracted columns...]]></content>
</cut>
```

---

## tr

```xml
<tr timestamp="1234567890" lines_processed="100" bytes_processed="5000">
  <content><![CDATA[transformed text...]]></content>
</tr>
```

---

## env

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

---

## system

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

---

## ps

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

---

## checksums

```xml
<checksums timestamp="1234567890" files="2">
  <file path="main.go" absolute="/project/main.go"
        size_bytes="2048"
        md5="d41d8cd98f00b204e9800998ecf8427e"
        sha1="da39a3ee5e6b4b0d3255bfef95601890afd80709"
        sha256="e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"/>
</checksums>
```

---

## git

```xml
