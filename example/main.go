package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/peterjliu/textproc"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var text_file = flag.String("infile", "", "some file with text")

func main() {
	flag.Parse()
	log.Printf("Reading from %s", *text_file)
	data, err := ioutil.ReadFile(*text_file)
	check(err)
	sentences := textproc.GetSentences(string(data))
	for i, s := range sentences {
		fmt.Printf("### %d ###", i)
		fmt.Println(s)
	}
}
