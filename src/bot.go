package main

// main bot class //
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

// status checks //
func batteryStatus() {
  log.inc(".battery")
    log.debug("voltage is at " + log.value(bot.battery.voltageString() + "v"))

    if (bot.battery.voltage() > 125) {
      log.warn("possible overvolting")
    } else if (bot.battery.voltage() < 70) {
      log.error("replace battery now")
    } else if (bot.battery.voltage() < 75) {
      log.warn("battery needs replacing")
    }
  log.dec()
}