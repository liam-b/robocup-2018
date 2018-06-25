package main

const RGB_GREEN_DIFFERENCE = 6
const RGB_SILVER_VALUE = 40
const RGB_BLACK_VALUE = 9

func GetColors() (string, string) {
  leftRed, leftGreen, leftBlue := bot.colorSensorLeft.Rgb()
  leftTotal := bot.colorSensorLeft.RgbIntensity()
  leftColor := WHITE

  if (leftTotal > RGB_SILVER_VALUE) {
    leftColor = SILVER
  } else if (leftGreen > leftBlue + RGB_GREEN_DIFFERENCE && leftGreen > leftRed + RGB_GREEN_DIFFERENCE && leftGreen < RGB_SILVER_VALUE) {
    leftColor = GREEN
  } else if (leftTotal < RGB_BLACK_VALUE) {
    leftColor = BLACK
  }

  rightRed, rightGreen, rightBlue := bot.colorSensorRight.Rgb()
  rightTotal := bot.colorSensorRight.RgbIntensity()
  rightColor := WHITE

  if (rightTotal > RGB_SILVER_VALUE) {
    rightColor = SILVER
  } else if (rightGreen > rightBlue + RGB_GREEN_DIFFERENCE && rightGreen > rightRed + RGB_GREEN_DIFFERENCE && rightGreen < RGB_SILVER_VALUE) {
    rightColor = GREEN
  } else if (rightTotal < RGB_BLACK_VALUE) {
    rightColor = BLACK
  }

  return leftColor, rightColor
}

func DetectedSilver() bool {
  left, right := GetColors()
  return left == SILVER && right == SILVER
}

func DetectedGreen(sensor int) bool {
  left, right := GetColors()
  return left == GREEN && right == GREEN
}

var liftedMatches = 0

func BotLifted(count int) bool {
  if bot.colorSensorLeft.RgbIntensity() == 0 && bot.colorSensorRight.RgbIntensity() == 0 {
    liftedMatches += 1
  } else {
    liftedMatches = 0
  }

  if liftedMatches > count {
    liftedMatches = 0
    return true
  }
  return false
}

func BotPlacedDown() bool {
  return bot.colorSensorLeft.RgbIntensity() > 1 && bot.colorSensorRight.RgbIntensity() > 1
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

func GyroTurnedToAngle(angle int, turnDirection int) bool {
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

func SpeedRatio(speed int, ratio float64, sign int) int {
  return speed + int(float64(speed) * ratio * float64(sign))
}

func ResetHelpers() {
  liftedMatches = 0
  waterTowerMatches = 0
  totalAngle = 0
}

func min(x, y int) int {
  if x < y { return x }
  return y
}

func max(x, y int) int {
  if x > y { return x }
  return y
}