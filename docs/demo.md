# Demo of stpl

A demo instance of stpl is running at: http://stpl-mdl.1d35.starter-us-east-1.openshiftapps.com/

## A simple example.

Let's say our have you are building a java vertx application. And you have the following pom.xml:

    $ curl -O https://github.com/luebken/stpl/blob/master/e2etests/example-1-vertx-web-outdated-rx-java-effective-pom.xml

You can get in analyzed by quering the ``analysis` endpoint:

    $ curl -X GET -d @example-1-vertx-web-outdated-rx-java-effective-pom.xml stpl-mdl.1d35.starter-us-east-1.openshiftapps.com/analysis

And would get result like this:
    $ {
    "Analysis": {
        "Recommendation": {
        "Similarity": 0,
        "ReferenceStack": {
            "Name": "",
            "Description": "",
            "Dependencies": null
        },
        "RecommendationItems": null
        },
        "LibrariesIoComponents": [
        {
            "name": "io.vertx:vertx-core",
            "platform": "Maven",
            "description": "Vert.x is a tool-kit for building reactive applications on the JVM",
            "homepage": "http://vertx.io",
            "repository_url": "scm:git:git@github.com:eclipse/vert.x.git",
            "rank": 18,
            "latest_release_published_at": "2017-03-15T13:46:59.000Z",
            "latest_release_number": "3.4.1",
            "language": "Java",
            "status": "",
            "package_manager_url": "http://search.maven.org/#search%7Cgav%7C1%7Cg%3A%22io.vertx%22%20AND%20a%3A%22vertx-core%22",
            "stars": 6055,
            "forks": 1129,
            "keywords": []
        },
        ....
