package maven

import (
	"encoding/xml"
	"log"

	"github.com/blang/semver"
)

//Projects where a pom.xml has <projects> as top level
type Projects struct {
	Projects []Project `xml:"project"`
}

//Project where a pom.xml has <project> as top level
type Project struct {
	GroupID              string               `xml:"groupId"`
	ArtifactID           string               `xml:"artifactId"`
	DependencyManagement dependencyManagement `xml:"dependencyManagement"`
	Dependencies         dependencies         `xml:"dependencies"`
	Description          string               `xml:"description"`
}
type dependencyManagement struct {
	Dependencies dependencies `xml:"dependencies"`
}
type dependencies struct {
	Dependencies []dependency `xml:"dependency"`
}
type dependency struct {
	GroupID       string `xml:"groupId"`
	ArtifactID    string `xml:"artifactId"`
	VersionString string `xml:"version"`
	Scope         string `xml:"scope"`
	SemVer        semver.Version
}

// Unmarshal pom.xmls with top level <projects> or <project>
func Unmarshal(data []byte) []Project {
	var result []Project

	//try with <projects> as top level
	var ps Projects
	xml.Unmarshal(data, &ps)
	if len(ps.Projects) > 0 {
		for _, p := range ps.Projects {
			result = append(result, p)
		}

	} else { // try with <project> as top level
		var p Project
		xml.Unmarshal(data, &p)
		result = append(result, p)
	}
	return result
}

func ParseSemVers(projects []Project) {
	for _, project := range projects {
		for _, dependency := range project.Dependencies.Dependencies {
			version, err := semver.ParseTolerant(dependency.VersionString)
			if err != nil {
				log.Printf("Error semver parsing %v %v", dependency.VersionString, err)
			} else {
				dependency.SemVer = version
			}
		}
	}
}
