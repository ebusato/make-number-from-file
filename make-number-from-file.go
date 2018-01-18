package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	inFile, _ := os.Open("file.txt")
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)

	for {
		word := ""
		for i := 0; i < 8; i++ {
			if !scanner.Scan() {
				return
			}
			s := scanner.Text()
			word += s
		}
		fmt.Printf("word = %v\n", word)
		number, _ := strconv.ParseUint(word, 2, 8)
		fmt.Printf("number = %v (%b)\n", number, number)
	}
}
