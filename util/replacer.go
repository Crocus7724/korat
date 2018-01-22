package util

import (
	"github.com/shurcooL/githubql"
	"strings"
)

func ReplaceBrackets(s githubql.String) string {
	r := strings.Replace(string(s), "[", "\\[", -1)
	return strings.Replace(r, "]", "\\]", -1)
}
