package Solutions

import "strings"

func reverseWords(s string) string {
	var sb strings.Builder

	substrings := strings.Split(s, " ")

	for i := len(substrings) - 1; i >= 0; i-- {
		if strings.TrimSpace(substrings[i]) != ""{
			if sb.Len() > 0 {
				sb.WriteString(" ")
			}
			sb.WriteString(substrings[i])
		}
	}

	return sb.String()
}