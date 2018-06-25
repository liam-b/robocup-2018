package main

const GREEN_FINISH_ANGLE = 700
const GREEN_OVERRIDE_COUNT = 80
const GREEN_COOLDOWN = 10

var greenOverride = 0
var greenCooldown = 0

func TurnOnGreen() string {
  if BEHAVIOUR == "turn_green:left" || BEHAVIOUR == "turn_green:right" {
    greenOverride += 1
    if greenOverride > GREEN_OVERRIDE_COUNT {
      greenOverride = 0
      return "follow_line"
    }
  }

  if BEHAVIOUR == "turn_green:left" {
    if GyroTurnedToAngle(GREEN_FINISH_ANGLE, LEFT) {
      greenOverride = 0
      greenCooldown = 0
      return "turn_green:cooldown"
    }
    OneSensorLineFollowing(LEFT)
  }

  if BEHAVIOUR == "turn_green:right" {
    if GyroTurnedToAngle(-GREEN_FINISH_ANGLE, RIGHT) {
      greenOverride = 0
      greenCooldown = 0
      return "turn_green:cooldown"
    }
    OneSensorLineFollowing(RIGHT)
  }

  if BEHAVIOUR == "turn_green:cooldown" {
    greenCooldown += 1
    if greenCooldown > GREEN_COOLDOWN {
      greenCooldown = 0
      return "follow_line"
    }
    FollowLine()
  }

  return BEHAVIOUR
}