package stacks

import (
	"testing"

	"github.com/luebken/stpl/pkg/stpl/maven"
)

func TestEmptyProject(t *testing.T) {
	p := maven.Project{}
	r := GetRecommendation(p)
	if len(r.RecommendationItems) != 0 {
		t.Error("Expected 0 RecommendationItems")
	}
}
