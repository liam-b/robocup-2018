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
  return (sensor == LEFT && left == GREEN) || (sensor == RIGHT && right == GREEN)
}
