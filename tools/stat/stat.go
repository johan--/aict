package stat

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"os/user"
	"strconv"
	"strings"
	"time"

	"github.com/synseqack/aict/internal/detect"
	"github.com/synseqack/aict/internal/format"
	"github.com/synseqack/aict/internal/meta"
	pathutil "github.com/synseqack/aict/internal/path"
	"github.com/synseqack/aict/internal/tool"
	xmlout "github.com/synseqack/aict/internal/xml"
)

func init() {
	tool.Register("stat", Run)
	tool.RegisterMeta("stat", tool.GenerateSchema("stat", "Display detailed file metadata including timestamps, permissions, and ownership", Config{}))
}

type Config struct {
	FollowSymlinks bool `flag:"" desc:"Follow symlinks and show target file info"`
	XML            bool
	JSON           bool
	Plain          bool
	Pretty         bool
}

type StatResult struct {
	XMLName     xml.Name    `xml:"stat"`
	Path        string      `xml:"path,attr"`
	Absolute    string      `xml:"absolute,attr"`
	Inode       uint64      `xml:"inode,attr"`
	Links       int         `xml:"links,attr"`
	Device      uint64      `xml:"device,attr"`
	Permissions string      `xml:"permissions,attr"`
	ModeOctal   string      `xml:"mode_octal,attr"`
	UID         uint32      `xml:"uid,attr"`
	GID         uint32      `xml:"gid,attr"`
	Owner       string      `xml:"owner,attr"`
	Group       string      `xml:"group,attr"`
	SizeBytes   int64       `xml:"size_bytes,attr"`
	SizeHuman   string      `xml:"size_human,attr"`
	Atime       int64       `xml:"atime,attr"`
	AtimeAgoS   int64       `xml:"atime_ago_s,attr"`
	Mtime       int64       `xml:"mtime,attr"`
	MtimeAgoS   int64       `xml:"mtime_ago_s,attr"`
	Ctime       int64       `xml:"ctime,attr"`
	CtimeAgoS   int64       `xml:"ctime_ago_s,attr"`
	Birth       int64       `xml:"birth,attr"`
	BirthAgoS   int64       `xml:"birth_ago_s,attr"`
	Type        string      `xml:"type,attr"`
	MIME        string      `xml:"mime,attr"`
	Language    string      `xml:"language,attr"`
	Timestamp   int64       `xml:"timestamp,attr"`
	Errors      []StatError `xml:"error,omitempty"`
}

func (*StatResult) isStatResult() {}

type StatError struct {
	XMLName xml.Name `xml:"error"`
	Code    int      `xml:"code,attr"`
	Msg     string   `xml:"msg,attr"`
	Path    string   `xml:"path,attr"`
}

func Run(args []string) error {
	cfg, paths := parseFlags(args)

	if len(paths) == 0 {
		paths = []string{"."}
	}

	for i, p := range paths {
		result, err := statPath(p, cfg)
		if err != nil {
			return err
		}
		if i > 0 {
			fmt.Println()
		}
		if err := outputResult(result, cfg); err != nil {
			return err
		}
	}
	return nil
}

func parseFlags(args []string) (Config, []string) {
	var cfg Config
	var positional []string

	for _, arg := range args {
		switch arg {
		case "-L", "--dereference":
			cfg.FollowSymlinks = true
		case "--xml", "-xml":
			cfg.XML = true
		case "--json", "-json":
			cfg.JSON = true
		case "--plain", "-plain":
			cfg.Plain = true
		case "--pretty", "-pretty":
			cfg.Pretty = true
		default:
			positional = append(positional, arg)
		}
	}

	if !cfg.XML && !cfg.JSON && !cfg.Plain {
		cfg.XML = xmlout.IsXMLMode()
	}

	return cfg, positional
}

