package main

import "math"
// import "fmt"
import "strconv"

const ERROR_CURVE = 100.0

const PROPORTIONAL = 5.0
const INTEGRAL = 0.0
const DOUBLE_INTEGRAL = 0.0
const DERIVATIVE = 0.0
const DOUBLE_DERIVATIVE = 0.0

const BASE_SPEED = 270

var lastError = 0.0
var lastDerivative = 0.0
var integral = 0.0
var doubleIntegral = 0.0

func PID() string {
  currentError := float64(LineSensorError())
  currentError += (currentError * math.Abs(currentError)) / ERROR_CURVE

  integral += currentError * (1 / LOOP_SPEED)
  doubleIntegral += integral * (1 / LOOP_SPEED)
  derivative := currentError - lastError
  doubleDerivative := derivative - lastDerivative

  motorSpeed := (PROPORTIONAL * currentError) + (INTEGRAL * integral) + (DOUBLE_INTEGRAL * doubleIntegral) + (DERIVATIVE * derivative) + (DOUBLE_DERIVATIVE * doubleDerivative)

  lastError = currentError;
  lastDerivative = derivative;

  go bot.motorLeft.RunForever(min(max(BASE_SPEED + int(motorSpeed), -1000), 1000))
  go bot.motorRight.RunForever(min(max(BASE_SPEED - int(motorSpeed), -1000), 1000))

  if BEHAVIOUR == "follow_line" {
    BehaviourDebug("following line with pid")
    return "follow_line:follow"
  }

  if STATE(":follow") {
    // BehaviourTrace("using pid to follow line")
    BehaviourTrace("p: " + strconv.Itoa(int(currentError)) + ", i: " + strconv.Itoa(int(integral)) + ", 2i: " + strconv.Itoa(int(doubleIntegral)) + ", d: " + strconv.Itoa(int(derivative)) + ", 2d: " + strconv.Itoa(int(doubleDerivative)))
  }
  return BEHAVIOUR
}

func ResetPID() {
  lastError = 0.0
  lastDerivative = 0.0
  integral = 0.0
  doubleIntegral = 0.0
}