package main

import "strconv"
import "strings"

type Sensor struct {
  Port string
  path string
  connection string
  name string

  cachedValues [8]int
  hasCached [8]bool

  manualDevice ManualDevice
}

func (sensor Sensor) New() Sensor {
  sensor.path = SENSOR_PATH
  sensor.manualDevice = ManualDevice{path: sensor.path, port: sensor.Port, connection: sensor.connection, name: sensor.name}
  sensor.manualDevice.findDeviceFromPort()
  return sensor
}

func (sensor *Sensor) resetCache() {
  sensor.cachedValues = [8]int{0, 0, 0, 0, 0, 0, 0, 0}
  sensor.hasCached = [8]bool{false, false, false, false, false, false, false, false}
}

func (sensor *Sensor) value(num int) int {
  if sensor.hasCached[num] { return sensor.cachedValues[num] }
  value := sensor.manualDevice.get("value" + strconv.Itoa(num))
  value = strings.TrimSuffix(value, "\n")
  result, _ := strconv.Atoi(value)
  sensor.cachedValues[num] = result
  sensor.hasCached[num] = true

  return result
}

func (sensor Sensor) mode(newMode string) {
  sensor.manualDevice.set("mode", newMode)
}

type ColorSensor struct {
  Port string
  REFLECT string
  COLOR string
  RGB string

  sensor Sensor
}

func (colorSensor ColorSensor) New() ColorSensor {
  colorSensor.sensor = Sensor{Port: colorSensor.Port, connection: "ev3-uart", name: "lego-ev3-color"}.New()
  colorSensor.REFLECT = "COL-REFLECT"
  colorSensor.COLOR = "COL-COLOR"
  colorSensor.RGB = "RGB-RAW"
  return colorSensor
}

func (colorSensor ColorSensor) GetCaches() [8]int {
  return colorSensor.sensor.cachedValues
}

func (colorSensor *ColorSensor) ResetCache() {
  colorSensor.sensor.resetCache()
}

func (colorSensor ColorSensor) Mode(newMode string) {
  colorSensor.sensor.mode(newMode)
}

func (colorSensor *ColorSensor) Intensity() int {
  return colorSensor.sensor.value(0)
}

func (colorSensor *ColorSensor) Color() int {
  return colorSensor.sensor.value(0)
}

func (colorSensor *ColorSensor) Rgb() (int, int, int) {
  return colorSensor.sensor.value(0) / 10, colorSensor.sensor.value(1) / 10, colorSensor.sensor.value(2) / 10
}

func (colorSensor *ColorSensor) RgbIntensity() int {
  return (colorSensor.sensor.value(0) + colorSensor.sensor.value(1) + colorSensor.sensor.value(2)) / 3 / 10
}

type UltrasonicSensor struct {
  Port string
  DISTANCE string

  sensor Sensor
}

func (ultrasonicSensor UltrasonicSensor) New() UltrasonicSensor {
  ultrasonicSensor.sensor = Sensor{Port: ultrasonicSensor.Port, connection: "ev3-uart", name: "lego-ev3-us"}.New()
  ultrasonicSensor.DISTANCE = "US-DIST-CM"
  return ultrasonicSensor
}

func (ultrasonicSensor *UltrasonicSensor) ResetCache() {
  ultrasonicSensor.sensor.resetCache()
}

func (ultrasonicSensor UltrasonicSensor) Mode(newMode string) {
  ultrasonicSensor.sensor.mode(newMode)
}

func (ultrasonicSensor *UltrasonicSensor) Distance() int {
  return int(ultrasonicSensor.sensor.value(0))
}

type TouchSensor struct {
  Port string
  TOUCH string

  sensor Sensor
}

func (touchSensor TouchSensor) New() TouchSensor {
  touchSensor.sensor = Sensor{Port: touchSensor.Port, connection: "ev3-analog", name: "lego-ev3-touch"}.New()
  touchSensor.TOUCH = "TOUCH"
  return touchSensor
}

func (touchSensor *TouchSensor) ResetCache() {
  touchSensor.sensor.resetCache()
}

func (touchSensor *TouchSensor) Mode(newMode string) {
  touchSensor.sensor.mode(newMode)
}

func (touchSensor *TouchSensor) Pressed() bool {
  return touchSensor.sensor.value(0) != 0
}
