package analysis

import (
	"github.com/luebken/stpl/pkg/stpl/maven"
)

type Analysis struct {
	Recommendation        Recommendation
	LibrariesIoComponents []LibrariesIoComponent
}

// @Title getAnalysis
// @Description Get an analysis for a given pom.xml
// @Accept  xml
// @Body	pom.xml
// @Router /analysis/ [get]
func GetAnalysis(project maven.Project) Analysis {
	r := GetRecommendation(project)
	libs := []LibrariesIoComponent{}
	for _, dep := range project.Dependencies.Dependencies {
		group := dep.GroupID
		artefact := dep.ArtifactID
		c := GetComponentInfo("maven", group, artefact)
		libs = append(libs, c.LibrariesIoComponent)
	}
	return Analysis{r, libs}
}
