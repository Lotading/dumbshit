package main

import (
	"fmt"
	"net/http"
)

func main() {
	file := http.FileServer(http.Dir("./"))
	http.Handle("/", file)
  fmt.Println("listening on http://localhost:8000/")
	http.ListenAndServe("8000", nil)
}
