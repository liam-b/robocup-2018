package main

import "strconv"

const SAVE_CAN_VERIFY_SPEED = 50
const SAVE_CAN_VERIFY_ATTEMPTS = 30
const SAVE_CAN_ENTER_SPEED = 300
const SAVE_CAN_ENTER_POSITION = 530
const SAVE_CAN_SEARCH_SPEED = 100
const SAVE_CAN_SEARCH_CAN_DISTANCE = 800
const SAVE_CAN_SEARCH_CAN_COUNT = 3

var chemicalSpillVerifyAttempts = 0
var searchGyroAngle = 0

func SaveCan() string {
  if STATE(":start") {
    BehaviourDebug("starting " + log.state(":verify") + " of chemical spill")
    chemicalSpillVerifyAttempts = 0
    go bot.motorRight.RunForever(SAVE_CAN_VERIFY_SPEED)
    go bot.motorLeft.RunForever(SAVE_CAN_VERIFY_SPEED)
    return "chemical_spill:verify"
  }

  if STATE(":verify") {
    BehaviourTrace("verifying chemical spill with color sensors")
    chemicalSpillVerifyAttempts += 1
    if chemicalSpillVerifyAttempts > SAVE_CAN_VERIFY_ATTEMPTS {
      BehaviourDebug("chemical spill wasn't verified within count, returning to " + log.state("follow_line"))
      chemicalSpillVerifyAttempts = 0
      go bot.motorRight.Stop()
      go bot.motorLeft.Stop()
      return "follow_line"
    }

    left, right := GetColors()
    if left == GREEN && right == GREEN {
      BehaviourDebug("verified chemical spill, switching to " + log.state(":enter"))
      bot.motorRight.SetPosition(0)
      bot.motorLeft.SetPosition(0)
      go bot.motorRight.RunForever(SAVE_CAN_ENTER_SPEED)
      go bot.motorLeft.RunForever(SAVE_CAN_ENTER_SPEED)
      return "chemical_spill:enter"
    }
  }

  if STATE(":enter") {
    BehaviourTrace("entering checmical spill")
    if bot.motorRight.GetPosition() > SAVE_CAN_ENTER_POSITION && bot.motorLeft.GetPosition() > SAVE_CAN_ENTER_POSITION {
      BehaviourDebug("motors have reached correct positions, moving onto " + log.state(":search") + " for the can")
      go bot.motorRight.RunForever(SAVE_CAN_SEARCH_SPEED)
      go bot.motorLeft.RunForever(-SAVE_CAN_SEARCH_SPEED)
      return "chemical_spill:search"
    }
  }

  if STATE(":search") {
    BehaviourTrace("searching for first instance of can")
    if DetectedCan(SAVE_CAN_SEARCH_CAN_DISTANCE, SAVE_CAN_SEARCH_CAN_COUNT) {
      BehaviourDebug("found first instance of can, starting " + log.state(":search:found"))
      ResetGyroTotalRotation()
      return "chemical_spill:search:found"
    }
  }

  if STATE(":search:found") {
    BehaviourTrace("searching for last instance of can")
    searchGyroAngle = GyroTotalRotation()
    if LostCan(SAVE_CAN_SEARCH_CAN_DISTANCE, SAVE_CAN_SEARCH_CAN_COUNT) {
      BehaviourDebug("found last instance of can, now starting " + log.state(":search:align"))
      go bot.motorRight.RunForever(-SAVE_CAN_SEARCH_SPEED)
      go bot.motorLeft.RunForever(SAVE_CAN_SEARCH_SPEED)
      totalAngle = 0
      return "chemical_spill:search:align"
    }
  }

  if STATE(":search:align") {
    BehaviourTrace("re-aligning with center of can")
    log.info(strconv.Itoa(searchGyroAngle))
    if GyroTurnedToAngle(-searchGyroAngle / 2, RIGHT) {
      BehaviourDebug("aligned to can and starting " + log.state(":save"))
      go bot.motorRight.Stop()
      go bot.motorLeft.Stop()
      return "chemical_spill:save"
    }
  }

  return BEHAVIOUR
}
