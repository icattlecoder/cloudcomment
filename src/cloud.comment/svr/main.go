package main

import (
	"fmt"
	"net/http"
)

func main() {
	run()
}
func run() {
	http.HandleFunc("/comment", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "lsdjflsjf")
	})
}
