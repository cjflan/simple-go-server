package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
    p := flag.String("port", "8080", "port for the web server")
    flag.Parse()

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello")
    })

    port := fmt.Sprint(":", *p)

    http.ListenAndServe(port, nil)
}
