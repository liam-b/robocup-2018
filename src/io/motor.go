package main

import "strconv"
import "strings"

type Motor struct {
  port string

  path string
  indexedDevice IndexedDevice
}

func (motor Motor) init() Motor {
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

func (driveMotors DriveMotors) init() DriveMotors {
  driveMotors.motorLeft = Motor{port: driveMotors.portLeft}.init()

  driveMotors.motorRight = Motor{port: driveMotors.portRight}.init()
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