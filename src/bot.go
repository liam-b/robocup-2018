package main

type Bot struct {
  battery Battery
  colorSensorLeft ColorSensor
  colorSensorRight ColorSensor
  ultrasonicSensor UltrasonicSensor
  imu IMU

  motorLeft Motor
  motorRight Motor

  ledshim Ledshim
}

func (bot *Bot) ResetAllCaches()() {
  bot.battery.ResetCache()
  bot.colorSensorLeft.ResetCache()
  bot.colorSensorRight.ResetCache()
  bot.ultrasonicSensor.ResetCache()
  bot.imu.ResetCache()

  // bot.motorLeft.ResetCache()
  // bot.motorRight.ResetCache()
}