package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
	"log"
	"flag"
)


const otherWord = "*"

func readTransformsFile(filename string) (transforms []string) {
	transforms = make([]string, 0)
	openedFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer openedFile.Close()

	scanner := bufio.NewScanner(openedFile)
	for scanner.Scan() {
		transforms = append(transforms, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return
}

func main() {
	var filename = flag.String("f", "transforms.ini", "Filename for reading transforms.")
	var verbose = flag.Bool("v", false, "Send possible transforms to stdin.")
	flag.Parse()

	transforms := readTransformsFile(*filename)

	if *verbose {
		log.Printf("Available transforms (read from the file '%s'):\n%s", *filename, transforms)
	}

	rand.Seed(time.Now().UTC().UnixNano())
	s := bufio.NewScanner(os.Stdin)
	log.Print("transformed =")
	for s.Scan() {
		t := transforms[rand.Intn(len(transforms))]
		transformed := strings.Replace(t, otherWord, s.Text(), -1)
		log.Print(transformed)
		fmt.Println(transformed)
	}
}