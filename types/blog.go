package types

import "strings"

type Blog struct {
	Name    string
	Url     string
	Content string
}

func (b *Blog) FormatName() string {
	if b == nil {
		return ""
	}

	formatName := b.Name
	for _, f := range formatFunc {
		formatName = f(formatName)
	}
	return formatName
}

var formatFunc = []func(string) string{
	func(s string) string { return strings.ReplaceAll(s, "-", " ") },
	func(s string) string { return strings.ReplaceAll(s, "_", " / ") },
	func(s string) string { return strings.TrimSuffix(s, ".md") },
}
