package main

import (
	"fmt"
	"strings"
)

func ToSnakeCase(str string) string {
	// step 0: handle empty string
	if str == "" {
		return ""
	}

	// step 1: CamelCase to snake_case
	for i := 0; i < len(str); i++ {
		if str[i] >= 'A' && str[i] <= 'Z' {
			if i == 0 {
				str = strings.ToLower(string(str[i])) + str[i+1:]
			} else {
				str = str[:i] + "_" + strings.ToLower(string(str[i])) + str[i+1:]
			}
		}
	}

	// step 2: remove spaces
	str = strings.ReplaceAll(str, " ", "")

	// step 3: remove double-underscores
	str = strings.ReplaceAll(str, "__", "_")

	return str
}

func main() {
	fmt.Println(ToSnakeCase("Hello World"))
}
