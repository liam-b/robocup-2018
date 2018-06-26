package main

func SpeedRatio(speed int, ratio float64, sign int) int {
  return speed + int(float64(speed) * ratio * float64(sign))
}

func ResetHelpers() {
  liftedMatches = 0
  waterTowerMatches = 0
  totalAngle = 0
}

func min(x, y int) int {
  if x < y { return x }
  return y
}

func max(x, y int) int {
  if x > y { return x }
  return y
}