package diff

import (
	"strings"
)

func CompareTexts(a, b string) string {
	wordsA := strings.Fields(a)
	wordsB := strings.Fields(b)
	return plainTextDiff(wordsA, wordsB)
}

func plainTextDiff(a, b []string) string {
	lcs := computeLCSMatrix(a, b)
	return buildDiffOutput(a, b, lcs)
}

func computeLCSMatrix(a, b []string) [][]int {
	m, n := len(a), len(b)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if a[i-1] == b[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp
}

func buildDiffOutput(a, b []string, dp [][]int) string {
	i, j := len(a), len(b)
	var chunks []string
	var deletions, insertions []string

	flushChanges := func() {
		var flush []string
		if len(deletions) > 0 {
			flush = append(flush, "[-"+strings.Join(deletions, " ")+"-]")
			deletions = nil
		}
		if len(insertions) > 0 {
			flush = append(flush, "{+"+strings.Join(insertions, " ")+"+}")
			insertions = nil
		}
		chunks = append(flush, chunks...)
	}

	for i > 0 || j > 0 {
		if i > 0 && j > 0 && a[i-1] == b[j-1] {
			flushChanges()
			chunks = append([]string{a[i-1]}, chunks...)
			i--
			j--
		} else if j > 0 && (i == 0 || dp[i][j-1] >= dp[i-1][j]) {
			insertions = append([]string{b[j-1]}, insertions...)
			j--
		} else if i > 0 {
			deletions = append([]string{a[i-1]}, deletions...)
			i--
		}
	}

	flushChanges()
	return strings.Join(chunks, " ")
}
