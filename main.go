package main

import (
	"fmt"
	"strings"
)

func main() {
	parts := []string{"a", "b", "c"}
	fmt.Println(Merge(parts))
	fmt.Println(Basic(parts))
	fmt.Println(Build(parts))
}

func Merge(parts []string) string {
	return strings.Join(parts, "")
}

func Basic(parts []string) string {
	var ret string
	for _, part := range parts {
		ret += part
	}
	return ret
}

func Build(parts []string) string {
	var sb strings.Builder
	for _, part := range parts {
		sb.WriteString(part)
	}
	return sb.String()
}
