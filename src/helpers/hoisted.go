package main

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