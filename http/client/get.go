package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://www.yahoo.co.jp/"
	res, _ := http.Get(url)
	defer res.Body.Close()

	byteArray, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(byteArray))
}
