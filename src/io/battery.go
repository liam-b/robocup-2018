package main

import "strconv"

type Battery struct {
  path string
  device Device

  cachedValue string
  hasCached bool
}

func (battery Battery) New() Battery {
  battery.path = BATTERY_PATH
  battery.device = Device{path: battery.path}
  return battery
}

func (battery *Battery) ResetCache() {
  battery.cachedValue = ""
  battery.hasCached = false
}

func (battery *Battery) VoltageString() string {
  voltage := battery.getVoltage()
  voltageString := string(voltage[0:len(voltage) - 6])
  if (len(voltageString) > 2) {
    return voltageString[:2] + "." + voltageString[2:]
  } else {
    return voltageString
  }
}

func (battery *Battery) Voltage() int {
  voltage := battery.getVoltage()
  output := string(voltage[0:len(voltage) - 6])
  value, _ := strconv.Atoi(output)
  return value
}

func (battery *Battery) getVoltage() string {
  if battery.hasCached { return battery.cachedValue }
  result := battery.device.get("voltage_now")
  battery.hasCached = true
  battery.cachedValue = result
  return result
}