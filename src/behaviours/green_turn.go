package main

const GREEN_FINISH_ANGLE = 700
const GREEN_OVERRIDE_COUNT = 80
const GREEN_COOLDOWN = 10

var greenOverride = 0
var greenCooldown = 0

func TurnOnGreen() string {
  if MINOR(":start") {
    greenOverride = 0
    return "turn_green:turn" // FIXME: won't pass params
  }

  if MINOR(":turn") {
    greenOverride += 1
    if greenOverride > GREEN_OVERRIDE_COUNT {
      greenOverride = 0
      return "follow_line"
    }

    if PARAM(".left") {
      if GyroTurnedToAngle(GREEN_FINISH_ANGLE, LEFT) {
        greenCooldown = 0
        return "turn_green:cooldown"
      }
      OneSensorLineFollowing(LEFT)
    }

    if PARAM(".right") {
      if GyroTurnedToAngle(-GREEN_FINISH_ANGLE, RIGHT) {
        greenCooldown = 0
        return "turn_green:cooldown"
      }
      OneSensorLineFollowing(RIGHT)
  }

  if MINOR(":cooldown") {
    greenCooldown += 1
    if greenCooldown > GREEN_COOLDOWN {
      greenCooldown = 0
      return "follow_line"
    }
    FollowLine()
  }

  return BEHAVIOUR
}