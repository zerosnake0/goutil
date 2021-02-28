package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	output := flag.String("o", "", "output file")
	flag.Parse()

	fp, err := os.Create(*output)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	fmt.Fprintf(fp, "package ptr\n")
	for _, s := range []string{
		"int", "int8", "int16", "int32", "int64",
		"uint", "uint8", "uint16", "uint32", "uint64",
		"float32", "float64", "string", "bool",
	} {
		upper := strings.ToUpper(s[:1]) + s[1:]
		fmt.Fprintf(fp, `
func %[2]s(v %[1]s) *%[1]s {
	return &v
}
`, s, upper)
	}
}
