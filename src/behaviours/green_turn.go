package main

const GREEN_FINISH_ANGLE = 700
const GREEN_COOLDOWN = 10

var greenCooldown = 0

func TurnOnGreen() string {
  if BEHAVIOUR == "turn_green:left" {
    if GyroTurnedToAngle(GREEN_FINISH_ANGLE, LEFT) {
      greenCooldown = 0
      return "turn_green:cooldown"
    }
    OneSensorLineFollowing(LEFT)
  }

  if BEHAVIOUR == "turn_green:right" {
    if GyroTurnedToAngle(-GREEN_FINISH_ANGLE, RIGHT) {
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