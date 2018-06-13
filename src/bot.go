package main

import "./io"

// main bot class //
type Bot struct {
  battery io.Battery
  colorSensorL io.ColorSensor
  colorSensorR io.ColorSensor
  ultrasonicSensor io.UltrasonicSensor

  motorL io.Motor
  motorR io.Motor
}

func (bot Bot) new() Bot {
  return bot
}

// status checks //
func batteryStatus() {
  log.inc(".battery")
    log.debug("voltage is at " + log.value(bot.battery.VoltageString() + "v"))
    currentVoltage := bot.battery.Voltage()

    if (currentVoltage > 125) {
      log.warn("possible overvolting")
    } else if (currentVoltage < 70) {
      log.error("replace battery now")
    } else if (currentVoltage < 75) {
      log.warn("battery needs replacing")
    }
  log.dec()
}