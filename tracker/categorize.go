package tracker

import "strings"

var categories = map[string][]string{
	"Writing": {"write", "document", "blog"},
	"Coding":  {"code", "bug", "dev", "fix"},
}

func Categorize(name string) string {

	name = strings.ToLower(name)
	for cat, keywords := range categories {

		for _, kw := range keywords {

			if strings.Contains(name, kw) {
				return cat
			}
		}
	}
	return "General"
}
