package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func waveUploadHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	outFile, err := os.Create("/tmp/hoge.bin")
	if err != nil {
		log.Println(err)
	}
	_, err = io.Copy(outFile, r.Body)
}

func main() {
	http.HandleFunc("/upload", waveUploadHandler)
	http.ListenAndServe(":8080", nil)
}
