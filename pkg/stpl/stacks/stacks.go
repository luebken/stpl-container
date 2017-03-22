package stacks

type Stack struct {
	Name         string
	Description  string
	Dependencies []dependency
}
type dependency struct {
	Name    string
	Version string
}

func (s Stack) containsDependencyName(name string) (bool, dependency) {
	for _, d := range s.Dependencies {
		if d.Name == name {
			return true, d
		}
	}
	return false, dependency{}
}

// ---

func AllStacks() []Stack {
	return stacks
}

var stacks = []Stack{vertxStackBareWeb, vertxStackBareMicroservice, vertxStackBareOpenShift}

var vertxStackBareWeb = Stack{
	Name:        "vertx-stack-bare-web",
	Description: "A bare stack to write web applications for Vert.x",
	Dependencies: []dependency{
		dependency{
			Name:    "io.vertx:vertx-core",
			Version: "3.4.1",
		},
		dependency{
			Name:    "io.vertx:vertx-jdbc-client",
			Version: "3.4.1",
		},
		dependency{
			Name:    "io.vertx:vertx-rx-java",
			Version: "3.4.1",
		},
		dependency{
			Name:    "io.vertx:vertx-web-client",
			Version: "3.4.1",
		},
		dependency{
			Name:    "io.vertx:vertx-web-templ-freemarker",
			Version: "3.4.1",
		},
		dependency{
			Name:    "io.vertx:vertx-web-templ-handlebars",
			Version: "3.4.1",
		},
		dependency{
			Name:    "io.vertx:io.vertx:vertx-web",
			Version: "3.4.1",
		},
	},
}

var vertxStackBareMicroservice = Stack{
	Name:        "vertx-stack-bare-microservice",
	Description: "A bare Vert.x stack for writing microservices.",
	Dependencies: []dependency{
		dependency{
			Name:    "io.vertx:vertx-circuit-breaker",
			Version: "3.4.1",
		},
		dependency{
			Name:    "io.vertx:vertx-config",
			Version: "3.4.1",
		},
		dependency{
			Name:    "io.vertx:vertx-core",
			Version: "3.4.1",
		},
		dependency{
			Name:    "vertx-rx-java",
			Version: "3.4.1",
		},
		dependency{
			Name:    "io.vertx:service-discovery",
			Version: "3.4.1",
		},
		dependency{
			Name:    "io.vertx:vertx-web-client",
			Version: "3.4.1",
		},
		dependency{
			Name:    "io.vertx:io.vertx:vertx-web",
			Version: "3.4.1",
		},
	},
}

var vertxStackBareOpenShift = Stack{
	Name:        "vertx-stack-bare-openshift",
	Description: "A bare Vert.x stack for writing OpenShift applications.",
	Dependencies: []dependency{
		dependency{
			Name:    "io.vertx:vertx-config",
			Version: "3.4.1",
		},
		dependency{
			Name:    "io.vertx:vertx-config-kubernetes-configmap",
			Version: "3.4.1",
		},
		dependency{
			Name:    "io.vertx:vertx-core",
			Version: "3.4.1",
		},
		dependency{
			Name:    "vertx-rx-java",
			Version: "3.4.1",
		},
		dependency{
			Name:    "io.vertx:service-discovery",
			Version: "3.4.1",
		},
		dependency{
			Name:    "io.vertx:vertx-service-discovery-bridge-kubernetes",
			Version: "3.4.1",
		},
		dependency{
			Name:    "io.vertx:vertx-web-client",
			Version: "3.4.1",
		},
		dependency{
			Name:    "io.vertx:io.vertx:vertx-web",
			Version: "3.4.1",
		},
	},
}
