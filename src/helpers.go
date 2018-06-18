package main

var waterTowerMatches = 0
var detectedWaterTowers = 0

func GetColors() (string, string) {
  leftRed, leftGreen, leftBlue := bot.colorSensorLeft.Rgb()
  leftTotal := int((leftRed + leftGreen + leftBlue) / 3)
  leftColor := WHITE

  if (leftTotal > 45) {
    leftColor = SILVER
  } else if (leftGreen > leftBlue + 6 && leftGreen > leftRed + 6) {
    leftColor = GREEN
  } else if (leftTotal < 6) {
    leftColor = BLACK
  }

  rightRed, rightGreen, rightBlue := bot.colorSensorRight.Rgb()
  rightTotal := int((rightRed + rightGreen + rightBlue) / 3)
  rightColor := WHITE

  if (rightTotal > 45) {
    rightColor = SILVER
  } else if (rightGreen > rightBlue + 6 && rightGreen > rightRed + 6) {
    rightColor = GREEN
  } else if (rightTotal < 6) {
    rightColor = BLACK
  }

  return leftColor, rightColor
}

func DetectedWaterTower(distance int, count int) bool {
  value := int(float64(2550 - bot.ultrasonicSensor.Distance()) / 2.55)

  if value > distance {
    waterTowerMatches += 1
  } else {
    waterTowerMatches = 0
  }

  return waterTowerMatches > count
}