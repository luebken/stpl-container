package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/luebken/stpl/pkg/stpl/maven"
)

func analytics(w http.ResponseWriter, req *http.Request) {
	log.Printf("%s %s %s", req.RemoteAddr, req.Method, req.URL)

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		//TODO error handling
		panic(err)
	}
	log.Println(string(body))

	projects := maven.Unmarshal(body)

	deps := projects[0].Dependencies.Dependencies
	artefact := projects[0].ArtifactID
	fmt.Fprintf(w, "%v has %d dependencies.\n", artefact, len(deps))
}

func main() {
	println("stplsrv listening on :8080")
	http.HandleFunc("/analytics", analytics)
	http.ListenAndServe(":8080", nil)
}
