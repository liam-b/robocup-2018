package main

const GREEN_FINISH_ANGLE = 700
const GREEN_OVERRIDE_COUNT = 60
const GREEN_COOLDOWN = 20
const GREEN_BOTH_COUNT = 5
const GREEN_BOTH_COOLDOWN = 6

var greenOverride = 0
var greenCooldown = 0
var greenBothCount = 0
var greenBothCooldown = 0

func TurnOnGreen() string {
  if STATE(":start") {
    BehaviourDebug("starting " + log.state(":turn"))
    greenOverride = 0
    greenBothCount = 0
    greenBothCooldown = 0
    bot.imu.ResetGyro()
    return "turn_green:turn" + PARAMS()
  }

  if STATE(":both") {
    greenBothCooldown += 1
    left, right := GetColors()
    if left == GREEN && right == GREEN {
      greenBothCount += 1
      if greenBothCount > GREEN_BOTH_COUNT {
        BehaviourDebug("found both sensors on green, switching to " + log.state("chemical_spill:verify"))
        return "chemical_spill:verify"
      }
    }

    if greenBothCooldown > GREEN_BOTH_COOLDOWN {
      return "follow_line"
    }
  }

  if STATE(":turn") {
    BehaviourTrace("turning on green with direction " + log.state(PARAMS()))
    greenOverride += 1
    if greenOverride > GREEN_OVERRIDE_COUNT {
      BehaviourDebug("green count override reached, returning to " + log.state("follow_line"))
      greenOverride = 0
      return "follow_line"
    }

    left, right := GetColors()
    if left == GREEN && right == GREEN {
      go bot.motorRight.Stop()
      go bot.motorLeft.Stop()
      BehaviourDebug("found both sensors on green, switching to " + log.state(":both"))
      return "turn_green:both"
    }

    if PARAM(".left") {
      if bot.imu.GyroValue() > GREEN_FINISH_ANGLE {
        BehaviourDebug("going into " + log.state(":cooldown") + " after turn")
        greenCooldown = 0
        integral = 0
        return "turn_green:cooldown"
      }
      OneSensorLineFollowing(LEFT)
    }

    if PARAM(".right") {
      if bot.imu.GyroValue() < -GREEN_FINISH_ANGLE {
        BehaviourDebug("going into " + log.state(":cooldown") + " after turn")
        greenCooldown = 0
        integral = 0
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
