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

  motorLefteft Motor
  motorRightight Motor
}

func (driveMotors DriveMotors) New() DriveMotors {
  driveMotors.motorLefteft = Motor{Port: driveMotors.PortLeft}.New()

  driveMotors.motorRightight = Motor{Port: driveMotors.PortRight}.New()
  return driveMotors
}

func (driveMotors DriveMotors) RunForever(speed int) {
  speedString := strconv.Itoa(speed)
  driveMotors.motorLefteft.indexedDevice.set("speed_sp", speedString)
  driveMotors.motorRightight.indexedDevice.set("speed_sp", speedString)

  driveMotors.motorLefteft.indexedDevice.set("command", "run-forever")
  driveMotors.motorRightight.indexedDevice.set("command", "run-forever")
}

func (driveMotors DriveMotors) RunRatioForever(ratio []int, speed int) {
  leftSpeed := ratio[0] * speed
  if leftSpeed > 1000 { leftSpeed = 1000 }
  rightSpeed := ratio[1] * speed
  if rightSpeed > 1000 { rightSpeed = 1000 }

  rightSpeedString := strconv.Itoa(leftSpeed)
  leftSpeedString := strconv.Itoa(rightSpeed)

  driveMotors.motorLefteft.indexedDevice.set("speed_sp", rightSpeedString)
  driveMotors.motorRightight.indexedDevice.set("speed_sp", leftSpeedString)

  driveMotors.motorLefteft.indexedDevice.set("command", "run-forever")
  driveMotors.motorRightight.indexedDevice.set("command", "run-forever")
}