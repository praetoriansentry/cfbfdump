package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/richardlehane/mscfb"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app      = kingpin.New("cfbfdump", "A dead simple tool to dump the contents of some cfbf file")
	cfbffile = kingpin.Arg("file", "The CFBF file to dump.").Required().File()
)

func main() {
	log.Print("Started")
	kingpin.Version("0.0.1")
	kingpin.Parse()
	doc, err := mscfb.New(*cfbffile)
	kingpin.FatalIfError(err, "Failed to parse file")
	log.Print("Read file")
	dir, err := ioutil.TempDir(".", (*cfbffile).Name())
	kingpin.FatalIfError(err, "Cant create temp directory")
	log.Printf("Creating output directory: %s", dir)
	for entry, err := doc.Next(); err == nil; entry, err = doc.Next() {
		p := strings.Join(entry.Path, string(os.PathSeparator))
		fullPath := dir + string(os.PathSeparator) + p
		_ = os.Mkdir(fullPath, os.ModePerm)
		buf := make([]byte, entry.Size)
		i, _ := doc.Read(buf)
		if i > 0 {
			ioutil.WriteFile(fullPath+string(os.PathSeparator)+entry.Name, buf, os.ModePerm)
			entryData, _ := json.Marshal(entry)
			log.Printf("Writing entry: %s", string(entryData[:]))

		}

	}

}
