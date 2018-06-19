package main

import "strings"

import "./io"

func Behave() {
  if BEHAVIOUR == "follow_line" && DO_GREEN_TURN && DetectedGreen(LEFT) { BEHAVIOUR = "turn_green:left" }
  if BEHAVIOUR == "follow_line" && DO_GREEN_TURN && DetectedGreen(RIGHT) { BEHAVIOUR = "turn_green:right" }
  if BEHAVIOUR == "follow_line" && DO_WATER_TOWER && DetectedWaterTower(WATER_TOWER_DETECT_DISTANCE, WATER_TOWER_DETECT_COUNT) { BEHAVIOUR = "water_tower:start" }
  // if BEHAVIOUR == "follow_line" && DO_CHEMICAL_SPILL && DetectedSilver() { BEHAVIOUR = "chemical_spill:verify" }

  // if modeContains("chemical_spill") { BEHAVIOUR = SaveCan() }
  if modeContains("water_tower") { BEHAVIOUR = AvoidWaterTower() }
  if modeContains("turn_green") { BEHAVIOUR = TurnOnGreen() }
  if modeContains("follow_line") { BEHAVIOUR = FollowLine(true, true) }

  if modeContains("water_tower") { bot.ledshim.SetPixel(io.BEHAVIOUR_PIXEL, io.COLOR_RED) }
  if modeContains("turn_green") { bot.ledshim.SetPixel(io.BEHAVIOUR_PIXEL, io.COLOR_GREEN) }
  if modeContains("follow_line") { bot.ledshim.SetPixel(io.BEHAVIOUR_PIXEL, io.COLOR_WHITE) }
}

func FollowLine(useLeftSensor bool, useRightSensor bool) string {
  intensityLeft := bot.colorSensorLeft.RgbIntensity()
  intensityRight := bot.colorSensorRight.RgbIntensity()

  if !useLeftSensor { intensityLeft = 30 }
  if !useRightSensor { intensityRight = 30 }

  if (intensityLeft < 20) {
    go bot.motorRight.RunForever(200)
    go bot.motorLeft.RunForever(-100)
  } else if (intensityRight < 20) {
    go bot.motorRight.RunForever(-100)
    go bot.motorLeft.RunForever(200)
  } else {
    go bot.motorRight.RunForever(180)
    go bot.motorLeft.RunForever(180)
  }

  return BEHAVIOUR
}

func TurnOnGreen() string {
  if BEHAVIOUR == "turn_green:left" {
    if GyroAtAngle(GREEN_FINISH_ANGLE, LEFT) {
      return "follow_line"
    }
    FollowLine(true, false)
  }

  if BEHAVIOUR == "turn_green:right" {
    if GyroAtAngle(-GREEN_FINISH_ANGLE, RIGHT) {
      return "follow_line"
    }
    FollowLine(false, true)
  }

  return BEHAVIOUR
}

var verifyAttempts = 0

func AvoidWaterTower() string {
  if BEHAVIOUR == "water_tower:start" {
    go bot.motorRight.RunForever(WATER_TOWER_VERIFY_SPEED)
    go bot.motorLeft.RunForever(WATER_TOWER_VERIFY_SPEED)
    verifyAttempts = 0
    return "water_tower:verify"
  }

  if BEHAVIOUR == "water_tower:verify" {
    if (DetectedWaterTower(WATER_TOWER_VERIFY_DISTANCE, WATER_TOWER_VERIFY_COUNT)) {
      go bot.motorRight.RunForever(WATER_TOWER_TURN_SPEED)
      go bot.motorLeft.RunForever(-WATER_TOWER_TURN_SPEED)
      return "water_tower:turn"
    } else {
      verifyAttempts += 1
      if verifyAttempts > WATER_TOWER_VERIFY_ATTEMPTS {
        return "follow_line"
      }
    }
  }

  if BEHAVIOUR == "water_tower:turn" {
    if GyroAtAngle(WATER_TOWER_TURN_ANGLE, LEFT) {
      go bot.motorLeft.Stop()
      go bot.motorRight.Stop()
      // run motors at correct turn ratio
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
    if GyroAtAngle(WATER_TOWER_RECAPTURE_ANGLE, LEFT) {
      // motors: stop
      return "follow_line"
    }
    FollowLine(true, false)
  }

  return BEHAVIOUR
}

func SaveCan() string {
  if BEHAVIOUR == "chemical_spill:verify" {
    // motors: slow
    left, right := GetColors()
    if left == GREEN && right == GREEN {
      // motors: stop
      return "water_tower:enter"
    }
  }

  return BEHAVIOUR
}

func modeContains(mode string) bool {
  return strings.Contains(BEHAVIOUR, mode)
}