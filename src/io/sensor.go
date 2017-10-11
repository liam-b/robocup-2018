package main

import "strconv"
import "strings"
// import "fmt"

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

  REFLECT string
  COLOR string
  RGB string

  sensor Sensor
}

func (colorSensor ColorSensor) new() ColorSensor {
  colorSensor.sensor = Sensor{port: colorSensor.port}.new()

  colorSensor.REFLECT = "COL-REFLECT"
  colorSensor.COLOR = "COL-COLOR"
  colorSensor.RGB = "RGB-RAW"

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

// type LineSensors struct {
//   leftPort string
//   rightPort string
//
//   REFLECT string
//   COLOR string
//   RGB string
//
//   leftSensor Sensor
//   rightSensor Sensor
// }
//
// func (colorSensor ColorSensor) new() ColorSensor {
//   colorSensor.sensor = Sensor{port: colorSensor.port}.new()
//
//   colorSensor.REFLECT = "COL-REFLECT"
//   colorSensor.COLOR = "COL-COLOR"
//   colorSensor.RGB = "RGB-RAW"
//
//   return colorSensor
// }

// for testing
// type TouchSensor struct {
//   port string
//
//   sensor Sensor
// }
//
// func (touchSensor TouchSensor) new() TouchSensor {
//   touchSensor.sensor = Sensor{port: touchSensor.port}.new()
//   return touchSensor
// }
//
// func (touchSensor TouchSensor) pressed() bool {
//   return touchSensor.sensor.value("0") == 1
// }

type GyroSensor struct {
  port string

  sensor Sensor
}

func (gyroSensor GyroSensor) new() GyroSensor {
  gyroSensor.sensor = Sensor{port: gyroSensor.port}.new()
  return gyroSensor
}

func (gyroSensor GyroSensor) angle() int {
  return gyroSensor.sensor.value("0")
}

type UltrasonicSensor struct {
  port string

  sensor Sensor
}

func (ultrasonicSensor UltrasonicSensor) new() UltrasonicSensor {
  ultrasonicSensor.sensor = Sensor{port: ultrasonicSensor.port}.new()
  ultrasonicSensor.sensor.mode("US-DIST-CM")
  return ultrasonicSensor
}

func (ultrasonicSensor UltrasonicSensor) distance() int {
  return ultrasonicSensor.sensor.value("0")
}