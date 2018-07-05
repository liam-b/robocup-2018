package main

func LineSensorError() float64 {
  return float64(bot.colorSensorLeft.RgbIntensity() - bot.colorSensorRight.RgbIntensity())
}
