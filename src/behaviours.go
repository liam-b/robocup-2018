package main

import "strings"
import "./io"

var behavioursFunctions = map[string]func()string{
  "chemical_spill": SaveCan,
  "water_tower": AvoidWaterTower,
  "turn_green": TurnOnGreen,
  "follow_line": FollowLine}

var behaviourLeds = map[string][3]int{
  "chemical_spill": io.COLOR_BLUE,
  "water_tower": io.COLOR_RED,
  "turn_green": io.COLOR_GREEN,
  "follow_line": io.COLOR_WHITE}

func Behave() {
  if BEHAVIOUR == "follow_line" {
    if DO_GREEN_TURN && DetectedGreen(LEFT) { BEHAVIOUR = "turn_green:left" }
    if DO_GREEN_TURN && DetectedGreen(RIGHT) { BEHAVIOUR = "turn_green:right" }
    if DO_WATER_TOWER && DetectedWaterTower(WATER_TOWER_DETECT_DISTANCE, WATER_TOWER_DETECT_COUNT) { BEHAVIOUR = "water_tower:start" }
    if DO_CHEMICAL_SPILL && DetectedSilver() { BEHAVIOUR = "chemical_spill:start" }
  }

  BEHAVIOUR = behavioursFunctions[strings.Split(BEHAVIOUR, ":")[0]]()
  bot.ledshim.SetPixel(io.BEHAVIOUR_PIXEL, behaviourLeds[BEHAVIOUR])
}

func FollowLine() string {
  intensityLeft := bot.colorSensorLeft.RgbIntensity()
  intensityRight := bot.colorSensorRight.RgbIntensity()

  if (intensityLeft < FOLLOW_LINE_HARD_TURN_VALUE) {
    go bot.motorRight.RunForever(SpeedRatio(FOLLOW_LINE_SPEED, FOLLOW_LINE_SOFT_TURN_RATIO, FAST))
    go bot.motorLeft.RunForever(SpeedRatio(FOLLOW_LINE_SPEED, FOLLOW_LINE_SOFT_TURN_RATIO, SLOW) - FOLLOW_LINE_SPEED)
  } else if (intensityRight < FOLLOW_LINE_HARD_TURN_VALUE) {
    go bot.motorRight.RunForever(SpeedRatio(FOLLOW_LINE_SPEED, FOLLOW_LINE_SOFT_TURN_RATIO, SLOW) - FOLLOW_LINE_SPEED)
    go bot.motorLeft.RunForever(SpeedRatio(FOLLOW_LINE_SPEED, FOLLOW_LINE_SOFT_TURN_RATIO, FAST))
  } else if (intensityLeft < FOLLOW_LINE_SOFT_TURN_VALUE) {
    go bot.motorRight.RunForever(SpeedRatio(FOLLOW_LINE_SPEED, FOLLOW_LINE_SOFT_TURN_RATIO, FAST))
    go bot.motorLeft.RunForever(SpeedRatio(FOLLOW_LINE_SPEED, FOLLOW_LINE_SOFT_TURN_RATIO, SLOW))
  } else if (intensityRight < FOLLOW_LINE_SOFT_TURN_VALUE) {
    go bot.motorRight.RunForever(SpeedRatio(FOLLOW_LINE_SPEED, FOLLOW_LINE_SOFT_TURN_RATIO, SLOW))
    go bot.motorLeft.RunForever(SpeedRatio(FOLLOW_LINE_SPEED, FOLLOW_LINE_SOFT_TURN_RATIO, FAST))
  } else {
    go bot.motorRight.RunForever(FOLLOW_LINE_SPEED)
    go bot.motorLeft.RunForever(FOLLOW_LINE_SPEED)
  }

  return BEHAVIOUR
}

func OneSensorLineFollowing(sensor int) string {
  redLeft, greenLeft, blueLeft := bot.colorSensorLeft.Rgb()
  redRight, greenRight, blueRight := bot.colorSensorRight.Rgb()

  if (sensor == LEFT && greenLeft < redLeft + FOLLOW_LINE_GREEN_DIFFERENCE && greenLeft < blueLeft + FOLLOW_LINE_GREEN_DIFFERENCE) {
    go bot.motorRight.RunForever(SpeedRatio(FOLLOW_LINE_SPEED, FOLLOW_LINE_SOFT_TURN_RATIO, FAST))
    go bot.motorLeft.RunForever(SpeedRatio(FOLLOW_LINE_SPEED, FOLLOW_LINE_SOFT_TURN_RATIO, SLOW) - FOLLOW_LINE_SPEED)
  } else if (sensor == RIGHT && greenRight < redRight + FOLLOW_LINE_GREEN_DIFFERENCE && greenRight < blueRight + FOLLOW_LINE_GREEN_DIFFERENCE) {
    go bot.motorRight.RunForever(SpeedRatio(FOLLOW_LINE_SPEED, FOLLOW_LINE_SOFT_TURN_RATIO, SLOW) - FOLLOW_LINE_SPEED)
    go bot.motorLeft.RunForever(SpeedRatio(FOLLOW_LINE_SPEED, FOLLOW_LINE_SOFT_TURN_RATIO, FAST))
  } else {
    go bot.motorRight.RunForever(FOLLOW_LINE_SPEED)
    go bot.motorLeft.RunForever(FOLLOW_LINE_SPEED)
  }

  return BEHAVIOUR
}

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

var chemicalSpillVerifyAttempts = 0

func SaveCan() string {
  if BEHAVIOUR == "chemical_spill:start" {
    chemicalSpillVerifyAttempts = 0
    go bot.motorRight.RunForever(SAVE_CAN_VERIFY_SPEED)
    go bot.motorLeft.RunForever(SAVE_CAN_VERIFY_SPEED)
    return "water_tower:verify"
  }

  if BEHAVIOUR == "chemical_spill:verify" {
    chemicalSpillVerifyAttempts += 1
    if chemicalSpillVerifyAttempts > SAVE_CAN_VERIFY_ATTEMPTS {
      chemicalSpillVerifyAttempts = 0
      go bot.motorRight.Stop()
      go bot.motorLeft.Stop()
      return "follow_line"
    }

    left, right := GetColors()
    if left == GREEN && right == GREEN {
      go bot.motorRight.Stop()
      go bot.motorLeft.Stop()
      return "chemical_spill:enter"
    }
  }

  return BEHAVIOUR
}