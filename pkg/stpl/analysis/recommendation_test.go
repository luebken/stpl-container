package analysis

import (
	"log"
	"testing"

	"github.com/luebken/stpl/pkg/stpl/maven"
)

var p90 maven.Project
var p30 maven.Project

func init() {

	ImportReferenceStacks()

	p90 = maven.Project{}
	deps90 := []maven.MavenDependency{
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
	p90.Dependencies.Dependencies = append(p90.Dependencies.Dependencies, deps90...)

	p30 = maven.Project{}
	deps30 := []maven.MavenDependency{
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
	}
	p30.Dependencies.Dependencies = append(p30.Dependencies.Dependencies, deps30...)
}

func TestEmptyProject(t *testing.T) {
	p := maven.Project{}
	r := GetRecommendation(p)
	if len(r.RecommendationItems) != 0 {
		t.Error("Expected 0 RecommendationItems")
	}
}

func TestAddDependency(t *testing.T) {
	r := GetRecommendation(p90)
	if r.Similarity != 0.9 {
		t.Errorf("Expected a similarity of 0.9. Got %v\n", r.Similarity)
	}
	if len(r.RecommendationItems) != 1 {
		t.Error("Expected 1 RecommendationItems")
	}
	if r.RecommendationItems[0].InputDependency.ArtefactID != "" {
		t.Errorf("Expected valid InputDependency")
	}
}

func TestToLittleSimilarity(t *testing.T) {

	log.Printf("ping")

	r := GetRecommendation(p30)
	if len(r.RecommendationItems) != 0 {
		t.Error("Expected 0 RecommendationItems")
	}
}
