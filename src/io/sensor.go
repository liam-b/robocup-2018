package main

import "strconv"
import "strings"

type Sensor struct {
  port string

  path string
  indexedDevice IndexedDevice
}

func (sensor Sensor) new() Sensor {
  sensor.path = SENSOR_PATH
  sensor.indexedDevice = IndexedDevice{path: sensor.path, port: sensor.port}
  sensor.indexedDevice.findDeviceFromPort()
  return sensor
}

func (sensor Sensor) value(num string) int {
  value := sensor.indexedDevice.get("value" + num)
  value = strings.TrimSuffix(value, "\n")
  result, _ := strconv.Atoi(value)
  return result
}

func (sensor Sensor) mode(newMode string) {
  sensor.indexedDevice.set("mode", newMode)
}

type ColorSensor struct {
  port string

  sensor Sensor
}

func (colorSensor ColorSensor) new() ColorSensor {
  colorSensor.sensor = Sensor{port: colorSensor.port}.new()
  return colorSensor
}

func (colorSensor ColorSensor) mode(newMode string) {
  colorSensor.sensor.mode(newMode)
}

func (colorSensor ColorSensor) intensity() int {
  return colorSensor.sensor.value("0")
}

func (colorSensor ColorSensor) color() int {
  return colorSensor.sensor.value("0")
}

func (colorSensor ColorSensor) rgb() (int, int, int) {
  return colorSensor.sensor.value("0"), colorSensor.sensor.value("1"), colorSensor.sensor.value("2")
}

type TouchSensor struct {
  port string

  sensor Sensor
}

func (touchSensor TouchSensor) new() TouchSensor {
  touchSensor.sensor = Sensor{port: touchSensor.port}.new()
  return touchSensor
}

func (touchSensor TouchSensor) pressed() bool {
  return touchSensor.sensor.value("0") == 1
}