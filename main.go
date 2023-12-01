package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

var (
	_utf8    = unicode.UTF8
	_utf8BOM = unicode.UTF8BOM
	_utf16BE = unicode.UTF16(unicode.BigEndian, unicode.UseBOM)
	_8859_6E = charmap.ISO8859_6E
	_8859_6I = charmap.ISO8859_6I
	_8859_8E = charmap.ISO8859_8E
	_8859_8I = charmap.ISO8859_8I

	specialNames = map[encoding.Encoding]string{
		charmap.CodePage858: "IBM Code Page 858",

		_8859_6E: "ISO 8859-6E",
		_8859_6I: "ISO 8859-6I",
		_8859_8E: "ISO 8859-8E",
		_8859_8I: "ISO 8859-8I",

		_utf8:    "UTF8",
		_utf8BOM: "UTF8BOM",
		_utf16BE: "UTF16BE",
	}
)

var allEncodings = []encoding.Encoding{
	charmap.CodePage037,
	charmap.CodePage437,
	charmap.CodePage850,
	charmap.CodePage852,
	charmap.CodePage855,
	charmap.CodePage858,
	charmap.CodePage860,
	charmap.CodePage862,
	charmap.CodePage863,
	charmap.CodePage865,
	charmap.CodePage866,
	charmap.CodePage1047,
	charmap.CodePage1140,
	charmap.ISO8859_1,
	charmap.ISO8859_2,
	charmap.ISO8859_3,
	charmap.ISO8859_4,
	charmap.ISO8859_5,
	charmap.ISO8859_6,
	_8859_6E,
	_8859_6I,
	charmap.ISO8859_7,
	charmap.ISO8859_8,
	_8859_8E,
	_8859_8I,
	charmap.ISO8859_9,
	charmap.ISO8859_10,
	charmap.ISO8859_13,
	charmap.ISO8859_14,
	charmap.ISO8859_15,
	charmap.ISO8859_16,
	charmap.KOI8R,
	charmap.KOI8U,
	charmap.Macintosh,
	charmap.MacintoshCyrillic,
	_utf8,
	_utf8BOM,
	_utf16BE,
	charmap.Windows874,
	charmap.Windows1250,
	charmap.Windows1251,
	charmap.Windows1252,
	charmap.Windows1253,
	charmap.Windows1254,
	charmap.Windows1255,
	charmap.Windows1256,
	charmap.Windows1257,
	charmap.Windows1258,
	charmap.XUserDefined,
}

var (
	flagInName, flagOutName string
	flagList                bool

	encodingNames    []string
	normNameEncoding = make(map[string]encoding.Encoding)
)

func normName(s string) string {
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ReplaceAll(s, "-", "")
	s = strings.ToLower(s)
	return s
}

func init() {
	flag.StringVar(&flagInName, "in", "UTF8", "input encoding name")
	flag.StringVar(&flagOutName, "out", "UTF8", "output encoding name")
	flag.BoolVar(&flagList, "list", false, "list available encoding names")

	for _, enc := range allEncodings {
		name, ok := specialNames[enc]
		if !ok {
			name = enc.(*charmap.Charmap).String()
		}
		encodingNames = append(encodingNames, name)
		normNameEncoding[normName(name)] = enc
	}
}

func main() {
	flag.Parse()

	if flagList {
		for _, name := range encodingNames {
			fmt.Println(name)
		}
		return
	}

	var inEnc, outEnc encoding.Encoding

	if inEnc = normNameEncoding[flagInName]; inEnc == nil {
		errorOut("invalid input encoding name: " + flagInName)
	}
	if outEnc = normNameEncoding[flagOutName]; outEnc == nil {
		errorOut("invalid output encoding name: " + flagOutName)
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

	r := transform.NewReader(f, inEnc.NewDecoder())
	w := transform.NewWriter(os.Stdout, outEnc.NewEncoder())

	n, err := io.Copy(w, r)
	if err != nil {
		errorOut(fmt.Sprintf("could not transcode, read input up to byte %d: %v", n+1, err))
	}
}

func errorOut(s string) {
	fmt.Fprintln(os.Stderr, "error: "+s)
	os.Exit(1)
}
