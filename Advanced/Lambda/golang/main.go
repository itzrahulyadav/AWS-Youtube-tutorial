package main

import (
	"fmt"
	"net/http"
    "io"
)
func welcome(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w,"Hello from Go!!")
}

func main() {

	// response,err := http.Get("https://pokeapi.co/api/v2/pokemon/ditto")
	response,err := http.Get("https://official-joke-api.appspot.com/random_joke")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response.Status)


	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println("Response body:", string(body))
	// http.Handle("/", http.FileServer(http.Dir("./")))
	// http.HandleFunc("/welcome", welcome)
	// if err := http.ListenAndServe(":8080", nil); err != nil {
	// 	fmt.Println(err)
	// }
}