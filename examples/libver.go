package main

import (
	"fmt"

	to_be "github.com/synesissoftware/to-be.Go"
	"github.com/synesissoftware/ver2go"
)

func main() {
	fmt.Printf("to_be v%s\n", to_be.VersionString())
	fmt.Printf("ver2go v%s\n", ver2go.VersionString())
}
