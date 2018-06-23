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

  /* if (intensityLeft < FOLLOW_LINE_HARD_TURN_VALUE && intensityRight < FOLLOW_LINE_HARD_TURN_VALUE) {
    go bot.motorRight.RunForever(FOLLOW_LINE_SPEED)
    go bot.motorLeft.RunForever(FOLLOW_LINE_SPEED)
  } else */ if (intensityLeft < FOLLOW_LINE_HARD_TURN_VALUE) {
    go bot.motorRight.RunForever(SpeedRatio(FOLLOW_LINE_SPEED, FOLLOW_LINE_SOFT_TURN_RATIO, FAST))
    go bot.motorLeft.RunForever(SpeedRatio(FOLLOW_LINE_SPEED, FOLLOW_LINE_SOFT_TURN_RATIO, SLOW) - FOLLOW_LINE_SPEED)
  } else if (intensityRight < FOLLOW_LINE_HARD_TURN_VALUE) {
    go bot.motorRight.RunForever(SpeedRatio(FOLLOW_LINE_SPEED, FOLLOW_LINE_SOFT_TURN_RATIO, SLOW) - FOLLOW_LINE_SPEED)
    go bot.motorLeft.RunForever(SpeedRatio(FOLLOW_LINE_SPEED, FOLLOW_LINE_SOFT_TURN_RATIO, FAST))
  } else if (intensityLeft < FOLLOW_LINE_SOFT_TURN_VALUE) {
    go bot.motorRight.RunForever(SpeedRatio(FOLLOW_LINE_SPEED, FOLLOW_LINE_SOFT_TURN_RATIO, FAST))
    go bot.motorLeft.RunForever(SpeedRatio(FOLLOW_LINE_SPEED, FOLLOW_LINE_SOFT_TURN_RATIO, SLOW))
  } else if (intensityRight < FOLLOW_LINE_SOFT_TURN_VALUE) {
    go bot.motorRight.RunForever(SpeedRatio(FOLLOW_LINE_SPEED, FOLLOW_LINE_SOFT_TURN_RATIO, SLOW))
    go bot.motorLeft.RunForever(SpeedRatio(FOLLOW_LINE_SPEED, FOLLOW_LINE_SOFT_TURN_RATIO, FAST))
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
    go bot.motorRight.RunForever(SpeedRatio(FOLLOW_LINE_SPEED, FOLLOW_LINE_SOFT_TURN_RATIO, FAST))
    go bot.motorLeft.RunForever(SpeedRatio(FOLLOW_LINE_SPEED, FOLLOW_LINE_SOFT_TURN_RATIO, SLOW) - FOLLOW_LINE_SPEED)
  } else if (sensor == RIGHT && greenRight < redRight + FOLLOW_LINE_GREEN_DIFFERENCE && greenRight < blueRight + FOLLOW_LINE_GREEN_DIFFERENCE) {
    go bot.motorRight.RunForever(SpeedRatio(FOLLOW_LINE_SPEED, FOLLOW_LINE_SOFT_TURN_RATIO, SLOW) - FOLLOW_LINE_SPEED)
    go bot.motorLeft.RunForever(SpeedRatio(FOLLOW_LINE_SPEED, FOLLOW_LINE_SOFT_TURN_RATIO, FAST))
  } else {
    go bot.motorRight.RunForever(FOLLOW_LINE_SPEED)
    go bot.motorLeft.RunForever(FOLLOW_LINE_SPEED)
  }

  return BEHAVIOUR
}