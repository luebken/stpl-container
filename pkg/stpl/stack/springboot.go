package stack

// generated
// git clone https://github.com/spring-projects/spring-boot
// git checkout v1.5.2.RELEASE
// cd springboot-starter
// mvn help:effective pom
// for f in *; do go run cmd/stplcli/main.go pom2referencestack $f; done

var springBootYamls = []string{
	springBootStarter,
	springBootActiveMQStarter,
	springBootActuatorStarter,
	springBootAMQPStarter,
	springBootAOPStarter,
	springBootArtemisStarter,
	springBootBatchStarter,
	springBootCacheStarter,
	springBootSpringCloudConnectorsStarter,
	springBootDataCassandraStarter,
	springBootDataCouchbaseStarter,
	springBootDataElasticsearchStarter,
	springBootDataJPAStarter,
	springBootDataMongoDBStarter,
	springBootDataMongoDBReactiveStarter,
	springBootDataNeo4jStarter,
	springBootDataRedisStarter,
	springBootDataRESTStarter,
	springBootDataSolrStarter,
	springBootFreeMarkerStarter,
	springBootGroovyTemplatesStarter,
	springBootHATEOASStarter,
	springBootIntegrationStarter,
	springBootJDBCStarter,
	springBootJerseyStarter,
	springBootJettyStarter,
	springBootJOOQStarter,
	springBootAtomikosJTAStarter,
	springBootBitronixJTAStarter,
	springBootNarayanaJTAStarter,
	springBootLog4j2Starter,
	springBootLoggingStarter,
	springBootMailStarter,
	springBootMobileStarter,
	springBootMustacheStarter,
	springBootSecurityStarter,
	springBootSocialFacebookStarter,
	springBootSocialLinkedInStarter,
	springBootSocialTwitterStarter,
	springBootTestStarter,
	springBootThymeleafStarter,
	springBootTomcatStarter,
	springBootUndertowStarter,
	springBootValidationStarter,
	springBootWebStarter,
	springBootWebServicesStarter,
	springBootWebFluxStarter,
	springBootWebSocketStarter,
}

var springBootStarter = `
name: Spring Boot Starter
description: "Core starter, including auto-configuration support, logging and YAML"
dependencies:
- org.springframework.boot:spring-boot:1.5.2.RELEASE
- org.springframework.boot:spring-boot-autoconfigure:1.5.2.RELEASE
- org.springframework.boot:spring-boot-starter-logging:1.5.2.RELEASE
- org.springframework:spring-core:4.3.7.RELEASE
`

var springBootActiveMQStarter = `
name: Spring Boot ActiveMQ Starter
description: "Starter for JMS messaging using Apache ActiveMQ"
dependencies:
- org.springframework.boot:spring-boot-starter:1.5.2.RELEASE
- org.springframework:spring-jms:4.3.7.RELEASE
- org.apache.activemq:activemq-broker:5.14.3
`

var springBootActuatorStarter = `
name: Spring Boot Actuator Starter
description: "Starter for using Spring Boot's Actuator which provides production ready features to help you monitor and manage your application"
dependencies:
- org.springframework.boot:spring-boot-starter:1.5.2.RELEASE
- org.springframework.boot:spring-boot-actuator:1.5.2.RELEASE
`

var springBootAMQPStarter = `
name: Spring Boot AMQP Starter
description: "Starter for using Spring AMQP and Rabbit MQ"
dependencies:
- org.springframework.boot:spring-boot-starter:1.5.2.RELEASE
- org.springframework:spring-messaging:4.3.7.RELEASE
- org.springframework.amqp:spring-rabbit:1.7.1.RELEASE
`

var springBootAOPStarter = `
name: Spring Boot AOP Starter
description: "Starter for aspect-oriented programming with Spring AOP and AspectJ"
dependencies:
- org.springframework.boot:spring-boot-starter:1.5.2.RELEASE
- org.springframework:spring-aop:4.3.7.RELEASE
- org.aspectj:aspectjweaver:1.8.9
`

var springBootArtemisStarter = `
name: Spring Boot Artemis Starter
description: "Starter for JMS messaging using Apache Artemis"
dependencies:
- org.springframework.boot:spring-boot-starter:1.5.2.RELEASE
- org.springframework:spring-jms:4.3.7.RELEASE
- org.apache.activemq:artemis-jms-client:1.5.3
`

