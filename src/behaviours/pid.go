package main

const KP = 3.0
const KI = 0.0
const KD = 20.0
const BASE_SPEED = 300

var lastError = 0.0
var integral = 0.0

func min(x, y int) int {
  if x < y {
    return x
  }
  return y
}

func max(x, y int) int {
  if x > y {
    return x
  }
  return y
}

func PID() string {
  currentError := float64(bot.colorSensorLeft.RgbIntensity() - bot.colorSensorRight.RgbIntensity())
  proportional := currentError
  integral := integral + currentError;
  derivative := currentError - lastError;

  motorSpeed := (KP * proportional) + (KI * integral) + (KD * derivative);
  lastError = currentError;

  leftMotorSpeed := min(max(BASE_SPEED + int(motorSpeed), -1000), 1000);
  rightMotorSpeed := min(max(BASE_SPEED - int(motorSpeed), -1000), 1000);

  go bot.motorLeft.RunForever(leftMotorSpeed)
  go bot.motorRight.RunForever(rightMotorSpeed)

  return BEHAVIOUR
}