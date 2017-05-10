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
	"github.com/luebken/stpl/pkg/stpl/stack"
)

func main() {
	//init
	redisUrl := os.Getenv("REDIS_URL")
	if redisUrl == "" {
		log.Info("Didn't find env REDIS_URL. Using default 6379")
		redisUrl = "localhost:6379"
	}
	analysis.SetRedisURL(redisUrl)
	librariesIOKey := os.Getenv("LIBRARIES_IO_API_KEY")
	if librariesIOKey == "" {
		log.Info("Didn't find env LIBRARIES_IO_API_KEY.")
		os.Exit(-1)
	}
	analysis.API_KEY = librariesIOKey

	ping, err := analysis.TestRedis()
	if err != nil || !ping {
		log.Panic("Couldn't ping redis. Err: ", err)
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
		panic(err)
	}
	projects := maven.Unmarshal(body)

	type result struct {
		Analysis analysis.Analysis
	}
	log.Printf("Analysing %v %v", projects[0].GroupID, projects[0].ArtifactID)
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
	fmt.Fprintf(w, "%v\n", string(b))
}

func getCachedComponents(w http.ResponseWriter, req *http.Request) {
	log.Printf("%s %s %s", req.RemoteAddr, req.Method, req.URL)

	stacks := stack.AllReferenceStacks()

	numberOfCalls := 0
	for _, stack := range stacks {
		for _, dep := range stack.Dependencies {
			if numberOfCalls < 50 {
				c := analysis.GetComponentInfo("maven", dep.GroupID, dep.ArtefactID)
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
Available endpoints:\n 
* POST /analysis\n 
* GET /referencestacks\n
\n\n
More info at https://github.com/luebken/stpl`
	fmt.Fprintf(w, s)
}
