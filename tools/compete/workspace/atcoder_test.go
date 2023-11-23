package workspace

import (
	"testing"
)

func TestAtCoderProblemParser_URLToDir(t *testing.T) {
	parser := newAtCoderProblemParser()

	url := "https://atcoder.jp/contests/abc001/tasks/abc001_a"
	if result := parser.SupportsURL(url); !result {
		t.Errorf("SupportsURL() = %v, want %v", result, true)
	}

	problem, _ := parser.ParseURL("https://atcoder.jp/contests/abc001/tasks/abc001_a")
	if problem.ContestID() != "abc001" {
		t.Errorf("ContestID() = %v, want %v", problem.ContestID(), "abc001")
	}
	if problem.TaskID() != "abc001_a" {
		t.Errorf("TaskID() = %v, want %v", problem.TaskID(), "abc001_a")
	}
	if problem.Dir() != "atcoder/abc001/abc001_a" {
		t.Errorf("Dir() = %v, want %v", problem.Dir(), "atcoder/abc001/abc001_a")
	}
	if problem.URL() != "https://atcoder.jp/contests/abc001/tasks/abc001_a" {
		t.Errorf("URL() = %v, want %v", problem.URL(), "https://atcoder.jp/contests/abc001/tasks/abc001_a")
	}
}

func TestAtCoderProblemParser_DirToURL(t *testing.T) {
	parser := newAtCoderProblemParser()

	dir := "compete/go/atcoder/abc001/abc001_a"
	if result := parser.SupportsDir(dir); !result {
		t.Errorf("SupportsDir() = %v, want %v", result, true)
	}

	problem, _ := parser.ParseDir(dir)
	if problem.ContestID() != "abc001" {
		t.Errorf("ContestID() = %v, want %v", problem.ContestID(), "abc001")
	}
	if problem.TaskID() != "abc001_a" {
		t.Errorf("TaskID() = %v, want %v", problem.TaskID(), "abc001_a")
	}
	if problem.Dir() != "atcoder/abc001/abc001_a" {
		t.Errorf("Dir() = %v, want %v", problem.Dir(), "atcoder/abc001/abc001_a")
	}
	if problem.URL() != "https://atcoder.jp/contests/abc001/tasks/abc001_a" {
		t.Errorf("URL() = %v, want %v", problem.URL(), "https://atcoder.jp/contests/abc001/tasks/abc001_a")
	}
}
