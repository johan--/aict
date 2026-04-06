package basename

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/synseqack/aict/internal/meta"
	"github.com/synseqack/aict/internal/tool"
	xmlout "github.com/synseqack/aict/internal/xml"
)

func init() {
	tool.Register("basename", Run)
	tool.RegisterMeta("basename", tool.GenerateSchema("basename", "Print filename portion of file paths", Config{}))
}

type Config struct {
	XML    bool
	JSON   bool
	Plain  bool
	Pretty bool
}

type BasenameResult struct {
	XMLName   xml.Name        `xml:"basename"`
	Paths     []BasenameEntry `xml:"entry,omitempty"`
	Timestamp int64           `xml:"timestamp,attr"`
	Errors    []BasenameError `xml:"error,omitempty"`
}

func (*BasenameResult) isBasenameResult() {}

type BasenameEntry struct {
	XMLName   xml.Name `xml:"entry"`
	Path      string   `xml:"path,attr"`
	Base      string   `xml:"base,attr"`
	Stem      string   `xml:"stem,attr"`
	Extension string   `xml:"extension,attr"`
}

type BasenameError struct {
	XMLName xml.Name `xml:"error"`
	Code    int      `xml:"code,attr"`
	Msg     string   `xml:"msg,attr"`
	Path    string   `xml:"path,attr"`
}

func Run(args []string) error {
	cfg, paths := parseFlags(args)

	if len(paths) == 0 {
		return outputResult(&BasenameResult{Timestamp: meta.Now()}, cfg)
	}

	suffix := ""
	if len(paths) > 1 && !strings.HasPrefix(paths[1], "-") {
		suffix = paths[1]
		paths = paths[1:]
	}

	result := &BasenameResult{Timestamp: meta.Now()}
	for _, p := range paths {
		base := filepath.Base(p)
		if suffix != "" && strings.HasSuffix(base, suffix) {
			base = strings.TrimSuffix(base, suffix)
		}

		ext := filepath.Ext(base)
		stem := strings.TrimSuffix(base, ext)

		result.Paths = append(result.Paths, BasenameEntry{
			Path:      p,
			Base:      base,
			Stem:      stem,
			Extension: ext,
		})
	}

	return outputResult(result, cfg)
}

func parseFlags(args []string) (Config, []string) {
	var cfg Config
	var positional []string

	for _, arg := range args {
		switch arg {
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

func outputResult(result *BasenameResult, cfg Config) error {
	if cfg.JSON {
		return xmlout.WriteJSON(os.Stdout, result)
	}
	if cfg.Plain {
		return writePlain(os.Stdout, result)
	}
	return xmlout.WriteXML(os.Stdout, result, cfg.Pretty)
}

func writePlain(w io.Writer, result *BasenameResult) error {
	for _, p := range result.Paths {
		fmt.Fprintln(w, p.Base)
	}
	return nil
}
