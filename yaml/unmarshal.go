package yaml

import (
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

func UnmarshalFromFile(filename string, obj interface{}) error {
	fp, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fp.Close()
	return UnmarshalFromReader(fp, obj)
}

func UnmarshalFromReader(reader io.Reader, obj interface{}) error {
	dec := yaml.NewDecoder(reader)
	dec.KnownFields(true)
	return dec.Decode(obj)
}

func MarshalToFile(filename string, obj interface{}) error {
	fp, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer fp.Close()
	return MarshalToWriter(fp, obj)
}

func MarshalToWriter(writer io.Writer, obj interface{}) error {
	enc := yaml.NewEncoder(writer)
	return enc.Encode(obj)
}