var springBootBatchStarter = `
name: Spring Boot Batch Starter
description: "Starter for using Spring Batch"
dependencies:
- org.springframework.boot:spring-boot-starter:1.5.2.RELEASE
- org.springframework.boot:spring-boot-starter-jdbc:1.5.2.RELEASE
- org.springframework.batch:spring-batch-core:3.0.7.RELEASE
`

var springBootCacheStarter = `
name: Spring Boot Cache Starter
description: "Starter for using Spring Framework's caching support"
dependencies:
- org.springframework.boot:spring-boot-starter:1.5.2.RELEASE
- org.springframework:spring-context:4.3.7.RELEASE
- org.springframework:spring-context-support:4.3.7.RELEASE
`

var springBootSpringCloudConnectorsStarter = `
name: Spring Boot Spring Cloud Connectors Starter
description: "Starter for using Spring Cloud Connectors which simplifies connecting to services in cloud platforms like Cloud Foundry and Heroku"
dependencies:
- org.springframework.boot:spring-boot-starter:1.5.2.RELEASE
- org.springframework.cloud:spring-cloud-spring-service-connector:1.2.3.RELEASE
- org.springframework.cloud:spring-cloud-cloudfoundry-connector:1.2.3.RELEASE
- org.springframework.cloud:spring-cloud-heroku-connector:1.2.3.RELEASE
- org.springframework.cloud:spring-cloud-localconfig-connector:1.2.3.RELEASE
`

var springBootDataCassandraStarter = `
name: Spring Boot Data Cassandra Starter
description: "Starter for using Cassandra distributed database and Spring Data Cassandra"
dependencies:
- org.springframework.boot:spring-boot-starter:1.5.2.RELEASE
- org.springframework:spring-tx:4.3.7.RELEASE
- org.springframework.data:spring-data-cassandra:1.5.1.RELEASE
`

var springBootDataCouchbaseStarter = `
name: Spring Boot Data Couchbase Starter
description: "Starter for using Couchbase document-oriented database and Spring Data Couchbase"
dependencies:
- org.springframework.boot:spring-boot-starter:1.5.2.RELEASE
- org.springframework.data:spring-data-couchbase:2.2.1.RELEASE
`

var springBootDataElasticsearchStarter = `
name: Spring Boot Data Elasticsearch Starter
description: "Starter for using Elasticsearch search and analytics engine and Spring Data Elasticsearch"
dependencies:
- org.springframework.boot:spring-boot-starter:1.5.2.RELEASE
- org.springframework.data:spring-data-elasticsearch:2.1.1.RELEASE
`

var springBootDataGemFireStarter = `
name: Spring Boot Data GemFire Starter
description: "Starter for using GemFire distributed data store and Spring Data GemFire"
dependencies:
- org.springframework.boot:spring-boot-starter:1.5.2.RELEASE
- org.springframework.data:spring-data-gemfire:1.9.1.RELEASE
`

var springBootDataJPAStarter = `
name: Spring Boot Data JPA Starter
description: "Starter for using Spring Data JPA with Hibernate"
dependencies:
- org.springframework.boot:spring-boot-starter:1.5.2.RELEASE
- org.springframework.boot:spring-boot-starter-aop:1.5.2.RELEASE
- org.springframework.boot:spring-boot-starter-jdbc:1.5.2.RELEASE
- org.hibernate:hibernate-core:5.0.12.Final
- org.hibernate:hibernate-entitymanager:5.0.12.Final
- javax.transaction:javax.transaction-api:1.2
- org.springframework.data:spring-data-jpa:1.11.1.RELEASE
- org.springframework:spring-aspects:4.3.7.RELEASE
`

var springBootDataLDAPStarter = `
name: Spring Boot Data LDAP Starter
description: "Starter for using Spring Data LDAP"
dependencies:
- org.springframework.boot:spring-boot-starter:1.5.2.RELEASE
- org.springframework.data:spring-data-ldap:1.0.1.RELEASE
`

var springBootDataMongoDBStarter = `
name: Spring Boot Data MongoDB Starter
description: "Starter for using MongoDB document-oriented database and Spring Data MongoDB"
dependencies:
- org.springframework.boot:spring-boot-starter:1.5.2.RELEASE
- org.mongodb:mongodb-driver:3.4.2
- org.springframework.data:spring-data-mongodb:1.10.1.RELEASE
`

