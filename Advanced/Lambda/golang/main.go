package main

import (
	"fmt"
	"net/http"
)
func welcome(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w,"Hello from Go!!")
}

func main() {

	http.HandleFunc("/", welcome)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}l