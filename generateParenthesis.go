package leetcode

import "fmt"

func generateParenthesis(n int) []string {
	if n == 0 {
		return nil
	}

	sub := generateParenthesis(n - 1)
	if len(sub) == 0 {
		return []string{"()"}
	}

	result := make([]string, 0, len(sub)*3)

	for i := range sub {
		result = append(result, fmt.Sprintf("()%s", sub[i]))
		result = append(result, fmt.Sprintf("(%s)", sub[i]))
		if i > 0 {
			result = append(result, fmt.Sprintf("%s()", sub[i]))
		}
	}

	return result
}
