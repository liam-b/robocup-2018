package main

type Bot struct {
  battery Battery
  colorSensorL ColorSensor
  colorSensorR ColorSensor
  ultrasonicSensor UltrasonicSensor

  motorL Motor
  motorR Motor
}

func (bot Bot) new() Bot {
  return bot
}

func checkBatteryVoltage() {
  log.inc(".battery")
    log.debug("voltage is at " + log.value(bot.battery.voltageString() + "v"))
    if (bot.battery.voltage() < 70) {
      log.warn("battery needs replacing now")
    } else if (bot.battery.voltage() < 75) {
      log.warn("current voltage is not fit for comp")
    }
  log.dec()
}