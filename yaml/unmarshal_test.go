package yaml

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func TestUnmarshalFromFile(t *testing.T) {
	t.Run("error", func(t *testing.T) {
		err := UnmarshalFromFile("", nil)
		require.True(t, os.IsNotExist(err))
	})
	t.Run("normal", func(t *testing.T) {
		must := require.New(t)

		fp, err := ioutil.TempFile("", "")
		must.NoError(err)

		filename := fp.Name()

		type st struct {
			A string `yaml:"a"`
		}
		obj := st{A: "dummy"}
		b, err := yaml.Marshal(obj)
		must.NoError(err)
		_, err = io.Copy(fp, bytes.NewReader(b))
		must.NoError(err)
		must.NoError(fp.Close())

		var obj2 st
		err = UnmarshalFromFile(filename, &obj2)
		must.NoError(err)
		must.Equal(obj, obj2)
	})
}

func TestUnmarshalFromReader(t *testing.T) {
	t.Run("error", func(t *testing.T) {
		var obj struct {
			A string `yaml:"a"`
		}
		err := UnmarshalFromReader(strings.NewReader(`
a: b
b: c`), &obj)
		require.Error(t, err)
	})
}
