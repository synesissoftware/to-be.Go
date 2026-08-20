package to_be_test

import (
	to_be "github.com/synesissoftware/to-be.Go"

	"github.com/stretchr/testify/require"

	"testing"
)

const (
	Expected_VersionMajor uint16 = 0
	Expected_VersionMinor uint16 = 3
	Expected_VersionPatch uint16 = 1
	Expected_VersionAB    uint16 = 0xFFFF
)

func Test_Version_Elements(t *testing.T) {
	require.Equal(t, Expected_VersionMajor, to_be.VersionMajor)
	require.Equal(t, Expected_VersionMinor, to_be.VersionMinor)
	require.Equal(t, Expected_VersionPatch, to_be.VersionPatch)
	require.Equal(t, Expected_VersionAB, to_be.VersionAB)
}

func Test_Version(t *testing.T) {
	require.Equal(t, uint64(0x0000_0003_0001_FFFF), to_be.Version())
}

func Test_Version_String(t *testing.T) {
	require.Equal(t, "0.3.1", to_be.VersionString())
}
