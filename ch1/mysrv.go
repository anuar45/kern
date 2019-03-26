package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World\n")
}
func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe("127.0.0.1:8888", nil)
}
