package main

import (
	"fmt"
	"net/http"

	"github.com/luebken/stpl/pkg/stpl/liba"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, liba.HelloWorld())
}
func main() {
	println("stplsrv listening on :8080")
	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":8080", nil)

}
