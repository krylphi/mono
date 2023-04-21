package utils

import "strings"

func Print(s string) {
	println(strings.Join([]string{s, "-v0-"}, "|"))
}
