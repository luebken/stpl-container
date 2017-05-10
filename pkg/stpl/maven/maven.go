package maven

import (
	"encoding/xml"

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
	Name                 string               `xml:"name"`
	Scope                string               `xml:"scope"`
	Description          string               `xml:"description"`
	DependencyManagement dependencyManagement `xml:"dependencyManagement"`
	Dependencies         dependencies         `xml:"dependencies"`
}
type dependencyManagement struct {
	Dependencies dependencies `xml:"dependencies"`
}
type dependencies struct {
	Dependencies []MavenDependency `xml:"dependency"`
}
type MavenDependency struct {
	GroupID       string `xml:"groupId"`
	ArtifactID    string `xml:"artifactId"`
	VersionString string `xml:"version"`
	Scope         string `xml:"scope"`
	SemVer        semver.Version
}

func (p Project) ContainsDependency(groupid string, artifactid string) (bool, MavenDependency) {
	for _, d := range p.Dependencies.Dependencies {
		if d.GroupID == groupid && d.ArtifactID == artifactid {
			return true, d
		}
	}
	return false, MavenDependency{}
}

// Unmarshal pom.xmls with top level <projects> or <project>
func Unmarshal(data []byte) ([]Project, error) {
	var result []Project

	//try with <projects> as top level
	var ps Projects
	_ = xml.Unmarshal(data, &ps)
	if len(ps.Projects) > 0 {
		for _, p := range ps.Projects {
			result = append(result, p)
		}

	} else { // try with <project> as top level
		var p Project
		err := xml.Unmarshal(data, &p)
		if err != nil {
			return nil, err
		}
		result = append(result, p)
	}
	return result, nil
}
