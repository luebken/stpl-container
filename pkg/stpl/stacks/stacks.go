package stacks

type ReferenceStack struct {
	Name         string
	Description  string
	Dependencies []Dependency
}
type Dependency struct {
	Name    string
	Version string
}

func (s ReferenceStack) containsDependencyName(name string) (bool, Dependency) {
	for _, d := range s.Dependencies {
		if d.Name == name {
			return true, d
		}
	}
	return false, Dependency{}
}

// ---

func AllReferenceStacks() []ReferenceStack {
	return stacks
}

var stacks = []ReferenceStack{vertxStackBareWeb, vertxStackBareMicroservice, vertxStackBareOpenShift}

var vertxStackBareWeb = ReferenceStack{
	Name:        "vertx-stack-bare-web",
	Description: "A bare stack to write web applications for Vert.x",
	Dependencies: []Dependency{
		Dependency{
			Name:    "io.vertx:vertx-core",
			Version: "3.4.1",
		},
		Dependency{
			Name:    "io.vertx:vertx-jdbc-client",
			Version: "3.4.1",
		},
		Dependency{
			Name:    "io.vertx:vertx-rx-java",
			Version: "3.4.1",
		},
		Dependency{
			Name:    "io.vertx:vertx-web-client",
			Version: "3.4.1",
		},
		Dependency{
			Name:    "io.vertx:vertx-web-templ-freemarker",
			Version: "3.4.1",
		},
		Dependency{
			Name:    "io.vertx:vertx-web-templ-handlebars",
			Version: "3.4.1",
		},
		Dependency{
			Name:    "io.vertx:io.vertx:vertx-web",
			Version: "3.4.1",
		},
	},
}

var vertxStackBareMicroservice = ReferenceStack{
	Name:        "vertx-stack-bare-microservice",
	Description: "A bare Vert.x stack for writing microservices.",
	Dependencies: []Dependency{
		Dependency{
			Name:    "io.vertx:vertx-circuit-breaker",
			Version: "3.4.1",
		},
		Dependency{
			Name:    "io.vertx:vertx-config",
			Version: "3.4.1",
		},
		Dependency{
			Name:    "io.vertx:vertx-core",
			Version: "3.4.1",
		},
		Dependency{
			Name:    "vertx-rx-java",
			Version: "3.4.1",
		},
		Dependency{
			Name:    "io.vertx:service-discovery",
			Version: "3.4.1",
		},
		Dependency{
			Name:    "io.vertx:vertx-web-client",
			Version: "3.4.1",
		},
		Dependency{
			Name:    "io.vertx:io.vertx:vertx-web",
			Version: "3.4.1",
		},
	},
}

var vertxStackBareOpenShift = ReferenceStack{
	Name:        "vertx-stack-bare-openshift",
	Description: "A bare Vert.x stack for writing OpenShift applications.",
	Dependencies: []Dependency{
		Dependency{
			Name:    "io.vertx:vertx-config",
			Version: "3.4.1",
		},
		Dependency{
			Name:    "io.vertx:vertx-config-kubernetes-configmap",
			Version: "3.4.1",
		},
		Dependency{
			Name:    "io.vertx:vertx-core",
			Version: "3.4.1",
		},
		Dependency{
			Name:    "vertx-rx-java",
			Version: "3.4.1",
		},
		Dependency{
			Name:    "io.vertx:service-discovery",
			Version: "3.4.1",
		},
		Dependency{
			Name:    "io.vertx:vertx-service-discovery-bridge-kubernetes",
			Version: "3.4.1",
		},
		Dependency{
			Name:    "io.vertx:vertx-web-client",
			Version: "3.4.1",
		},
		Dependency{
			Name:    "io.vertx:io.vertx:vertx-web",
			Version: "3.4.1",
		},
	},
}
