package main

import (
	"bytes"
	"debug/buildinfo"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

var (
	j    bool
	file string
)

func init() {
	flag.BoolVar(&j, "json", false, "json output")
	flag.StringVar(&file, "file", "-", "file")
}

func main() {
	flag.Parse()
	var r io.ReaderAt
	if file == "-" {
		b, err := io.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}
		r = bytes.NewReader(b)
	} else {
		f, err := os.Open(file)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		r = f
	}
	info, err := buildinfo.Read(r)
	if err != nil {
		log.Fatal(err)
	}
	if j {
		b, err := json.MarshalIndent(info, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(b))
	} else {
		fmt.Print(info.String())
	}
}
