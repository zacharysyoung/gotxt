package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"sort"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

var encodings = map[string]encoding.Encoding{
	"UTF8":        unicode.UTF8,
	"UTF8BOM":     unicode.UTF8BOM,
	"UTF16BE":     unicode.UTF16(unicode.BigEndian, unicode.UseBOM),
	"ISO 8859-6E": charmap.ISO8859_6E,
	"ISO 8859-6I": charmap.ISO8859_6I,
	"ISO 8859-8E": charmap.ISO8859_8E,
	"ISO 8859-8I": charmap.ISO8859_8I,
}

var decoderName, encoderName string
var list bool

func init() {
	flag.StringVar(&decoderName, "decoder", "UTF8", "name of decoder")
	flag.StringVar(&encoderName, "encoder", "UTF8", "name of encoder")
	flag.BoolVar(&list, "list", false, "list encoder/decoder names")

	for _, enc := range charmap.All {
		cmap, ok := enc.(*charmap.Charmap)
		if !ok {
			continue
		}
		encodings[cmap.String()] = enc
	}
}

func main() {
	flag.Parse()

	if list {
		var names []string
		for name := range encodings {
			names = append(names, name)
		}
		sort.Strings(names)
		for _, name := range names {
			fmt.Println(name)
		}
		return
	}

	decoder, ok := encodings[decoderName]
	if !ok {
		errorOut(decoderName + " is not a valid encoding")
	}

	encoder, ok := encodings[encoderName]
	if !ok {
		errorOut(encoderName + " is not a valid encoding")
	}

	var f *os.File
	var err error
	if len(flag.Args()) < 1 {
		f = os.Stdin
	} else {
		fname := flag.Args()[0]
		f, err = os.Open(fname)
		if err != nil {
			errorOut(fmt.Sprintf("could not open file %s: %v", fname, err))
		}
	}
	defer f.Close()

	td := transform.NewReader(f, decoder.NewDecoder())
	b, err := io.ReadAll(td)
	if err != nil {
		panic(err)
	}

	w := transform.NewWriter(os.Stdout, encoder.NewEncoder())
	w.Write(b)
}

func errorOut(s string) {
	fmt.Fprintln(os.Stderr, "error: "+s)
	os.Exit(1)
}
