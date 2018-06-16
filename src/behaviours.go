package main

func findColors() (string, string) {
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

func FollowLine() {
  intensityL := bot.colorSensorLeft.Intensity()
  intensityR := bot.colorSensorRight.Intensity()

  if (intensityL < 72) {
    go bot.motorRight.RunForever(440)
    go bot.motorLeft.RunForever(230)
  }

  if (intensityR < 72) {
    go bot.motorRight.RunForever(230)
    go bot.motorLeft.RunForever(440)
  }

  if (intensityL < 16) {
    go bot.motorRight.RunForever(500)
    go bot.motorLeft.RunForever(80)
  }

  if (intensityR < 16) {
    go bot.motorRight.RunForever(80)
    go bot.motorLeft.RunForever(500)
  }

  if (intensityR > 60 && intensityL > 60) {
    go bot.motorRight.RunForever(300)
    go bot.motorLeft.RunForever(300)
  }
}