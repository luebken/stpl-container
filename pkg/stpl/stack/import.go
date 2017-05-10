package stack

import (
	"fmt"
	"regexp"

	log "github.com/Sirupsen/logrus"
	"github.com/blang/semver"
	yaml "gopkg.in/yaml.v2"
)

var stacks = []ReferenceStack{}

type ReferenceStack struct {
	Name            string       `yaml:"name"`
	Description     string       `yaml:"description"`
	DependenciesRaw []string     `yaml:"dependencies" json:"-"`
	Dependencies    []Dependency `yaml:"-"`
}

type Dependency struct {
	GroupID       string
	ArtefactID    string
	VersionString string
	SemVer        semver.Version
}

func NewDependency(groupid string, artefactid string, versionString string) (Dependency, error) {
	version, err := semver.ParseTolerant(versionString)
	if err != nil {
		err2 := fmt.Errorf("Error semver parsing %v %v", versionString, err)
		return Dependency{}, err2
	}

	d2 := Dependency{GroupID: groupid, ArtefactID: artefactid, VersionString: versionString, SemVer: version}
	return d2, nil
}

func (s ReferenceStack) ContainsDependencyName(groupid string, artefactid string) (bool, Dependency) {
	for _, d := range s.Dependencies {
		if d.GroupID == groupid && d.ArtefactID == artefactid {
			return true, d
		}
	}
	return false, Dependency{}
}

// @Title getReferenceStacks
// @Description Get all reference stacks known to the system
// @Produce	json
// @Success 200 {array} ReferenceStack
// @Router /referencestacks/ [get]
func AllReferenceStacks() []ReferenceStack {
	return stacks
}

func ImportReferenceStacks() {

	yamls := append(vertxYamls, springBootYamls...)
	for _, y := range yamls {
		s := ReferenceStack{}
		err := yaml.Unmarshal([]byte(y), &s)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		for _, depRaw := range s.DependenciesRaw {
			//re, _ := regexp.Compile("([\\w\\.\\-]*):([\\w\\.\\-]*):([\\w\\.\\-]*)")
			re, _ := regexp.Compile("([\\w\\.\\-]*):([\\w\\.\\-]*):([\\w*\\.\\w*\\.\\w]).*")
			result := re.FindAllStringSubmatch(depRaw, -1)
			versionString := result[0][3]
			d2, err := NewDependency(result[0][1], result[0][2], versionString)
			if err != nil {
				log.Printf("Error importing %v", err)
			} else {
				s.Dependencies = append(s.Dependencies, d2)
			}
		}
		stacks = append(stacks, s)
	}

	log.Printf("Imported %v reference stacks.", len(stacks))
}
