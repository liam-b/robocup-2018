package main

const WATER_TOWER_DETECT_DISTANCE = 920
const WATER_TOWER_DETECT_COUNT = 4
const WATER_TOWER_VERIFY_DISTANCE = WATER_TOWER_DETECT_DISTANCE + 10
const WATER_TOWER_VERIFY_COUNT = 5
const WATER_TOWER_VERIFY_ATTEMPTS = 150
const WATER_TOWER_VERIFY_SPEED = 40
const WATER_TOWER_TURN_ANGLE = 900
const WATER_TOWER_TURN_SPEED = 100
const WATER_TOWER_AVOID_RATIO = 0.65
const WATER_TOWER_AVOID_SPEED = 300
const WATER_TOWER_RECAPTURE_SPEED = 250
const WATER_TOWER_RECAPTURE_ANGLE = 650

var waterTowerVerifyAttempts = 0

func AvoidWaterTower() string {
  if STATE(":start") {
    BehaviourDebug("starting " + log.state(":verify") + " of water tower")
    go bot.motorRight.RunForever(WATER_TOWER_VERIFY_SPEED)
    go bot.motorLeft.RunForever(WATER_TOWER_VERIFY_SPEED)
    waterTowerVerifyAttempts = 0
    return "water_tower:verify"
  }

  if STATE(":verify") {
    BehaviourTrace("verifying water tower with ultrasonic sensor")
    if (DetectedWaterTower(WATER_TOWER_VERIFY_DISTANCE, WATER_TOWER_VERIFY_COUNT)) {
      BehaviourDebug("starting " + log.state(":turn") + " to line up for " + log.state(":avoid"))
      go bot.motorRight.RunForever(WATER_TOWER_TURN_SPEED)
      go bot.motorLeft.RunForever(-WATER_TOWER_TURN_SPEED)
      bot.imu.ResetGyro()
      return "water_tower:turn"
    } else {
      waterTowerVerifyAttempts += 1
      if waterTowerVerifyAttempts > WATER_TOWER_VERIFY_ATTEMPTS {
        BehaviourDebug("water tower not verified within count, returning to " + log.state("follow_line"))
        return "follow_line"
      }
    }
  }

  if STATE(":turn") {
    BehaviourTrace("turning to line up for " + log.state(":avoid"))
    if bot.imu.GyroValue() > WATER_TOWER_TURN_ANGLE {
      BehaviourDebug("turned to correct gyro angle, starting to " + log.state(":avoid"))
      go bot.motorLeft.RunForever(int(float64(WATER_TOWER_AVOID_SPEED) * 1.0))
      go bot.motorRight.RunForever(int(float64(WATER_TOWER_AVOID_SPEED) * WATER_TOWER_AVOID_RATIO))
      return "water_tower:avoid"
    }
  }

  if STATE(":avoid") {
    BehaviourTrace("avoiding water tower")
    color, _ :=  GetColors()
    if color == BLACK {
      BehaviourDebug("found line after " + log.state(":avoid") + ", moving to " + log.state(":recapture"))
      go bot.motorRight.RunForever(WATER_TOWER_RECAPTURE_SPEED)
      go bot.motorLeft.RunForever(-int(WATER_TOWER_RECAPTURE_SPEED / 2))
      bot.imu.ResetGyro()
      return "water_tower:recapture"
    }
  }

  if STATE(":recapture") {
    BehaviourTrace("recapturing line")
    if bot.imu.GyroValue() > WATER_TOWER_RECAPTURE_ANGLE {
      BehaviourDebug("finished recapturing line, returning to " + log.state("follow_line"))
      return "follow_line"
    }
  }

  return BEHAVIOUR
}
