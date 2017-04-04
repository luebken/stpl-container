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

func TestDowngradeDepdency(t *testing.T) {

	ImportReferenceStacks()

	p := maven.Project{}
	deps := []maven.MavenDependency{
		maven.MavenDependency{
			GroupID:       "io.vertx",
			ArtifactID:    "vertx-config-kubernetes-configmap",
			VersionString: "3.4.1",
		},
		maven.MavenDependency{
			GroupID:       "io.vertx",
			ArtifactID:    "vertx-config-yaml",
			VersionString: "3.4.1",
		},
		maven.MavenDependency{
			GroupID:       "io.vertx",
			ArtifactID:    "vertx-config",
			VersionString: "3.4.1",
		},
		maven.MavenDependency{
			GroupID:       "io.vertx",
			ArtifactID:    "vertx-core",
			VersionString: "3.4.1",
		},
		maven.MavenDependency{
			GroupID:       "io.vertx",
			ArtifactID:    "vertx-rx-java",
			VersionString: "3.4.1",
		},
		maven.MavenDependency{
			GroupID:       "io.vertx",
			ArtifactID:    "vertx-service-discovery-bridge-kubernetes",
			VersionString: "3.4.1",
		},
		maven.MavenDependency{
			GroupID:       "io.vertx",
			ArtifactID:    "vertx-service-discovery",
			VersionString: "3.4.1",
		},
		maven.MavenDependency{
			GroupID:       "io.vertx",
			ArtifactID:    "vertx-web-client",
			VersionString: "3.4.1",
		},
		maven.MavenDependency{
			GroupID:       "io.vertx",
			ArtifactID:    "vertx-web",
			VersionString: "3.4.1",
		},
	}
	p.Dependencies.Dependencies = append(p.Dependencies.Dependencies, deps...)

	r := GetRecommendation(p)
	if r.Similarity != 0.9 {
		t.Errorf("Expected a similarity of 0.9. Got %v\n", r.Similarity)
	}
	if len(r.RecommendationItems) != 1 {
		t.Error("Expected 1 RecommendationItems")
	}
}
