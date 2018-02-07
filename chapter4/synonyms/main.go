package main

import (
	"os"
	"bufio"
	"log"
	"fmt"
	"go_blueprints/chapter4/thesaurus"
)

func main() {
	apiKey := os.Getenv("BHT_APIKEY")
	thesaurusVal := &thesaurus.BigHuge{APIKey: apiKey}
	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		word := s.Text()
		syns, err := thesaurusVal.Synonyms(word)
		if err != nil {
			log.Fatalln("Failed when looking for synonyms for \"" + word + "\"", err)
		}
		if len(syns) == 0 {
			log.Fatalln("Couldn't find any synonyms for \"" + word + "\"")
		}
		log.Print("synomyms = ")
		for _, syn := range syns {
			log.Print(syn)
			fmt.Println(syn)
		}
	}
}