#!/usr/bin/env bats
# https://github.com/sstephenson/bats
@test "example-1-pom list dependencies" {
  result=`curl -X POST -d @e2etests/vertx-web-outdated-rx-java-effective-pom.xml localhost:8080/advice`
  echo $result
  [ "$result" = "circuit-breaker-examples has 3 dependencies." ]
}