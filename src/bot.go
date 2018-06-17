package main

import "./io"

type Bot struct {
  battery io.Battery
  colorSensorLeft io.ColorSensor
  colorSensorRight io.ColorSensor
  ultrasonicSensor io.UltrasonicSensor
  imu io.IMU

  motorLeft io.Motor
  motorRight io.Motor

  ledshim io.Ledshim
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

func batteryStatus() {
  log.inc(".battery")
    log.debug("voltage is at " + log.value(bot.battery.VoltageString() + "v"))
    currentVoltage := bot.battery.Voltage()

    if (currentVoltage > 125) {
      log.warn("possible overvolting")
    } else if (currentVoltage < 105) {
      log.error("replace battery now")
      bot.ledshim.SetPixel(io.BATTERY_PIXEL, io.COLOR_RED)
    } else if (currentVoltage < 110) {
      log.warn("battery needs replacing")
      bot.ledshim.SetPixel(io.BATTERY_PIXEL, io.COLOR_BLUE)
    } else {
      bot.ledshim.SetPixel(io.BATTERY_PIXEL, io.COLOR_GREEN)
    }
  log.dec()
}