package main

import (
	"fmt"

	to_be "github.com/synesissoftware/to-be.Go"

	"os"
)

func reportOnVar(key string) {

	if value, exists := os.LookupEnv(key); exists {

		bool_strings := []string{
			"is not",
			"is",
		}

		bool2index := func(b bool) uint {
			if b {
				return 1
			} else {
				return 0
			}
		}

		fmt.Printf("environment variable '%s' exists and has the value '%s':\n", key, value)
		fmt.Printf("\t%s falsey\n", bool_strings[bool2index(to_be.StringIsFalsey(value))])
		fmt.Printf("\t%s truey\n", bool_strings[bool2index(to_be.StringIsTruey(value))])
		fmt.Printf("\t%s truthy\n", bool_strings[bool2index(to_be.StringIsTruthy(value))])
	} else {

		fmt.Printf("environment variable '%s' does not exist\n", key)
	}
}

func main() {
	reportOnVar("DEBUG_MODE")
	reportOnVar("ENABLE_LOGGING")
	reportOnVar("ENABLE_TRACING")
}
