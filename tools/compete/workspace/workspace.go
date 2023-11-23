package workspace

import (
	"fmt"

	"github.com/kzmshx/compete/tools/compete/language"
)

type Parser interface {
	ParseDir(dir string) (Problem, error)
	ParseURL(url string) (Problem, error)
	SupportsDir(dir string) bool
	SupportsURL(url string) bool
}

type Problem interface {
	Dir() string
	URL() string
}

var parsers = []Parser{
	newAtCoderProblemParser(),
}

func Create(url string, lang string) (Problem, error) {
	for _, parser := range parsers {
		if parser.SupportsURL(url) {
			return parser.ParseURL(url)
		}
	}
	return nil, fmt.Errorf("unsupported url: %s", url)
}

func CreateFromDir(dir string) (Problem, error) {
	for _, parser := range parsers {
		if parser.SupportsDir(dir) {
			return parser.ParseDir(dir)
		}
	}
	return nil, fmt.Errorf("unsupported dir: %s", dir)
}

type Workspace struct {
	URL      string
	Language language.Language
	Dir      string
	Version  int
}

// compete new <url> <language>
// compete test <dir>
// compete submit <dir>
