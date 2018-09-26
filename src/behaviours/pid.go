package main

import "math"
// import "fmt"
import "strconv"

const MOTOR_RATE = 2.8
const MOTOR_CURVE = 0.6

const PROPORTIONAL = 9 // 9
const INTEGRAL = 0.40 // 0.40
const DOUBLE_INTEGRAL = 0.0
const DERIVATIVE = 12.0 // 12.0
const DOUBLE_DERIVATIVE = 15.0 // 15.0

const BASE_SPEED = 400 // 400
const INTEGRAL_SLOW = 190.0 // 200.0
const DERIVATIVE_SLOW = 150.0 // 200.0

var lastError = 0.0
var lastDerivative = 0.0
var integral = 0.0
var doubleIntegral = 0.0
var baseSpeed = 0

func PID() string {
  currentError := float64(LineSensorError())

  integral += currentError //* (1 / LOOP_SPEED)
  doubleIntegral += integral * (1.0 / LOOP_SPEED)
  derivative := currentError - lastError
  doubleDerivative := derivative - lastDerivative

  motorSpeed := (PROPORTIONAL * currentError) + (INTEGRAL * integral) + (DOUBLE_INTEGRAL * doubleIntegral) + (DERIVATIVE * derivative) + (DOUBLE_DERIVATIVE * doubleDerivative)
  motorSpeed = (motorSpeed + motorSpeed * ((math.Abs(motorSpeed) - 1000) / 1000) * MOTOR_CURVE) * MOTOR_RATE

  lastError = currentError;
  lastDerivative = derivative;

  baseSpeed = BASE_SPEED - int((math.Abs(integral) / 350.0) * INTEGRAL_SLOW) - int((math.Abs(derivative) / 20.0) * DERIVATIVE_SLOW)

  go bot.motorLeft.RunForever(min(max(baseSpeed + int(motorSpeed), -1000), 1000))
  go bot.motorRight.RunForever(min(max(baseSpeed - int(motorSpeed), -1000), 1000))

  if BEHAVIOUR == "follow_line" {
    BehaviourDebug("following line with pid")
    return "follow_line:follow"
  }

  if STATE(":follow") {
    // BehaviourTrace("using pid to follow line")
    BehaviourTrace("r:" + strconv.Itoa(bot.colorSensorRight.RgbIntensity()) + ", l: " + strconv.Itoa(bot.colorSensorLeft.RgbIntensity()))
    // BehaviourTrace("p: " + strconv.Itoa(int(currentError)) + ", i: " + strconv.Itoa(int(integral)) + ", 2i: " + strconv.Itoa(int(doubleIntegral)) + ", d: " + strconv.Itoa(int(derivative)) + ", 2d: " + strconv.Itoa(int(doubleDerivative)))
  }
  return BEHAVIOUR
}

func ResetPID() {
  lastError = 0.0
  lastDerivative = 0.0
  integral = 0.0
  doubleIntegral = 0.0
}