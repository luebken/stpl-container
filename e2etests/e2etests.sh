#!/bin/bash

echo "Example 1: Update version of one outdated dependency"
printf "Similarity: "
curl -s -X GET -d @e2etests/example-1-vertx-web-outdated-rx-java-effective-pom.xml localhost:8088/analysis | jq .Analysis.Recommendation.Similarity
curl -s -X GET -d @e2etests/example-1-vertx-web-outdated-rx-java-effective-pom.xml localhost:8088/analysis | jq .Analysis.Recommendation.RecommendationItems

echo "Example 2: Add missing dependency"
printf "Similarity: "
curl -s -X GET -d @e2etests/example-2-vertx-openshift-missing-dependency-effective-pom.xml localhost:8088/analysis | jq .Analysis.Recommendation.Similarity
curl -s -X GET -d @e2etests/example-2-vertx-openshift-missing-dependency-effective-pom.xml localhost:8088/analysis | jq .Analysis.Recommendation.RecommendationItems

echo "Example 3: Downgrade version of one too new dependency"
printf "Similarity: "
curl -s -X GET -d @e2etests/example-3-vertx-web-too-new-rx-java-effective-pom.xml localhost:8088/analysis | jq .Analysis.Recommendation.Similarity
curl -s -X GET -d @e2etests/example-3-vertx-web-too-new-rx-java-effective-pom.xml localhost:8088/analysis | jq .Analysis.Recommendation.RecommendationItems

echo "Example 4: Two recommendations (Add / Downgrade)"
printf "Similarity: "
curl -s -X GET -d @e2etests/example-4-vertx-openshift-missing-and-new-dependency-effective-pom.xml localhost:8088/analysis | jq .Analysis.Recommendation.Similarity
curl -s -X GET -d @e2etests/example-4-vertx-openshift-missing-and-new-dependency-effective-pom.xml localhost:8088/analysis | jq .Analysis.Recommendation.RecommendationItems

# TODO discuss:
echo "Example 5: No recommendations"
printf "Similarity: "
curl -s -X GET -d @e2etests/example-5-vertx-no-reference-stack-pom.xml localhost:8088/analysis | jq .Analysis.Recommendation.Similarity
curl -s -X GET -d @e2etests/example-5-vertx-no-reference-stack-pom.xml localhost:8088/analysis | jq .Analysis.Recommendation.RecommendationItems