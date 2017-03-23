package stacks

import (
	"log"
	"regexp"

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
	GroupID    string
	ArtefactID string
	Version    string
}

func (s ReferenceStack) containsDependencyName(groupid string, artefactid string) (bool, Dependency) {
	for _, d := range s.Dependencies {
		if d.GroupID == groupid && d.ArtefactID == artefactid {
			return true, d
		}
	}
	return false, Dependency{}
}

func AllReferenceStacks() []ReferenceStack {
	return stacks
}

func ImportReferenceStacks() {
	for _, y := range yamls {
		s := ReferenceStack{}
		err := yaml.Unmarshal([]byte(y), &s)
		if err != nil {
			log.Fatalf("error: %v", err)
		}
		for _, d := range s.DependenciesRaw {
			re, _ := regexp.Compile("([\\w\\.\\-]*):([\\w\\.\\-]*):([\\w\\.\\-]*)")
			result := re.FindAllStringSubmatch(d, -1)
			d2 := Dependency{GroupID: result[0][1], ArtefactID: result[0][2], Version: result[0][3]}
			s.Dependencies = append(s.Dependencies, d2)
		}
		stacks = append(stacks, s)
	}

	log.Printf("Imported %v reference stacks.", len(stacks))
}

var yamls = []string{
	vertxstackbareweb,
	vertxstackbaremicroservice,
	vertxstackbaremicroservicegrpc,
	vertxstackbareopenshift,
	vertxstackfullweb,
	vertxstackfullwebopenshift,
}

var vertxstackbareweb = `
name: vertx-stack-bare-web
description: A bare Vert.x stack for writing web applications.
dependencies: 
- io.vertx:vertx-core:3.4.1
- io.vertx:vertx-jdbc-client:3.4.1
- io.vertx:vertx-rx-java:3.4.1
- io.vertx:vertx-web-client:3.4.1
- io.vertx:vertx-web-templ-freemarker:3.4.1
- io.vertx:vertx-web-templ-handlebars:3.4.1
- io.vertx:vertx-web:3.4.1
`
var vertxstackbaremicroservice = `
name: vertx-stack-bare-microservice
description: A bare Vert.x stack for writing Rest based microservices.
dependencies: 
- io.vertx:vertx-circuit-breaker:3.4.1
- io.vertx:vertx-config:3.4.1
- io.vertx:vertx-core:3.4.1
- io.vertx:vertx-health-check:3.4.1
- io.vertx:vertx-rx-java:3.4.1
- io.vertx:vertx-service-discovery:3.4.1
- io.vertx:vertx-web-client:3.4.1
- io.vertx:vertx-web:3.4.1
`

var vertxstackbaremicroservicegrpc = `
name: vertx-stack-bare-microservice-grpc
description: A bare Vert.x stack for writing gRPC based microservices.
dependencies: 
- io.vertx:vertx-circuit-breaker:3.4.1
- io.vertx:vertx-config:3.4.1
- io.vertx:vertx-core:3.4.1
- io.vertx:vertx-grpc:3.4.1
- io.vertx:vertx-health-check:3.4.1
- io.vertx:vertx-rx-java:3.4.1
- io.vertx:vertx-service-discovery:3.4.1
`

var vertxstackbareopenshift = `
name: vertx-stack-bare-openshift
description: A bare Vert.x stack for writing OpenShift applications.
dependencies: 
- io.vertx:vertx-config-kubernetes-configmap:3.4.1
- io.vertx:vertx-config-yaml:3.4.1
- io.vertx:vertx-config:3.4.1
- io.vertx:vertx-core:3.4.1
- io.vertx:vertx-health-check:3.4.1
- io.vertx:vertx-rx-java:3.4.1
- io.vertx:vertx-service-discovery-bridge-kubernetes:3.4.1
- io.vertx:vertx-service-discovery:3.4.1
- io.vertx:vertx-web-client:3.4.1
- io.vertx:vertx-web:3.4.1
`

var vertxstackfullweb = `
name: vertx-stack-full-web
description: A full Vert.x stack for writing production ready web application.
dependencies:
- ch.qos.logback:logback-classic:1.2.1
- ch.qos.logback:logback-core:1.2.1
- com.google.guava:guava:20.0
- commons-io:commons-io:2.5
- io.vertx:vertx-core:3.4.1
- io.vertx:vertx-jdbc-client:3.4.1
- io.vertx:vertx-rx-java:3.4.1
- io.vertx:vertx-web-client:3.4.1
- io.vertx:vertx-web-templ-freemarker:3.4.1
- io.vertx:vertx-web-templ-handlebars:3.4.1
- io.vertx:vertx-web:3.4.1
- javax.validation:validation-api:1.1.0.Final
- joda-time:joda-time:2.9.7
- org.apache.commons:commons-lang3:3.4
- org.hibernate:hibernate-validator:5.4.0.Final
- org.jooq:jooq:3.6.0
- org.liquibase:liquibase-core:3.5.3
- org.postgresql:postgresql:9.4.1212
- org.slf4j:slf4j-api:1.7.22
`

var vertxstackfullwebopenshift = `
name:  vertx-stack-full-web-openshift 
description: A full Vert.x stack for writing production ready web application on OpenShift
dependencies: 
- ch.qos.logback:logback-classic:1.2.1
- ch.qos.logback:logback-core:1.2.1
- com.google.guava:guava:20.0
- commons-io:commons-io:2.5
- io.vertx:vertx-config-kubernetes-configmap:3.4.1
- io.vertx:vertx-config-yaml:3.4.1
- io.vertx:vertx-config:3.4.1
- io.vertx:vertx-core:3.4.1
- io.vertx:vertx-health-check:3.4.1
- io.vertx:vertx-jdbc-client:3.4.1
- io.vertx:vertx-rx-java:3.4.1
- io.vertx:vertx-service-discovery-bridge-kubernetes:3.4.1
- io.vertx:vertx-service-discovery:3.4.1
- io.vertx:vertx-web-client:3.4.1
- io.vertx:vertx-web-templ-freemarker:3.4.1
- io.vertx:vertx-web-templ-handlebars:3.4.1
- io.vertx:vertx-web:3.4.1
- javax.validation:validation-api:1.1.0.Final
- joda-time:joda-time:2.9.7
- org.apache.commons:commons-lang3:3.4
- org.hibernate:hibernate-validator:5.4.0.Final
- org.jooq:jooq:3.6.0
- org.liquibase:liquibase-core:3.5.3
- org.postgresql:postgresql:9.4.1212
- org.slf4j:slf4j-api:1.7.22
`