var springBootDataMongoDBReactiveStarter = `
name: Spring Boot Data MongoDB Reactive Starter
description: "Starter for using MongoDB document-oriented database and Spring Data MongoDB Reactive"
dependencies:
- org.springframework.boot:spring-boot-starter:2.0.0.BUILD-SNAPSHOT
- org.springframework.data:spring-data-mongodb:2.0.0.BUILD-SNAPSHOT
- org.mongodb:mongodb-driver:3.4.2
- org.mongodb:mongodb-driver-async:3.4.2
- org.mongodb:mongodb-driver-reactivestreams:1.3.0
- io.projectreactor:reactor-core:3.0.6.RELEASE
`

var springBootDataNeo4jStarter = `
name: Spring Boot Data Neo4j Starter
description: "Starter for using Neo4j graph database and Spring Data Neo4j"
dependencies:
- org.springframework.boot:spring-boot-starter:1.5.2.RELEASE
- org.springframework.data:spring-data-neo4j:4.2.1.RELEASE
`

var springBootDataRedisStarter = `
name: Spring Boot Data Redis Starter
description: "Starter for using Redis key-value data store with Spring Data Redis and the Jedis client"
dependencies:
- org.springframework.boot:spring-boot-starter:1.5.2.RELEASE
- org.springframework.data:spring-data-redis:1.8.1.RELEASE
- redis.clients:jedis:2.9.0
`

var springBootDataRESTStarter = `
name: Spring Boot Data REST Starter
description: "Starter for exposing Spring Data repositories over REST using Spring Data REST"
dependencies:
- org.springframework.boot:spring-boot-starter:1.5.2.RELEASE
- org.springframework.boot:spring-boot-starter-web:1.5.2.RELEASE
- com.fasterxml.jackson.core:jackson-annotations:2.8.0
- com.fasterxml.jackson.core:jackson-databind:2.8.7
- org.springframework.data:spring-data-rest-webmvc:2.6.1.RELEASE
`

var springBootDataSolrStarter = `
name: Spring Boot Data Solr Starter
description: "Starter for using the Apache Solr search platform with Spring Data Solr"
dependencies:
- org.springframework.boot:spring-boot-starter:1.5.2.RELEASE
- org.apache.solr:solr-solrj:5.5.4
- org.springframework.data:spring-data-solr:2.1.1.RELEASE
- org.apache.httpcomponents:httpmime:4.5.3
`

var springBootFreeMarkerStarter = `
name: Spring Boot FreeMarker Starter
description: "Starter for building MVC web applications using FreeMarker views"
dependencies:
- org.springframework.boot:spring-boot-starter:1.5.2.RELEASE
- org.springframework.boot:spring-boot-starter-web:1.5.2.RELEASE
- org.freemarker:freemarker:2.3.25-incubating
- org.springframework:spring-context-support:4.3.7.RELEASE
`

var springBootGroovyTemplatesStarter = `
name: Spring Boot Groovy Templates Starter
description: "Starter for building MVC web applications using Groovy Templates views"
dependencies:
- org.springframework.boot:spring-boot-starter:1.5.2.RELEASE
- org.springframework.boot:spring-boot-starter-web:1.5.2.RELEASE
- org.codehaus.groovy:groovy-templates:2.4.9
`

var springBootHATEOASStarter = `
name: Spring Boot HATEOAS Starter
description: "Starter for building hypermedia-based RESTful web application with Spring MVC and Spring HATEOAS"
dependencies:
- org.springframework.boot:spring-boot-starter-web:1.5.2.RELEASE
- org.springframework.hateoas:spring-hateoas:0.23.0.RELEASE
- org.springframework.plugin:spring-plugin-core:1.2.0.RELEASE
`

var springBootIntegrationStarter = `
name: Spring Boot Integration Starter
description: "Starter for using Spring Integration"
dependencies:
- org.springframework.boot:spring-boot-starter:1.5.2.RELEASE
- org.springframework.boot:spring-boot-starter-aop:1.5.2.RELEASE
- org.springframework.integration:spring-integration-core:4.3.8.RELEASE
- org.springframework.integration:spring-integration-java-dsl:1.2.1.RELEASE
`

var springBootJDBCStarter = `
name: Spring Boot JDBC Starter
description: "Starter for using JDBC with the Tomcat JDBC connection pool"
dependencies:
- org.springframework.boot:spring-boot-starter:1.5.2.RELEASE
- org.apache.tomcat:tomcat-jdbc:8.5.11
- org.springframework:spring-jdbc:4.3.7.RELEASE
`

