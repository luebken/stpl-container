package stacks

import (
	"sort"

	"github.com/luebken/stpl/pkg/stpl/maven"
)

type Advices []Advice

func (slice Advices) Len() int {
	return len(slice)
}

func (slice Advices) Less(i, j int) bool {
	return slice[i].Similarity < slice[j].Similarity
}

func (slice Advices) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

type Advice struct {
	Similarity     float32
	ReferenceStack ReferenceStack
	AdviceItems    []AdviceItem
}

type AdviceItem struct {
	InputDependency       string
	RecommendedDependency string
}

func GetAdvice(project maven.Project) Advice {

	advices := []Advice{}

	for _, stack := range AllReferenceStacks() {
		similarity := calculateSimilarity(project, stack)
		if similarity > 0 {
			items := calculateAdviceItems(project, stack)
			advice := Advice{Similarity: similarity, ReferenceStack: stack, AdviceItems: items}
			advices = append(advices, advice)
		}
	}

	as := Advices(advices)
	sort.Sort(sort.Reverse(as))

	if len(as) > 0 {
		return advices[0]
	}
	return Advice{}
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
func calculateAdviceItems(project maven.Project, stack ReferenceStack) []AdviceItem {
	items := []AdviceItem{}

	for _, mavenDependecy := range project.Dependencies.Dependencies {
		mavenDependencyName := mavenDependecy.GroupID + ":" + mavenDependecy.ArtifactID
		contains, dep := stack.containsDependencyName(mavenDependencyName)

		// recommend switching version
		if contains {
			if mavenDependecy.Version != dep.Version {
				items = append(items, AdviceItem{
					InputDependency:       mavenDependencyName + ":" + mavenDependecy.Version,
					RecommendedDependency: dep.Name + ":" + dep.Version})
			}
		}
	}

	return items
}
