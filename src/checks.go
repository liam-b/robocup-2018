package main

func BatteryStatus() {
  log.inc(".battery")
    log.debug("voltage is at " + log.value(bot.battery.VoltageString() + "v"))
    currentVoltage := bot.battery.Voltage()

    if (currentVoltage > 125) {
      log.warn("possible overvolting")
    } else if (currentVoltage < 105) {
      log.error("replace battery now")
      bot.ledshim.SetPixel(BATTERY_PIXEL, COLOR_RED)
    } else if (currentVoltage < 110) {
      log.warn("battery needs replacing")
      bot.ledshim.SetPixel(BATTERY_PIXEL, COLOR_BLUE)
    } else {
      bot.ledshim.SetPixel(BATTERY_PIXEL, COLOR_GREEN)
    }
  log.dec()
}