func statPath(path string, cfg Config) (*StatResult, error) {
	resolved, err := pathutil.Resolve(path)
	if err != nil {
		return &StatResult{
			Path:      path,
			Timestamp: meta.Now(),
			Errors:    []StatError{{Code: 1, Msg: err.Error(), Path: path}},
		}, nil
	}

	var info os.FileInfo
	var errStat error

	if cfg.FollowSymlinks {
		info, errStat = os.Stat(resolved.Absolute)
	} else {
		info, errStat = os.Lstat(resolved.Absolute)
	}

	if errStat != nil {
		code := 1
		if os.IsNotExist(errStat) {
			code = 2
		}
		return &StatResult{
			Path:      resolved.Given,
			Absolute:  resolved.Absolute,
			Timestamp: meta.Now(),
			Errors:    []StatError{{Code: code, Msg: "no such file or directory", Path: resolved.Absolute}},
		}, nil
	}

	result := &StatResult{
		Path:      resolved.Given,
		Absolute:  resolved.Absolute,
		Timestamp: meta.Now(),
	}

	sysInfo := info.Sys()
	if sysInfo != nil {
		result.Inode = getIno(sysInfo)
		result.Links = int(getNlink(sysInfo))
		result.Device = getDev(sysInfo)
		result.UID = getUID(sysInfo)
		result.GID = getGID(sysInfo)

		if owner, err := user.LookupId(strconv.FormatUint(uint64(result.UID), 10)); err == nil {
			result.Owner = owner.Username
		} else {
			result.Owner = "unknown"
		}

		if group, err := user.LookupGroupId(strconv.FormatUint(uint64(result.GID), 10)); err == nil {
			result.Group = group.Name
		} else {
			result.Group = "unknown"
		}

		sec := getATimSec(sysInfo)
		if sec > 0 {
			result.Atime = sec
			result.AtimeAgoS = meta.AgoSeconds(sec)
		} else {
			result.Atime = info.ModTime().Unix()
			result.AtimeAgoS = meta.AgoSeconds(info.ModTime().Unix())
		}

		sec = getMTimSec(sysInfo)
		if sec > 0 {
			result.Mtime = sec
			result.MtimeAgoS = meta.AgoSeconds(sec)
		} else {
			result.Mtime = info.ModTime().Unix()
			result.MtimeAgoS = meta.AgoSeconds(info.ModTime().Unix())
		}

		sec = getCTimSec(sysInfo)
		if sec > 0 {
			result.Ctime = sec
			result.CtimeAgoS = meta.AgoSeconds(sec)
		} else {
			result.Ctime = info.ModTime().Unix()
			result.CtimeAgoS = meta.AgoSeconds(info.ModTime().Unix())
		}

		result.Birth = 0
		result.BirthAgoS = 0
	} else {
		result.Atime = info.ModTime().Unix()
		result.AtimeAgoS = meta.AgoSeconds(info.ModTime().Unix())
		result.Mtime = info.ModTime().Unix()
		result.MtimeAgoS = meta.AgoSeconds(info.ModTime().Unix())
		result.Ctime = info.ModTime().Unix()
		result.CtimeAgoS = meta.AgoSeconds(info.ModTime().Unix())
		result.Birth = 0
		result.BirthAgoS = 0
	}

	result.SizeBytes = info.Size()
	result.SizeHuman = format.Size(uint64(result.SizeBytes))
	result.Permissions = formatPermissions(info.Mode())
	result.ModeOctal = "0" + strconv.FormatUint(uint64(info.Mode().Perm()), 8)
	result.Type = getFileType(info)

	mime := "application/octet-stream"
	language := ""
	if !info.IsDir() {
		mime, _, _ = detect.DetectFromFile(resolved.Absolute)
		language = detect.LanguageFromFile(resolved.Absolute)
	}
	result.MIME = mime
	result.Language = language

	return result, nil
}

func formatPermissions(mode os.FileMode) string {
	var b strings.Builder
	b.Grow(10)

	if mode&os.ModeSymlink != 0 {
		b.WriteByte('l')
	} else if mode.IsDir() {
		b.WriteByte('d')
	} else {
		b.WriteByte('-')
	}

	for i := 8; i >= 0; i-- {
		bit := uint(1) << uint(i)
		switch {
		case mode&os.FileMode(bit) != 0:
			switch i % 3 {
			case 0:
				b.WriteByte('x')
			case 1:
				b.WriteByte('w')
			case 2:
				b.WriteByte('r')
			}
		default:
			b.WriteByte('-')
		}
	}

	return b.String()
}

func getFileType(info os.FileInfo) string {
	mode := info.Mode()
	if mode&os.ModeSymlink != 0 {
		return "symlink"
	}
	if mode.IsDir() {
		return "directory"
	}
	if mode.IsRegular() {
		return "file"
	}
	if mode&os.ModeDevice != 0 {
		return "block"
	}
	if mode&os.ModeCharDevice != 0 {
		return "character"
	}
	if mode&os.ModeNamedPipe != 0 {
		return "pipe"
	}
	if mode&os.ModeSocket != 0 {
		return "socket"
	}
	return "unknown"
}

