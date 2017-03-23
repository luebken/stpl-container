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

	stacks.ImportReferenceStacks()

	println("stplsrv listening on :8080")
	http.HandleFunc("/recommendation", getRecommendation)
	http.HandleFunc("/referencestacks", getReferenceStacks)
	http.ListenAndServe(":8080", nil)
}

func getRecommendation(w http.ResponseWriter, req *http.Request) {
	log.Printf("%s %s %s", req.RemoteAddr, req.Method, req.URL)

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	projects := maven.Unmarshal(body)

	advice := stacks.GetRecommendation(projects[0])
	b, err := json.Marshal(advice)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println()
	fmt.Fprintf(w, "%v\n", string(b))
}

func getReferenceStacks(w http.ResponseWriter, req *http.Request) {
	log.Printf("%s %s %s", req.RemoteAddr, req.Method, req.URL)

	stacks := stacks.AllReferenceStacks()
	b, err := json.Marshal(stacks)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println()
	fmt.Fprintf(w, "%v\n", string(b))
}
