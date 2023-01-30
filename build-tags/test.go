//go:build (fixer && ignore) || test
// +build fixer,ignore test

package main

func init() {
	features = append(features,
		"This is a test plugin of the app",
	)
}
