package main

import (
	"log"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app  = kingpin.New("cfbfdump", "A dead simple tool to dump the contents of some cfbf file")
	file = kingpin.Arg("file", "The CFBF file to dump.").Required().File()
)

func main() {
	kingpin.Version("0.0.1")
	kingpin.Parse()
	log.Print("started")
}
