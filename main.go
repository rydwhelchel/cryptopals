package main

import (
	"log"
	"os"
	"strings"

	"github.com/rydwhelchel/cryptopals/problemset1"
)

func main() {
	context := problemset1.Context{
		BookPath: "./throughthelookingglass.txt",
		S1c4Path: "./s1c4.txt",
	}

	log.Println("~~~~~~~~ Problem set 1, challenge 4 ~~~~~~~~")
	file, err := os.ReadFile(context.S1c4Path)
	if err != nil {
		log.Panicln("Unable to read file")
	}
	lines := strings.Split(string(file), "\n")
	result, err := context.FindSingleByteEncryption(lines)
	if err != nil {
		log.Panicf("Failed to find single byte encryption - %v\n", err)
	}
	log.Printf("Single byte encryption result: %+v", result)
}
