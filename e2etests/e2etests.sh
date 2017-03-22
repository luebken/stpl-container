#!/bin/bash

echo "Recommend upgrading outdated version"
curl -s -X POST -d @e2etests/vertx-web-outdated-rx-java-effective-pom.xml localhost:8080/recommendation | jq .RecommendationItems