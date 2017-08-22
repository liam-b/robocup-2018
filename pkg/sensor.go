package main

import "strconv"

type Sensor struct {
  port string

  path string
  indexedDevice IndexedDevice
}

func (sensor *Sensor) init() {
  sensor.path = SENSOR_PATH
  sensor.indexedDevice = IndexedDevice{path: sensor.path, port: sensor.port}
  sensor.indexedDevice.findDeviceFromPort()
}

func (sensor Sensor) value(num string) int {
  result, _ := strconv.ParseInt(sensor.indexedDevice.get("value" + num), 10, 0)
  return int(result)
}

func (sensor Sensor) mode(newMode string) {
  sensor.indexedDevice.set("mode", newMode)
}

type ColorSensor struct {
  port string

  sensor Sensor
}

func (colorSensor *ColorSensor) init() {
  colorSensor.sensor = Sensor{port: colorSensor.port}
  colorSensor.sensor.init()
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