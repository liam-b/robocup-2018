package main

import "strconv"

const SAVE_CAN_VERIFY_SPEED = 50
const SAVE_CAN_VERIFY_ATTEMPTS = 30

const SAVE_CAN_ENTER_SPEED = 300
const SAVE_CAN_ENTER_POSITION = 570

const SAVE_CAN_SEARCH_SPEED = 40
const SAVE_CAN_SEARCH_CAN_DISTANCE = 800
const SAVE_CAN_SEARCH_CAN_COUNT = 7

const SAVE_CAN_SAVE_SPEED = 130

const SAVE_CAN_ESCAPE_TURN_SPEED = 70
const SAVE_CAN_ESCAPE_TURN_FUDGE_ANGLE = 0.07

const SAVE_CAN_SAVE_POSITION = 200
const SAVE_CAN_SAVE_REMOVE_POSITION = 520

var chemicalSpillVerifyAttempts = 0
var searchGyroAngle = 0
var saveCanIter = 0
var saveCanCount = 80

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
      OpenClaw()
      return "chemical_spill:enter"
    }
  }

  if STATE(":enter") {
    BehaviourTrace("entering checmical spill")
    if bot.motorRight.GetPosition() > SAVE_CAN_ENTER_POSITION && bot.motorLeft.GetPosition() > SAVE_CAN_ENTER_POSITION {
      BehaviourDebug("motors have reached correct positions, moving onto " + log.state(":search") + " for the can")
      go bot.motorRight.RunForever(SAVE_CAN_SEARCH_SPEED)
      go bot.motorLeft.RunForever(-SAVE_CAN_SEARCH_SPEED)
      bot.imu.ResetGyro()
      return "chemical_spill:search"
    }
  }

  if STATE(":search") {
    BehaviourTrace("searching for first instance of can")
    // log.info(strconv.Itoa(bot.ultrasonicSensor.Distance()))
    if DetectedCan(SAVE_CAN_SEARCH_CAN_DISTANCE, SAVE_CAN_SEARCH_CAN_COUNT) {
      BehaviourDebug("found first instance of can, starting " + log.state(":search:found"))
      searchGyroAngle = bot.imu.GyroValue()
      return "chemical_spill:search:found"
    }
  }

  if STATE(":search:found") {
    BehaviourTrace("searching for last instance of can")
    // log.info(strconv.Itoa(bot.ultrasonicSensor.Distance()))
    if LostCan(SAVE_CAN_SEARCH_CAN_DISTANCE, SAVE_CAN_SEARCH_CAN_COUNT) {
      BehaviourDebug("found last instance of can, now starting " + log.state(":search:align"))
      go bot.motorRight.RunForever(-SAVE_CAN_SEARCH_SPEED)
      go bot.motorLeft.RunForever(SAVE_CAN_SEARCH_SPEED)
      searchGyroAngle = bot.imu.GyroValue() - int(float64(bot.imu.GyroValue() - searchGyroAngle) / 2.0)
      return "chemical_spill:search:align"
    }
  }

  if STATE(":search:align") {
    BehaviourTrace("re-aligning with center of can")
    log.info(strconv.Itoa(bot.imu.GyroValue()) + ", " + strconv.Itoa(searchGyroAngle))
    if bot.imu.GyroValue() < searchGyroAngle {
      BehaviourDebug("aligned to can and starting " + log.state(":save"))
      bot.motorRight.SetPosition(0)
      bot.motorLeft.SetPosition(0)
      go bot.motorRight.RunForever(SAVE_CAN_SAVE_SPEED)
      go bot.motorLeft.RunForever(SAVE_CAN_SAVE_SPEED)
      saveCanIter = 0
      return "chemical_spill:save"
    }
  }

  if STATE(":save") {
    BehaviourTrace("moving to infront of can")
    if bot.motorRight.GetPosition() > SAVE_CAN_SAVE_POSITION && bot.motorLeft.GetPosition() > SAVE_CAN_SAVE_POSITION {
      BehaviourDebug("infront of can, grabbing and starting " + log.state(":save:remove"))
      CloseClaw()
      go bot.motorRight.RunForever(SAVE_CAN_SAVE_SPEED)
      go bot.motorLeft.RunForever(SAVE_CAN_SAVE_SPEED)
      return "chemical_spill:save:remove"
    }
  }

  if STATE(":save:remove") {
    BehaviourTrace("removing can from chemical spill")
    if bot.motorRight.GetPosition() > SAVE_CAN_SAVE_REMOVE_POSITION && bot.motorLeft.GetPosition() > SAVE_CAN_SAVE_REMOVE_POSITION {
      BehaviourDebug("removed can from chemical spill, moving to " + log.state(":save:return"))
      go bot.motorRight.RunForever(-SAVE_CAN_SAVE_SPEED)
      go bot.motorLeft.RunForever(-SAVE_CAN_SAVE_SPEED)
      OpenClaw()
      return "chemical_spill:save:return"
    }
  }

  if STATE(":save:return") {
    BehaviourTrace("returning to center of spill after saving")
    if bot.motorRight.GetPosition() < 0 && bot.motorLeft.GetPosition() < 0 {
      BehaviourDebug("back at middle, starting " + log.state(":escape"))
      go bot.motorRight.RunForever(-SAVE_CAN_ESCAPE_TURN_SPEED)
      go bot.motorLeft.RunForever(SAVE_CAN_ESCAPE_TURN_SPEED)
      OpenClaw()
      return "chemical_spill:escape:turn"
    }
  }

  if STATE(":escape:turn") {
    BehaviourTrace("turning to exit angle")
    if bot.imu.GyroValue() < int(float64(searchGyroAngle) * SAVE_CAN_ESCAPE_TURN_FUDGE_ANGLE) {
      BehaviourDebug("correct alignment with exit, moving to " + log.state(":escape:exit"))
      go bot.motorRight.RunForever(-SAVE_CAN_ENTER_SPEED)
      go bot.motorLeft.RunForever(-SAVE_CAN_ENTER_SPEED)

      return "chemical_spill:escape:exit"
    }
  }

  if STATE(":escape:exit") {
    left := bot.colorSensorLeft.RgbIntensity()
    right := bot.colorSensorRight.RgbIntensity()
    if left > 30 && right > 30 {
      go bot.motorRight.RunForever(150)
      go bot.motorLeft.RunForever(-150)
      return "chemical_spill:escape:align"
    }
  }

  if STATE(":escape:align") {
    _, right := GetColors()
    if right == BLACK {
      go bot.motorRight.Stop()
      go bot.motorLeft.Stop()
      return "follow_line"
    }
  }

  return BEHAVIOUR
}
