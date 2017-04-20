package analysis

import (
	"log"

	"github.com/luebken/stpl/pkg/stpl/maven"
)

type Analysis struct {
	Recommendation        Recommendation
	LibrariesIoComponents []LibrariesIoComponent
}

func GetAnalysis(project maven.Project) Analysis {
	r := GetRecommendation(project)
	libs := []LibrariesIoComponent{}
	for _, dep := range project.Dependencies.Dependencies {
		group := dep.GroupID
		artefact := dep.ArtifactID
		c := GetComponentInfo("maven", group, artefact)
		log.Printf("Print component name %v", c.LibrariesIoComponent.Name)
		log.Printf("Print component RepositoryUrl %v", c.LibrariesIoComponent.RepositoryURL)
		log.Printf("Print component Status %v", c.LibrariesIoComponent.Status)
		libs = append(libs, c.LibrariesIoComponent)
	}
	return Analysis{r, libs}
}
