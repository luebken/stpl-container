package analysis

import (
	"github.com/luebken/stpl/pkg/stpl/maven"
	"github.com/luebken/stpl/pkg/stpl/npm"
)

type Analysis struct {
	Recommendation        Recommendation
	LibrariesIoComponents []LibrariesIoComponent
}

// @Title getAnalysis
// @Description Get an analysis for a given pom.xml
// @Accept  xml
// @Produce  json
// @Body	pom.xml
// @Router /analysis/ [get]
func GetAnalysis(project maven.Project) Analysis {
	r := GetRecommendation(project)
	libs := []LibrariesIoComponent{}
	for _, dep := range project.Dependencies.Dependencies {
		group := dep.GroupID
		artefact := dep.ArtifactID
		c := GetComponentInfo("maven", group+":"+artefact)
		libs = append(libs, c.LibrariesIoComponent)
	}
	return Analysis{r, libs}
}

func GetAnalysis2(pkg npm.Package) Analysis {
	libs := []LibrariesIoComponent{}
	for key, _ := range pkg.Dependencies {
		c := GetComponentInfo("npm", key)
		libs = append(libs, c.LibrariesIoComponent)
	}
	return Analysis{Recommendation{}, libs}
}
