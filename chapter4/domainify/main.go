package main

import (
	"time"
	"bufio"
	"os"
	"unicode"
	"fmt"
	"math/rand"
	"strings"
	"flag"
	"regexp"
)


const allowedChars = "abcdefghijklmnopqrstuvwxyz0123456789_-"

func main() {
	var tldsArguments = flag.String("tlds", "com net", "TLDs list.")
	var verbose = flag.Bool("v", false, "Print possible TLDs.")
	flag.Parse()

	r, _ := regexp.Compile("\\w+")
	tlds := r.FindAllString(*tldsArguments, -1)

	if *verbose {
		fmt.Printf("Possible TLDs:\n%s\n", tlds)
	}

	rand.Seed(time.Now().UTC().UnixNano())
	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		text := strings.ToLower(s.Text())
		var newText []rune

		firstSpace := true

		for _, r := range text {
			if unicode.IsSpace(r) {
				if firstSpace {
					r = '-'
					firstSpace = false
				} else {
					continue
				}
			} else {
				firstSpace = true
			}

			if !strings.ContainsRune(allowedChars, r) {
				continue
			}
			newText = append(newText, r)
		}
		fmt.Println(string(newText) + "." +
			tlds[rand.Intn(len(tlds))])
	}
}