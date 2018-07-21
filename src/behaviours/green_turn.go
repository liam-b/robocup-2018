package main

const GREEN_FINISH_ANGLE = 700
const GREEN_OVERRIDE_COUNT = 60
const GREEN_COOLDOWN = 20

var greenOverride = 0
var greenCooldown = 0

func TurnOnGreen() string {
  if STATE(":start") {
    BehaviourDebug("starting " + log.state(":turn"))
    greenOverride = 0
    bot.imu.ResetGyro()
    return "turn_green:turn" + PARAMS()
  }

  if STATE(":turn") {
    BehaviourTrace("turning on green with direction " + log.state(PARAMS()))
    greenOverride += 1
    if greenOverride > GREEN_OVERRIDE_COUNT {
      BehaviourDebug("green count override reached, returning to " + log.state("follow_line"))
      greenOverride = 0
      return "follow_line"
    }

    if PARAM(".left") {
      if bot.imu.GyroValue() > GREEN_FINISH_ANGLE {
        BehaviourDebug("going into " + log.state(":cooldown") + " after turn")
        greenCooldown = 0
        return "turn_green:cooldown"
      }
      OneSensorLineFollowing(LEFT)
    }

    if PARAM(".right") {
      if bot.imu.GyroValue() < -GREEN_FINISH_ANGLE {
        BehaviourDebug("going into " + log.state(":cooldown") + " after turn")
        greenCooldown = 0
        return "turn_green:cooldown"
      }
      OneSensorLineFollowing(RIGHT)
    }
  }

  if STATE(":cooldown") {
    BehaviourTrace("following line with pid while waiting on cooldown")
    greenCooldown += 1
    if greenCooldown > GREEN_COOLDOWN {
      BehaviourDebug("cooldown count reached, returning to normal " + log.state("follow_line"))
      greenCooldown = 0
      return "follow_line"
    }
    PID()
  }

  return BEHAVIOUR
}
