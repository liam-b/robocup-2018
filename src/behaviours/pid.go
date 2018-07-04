package main

import "math"

const KP = 7.5
const KI = 0.0
const KD = 8.0
const BASE_SPEED = 300

var lastError = 0.0
var integral = 0.0

func PID() string {
  currentError := colorSensorError()
  currentError += currentError * (math.Abs(currentError) / 50)
  
  integral := integral + currentError;
  derivative := currentError - lastError

  motorSpeed := (KP * currentError) + (KI * integral) + (KD * derivative);
  lastError = currentError;

  leftMotorSpeed := min(max(BASE_SPEED + int(motorSpeed), -1000), 1000);
  rightMotorSpeed := min(max(BASE_SPEED - int(motorSpeed), -1000), 1000);

  go bot.motorLeft.RunForever(leftMotorSpeed)
  go bot.motorRight.RunForever(rightMotorSpeed)

  return BEHAVIOUR
}

func colorSensorError() float64 {
  return float64(bot.colorSensorLeft.RgbIntensity() - bot.colorSensorRight.RgbIntensity())
}
