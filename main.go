package main

import (
	"fmt"
	"net/http"

	"github.com/takama/router"
)

// Run server: go run main.go
// Try few requests:
// - curl http://localhost:8000
// - curl http://localhost:8000/test-some-path
func main() {
	r := router.New()
	http.HandleFunc("/", handler)
	http.ListenAndServe("localhost:8000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
