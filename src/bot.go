package main

type Bot struct {
  battery Battery
  colorSensorL ColorSensor
  colorSensorR ColorSensor
  ultrasonicSensor UltrasonicSensor
  gyroSensor GyroSensor
  speaker Speaker

  motorL Motor
  motorR Motor

  button Button
}

func (bot Bot) new() Bot {
  return bot
}
