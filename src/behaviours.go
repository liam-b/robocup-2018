package main

import "strings"

func Behave() {
  if !mode("water_tower") && DetectedWaterTower(WATER_TOWER_DETECT_DISTANCE, WATER_TOWER_DETECT_COUNT) { MODE = "water_tower" }

  if mode("water_tower") { MODE = AvoidWaterTower() }
  if mode("follow_line") { MODE = FollowLine() }
}

func FollowLine() string {
  intensityL := bot.colorSensorLeft.Intensity()
  intensityR := bot.colorSensorRight.Intensity()

  if (intensityL < 72) {
    go bot.motorRight.RunForever(440)
    go bot.motorLeft.RunForever(230)
  }

  if (intensityR < 72) {
    go bot.motorRight.RunForever(230)
    go bot.motorLeft.RunForever(440)
  }

  if (intensityL < 16) {
    go bot.motorRight.RunForever(500)
    go bot.motorLeft.RunForever(80)
  }

  if (intensityR < 16) {
    go bot.motorRight.RunForever(80)
    go bot.motorLeft.RunForever(500)
  }

  if (intensityR > 60 && intensityL > 60) {
    go bot.motorRight.RunForever(300)
    go bot.motorLeft.RunForever(300)
  }

  return MODE
}

var verifyAttempts = 0

func AvoidWaterTower() string {
  if MODE == "water_tower" {
    log.debug("start moving forwards")
    verifyAttempts = 0
    return "water_tower:verify"
  }

  if MODE == "water_tower:verify" {
    log.debug("verifying water tower")
    if (DetectedWaterTower(WATER_TOWER_VERIFY_DISTANCE, WATER_TOWER_VERIFY_COUNT)) {
      log.debug("stopping motors, verified water tower")
      return "water_tower:avoid"
    } else {
      verifyAttempts += 1
      if verifyAttempts > WATER_TOWER_VERIFY_ATTEMPTS {
        log.debug("stopping motors, lost water tower")
        return "follow_line"
      }
    }
  }

  return MODE
}

func mode(mode string) bool {
  return strings.Contains(MODE, mode)
}