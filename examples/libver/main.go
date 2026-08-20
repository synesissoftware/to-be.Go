package main

import (
	to_be "github.com/synesissoftware/to-be.Go"
	ver2go "github.com/synesissoftware/ver2go"

	"fmt"
)

func main() {
	fmt.Printf("to-be.Go v%s\n", to_be.VersionString())
	fmt.Printf("ver2go v%s\n", ver2go.VersionString())
}
