package main

import (
	"fmt"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func main() {

	http.ListenAndServe("http://127.0.0.1:9500", nil)
	http.HandleFunc("/aa", f1)
}
