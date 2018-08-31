package main

import (
	"fmt"
	"net/http"
)

func handler( w http.ResponseWriter, r *http.Request ) {
	fmt.Fprintf( w, "Hello, World!\n" )
}

func testHandler( w http.ResponseWriter, r *http.Request ) {
	fmt.Fprintf( w, "Hello, World! for /test/\n" )
}

func testIndexHandler( w http.ResponseWriter, r *http.Request ) {
	fmt.Fprintf( w, "Hello, World! for /test/index.html\n" )
}

func testTestHandler( w http.ResponseWriter, r *http.Request ) {
	fmt.Fprintf( w, "Hello, World! for /test/test.html\n" )
}

func main() {
	http.HandleFunc( "/", handler )
	http.HandleFunc( "/test/", testHandler )
	http.HandleFunc( "/test/index.html", testIndexHandler )
	http.HandleFunc( "/test/test.html", testTestHandler )
	http.ListenAndServe( ":8080", nil )
}