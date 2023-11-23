package oj

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
)

const cmd = "oj"

func Exists() bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}

func Download(url string, targetDir string) error {
	return execCommand([]string{cmd, "d", url, "-d", targetDir})
}

func Submit(url string, targetFile string) error {
	return execCommand([]string{cmd, "s", url, targetFile, "--yes"})
}

func execCommand(commandLine []string) error {
	cmd := exec.Command(commandLine[0], commandLine[1:]...)

	var stdout, stderr io.Reader
	var err error

	stdout, err = cmd.StdoutPipe()
	if err != nil {
		return err
	}

	stderr, err = cmd.StderrPipe()
	if err != nil {
		return err
	}

	if err = cmd.Start(); err != nil {
		return err
	}

	scanner := bufio.NewScanner(stdout)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	scanner = bufio.NewScanner(stderr)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err = cmd.Wait(); err != nil {
		return err
	}

	return err
}
