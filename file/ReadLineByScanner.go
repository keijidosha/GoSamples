package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fp, err := os.Open("hoge.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
