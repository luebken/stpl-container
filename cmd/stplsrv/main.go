package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	log "github.com/Sirupsen/logrus"

	"github.com/luebken/stpl/pkg/stpl/analysis"
	"github.com/luebken/stpl/pkg/stpl/maven"
	"github.com/luebken/stpl/pkg/stpl/npm"
	"github.com/luebken/stpl/pkg/stpl/stack"
)

// @APIVersion 0.0.1
// @APITitle Stpl API
// @APIDescription https://github.com/luebken/stpl
// @SubApi Analysis [/analysis]
// @SubApi ReferenceStacks [/referencestacks]
func main() {

	//init & prerequisites
	err := analysis.GetEnvRedisURL()
	if err != nil {
		log.Error("Couldn't get REDIS_URL.", err)
		os.Exit(-1)
	}
	err = analysis.GetLibrariesIoAPIKey()
	if err != nil {
		log.Error("Couldn't get LIBRARIES_IO_API_KEY.", err)
		os.Exit(-1)
	}

	keys, err := analysis.GetNumberOfCachedComponents()
	if err != nil {
		log.Panic("Couldn't get keys. Err: ", err)
	} else {
		log.Info("Connected to Redis. Number of cached components: ", keys)
	}

	stack.ImportReferenceStacks()

	//start server
	log.Print("stplsrv listening on :8088")
	http.HandleFunc("/", getHelp)
	http.HandleFunc("/analysis", getAnalysis)
	http.HandleFunc("/referencestacks", getReferenceStacks)
	http.HandleFunc("/admin/updatecomponentcache", getCachedComponents)
	http.ListenAndServe(":8088", nil)
}

func getAnalysis(w http.ResponseWriter, req *http.Request) {
	log.Printf("%s %s %s", req.RemoteAddr, req.Method, req.URL)

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Errorf("Error reading body: %v", err)
	}

	type result struct {
		Analysis analysis.Analysis
	}
	r := result{}

	// Maven
	projects, err := maven.Unmarshal(body)
	if err != nil {
		log.Infof("Couldn't parse Maven (pom.xml) %v", err)
	} else {
		log.Printf("Analysing %v %v", projects[0].GroupID, projects[0].ArtifactID)
		r = result{analysis.GetAnalysis(projects[0])}
	}

	// NodeJS
	pkg, err := npm.Unmarshal(body)
	log.Infof("NPM package (package.json) %v", pkg)
	r = result{analysis.GetAnalysis2(pkg)}

	// Result
	b, err := json.Marshal(r)
	if err != nil {
		log.Error(err)
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
	fmt.Fprintf(w, "%v\n", string(b))
}

func getCachedComponents(w http.ResponseWriter, req *http.Request) {
	log.Printf("%s %s %s", req.RemoteAddr, req.Method, req.URL)

	stacks := stack.AllReferenceStacks()

	numberOfCalls := 0
	for _, stack := range stacks {
		for _, dep := range stack.Dependencies {
			if numberOfCalls < 50 {
				c := analysis.GetComponentInfo("maven", dep.GroupID+":"+dep.ArtefactID)
				if c.UsedAPI {
					numberOfCalls++
				}
			}
		}
	}
	nrCachedComponents, err := analysis.GetNumberOfCachedComponents()
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Infof("Remote calls: %v", numberOfCalls)
	msg := fmt.Sprintf("Cached components: %v", nrCachedComponents)
	log.Info(msg)
	fmt.Fprintf(w, "%v\n", msg)
}

func getHelp(w http.ResponseWriter, req *http.Request) {
	s := `
<html>
<body>
Example and API docs: <a href="https://github.com/luebken/stpl/blob/master/docs">github.com/luebken/stpl/docs</a>
</body>
</html>
`
	fmt.Fprintf(w, s)
}
