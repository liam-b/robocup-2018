package main

import "io/ioutil"
import "os"
import "strings"
import "errors"
import "strconv"

const S1 = "spi0.1:S1"
const S2 = "spi0.1:S2"
const S3 = "spi0.1:S3"
const S4 = "spi0.1:S4"

const MA = "spi0.1:MA"
const MB = "spi0.1:MB"
const MC = "spi0.1:MC"
const MD = "spi0.1:MD"

var PORT = map[string]string{
  "spi0.1:S1": "port0",
  "spi0.1:S2": "port1",
  "spi0.1:S3": "port2",
  "spi0.1:S4": "port3",

  "spi0.1:MA": "port4",
  "spi0.1:MB": "port5",
  "spi0.1:MC": "port6",
  "spi0.1:MD": "port7"}

var MOTOR_PATH string = "/sys/class/tacho-motor/"
var SENSOR_PATH string = "/sys/class/lego-sensor/"
var BATTERY_PATH string = "/sys/class/power_supply/brickpi3-battery/"
var LEGO_PORT string = "/sys/class/lego-port/"

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func read(path string) string {
  dat, err := ioutil.ReadFile(path)
  check(err)
  return string(dat)
}

func readBytes(path string) []byte {
  dat, err := ioutil.ReadFile(path)
  check(err)
  return dat
}

func write(path string, data string) {
  dat := []byte(data)
  err := ioutil.WriteFile(path, dat, 0644)
  check(err)
}

func list(path string) []string {
  files, _ := ioutil.ReadDir(path)
  stringFiles := make([]string, len(files))
  for i, file := range files {
    stringFiles[i] = file.Name()
  }
  return stringFiles
}

func exists(path string) bool {
  _, err := os.Stat(path)
  return os.IsNotExist(err)
}

type Device struct {
  path string
}

func (device Device) get(attribute string) string {
  return read(device.path + attribute)
}

func (device Device) set(attribute string, data string) {
  write(device.path + attribute, data)
}

func (device Device) bytes(file string) []byte {
  return readBytes(device.path + file)
}

type IndexedDevice struct {
  port string
  path string
}

func (device *IndexedDevice) findDeviceFromPort() {
  files := list(device.path)
  foundDevice := false

  for _, file := range files {
    if strings.Contains(read(device.path + "/" + file + "/address"), device.port) {
      device.path = device.path + "/" + file + "/"
      foundDevice = true
      break
    }
  }

  if !foundDevice {
    check(errors.New("no device in address " + device.port))
  }
}

func (device IndexedDevice) get(attribute string) string {
  return read(device.path + attribute)
}

func (device IndexedDevice) set(attribute string, data string) {
  write(device.path + attribute, data)
}

type ManualDevice struct {
  port string
  path string

  connection string
  name string
}

func (device *ManualDevice) findDeviceFromPort() {
  files := list(device.path)
  foundInitialDevice := false

  for _, file := range files {
    if strings.Contains(read(device.path + "/" + file + "/address"), device.port) {
      device.path = device.path + "/" + file + "/"
      foundInitialDevice = true
      break
    }
  }

  if !foundInitialDevice {
    portPath := LEGO_PORT + PORT[device.port] + "/"
    write(portPath + "mode", device.connection)
    write(portPath + "set_device", device.name)

    files := list(device.path)
    foundDevice := false

    for _, file := range files {
      if strings.Contains(read(device.path + "/" + file + "/address"), device.port) {
        device.path = device.path + "/" + file + "/"
        foundDevice = true
        break
      }
    }

    if !foundDevice {
      check(errors.New("no device in address " + device.port))
    }
  }
}

func (device ManualDevice) get(attribute string) string {
  return read(device.path + attribute)
}

func (device ManualDevice) set(attribute string, data string) {
  write(device.path + attribute, data)
}

type Sensor struct {
  port string
  path string
  connection string
  name string

  manualDevice ManualDevice
}

func (sensor Sensor) new() Sensor {
  sensor.path = SENSOR_PATH
  sensor.manualDevice = ManualDevice{path: sensor.path, port: sensor.port, connection: sensor.connection, name: sensor.name}
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
  port string

  REFLECT string
  COLOR string
  RGB string

  sensor Sensor
}

