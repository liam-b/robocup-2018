package main

var totalAngle = 0

func GyroTurnedToAngle(angle int, turnDirection int) bool {
  totalAngle += bot.imu.ReadGyro()

  if turnDirection == LEFT && totalAngle > angle {
    totalAngle = 0
    return true
  }

  if turnDirection == RIGHT && totalAngle < angle {
    totalAngle = 0
    return true
  }

  return false
}