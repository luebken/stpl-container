package stacks

import (
	"sort"

	"github.com/luebken/stpl/pkg/stpl/maven"
)

type Recommendations []Recommendation

func (slice Recommendations) Len() int { return len(slice) }

func (slice Recommendations) Less(i, j int) bool { return slice[i].Similarity < slice[j].Similarity }

func (slice Recommendations) Swap(i, j int) { slice[i], slice[j] = slice[j], slice[i] }

type Recommendation struct {
	Similarity          float32
	ReferenceStack      ReferenceStack
	RecommendationItems []RecommendationItem
}

type RecommendationItem struct {
	InputDependency       string
	RecommendedDependency string
}

func GetRecommendation(project maven.Project) Recommendation {

	advices := []Recommendation{}

	for _, stack := range AllReferenceStacks() {
		similarity := calculateSimilarity(project, stack)
		if similarity > 0 {
			items := calculateRecommendationItems(project, stack)
			advice := Recommendation{Similarity: similarity, ReferenceStack: stack, RecommendationItems: items}
			advices = append(advices, advice)
		}
	}

	as := Recommendations(advices)
	sort.Sort(sort.Reverse(as))

	if len(as) > 0 {
		return advices[0]
	}
	return Recommendation{}
}

func calculateSimilarity(project maven.Project, stack ReferenceStack) float32 {
	nrOfSameDependencies := 0
	for _, mavenDependecy := range project.Dependencies.Dependencies {
		contains, _ := stack.containsDependencyName(mavenDependecy.GroupID + ":" + mavenDependecy.ArtifactID)
		if contains {
			nrOfSameDependencies++
		}
	}

	return float32(nrOfSameDependencies) / float32(len(project.Dependencies.Dependencies))
}

func calculateRecommendationItems(project maven.Project, stack ReferenceStack) []RecommendationItem {
	items := []RecommendationItem{}

	for _, mavenDependecy := range project.Dependencies.Dependencies {
		mavenDependencyName := mavenDependecy.GroupID + ":" + mavenDependecy.ArtifactID
		contains, dep := stack.containsDependencyName(mavenDependencyName)

		// recommend switching version
		if contains {
			if mavenDependecy.Version != dep.Version {
				items = append(items, RecommendationItem{
					InputDependency:       mavenDependencyName + ":" + mavenDependecy.Version,
					RecommendedDependency: dep.Name + ":" + dep.Version,
				})
			}
		}
	}

	return items
}
