package analysis

import (
	"github.com/luebken/stpl/pkg/stpl/maven"
)

type Analysis struct {
	Recommendation Recommendation
	ComponentInfo  ComponentInfo
}

func GetAnalysis(project maven.Project) Analysis {
	r := GetRecommendation(project)
	group := project.Dependencies.Dependencies[0].GroupID
	artefact := project.Dependencies.Dependencies[0].ArtifactID
	c := GetComponentInfo("Maven", group, artefact)
	return Analysis{r, c}
}
