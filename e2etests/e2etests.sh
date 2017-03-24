#!/bin/bash

echo "Example 1: Update version of one outdated depdendecy"
curl -s -X POST -d @e2etests/example-1-vertx-web-outdated-rx-java-effective-pom.xml localhost:8080/recommendation | jq .RecommendationItems

