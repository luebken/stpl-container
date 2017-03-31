package stacks

import (
	"log"
	"math"
	"sort"

	"fmt"

	"github.com/luebken/stpl/pkg/stpl/maven"
)

//Recommendations a list of possible recommendations per project. Sortable by similarity.
type Recommendations []Recommendation

func (slice Recommendations) Len() int { return len(slice) }

func (slice Recommendations) Less(i, j int) bool { return slice[i].Similarity < slice[j].Similarity }

func (slice Recommendations) Swap(i, j int) { slice[i], slice[j] = slice[j], slice[i] }

// Recommendation A single recommendation including the reference stack and indiviual recommendation items
type Recommendation struct {
	Similarity          float64
	ReferenceStack      ReferenceStack
	RecommendationItems []RecommendationItem
}

// RecommendationItem shows InputDependency, RecommendedDependency and a RecommendationText
type RecommendationItem struct {
	InputDependency       Dependency
	RecommendedDependency Dependency
	RecommendationText    string
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

	for _, referenceDep := range referenceStack.Dependencies {

		mavenContainsDependency, mavenDep := project.ContainsDependencyName(referenceDep.GroupID, referenceDep.ArtefactID)

		inputDependency, err := NewDependency(mavenDep.GroupID, mavenDep.ArtifactID, mavenDep.VersionString)
		if err != nil {
			// log.Printf("Error inputDependency artefact: '%v' %v", mavenDep.ArtifactID, err)
		}
		recommendedDependency, err := NewDependency(referenceDep.GroupID, referenceDep.ArtefactID, referenceDep.VersionString)
		if err != nil {
			log.Printf("Error recommendedDependency %v", err)
		}
		recommendedText := ""

		// recommend switching version
		if mavenContainsDependency {
			if mavenDep.VersionString != referenceDep.VersionString {
				// Upgrade
				if recommendedDependency.SemVer.GT(inputDependency.SemVer) {
					recommendedText = fmt.Sprintf("Recommended: Upgrade %v:%v to %v. %v applications include a newer version of %v.",
						mavenDep.GroupID,
						mavenDep.ArtifactID,
						referenceDep.VersionString,
						referenceStack.Name,
						mavenDep.ArtifactID,
					)
				}
				// Downgrade
				if recommendedDependency.SemVer.LT(inputDependency.SemVer) {
					recommendedText = fmt.Sprintf("Recommended: Downgrade %v:%v to %v. %v applications include an older version of %v.",
						mavenDep.GroupID,
						mavenDep.ArtifactID,
						referenceDep.VersionString,
						referenceStack.Name,
						mavenDep.ArtifactID,
					)
				}

			}
		} else { // recommend add dependency
			recommendedText = fmt.Sprintf("Recommended: Add %v. %v applications include %v.",
				referenceDep.VersionString,
				referenceStack.Name,
				referenceDep.ArtefactID,
			)
		}

		if recommendedText != "" {
			items = append(items, RecommendationItem{
				InputDependency:       inputDependency,
				RecommendedDependency: recommendedDependency,
				RecommendationText:    recommendedText,
			})
		}
	}

	return items
}
