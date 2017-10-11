package main

func checkBatteryVoltage() {
  log.debug("voltage is at " + log.value(bot.battery.voltageString() + "v"))
  if (bot.battery.voltage() < 70) {
    log.warn("battery needs replacing now")
  } else if (bot.battery.voltage() < 75) {
    log.warn("current voltage is not fit for comp")
  }
}