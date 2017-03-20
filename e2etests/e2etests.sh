#!/usr/bin/env bats
# https://github.com/sstephenson/bats
@test "example-1-pom list dependencies" {
  result=`curl -X POST -d @example-1-pom.xml localhost:8080/analytics`
  echo $result
  [ "$result" = "circuit-breaker-examples has 3 dependencies." ]
}