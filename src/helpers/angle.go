package main

// func UpdateGyroAngle(int gyroValue) {
//
// }

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

func GyroTotalRotation() int {
  totalAngle += bot.imu.ReadGyro()
  return totalAngle
}

func ResetGyroTotalRotation() {
  totalAngle = 0
}