var springBootJerseyStarter = `
name: Spring Boot Jersey Starter
description: "Starter for building RESTful web applications using JAX-RS and Jersey. An alternative to spring-boot-starter-web"
dependencies:
- org.springframework.boot:spring-boot-starter:1.5.2.RELEASE
- org.springframework.boot:spring-boot-starter-tomcat:1.5.2.RELEASE
- org.springframework.boot:spring-boot-starter-validation:1.5.2.RELEASE
- com.fasterxml.jackson.core:jackson-databind:2.8.7
- org.springframework:spring-web:4.3.7.RELEASE
- org.glassfish.jersey.core:jersey-server:2.25.1
- org.glassfish.jersey.containers:jersey-container-servlet-core:2.25.1
- org.glassfish.jersey.containers:jersey-container-servlet:2.25.1
- org.glassfish.jersey.ext:jersey-bean-validation:2.25.1
- org.glassfish.jersey.ext:jersey-spring3:2.25.1
- org.glassfish.jersey.media:jersey-media-json-jackson:2.25.1
`

var springBootJettyStarter = `
name: Spring Boot Jetty Starter
description: "Starter for using Jetty as the embedded servlet container. An alternative to spring-boot-starter-tomcat"
dependencies:
- org.eclipse.jetty:jetty-servlets:9.4.2.v20170220
- org.eclipse.jetty:jetty-webapp:9.4.2.v20170220
- org.eclipse.jetty.websocket:websocket-server:9.4.2.v20170220
- org.eclipse.jetty.websocket:javax-websocket-server-impl:9.4.2.v20170220
- org.mortbay.jasper:apache-el:8.0.33
`

var springBootJOOQStarter = `
name: Spring Boot JOOQ Starter
description: "Starter for using jOOQ to access SQL databases. An alternative to spring-boot-starter-data-jpa or spring-boot-starter-jdbc"
dependencies:
- org.springframework.boot:spring-boot-starter:1.5.2.RELEASE
- org.springframework.boot:spring-boot-starter-jdbc:1.5.2.RELEASE
- org.springframework:spring-tx:4.3.7.RELEASE
- org.jooq:jooq:3.9.1
`

var springBootAtomikosJTAStarter = `
name: Spring Boot Atomikos JTA Starter
description: "Starter for JTA transactions using Atomikos"
dependencies:
- org.springframework.boot:spring-boot-starter:1.5.2.RELEASE
- com.atomikos:transactions-jms:3.9.3
- com.atomikos:transactions-jta:3.9.3
- com.atomikos:transactions-jdbc:3.9.3
- javax.transaction:javax.transaction-api:1.2
`

var springBootBitronixJTAStarter = `
name: Spring Boot Bitronix JTA Starter
description: "Starter for JTA transactions using Bitronix"
dependencies:
- org.springframework.boot:spring-boot-starter:1.5.2.RELEASE
- javax.jms:jms-api:1.1-rev-1
- javax.transaction:javax.transaction-api:1.2
- org.codehaus.btm:btm:2.1.4
`

var springBootNarayanaJTAStarter = `
name: Spring Boot Narayana JTA Starter
description: "Spring Boot Narayana JTA Starter"
dependencies:
- org.springframework.boot:spring-boot-starter:1.5.2.RELEASE
- org.jboss:jboss-transaction-spi:7.5.1.Final
- org.jboss.narayana.jta:jdbc:5.5.3.Final
- org.jboss.narayana.jta:jms:5.5.3.Final
- org.jboss.narayana.jta:jta:5.5.3.Final
- org.jboss.narayana.jts:narayana-jts-integration:5.5.3.Final
- javax.transaction:javax.transaction-api:1.2
`

var springBootLog4j2Starter = `
name: Spring Boot Log4j 2 Starter
description: "Starter for using Log4j2 for logging. An alternative to spring-boot-starter-logging"
dependencies:
- org.apache.logging.log4j:log4j-slf4j-impl:2.7
- org.apache.logging.log4j:log4j-api:2.7
- org.apache.logging.log4j:log4j-core:2.7
- org.slf4j:jcl-over-slf4j:1.7.24
- org.slf4j:jul-to-slf4j:1.7.24
`

var springBootLoggingStarter = `
name: Spring Boot Logging Starter
description: "Starter for logging using Logback. Default logging starter"
dependencies:
- ch.qos.logback:logback-classic:1.2.2
- org.slf4j:jcl-over-slf4j:1.7.25
- org.slf4j:jul-to-slf4j:1.7.25
- org.slf4j:log4j-over-slf4j:1.7.25
`

