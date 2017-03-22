package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/luebken/stpl/pkg/stpl/maven"
	"github.com/luebken/stpl/pkg/stpl/stacks"
)

func main() {
	println("stplsrv listening on :8080")
	http.HandleFunc("/analytics", analytics)
	http.HandleFunc("/stacks", getStacks)
	http.ListenAndServe(":8080", nil)
}

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

func getStacks(w http.ResponseWriter, req *http.Request) {
	log.Printf("%s %s %s", req.RemoteAddr, req.Method, req.URL)

	stacks := stacks.Stacks()
	b, err := json.Marshal(stacks)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println()
	fmt.Fprintf(w, "%v\n", string(b))
}
