package io

import "strconv"
import "strings"

type Sensor struct {
  Port string
  path string
  connection string
  name string

  manualDevice ManualDevice
}

func (sensor Sensor) New() Sensor {
  sensor.path = SENSOR_PATH
  sensor.manualDevice = ManualDevice{path: sensor.path, port: sensor.Port, connection: sensor.connection, name: sensor.name}
  sensor.manualDevice.findDeviceFromPort()
  return sensor
}

func (sensor Sensor) value(num string) int {
  value := sensor.manualDevice.get("value" + num)
  value = strings.TrimSuffix(value, "\n")
  result, _ := strconv.Atoi(value)
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

func (colorSensor ColorSensor) Mode(newMode string) {
  colorSensor.sensor.mode(newMode)
}

func (colorSensor ColorSensor) Intensity() int {
  return colorSensor.sensor.value("0")
}

func (colorSensor ColorSensor) Color() int {
  return colorSensor.sensor.value("0")
}

func (colorSensor ColorSensor) Rgb() (int, int, int) {
  return colorSensor.sensor.value("0") / 10, colorSensor.sensor.value("1") / 10, colorSensor.sensor.value("2") / 10
}

func (colorSensor ColorSensor) RgbIntensity() int {
  return (colorSensor.sensor.value("0") + colorSensor.sensor.value("1") + colorSensor.sensor.value("2")) / 3 / 10
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

func (ultrasonicSensor UltrasonicSensor) Mode(newMode string) {
  ultrasonicSensor.sensor.mode(newMode)
}

func (ultrasonicSensor UltrasonicSensor) Distance() int {
  return ultrasonicSensor.sensor.value("0")
}

type ButtonSensor struct {
  Port string

  TOUCH string

  sensor Sensor
}

func (buttonSensor ButtonSensor) New() ButtonSensor {
  buttonSensor.sensor = Sensor{Port: buttonSensor.Port, connection: "ev3-uart", name: "lego-ev3-touch"}.New()
  buttonSensor.TOUCH = "TOUCH"

  return buttonSensor
}

func (buttonSensor ButtonSensor) Mode(newMode string) {
  buttonSensor.sensor.mode(newMode)
}

func (buttonSensor ButtonSensor) Pressed() bool {
  return buttonSensor.sensor.value("0") == 1
}