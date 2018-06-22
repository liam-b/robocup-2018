package io

import "strconv"
import "strings"

type Motor struct {
  Port string

  path string
  indexedDevice IndexedDevice
}

func (motor Motor) New() Motor {
  motor.path = MOTOR_PATH
  motor.indexedDevice = IndexedDevice{path: motor.path, port: motor.Port}
  motor.indexedDevice.findDeviceFromPort()
  return motor
}

func (motor Motor) RunForever(speed int) {
  speedString := strconv.Itoa(speed)
  motor.indexedDevice.set("speed_sp", speedString)
  motor.indexedDevice.set("command", "run-forever")
}

func (motor Motor) RunToPosition(speed int, position int) {
  speedString := strconv.Itoa(speed)
  positionString := strconv.Itoa(position)
  motor.indexedDevice.set("speed_sp", speedString)
  motor.indexedDevice.set("position_sp", positionString)
  motor.indexedDevice.set("command", "run-to-abs-pos")
}

func (motor Motor) RunToRelativePosition(speed int, position int) {
  speedString := strconv.Itoa(speed)
  positionString := strconv.Itoa(position)
  motor.indexedDevice.set("speed_sp", speedString)
  motor.indexedDevice.set("position_sp", positionString)
  motor.indexedDevice.set("command", "run-to-rel-pos")
}

func (motor Motor) RunTimed(speed int, time int) {
  speedString := strconv.Itoa(speed)
  timeString := strconv.Itoa(time)
  motor.indexedDevice.set("speed_sp", speedString)
  motor.indexedDevice.set("time_sp", timeString)
  motor.indexedDevice.set("command", "run-timed")
}

func (motor Motor) Stop() {
  motor.indexedDevice.set("command", "stop")
}

func (motor Motor) State() []string {
  states := motor.indexedDevice.get("state")
  return strings.Split(states, " ")
}

type DriveMotors struct {
  PortLeft string
  PortRight string

  motorLeft Motor
  motorRight Motor
}

func (driveMotors DriveMotors) New() DriveMotors {
  driveMotors.motorLeft = Motor{Port: driveMotors.PortLeft}.New()
  driveMotors.motorRight = Motor{Port: driveMotors.PortRight}.New()
  return driveMotors
}

func (driveMotors DriveMotors) Stop() {
  driveMotors.motorLeft.indexedDevice.set("command", "stop")
  driveMotors.motorRight.indexedDevice.set("command", "stop")
}

func (driveMotors DriveMotors) RunForever(speed int) {
  speedString := strconv.Itoa(speed)
  driveMotors.motorLeft.indexedDevice.set("speed_sp", speedString)
  driveMotors.motorRight.indexedDevice.set("speed_sp", speedString)

  driveMotors.motorLeft.indexedDevice.set("command", "run-forever")
  driveMotors.motorRight.indexedDevice.set("command", "run-forever")
}

func (driveMotors DriveMotors) RunRatioForever(speed int, ratio [2]int) {
  leftSpeed := speed + int(float64(speed) * float64(ratio[0]))
  rightSpeed := speed + int(float64(speed) * float64(ratio[1]))

  leftSpeedString := strconv.Itoa(leftSpeed)
  rightSpeedString := strconv.Itoa(rightSpeed)

  driveMotors.motorLeft.indexedDevice.set("speed_sp", rightSpeedString)
  driveMotors.motorRight.indexedDevice.set("speed_sp", leftSpeedString)

  driveMotors.motorLeft.indexedDevice.set("command", "run-forever")
  driveMotors.motorRight.indexedDevice.set("command", "run-forever")
}