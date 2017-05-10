package npm

import (
	"encoding/json"

	log "github.com/Sirupsen/logrus"
)

type Package struct {
	Name         string            `json:"name"`
	Description  string            `json:"description"`
	Dependencies map[string]string `json:"dependencies"`
}

// Unmarshal pom.xmls with top level <projects> or <project>
func Unmarshal(data []byte) (Package, error) {
	var result Package

	err := json.Unmarshal(data, &result)
	if err != nil {
		log.Error(err)
		return Package{}, err
	}
	return result, nil
}
