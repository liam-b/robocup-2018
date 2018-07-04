package main

const FOLLOW_LINE_SPEED = 300
const FOLLOW_LINE_HARD_TURN_RATIO = 1.3
const FOLLOW_LINE_SOFT_TURN_RATIO = 1.2
const FOLLOW_LINE_HARD_TURN_VALUE = 10
const FOLLOW_LINE_SOFT_TURN_VALUE = 27
const FOLLOW_LINE_GREEN_DIFFERENCE = 14

func FollowLine() string {
  intensityLeft := bot.colorSensorLeft.RgbIntensity()
  intensityRight := bot.colorSensorRight.RgbIntensity()

  if (intensityLeft < FOLLOW_LINE_HARD_TURN_VALUE) {
    go bot.motorRight.RunForever(FOLLOW_LINE_SPEED + 200)
    go bot.motorLeft.RunForever(FOLLOW_LINE_SPEED - 200)
  } else if (intensityRight < FOLLOW_LINE_HARD_TURN_VALUE) {
    go bot.motorRight.RunForever(FOLLOW_LINE_SPEED - 200)
    go bot.motorLeft.RunForever(FOLLOW_LINE_SPEED + 200)
  } else if (intensityLeft < FOLLOW_LINE_SOFT_TURN_VALUE) {
    go bot.motorRight.RunForever(FOLLOW_LINE_SPEED + 50)
    go bot.motorLeft.RunForever(FOLLOW_LINE_SPEED - 50)
  } else if (intensityRight < FOLLOW_LINE_SOFT_TURN_VALUE) {
    go bot.motorRight.RunForever(FOLLOW_LINE_SPEED - 50)
    go bot.motorLeft.RunForever(FOLLOW_LINE_SPEED + 50)
  } else {
    go bot.motorRight.RunForever(FOLLOW_LINE_SPEED)
    go bot.motorLeft.RunForever(FOLLOW_LINE_SPEED)
  }

  return BEHAVIOUR
}

func OneSensorLineFollowing(sensor int) string {
  redLeft, greenLeft, blueLeft := bot.colorSensorLeft.Rgb()
  redRight, greenRight, blueRight := bot.colorSensorRight.Rgb()

  if (sensor == LEFT && greenLeft < redLeft + FOLLOW_LINE_GREEN_DIFFERENCE && greenLeft < blueLeft + FOLLOW_LINE_GREEN_DIFFERENCE) {
    go bot.motorRight.RunForever(FOLLOW_LINE_SPEED + 100)
    go bot.motorLeft.RunForever(FOLLOW_LINE_SPEED - 100)
  } else if (sensor == RIGHT && greenRight < redRight + FOLLOW_LINE_GREEN_DIFFERENCE && greenRight < blueRight + FOLLOW_LINE_GREEN_DIFFERENCE) {
    go bot.motorRight.RunForever(FOLLOW_LINE_SPEED + 100)
    go bot.motorLeft.RunForever(FOLLOW_LINE_SPEED - 100)
  } else {
    go bot.motorRight.RunForever(FOLLOW_LINE_SPEED)
    go bot.motorLeft.RunForever(FOLLOW_LINE_SPEED)
  }

  return BEHAVIOUR
}
