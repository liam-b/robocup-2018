package main

import "math"
// import "fmt"
// import "strconv"

const KS = 5000.0
const KE = 85.0

const KP = 7.0
const KI = 0.32
const KD = 11.0
const BASE_SPEED = 270

var lastError = 0.0
var integral = 0.0

func PID() string {
  currentError := LineSensorError()
  currentError += (currentError * math.Abs(currentError)) / KE

  integral = integral + currentError;
  derivative := currentError - lastError

  // if int(currentError) > -3 && int(currentError) < 3 {
  //   integral = 0
  // }

  motorSpeed := (KP * currentError) + (KI * integral) + (KD * derivative)
  // motorSpeed += (motorSpeed * math.Abs(motorSpeed)) / KS
  lastError = currentError;

  leftMotorSpeed := min(max(BASE_SPEED + int(motorSpeed), -1000), 1000)
  rightMotorSpeed := min(max(BASE_SPEED - int(motorSpeed), -1000), 1000)

  go bot.motorLeft.RunForever(leftMotorSpeed)
  go bot.motorRight.RunForever(rightMotorSpeed)

  if BEHAVIOUR == "follow_line" {
    BehaviourDebug("following line with pid")
    return "follow_line:follow"
  }

  if STATE(":follow") {
    BehaviourTrace("using pid to follow line")
  }
  return BEHAVIOUR
}
