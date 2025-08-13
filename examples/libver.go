package main

import (
	to_be "github.com/synesissoftware/to-be.Go"
	"github.com/synesissoftware/ver2go"

	"fmt"
)

func main() {
	fmt.Printf("to_be v%s\n", ver2go.CalcVersionString(to_be.VersionMajor, to_be.VersionMinor, to_be.VersionPatch, to_be.VersionAB))
	fmt.Printf("ver2go v%s\n", ver2go.CalcVersionString(ver2go.VersionMajor, ver2go.VersionMinor, ver2go.VersionPatch, ver2go.VersionAB))
}
