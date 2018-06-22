package main

var waterTowerVerifyAttempts = 0

func AvoidWaterTower() string {
  if BEHAVIOUR == "water_tower:start" {
    go bot.motorRight.RunForever(WATER_TOWER_VERIFY_SPEED)
    go bot.motorLeft.RunForever(WATER_TOWER_VERIFY_SPEED)
    waterTowerVerifyAttempts = 0
    return "water_tower:verify"
  }

  if BEHAVIOUR == "water_tower:verify" {
    if (DetectedWaterTower(WATER_TOWER_VERIFY_DISTANCE, WATER_TOWER_VERIFY_COUNT)) {
      go bot.motorRight.RunForever(WATER_TOWER_TURN_SPEED)
      go bot.motorLeft.RunForever(-WATER_TOWER_TURN_SPEED)
      return "water_tower:turn"
    } else {
      waterTowerVerifyAttempts += 1
      if waterTowerVerifyAttempts > WATER_TOWER_VERIFY_ATTEMPTS {
        return "follow_line"
      }
    }
  }

  if BEHAVIOUR == "water_tower:turn" {
    if GyroTurnedToAngle(WATER_TOWER_TURN_ANGLE, LEFT) {
      go bot.motorLeft.RunForever(int(float64(WATER_TOWER_AVOID_SPEED) * 1.0))
      go bot.motorRight.RunForever(int(float64(WATER_TOWER_AVOID_SPEED) * WATER_TOWER_AVOID_RATIO))
      return "water_tower:avoid"
    }
  }

  if BEHAVIOUR == "water_tower:avoid" {
    color, _ :=  GetColors()
    if color == BLACK {
      return "water_tower:recapture"
    }
  }

  if BEHAVIOUR == "water_tower:recapture" {
    if GyroTurnedToAngle(WATER_TOWER_RECAPTURE_ANGLE, LEFT) {
      return "follow_line"
    }
    FollowLine()
  }

  return BEHAVIOUR
}