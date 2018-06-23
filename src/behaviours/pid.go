package main

// base
const KP = 4.2
const KI = 2.0
const KD = 7.0
const BASE_SPEED = 400

var lastError = 0.0
var integral = 0.0

func min(x, y int) int {
  if x < y { return x }
  return y
}

func max(x, y int) int {
  if x > y { return x }
  return y
}

func PID() string {
  currentError := float64(bot.colorSensorLeft.RgbIntensity() - bot.colorSensorRight.RgbIntensity())
  integral := integral + currentError;
  derivative := currentError - lastError;

  motorSpeed := (KP * currentError) + (KI * integral) + (KD * derivative);
  lastError = currentError;

  leftMotorSpeed := min(max(BASE_SPEED + int(motorSpeed), -1000), 1000);
  rightMotorSpeed := min(max(BASE_SPEED - int(motorSpeed), -1000), 1000);

  go bot.motorLeft.RunForever(leftMotorSpeed)
  go bot.motorRight.RunForever(rightMotorSpeed)

  return BEHAVIOUR
}