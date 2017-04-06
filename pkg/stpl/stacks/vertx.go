package stacks

var vertxYamls = []string{
	vertxstackbareweb,
	vertxstackbaremicroservice,
	vertxstackbaremicroservicegrpc,
	vertxstackbareopenshift,
	vertxstackfullweb,
	vertxstackfullwebopenshift,
}

var vertxstackbareweb = `
name: Vert.x Bare Web
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
name: Vert.x Bare Microservice
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
name: Vert.x Bare GRPC
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
name: Vert.x Bare OpenShift
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
name: Vert.x Full Web
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
name:  Vert.x Full OpenShift 
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
