package main

import (
	"flag"
	"log"
	"strings"
)

func main() {
	var inputFile string
	var xxxTags string
	var defaultTag string
	flag.StringVar(&inputFile, "input", "", "path to input file")
	flag.StringVar(&xxxTags, "XXX_skip", "", "skip tags to inject on XXX fields")
	flag.BoolVar(&verbose, "verbose", false, "verbose logging")
	flag.StringVar(&defaultTag, "add_tag", "", "global add tag")

	flag.Parse()

	var xxxSkipSlice []string
	if len(xxxTags) > 0 {
		xxxSkipSlice = strings.Split(xxxTags, ",")
	}

	if len(inputFile) == 0 {
		log.Fatal("input file is mandatory")
	}

	areas, err := parseFile(inputFile, xxxSkipSlice, strings.TrimSpace(defaultTag))
	if err != nil {
		log.Fatal(err)
	}
	if err = writeFile(inputFile, areas); err != nil {
		log.Fatal(err)
	}
}
