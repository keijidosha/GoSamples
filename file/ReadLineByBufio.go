package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fp, err := os.Open("ReadLine.go")
	if err != nil {
		log.Fatal(err)
	}

	defer fp.Close()

	reader := bufio.NewReaderSize(fp, 8192)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(line))
	}
}
