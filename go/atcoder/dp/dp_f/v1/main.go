package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func readLines() []string {
	stdin, _ := io.ReadAll(os.Stdin)
	return strings.Split(string(stdin), "\n")
}

func max(s, t int) int {
	if s > t {
		return s
	}
	return t
}

func main() {
	lines := readLines()
	s := lines[0]
	t := lines[1]
	dp := make([][]int, len(s)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(t)+1)
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = 0
		}
	}
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(t); j++ {
			if s[i] == t[j] {
				dp[i+1][j+1] = max(dp[i][j]+1, dp[i+1][j+1])
			}
			dp[i+1][j+1] = max(dp[i+1][j+1], dp[i+1][j])
			dp[i+1][j+1] = max(dp[i+1][j+1], dp[i][j+1])
		}
	}

	out := ""
	i := len(s)
	j := len(t)
	for i > 0 && j > 0 {
		if dp[i][j] == dp[i-1][j] {
			i--
		} else if dp[i][j] == dp[i][j-1] {
			j--
		} else {
			out = string(s[i-1]) + out
			i--
			j--
		}
	}
	fmt.Println(out)
}
