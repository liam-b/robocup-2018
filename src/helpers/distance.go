package main

var waterTowerMatches = 0
var canMatches = 0

func DetectedWaterTower(distance int, count int) bool {
  value := bot.ultrasonicSensor.Distance()

  if waterTowersAvoided > 3 { return false }

  if value > distance {
    waterTowerMatches += 1
  } else {
    waterTowerMatches = 0
  }

  if waterTowerMatches > count {
    waterTowerMatches = 0
    return true
  }
  return false
}

func DetectedCan(distance int, count int) bool {
  value := bot.ultrasonicSensor.Distance()

  if value > distance {
    canMatches += 1
  }  else {
    canMatches = 0
  }

  if canMatches > count {
    canMatches = 0
    return true
  }
  return false
}

func LostCan(distance int, count int) bool {
  value := bot.ultrasonicSensor.Distance()

  if value < distance {
    canMatches += 1
  }  else {
    canMatches = 0
  }

  if canMatches > count {
    canMatches = 0
    return true
  }
  return false
}
