package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

func main() {
	var list bool
	var decoderName, encoderName string
	flag.StringVar(&decoderName, "decoder", "ISO 8859-1", "name of decoder")
	flag.StringVar(&encoderName, "encoder", "ISO 8859-1", "name of encoder")
	flag.BoolVar(&list, "list", false, "list encoder/decoer names")
	flag.Parse()

	if list {
		for _, enc := range charmap.All {
			fmt.Println(enc)
		}
		return
	}

	var fname string
	if len(flag.Args()) < 1 {
		panic("include file name")
	} else {
		fname = flag.Args()[0]
	}

	f, err := os.Open(fname)
	if err != nil {
		panic(err)
	}

	var decoder, encoder charmap.Charmap
	for _, enc := range charmap.All {
		cmap, ok := enc.(*charmap.Charmap)
		if !ok {
			continue
		}
		if cmap.String() == decoderName {
			decoder = *cmap
		}
		if cmap.String() == encoderName {
			encoder = *cmap
		}
	}

	r := bufio.NewReader(f)
	td := transform.NewReader(r, decoder.NewDecoder())
	b, _ := ioutil.ReadAll(td)
	w := bufio.NewWriter(os.Stdout)
	tr := transform.NewWriter(w, encoder.NewEncoder())
	tr.Write(b)
	w.Flush()
}