var springBootMailStarter = `
name: Spring Boot Mail Starter
description: "Starter for using Java Mail and Spring Framework's email sending support"
dependencies:
- org.springframework.boot:spring-boot-starter:1.5.2.RELEASE
- org.springframework:spring-context:4.3.7.RELEASE
- org.springframework:spring-context-support:4.3.7.RELEASE
- com.sun.mail:javax.mail:1.5.6
`

var springBootMobileStarter = `
name: Spring Boot Mobile Starter
description: "Starter for building web applications using Spring Mobile"
dependencies:
- org.springframework.boot:spring-boot-starter:1.5.2.RELEASE
- org.springframework.boot:spring-boot-starter-web:1.5.2.RELEASE
- org.springframework.mobile:spring-mobile-device:1.1.5.RELEASE
`

var springBootMustacheStarter = `
name: Spring Boot Mustache Starter
description: "Starter for building MVC web applications using Mustache views"
dependencies:
- org.springframework.boot:spring-boot-starter:1.5.2.RELEASE
- org.springframework.boot:spring-boot-starter-web:1.5.2.RELEASE
- com.samskivert:jmustache:1.13
`

var springBootStarterParent = `
name: Spring Boot Starter Parent
description: "Parent pom providing dependency and plugin management for applications built with Maven"
dependencies:
`

var springBootReactorNettyStarter = `
name: Spring Boot Reactor Netty Starter
description: "Starter for using Reactor Netty as the embedded reactive HTTP server."
dependencies:
- io.projectreactor.ipc:reactor-netty:0.6.2.RELEASE
`

var springBootSecurityStarter = `
name: Spring Boot Security Starter
description: "Starter for using Spring Security"
dependencies:
- org.springframework.boot:spring-boot-starter:1.5.2.RELEASE
- org.springframework:spring-aop:4.3.7.RELEASE
- org.springframework.security:spring-security-config:4.2.2.RELEASE
- org.springframework.security:spring-security-web:4.2.2.RELEASE
`

var springBootSocialFacebookStarter = `
name: Spring Boot Social Facebook Starter
description: "Starter for using Spring Social Facebook"
dependencies:
- org.springframework.boot:spring-boot-starter:1.5.2.RELEASE
- org.springframework.boot:spring-boot-starter-web:1.5.2.RELEASE
- org.springframework.social:spring-social-config:1.1.4.RELEASE
- org.springframework.social:spring-social-core:1.1.4.RELEASE
- org.springframework.social:spring-social-web:1.1.4.RELEASE
- org.springframework.social:spring-social-facebook:2.0.3.RELEASE
`

var springBootSocialLinkedInStarter = `
name: Spring Boot Social LinkedIn Starter
description: "Stater for using Spring Social LinkedIn"
dependencies:
- org.springframework.boot:spring-boot-starter:1.5.2.RELEASE
- org.springframework.boot:spring-boot-starter-web:1.5.2.RELEASE
- org.springframework.social:spring-social-config:1.1.4.RELEASE
- org.springframework.social:spring-social-core:1.1.4.RELEASE
- org.springframework.social:spring-social-web:1.1.4.RELEASE
- org.springframework.social:spring-social-linkedin:1.0.2.RELEASE
`

var springBootSocialTwitterStarter = `
name: Spring Boot Social Twitter Starter
description: "Starter for using Spring Social Twitter"
dependencies:
- org.springframework.boot:spring-boot-starter:1.5.2.RELEASE
- org.springframework.boot:spring-boot-starter-web:1.5.2.RELEASE
- org.springframework.social:spring-social-config:1.1.4.RELEASE
- org.springframework.social:spring-social-core:1.1.4.RELEASE
- org.springframework.social:spring-social-web:1.1.4.RELEASE
- org.springframework.social:spring-social-twitter:1.1.2.RELEASE
`

var springBootTestStarter = `
name: Spring Boot Test Starter
description: "Starter for testing Spring Boot applications with libraries including JUnit, Hamcrest and Mockito"
dependencies:
- org.springframework.boot:spring-boot-test:1.5.2.RELEASE
- org.springframework.boot:spring-boot-test-autoconfigure:1.5.2.RELEASE
- com.jayway.jsonpath:json-path:2.2.0
- junit:junit:4.12
- org.assertj:assertj-core:2.6.0
- org.mockito:mockito-core:1.10.19
- org.hamcrest:hamcrest-core:1.3
- org.hamcrest:hamcrest-library:1.3
- org.skyscreamer:jsonassert:1.4.0
- org.springframework:spring-core:4.3.7.RELEASE
- org.springframework:spring-test:4.3.7.RELEASE
`

