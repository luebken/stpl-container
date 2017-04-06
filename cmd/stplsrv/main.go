package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/Sirupsen/logrus"

	"github.com/luebken/stpl/pkg/stpl/analysis"
	"github.com/luebken/stpl/pkg/stpl/maven"
	"github.com/luebken/stpl/pkg/stpl/stack"
)

func main() {
	stack.ImportReferenceStacks()

	log.Print("stplsrv listening on :8088")
	http.HandleFunc("/", getHelp)
	http.HandleFunc("/analysis", getAnalysis)
	http.HandleFunc("/referencestacks", getReferenceStacks)
	http.ListenAndServe(":8088", nil)
}

func getAnalysis(w http.ResponseWriter, req *http.Request) {
	log.Printf("%s %s %s", req.RemoteAddr, req.Method, req.URL)

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	projects := maven.Unmarshal(body)

	type result struct {
		Analysis analysis.Analysis
	}
	r := result{analysis.GetAnalysis(projects[0])}
	b, err := json.Marshal(r)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprintf(w, "%v\n", string(b))
}

func getReferenceStacks(w http.ResponseWriter, req *http.Request) {
	log.Printf("%s %s %s", req.RemoteAddr, req.Method, req.URL)

	stacks := stack.AllReferenceStacks()
	b, err := json.Marshal(stacks)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println()
	fmt.Fprintf(w, "%v\n", string(b))
}

func getHelp(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Available endpoints:\n /analysis\n /referencestacks\n")

}
