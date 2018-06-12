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
    // log.trace(string(bot.battery.voltage()))
    else if (bot.battery.voltage() < 700) {
      log.error("replace battery now")
    } else if (bot.battery.voltage() < 750) {
      log.warn("battery needs replacing")
    }
  log.dec()
}