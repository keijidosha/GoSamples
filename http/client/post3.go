package main

import (
	"net/http"
	"net/url"
	"io/ioutil"
	"fmt"
	"strings"
	"time"
)

func post( ch chan int, urlStr string ) {
	values := url.Values{}
	values.Add( "name", "iPhone")
	values.Add( "price", "10000")

	req, err := http.NewRequest( "POST", urlStr, strings.NewReader( values.Encode()))
	if err != nil {
		fmt.Println( err )
		return
	}

	req.Header.Add( "User-Agent", "Go" )
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	req.SetBasicAuth( "user", "password" )

	client := &http.Client{ Timeout: time.Duration(10 * time.Second) }
	res, err := client.Do( req )

	if err != nil {
		fmt.Println( err )
		return
	}

	defer res.Body.Close()

	byteArray, _ := ioutil.ReadAll( res.Body )
	fmt.Println( string( byteArray ))

	ch <- 0
}

func main() {
	ch := make( chan int )

	urlStr := "http://localhost:8000/post_server.php"
	go post( ch, urlStr )

	ret := <- ch
	fmt.Printf( "ret: %d\n", ret )
}