func (colorSensor ColorSensor) new() ColorSensor {
  colorSensor.sensor = Sensor{port: colorSensor.port, connection: "ev3-uart", name: "lego-ev3-color"}.new()

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

func (colorSensor ColorSensor) rgbIntensity() int {
  return (colorSensor.sensor.value("0") + colorSensor.sensor.value("1") + colorSensor.sensor.value("2")) / 3
}

type UltrasonicSensor struct {
  port string

  US_DIST_CM string

  sensor Sensor
}

func (ultrasonicSensor UltrasonicSensor) new() UltrasonicSensor {
  ultrasonicSensor.sensor = Sensor{port: ultrasonicSensor.port, connection: "ev3-uart", name: "lego-ev3-us"}.new()

  ultrasonicSensor.US_DIST_CM = "US-DIST-CM"

  return ultrasonicSensor
}

func (ultrasonicSensor UltrasonicSensor) distance() int {
  return ultrasonicSensor.sensor.value("0")
}

func (ultrasonicSensor UltrasonicSensor) mode(newMode string) {
  ultrasonicSensor.sensor.mode(newMode)
}

type Motor struct {
  port string

  path string
  indexedDevice IndexedDevice
}

func (motor Motor) new() Motor {
  motor.path = MOTOR_PATH
  motor.indexedDevice = IndexedDevice{path: motor.path, port: motor.port}
  motor.indexedDevice.findDeviceFromPort()
  return motor
}

func (motor Motor) runForever(speed int) {
  speedString := strconv.Itoa(speed)
  motor.indexedDevice.set("speed_sp", speedString)
  motor.indexedDevice.set("command", "run-forever")
}

func (motor Motor) runToPosition(speed int, position int) {
  speedString := strconv.Itoa(speed)
  positionString := strconv.Itoa(position)
  motor.indexedDevice.set("speed_sp", speedString)
  motor.indexedDevice.set("position_sp", positionString)
  motor.indexedDevice.set("command", "run-to-abs-pos")
}

func (motor Motor) runToRelativePosition(speed int, position int) {
  speedString := strconv.Itoa(speed)
  positionString := strconv.Itoa(position)
  motor.indexedDevice.set("speed_sp", speedString)
  motor.indexedDevice.set("position_sp", positionString)
  motor.indexedDevice.set("command", "run-to-rel-pos")
}

func (motor Motor) runTimed(speed int, time int) {
  speedString := strconv.Itoa(speed)
  timeString := strconv.Itoa(time)
  motor.indexedDevice.set("speed_sp", speedString)
  motor.indexedDevice.set("time_sp", timeString)
  motor.indexedDevice.set("command", "run-timed")
}

func (motor Motor) stop() {
  motor.indexedDevice.set("command", "stop")
}

func (motor Motor) state() []string {
  states := motor.indexedDevice.get("state")
  return strings.Split(states, " ")
}

type DriveMotors struct {
  portLeft string
  portRight string

  motorLeft Motor
  motorRight Motor
}

func (driveMotors DriveMotors) new() DriveMotors {
  driveMotors.motorLeft = Motor{port: driveMotors.portLeft}.new()

  driveMotors.motorRight = Motor{port: driveMotors.portRight}.new()
  return driveMotors
}

func (driveMotors DriveMotors) runForever(speed int) {
  speedString := strconv.Itoa(speed)
  driveMotors.motorLeft.indexedDevice.set("speed_sp", speedString)
  driveMotors.motorRight.indexedDevice.set("speed_sp", speedString)

  driveMotors.motorLeft.indexedDevice.set("command", "run-forever")
  driveMotors.motorRight.indexedDevice.set("command", "run-forever")
}

func (driveMotors DriveMotors) runRatioForever(ratio []int, speed int) {
  leftSpeed := ratio[0] * speed
  if leftSpeed > 1000 { leftSpeed = 1000 }
  rightSpeed := ratio[1] * speed
  if rightSpeed > 1000 { rightSpeed = 1000 }

  rightSpeedString := strconv.Itoa(leftSpeed)
  leftSpeedString := strconv.Itoa(rightSpeed)

  driveMotors.motorLeft.indexedDevice.set("speed_sp", rightSpeedString)
  driveMotors.motorRight.indexedDevice.set("speed_sp", leftSpeedString)

  driveMotors.motorLeft.indexedDevice.set("command", "run-forever")
  driveMotors.motorRight.indexedDevice.set("command", "run-forever")
}

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