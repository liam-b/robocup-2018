package main

var waterTowerMatches = 0

func DetectedWaterTower(distance int, count int) bool {
  value := int(float64(2550 - bot.ultrasonicSensor.Distance()) / 2.55)

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