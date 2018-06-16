package main

import "./io"

type Bot struct {
  battery io.Battery
  colorSensorLeft io.ColorSensor
  colorSensorRight io.ColorSensor
  ultrasonicSensor io.UltrasonicSensor

  imu io.IMU
  ledshim io.Ledshim

  motorLeft io.Motor
  motorRight io.Motor
}

func (bot Bot) new() Bot {
  return bot
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