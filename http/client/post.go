package main

import (
	"net/http"
	"net/url"
	"io/ioutil"
	"fmt"
)

func main() {
	urlStr := "http://localhost:8000/post_server.php"
	values := url.Values{}
	values.Add( "name", "iPhone")
	values.Add( "price", "10000")
	res, _ := http.PostForm( urlStr, values )
	defer res.Body.Close()

	byteArray, _ := ioutil.ReadAll( res.Body )
	fmt.Println( string( byteArray ))
}