var springBootThymeleafStarter = `
name: Spring Boot Thymeleaf Starter
description: "Starter for building MVC web applications using Thymeleaf views"
dependencies:
- org.springframework.boot:spring-boot-starter:1.5.2.RELEASE
- org.springframework.boot:spring-boot-starter-web:1.5.2.RELEASE
- org.thymeleaf:thymeleaf-spring4:2.1.5.RELEASE
- nz.net.ultraq.thymeleaf:thymeleaf-layout-dialect:1.4.0
`

var springBootTomcatStarter = `
name: Spring Boot Tomcat Starter
description: "Starter for using Tomcat as the embedded servlet container. Default servlet container starter used by spring-boot-starter-web"
dependencies:
- org.apache.tomcat.embed:tomcat-embed-core:8.5.11
- org.apache.tomcat.embed:tomcat-embed-el:8.5.11
- org.apache.tomcat.embed:tomcat-embed-websocket:8.5.11
`

var springBootUndertowStarter = `
name: Spring Boot Undertow Starter
description: "Starter for using Undertow as the embedded servlet container. An alternative to spring-boot-starter-tomcat"
dependencies:
- io.undertow:undertow-core:1.4.11.Final
- io.undertow:undertow-servlet:1.4.11.Final
- io.undertow:undertow-websockets-jsr:1.4.11.Final
- javax.servlet:javax.servlet-api:3.1.0
- org.glassfish:javax.el:3.0.0
`

var springBootValidationStarter = `
name: Spring Boot Validation Starter
description: "Starter for using Java Bean Validation with Hibernate Validator"
dependencies:
- org.springframework.boot:spring-boot-starter:1.5.2.RELEASE
- org.apache.tomcat.embed:tomcat-embed-el:8.5.11
- org.hibernate:hibernate-validator:5.3.4.Final
`

var springBootWebStarter = `
name: Spring Boot Web Starter
description: "Starter for building web, including RESTful, applications using Spring MVC. Uses Tomcat as the default embedded container"
dependencies:
- org.springframework.boot:spring-boot-starter:1.5.2.RELEASE
- org.springframework.boot:spring-boot-starter-tomcat:1.5.2.RELEASE
- org.hibernate:hibernate-validator:5.3.4.Final
- com.fasterxml.jackson.core:jackson-databind:2.8.7
- org.springframework:spring-web:4.3.7.RELEASE
- org.springframework:spring-webmvc:4.3.7.RELEASE
`

var springBootWebServicesStarter = `
name: Spring Boot Web Services Starter
description: "Starter for using Spring Web Services"
dependencies:
- org.springframework.boot:spring-boot-starter:1.5.2.RELEASE
- org.springframework.boot:spring-boot-starter-web:1.5.2.RELEASE
- org.springframework:spring-oxm:4.3.7.RELEASE
- org.springframework.ws:spring-ws-core:2.4.0.RELEASE
`

var springBootWebFluxStarter = `
name: Spring Boot WebFlux Starter
description: "Starter for building WebFlux applications using Spring Framework's Reactive Web support"
dependencies:
- org.springframework.boot:spring-boot-starter:2.0.0.BUILD-SNAPSHOT
- org.springframework.boot:spring-boot-starter-reactor-netty:2.0.0.BUILD-SNAPSHOT
- com.fasterxml.jackson.core:jackson-databind:2.9.0.pr2
- org.hibernate:hibernate-validator:5.4.1.Final
- org.springframework:spring-web:5.0.0.BUILD-SNAPSHOT
- org.springframework:spring-webflux:5.0.0.BUILD-SNAPSHOT
`

var springBootWebSocketStarter = `
name: Spring Boot WebSocket Starter
description: "Starter for building WebSocket applications using Spring Framework's WebSocket support"
dependencies:
- org.springframework.boot:spring-boot-starter:1.5.2.RELEASE
- org.springframework.boot:spring-boot-starter-web:1.5.2.RELEASE
- org.springframework:spring-messaging:4.3.7.RELEASE
- org.springframework:spring-websocket:4.3.7.RELEASE
`
