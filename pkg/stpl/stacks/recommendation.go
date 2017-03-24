package stacks

import (
	"math"
	"sort"

	"github.com/luebken/stpl/pkg/stpl/maven"
)

type Recommendations []Recommendation

func (slice Recommendations) Len() int { return len(slice) }

func (slice Recommendations) Less(i, j int) bool { return slice[i].Similarity < slice[j].Similarity }

func (slice Recommendations) Swap(i, j int) { slice[i], slice[j] = slice[j], slice[i] }

type Recommendation struct {
	Similarity          float64
	ReferenceStack      ReferenceStack
	RecommendationItems []RecommendationItem
}

type RecommendationItem struct {
	InputDependency       Dependency
	RecommendedDependency Dependency
}

func GetRecommendation(project maven.Project) Recommendation {

	advices := []Recommendation{}

	for _, stack := range AllReferenceStacks() {
		similarity := calculateSimilarity(project, stack)
		if similarity > 0 {
			items := calculateRecommendationItems(project, stack)
			advice := Recommendation{
				Similarity:          similarity,
				ReferenceStack:      stack,
				RecommendationItems: items,
			}
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

func calculateSimilarity(input maven.Project, reference ReferenceStack) float64 {
	lenInputDeps := float64(len(input.Dependencies.Dependencies))
	lenReferenceDeps := float64(len(reference.Dependencies))
	maxLenDependencies := math.Max(lenInputDeps, lenReferenceDeps)

	intersect := []Dependency{}

	for _, mavenDependecy := range input.Dependencies.Dependencies {
		contains, _ := reference.containsDependencyName(mavenDependecy.GroupID, mavenDependecy.ArtifactID)

		if contains {
			intersect = append(intersect, Dependency{GroupID: mavenDependecy.GroupID, ArtefactID: mavenDependecy.ArtifactID})
		}
	}

	nrOfSameDependencies := float64(len(intersect))
	return nrOfSameDependencies / maxLenDependencies
}

func calculateRecommendationItems(project maven.Project, referenceStack ReferenceStack) []RecommendationItem {
	items := []RecommendationItem{}

	for _, mavenDependecy := range project.Dependencies.Dependencies {
		contains, dep := referenceStack.containsDependencyName(mavenDependecy.GroupID, mavenDependecy.ArtifactID)

		// recommend switching version
		if contains {
			if mavenDependecy.Version != dep.Version {
				items = append(items, RecommendationItem{
					InputDependency:       Dependency{GroupID: mavenDependecy.GroupID, ArtefactID: mavenDependecy.ArtifactID, Version: mavenDependecy.Version},
					RecommendedDependency: Dependency{GroupID: dep.GroupID, ArtefactID: dep.ArtefactID, Version: dep.Version},
				})
			}
		}
	}

	return items
}
