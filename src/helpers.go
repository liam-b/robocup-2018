package main

func GetColors() (string, string) {
  leftRed, leftGreen, leftBlue := bot.colorSensorLeft.Rgb()
  leftTotal := bot.colorSensorLeft.RgbIntensity()
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

func DetectedSilver() bool {
  left, right := GetColors()
  return left == SILVER && right == SILVER
}

func DetectedGreen(sensor int) bool {
  red, green, blue := 0, 0, 0

  if sensor == LEFT {
    red, green, blue = bot.colorSensorLeft.Rgb()
  }
  if sensor == RIGHT {
    red, green, blue = bot.colorSensorRight.Rgb()
  }

  return green > blue + GREEN_DETECT_RGB_DIFFERENCE && green > red + GREEN_DETECT_RGB_DIFFERENCE
}

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

var totalAngle = 0

func GyroAtAngle(angle int, turnDirection int) bool {
  totalAngle += bot.imu.ReadGyro()

  if turnDirection == LEFT && totalAngle > angle {
    totalAngle = 0
    return true
  }

  if turnDirection == RIGHT && totalAngle < angle {
    totalAngle = 0
    return true
  }

  return false
}