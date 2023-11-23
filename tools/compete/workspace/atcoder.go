package workspace

import (
	"fmt"
	"net/url"
	"path"
	"regexp"
)

var atCoderProblemDirRegexp = regexp.MustCompile(`atcoder/([^/]+)/([^/]+)$`)
var atCoderProblemURLRegexp = regexp.MustCompile(`^https://atcoder.jp/contests/([^/]+)/tasks/([^/]+)$`)

type AtCoderProblemParser struct{}

func newAtCoderProblemParser() *AtCoderProblemParser {
	return &AtCoderProblemParser{}
}

func (p *AtCoderProblemParser) ParseDir(dir string) (Problem, error) {
	matches := atCoderProblemDirRegexp.FindStringSubmatch(dir)
	if len(matches) != 3 {
		return nil, fmt.Errorf("invalid dir: %s", dir)
	}
	return &AtCoderProblem{
		contestID: matches[1],
		taskID:    matches[2],
	}, nil
}

func (p *AtCoderProblemParser) ParseURL(url string) (Problem, error) {
	matches := atCoderProblemURLRegexp.FindStringSubmatch(url)
	if len(matches) != 3 {
		return nil, fmt.Errorf("invalid url: %s", url)
	}
	return &AtCoderProblem{
		contestID: matches[1],
		taskID:    matches[2],
	}, nil
}

func (p *AtCoderProblemParser) SupportsDir(dir string) bool {
	return atCoderProblemDirRegexp.MatchString(dir)
}

func (p *AtCoderProblemParser) SupportsURL(url string) bool {
	return atCoderProblemURLRegexp.MatchString(url)
}

type AtCoderProblem struct {
	contestID string
	taskID    string
}

func (p *AtCoderProblem) ContestID() string {
	return p.contestID
}

func (p *AtCoderProblem) TaskID() string {
	return p.taskID
}

func (p *AtCoderProblem) Dir() string {
	return path.Join("atcoder", p.contestID, p.taskID)
}

func (p *AtCoderProblem) URL() string {
	result, _ := url.JoinPath("https://atcoder.jp/contests", p.contestID, "tasks", p.taskID)
	return result
}
