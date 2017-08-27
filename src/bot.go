package main

type Bot struct {
  battery Battery
  colorSensor ColorSensor
  speaker Speaker
  touchSensor TouchSensor
  gyroSensor GyroSensor
  button Button
}

func (bot Bot) new() Bot {
  return bot
}
