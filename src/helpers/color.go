package main

const RGB_GREEN_DIFFERENCE = 6
const RGB_SILVER_VALUE = 32
const RGB_BLACK_VALUE = 9

var liftedMatches = 0

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
  return (sensor == LEFT && left == GREEN) || (sensor == RIGHT && right == GREEN)
}

func LineSensorError() float64 {
  // return float64((bot.colorSensorLeft.RgbIntensity() + 5) - bot.colorSensorRight.RgbIntensity())
  return float64(min(bot.colorSensorLeft.RgbIntensity() + 2, 30) - min(bot.colorSensorRight.RgbIntensity(), 30))
}

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