func formatSize(size int64) string {
	const unit = 1024
	if size < unit {
		return fmt.Sprintf("%dB", size)
	}
	div, exp := int64(unit), 0
	for n := size / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f%c", float64(size)/float64(div), "KMGTPE"[exp])
}

func outputResult(result *StatResult, cfg Config) error {
	if cfg.JSON {
		return xmlout.WriteJSON(os.Stdout, result)
	}
	if cfg.Plain {
		return writePlain(os.Stdout, result)
	}
	return xmlout.WriteXML(os.Stdout, result, cfg.Pretty)
}

func writePlain(w io.Writer, result *StatResult) error {
	if len(result.Errors) > 0 {
		for _, e := range result.Errors {
			fmt.Fprintf(w, "stat: %s: %s\n", e.Path, e.Msg)
		}
		return nil
	}

	fmt.Fprintf(w, "  File: %s\n", result.Path)
	fmt.Fprintf(w, "  Size: %d\t\tBlocks: %d\tIO Block: %d\t%s\n",
		result.SizeBytes, result.Links, result.Device, result.Type)
	fmt.Fprintf(w, "Device: %d\t\tInode: %d\tLinks: %d\n",
		result.Device, result.Inode, result.Links)
	fmt.Fprintf(w, "Access: %s (%s)\n", result.Permissions, result.ModeOctal)
	fmt.Fprintf(w, "Uid: %d\t(%s)\tGid: %d\t(%s)\n",
		result.UID, result.Owner, result.GID, result.Group)
	fmt.Fprintf(w, "Access: %s\n", time.Unix(result.Atime, 0).Format(time.RubyDate))
	fmt.Fprintf(w, "Modify: %s\n", time.Unix(result.Mtime, 0).Format(time.RubyDate))
	fmt.Fprintf(w, "Change: %s\n", time.Unix(result.Ctime, 0).Format(time.RubyDate))

	return nil
}

func getIno(sysInfo any) uint64 {
	switch v := sysInfo.(type) {
	case interface{ Ino() uint64 }:
		return v.Ino()
	case interface{ Ino() uint32 }:
		return uint64(v.Ino())
	default:
		return 0
	}
}

func getNlink(sysInfo any) uint64 {
	switch v := sysInfo.(type) {
	case interface{ Nlink() uint64 }:
		return v.Nlink()
	case interface{ Nlink() uint32 }:
		return uint64(v.Nlink())
	default:
		return 0
	}
}

func getDev(sysInfo any) uint64 {
	switch v := sysInfo.(type) {
	case interface{ Dev() uint64 }:
		return v.Dev()
	default:
		return 0
	}
}

func getUID(sysInfo any) uint32 {
	switch v := sysInfo.(type) {
	case interface{ Uid() uint32 }:
		return v.Uid()
	case interface{ UID() uint32 }:
		return v.UID()
	default:
		return 0
	}
}

func getGID(sysInfo any) uint32 {
	switch v := sysInfo.(type) {
	case interface{ Gid() uint32 }:
		return v.Gid()
	case interface{ GID() uint32 }:
		return v.GID()
	default:
		return 0
	}
}

func getATimSec(sysInfo any) int64 {
	switch v := sysInfo.(type) {
	case interface {
		Atim() interface{ Sec() int64 }
	}:
		return v.Atim().Sec()
	case interface {
		Atime() interface{ Sec() int64 }
	}:
		return v.Atime().Sec()
	default:
		return 0
	}
}

func getMTimSec(sysInfo any) int64 {
	switch v := sysInfo.(type) {
	case interface {
		Mtim() interface{ Sec() int64 }
	}:
		return v.Mtim().Sec()
	case interface {
		Mtime() interface{ Sec() int64 }
	}:
		return v.Mtime().Sec()
	default:
		return 0
	}
}

func getCTimSec(sysInfo any) int64 {
	switch v := sysInfo.(type) {
	case interface {
		Ctim() interface{ Sec() int64 }
	}:
		return v.Ctim().Sec()
	case interface {
		Ctime() interface{ Sec() int64 }
	}:
		return v.Ctime().Sec()
	default:
		return 0
	}
}
