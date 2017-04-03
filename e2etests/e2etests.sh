#!/bin/bash

echo "Example 1: Update version of one outdated dependency"
printf "Similarity: "
curl -s -X GET -d @e2etests/example-1-vertx-web-outdated-rx-java-effective-pom.xml localhost:8088/recommendation | jq .Similarity
curl -s -X GET -d @e2etests/example-1-vertx-web-outdated-rx-java-effective-pom.xml localhost:8088/recommendation | jq .RecommendationItems

echo "Example 2: Add missing dependency"
printf "Similarity: "
curl -s -X GET -d @e2etests/example-2-vertx-web-missing-dependency-effective-pom.xml localhost:8088/recommendation | jq .Similarity
curl -s -X GET -d @e2etests/example-2-vertx-web-missing-dependency-effective-pom.xml localhost:8088/recommendation | jq .RecommendationItems

echo "Example 3: Downgrade version of one too new dependency"
printf "Similarity: "
curl -s -X GET -d @e2etests/example-3-vertx-web-too-new-rx-java-effective-pom.xml localhost:8088/recommendation | jq .Similarity
curl -s -X GET -d @e2etests/example-3-vertx-web-too-new-rx-java-effective-pom.xml localhost:8088/recommendation | jq .RecommendationItems
