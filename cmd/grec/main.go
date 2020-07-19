package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/maxhawkins/grec"
)

func main() {
	flag.Parse()
	filename := flag.Arg(0)

	if flag.NArg() < 1 {
		fmt.Fprintln(os.Stderr, "usage: grec <infile.m4a>")
		os.Exit(1)
	}

	tran, err := grec.ParseFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	if err := enc.Encode(tran); err != nil {
		log.Fatal(err)
	}
}
