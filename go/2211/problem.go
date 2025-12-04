package problem

import "strings"

func solve(directions string) int {

  directions = strings.TrimLeft(directions, "L")
  directions = strings.TrimRight(directions,"R")
  stationaryCount := strings.Count(directions, "S")

  collisionCount := len(directions) - stationaryCount

  return collisionCount 
}
