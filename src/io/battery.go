package main

import "strconv"

type Battery struct {
  path string
  device Device
}

func (battery Battery) new() Battery {
  battery.path = BATTERY_PATH
  battery.device = Device{path: battery.path}
  return battery
}

func (battery Battery) voltageString() string {
  voltage := battery.device.get("voltage_now")
  return string(voltage[0]) + "." + string(voltage[1])
}

func (battery Battery) voltage() int {
  voltage := battery.device.get("voltage_now")
  output := string(voltage[0]) + string(voltage[1])
  value, _ := strconv.Atoi(output)
  return value
}