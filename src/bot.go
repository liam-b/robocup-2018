package main

type Bot struct {
  battery Battery
  colorSensorL ColorSensor
  colorSensorR ColorSensor
  ultrasonicSensor UltrasonicSensor
  gyroSensor GyroSensor
  speaker Speaker

  button Button
}

func (bot Bot) new() Bot {
  return bot
}
