package cmd

import "fmt"

var version string = "v0.1.2"

func Version(args []string) int {
	fmt.Printf("volt version: %s\n", version)

	return 0
}
