package analysis

import (
	"github.com/luebken/stpl/pkg/stpl/maven"
)

type Analysis struct {
	Recommendation Recommendation
}

func GetAnalysis(project maven.Project) Analysis {
	r := GetRecommendation(project)
	return Analysis{r}
}
