package io

import "strconv"

type Battery struct {
  path string
  device Device
}

func (battery Battery) New() Battery {
  battery.path = BATTERY_PATH
  battery.device = Device{path: battery.path}
  return battery
}

func (battery Battery) VoltageString() string {
  voltage := battery.device.get("voltage_now")
  voltageString := string(voltage[0:len(voltage) - 6])
  if (len(voltageString) > 2) {
    return voltageString[:2] + "." + voltageString[2:]
  } else {
    return voltageString
  }
}

func (battery Battery) Voltage() int {
  voltage := battery.device.get("voltage_now")
  output := string(voltage[0:len(voltage) - 6])
  value, _ := strconv.Atoi(output)
  return value
}