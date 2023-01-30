//go:build fixer
// +build fixer

package main

func init() {
	features = append(features,
		"This is a fixer plugin of the app",
	)
}
