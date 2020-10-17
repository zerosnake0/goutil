package convert

import (
	"runtime/debug"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLocalByteToString(t *testing.T) {
	must := require.New(t)

	s := "test string"
	b := LocalStringToBytes(s)

	debug.FreeOSMemory()

	must.Equal([]byte(s), b)
	must.Equal(len(s), len(b))
	must.Equal(len(s), cap(b))
}

func TestLocalByteToString2(t *testing.T) {
	must := require.New(t)

	b := []byte{'t', 'e', 's', 't'}
	s := LocalByteToString(b)

	debug.FreeOSMemory()

	must.Equal(string(b), s)
}
