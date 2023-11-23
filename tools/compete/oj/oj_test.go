package oj

import (
	"os"
	"testing"
)

func TestExecOjDownload(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "test")
	if err != nil {
		t.Errorf("os.MkdirTemp() error = %v", err)
	}
	defer os.RemoveAll(tmpDir)

	if err := Download("https://atcoder.jp/contests/abc001/tasks/abc001_1", tmpDir); err != nil {
		t.Errorf("Download() error = %v", err)
	}
}

func TestExecOjSubmit(t *testing.T) {
	testFilename := "testdata/main.go"
	if err := Submit("https://atcoder.jp/contests/abc001/tasks/abc001_1", testFilename); err != nil {
		t.Errorf("Submit() error = %v", err)
	}
}
