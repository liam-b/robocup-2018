package main

const SAVE_CAN_VERIFY_SPEED = 50
const SAVE_CAN_VERIFY_ATTEMPTS = 30
const SAVE_CAN_ENTER_SPEED = 300
const SAVE_CAN_ENTER_COUNT = 60

var chemicalSpillVerifyAttempts = 0
var chemicalSpillEnterCount = 0

func SaveCan() string {
  if STATE(":start") {
    chemicalSpillVerifyAttempts = 0
    go bot.motorRight.RunForever(SAVE_CAN_VERIFY_SPEED)
    go bot.motorLeft.RunForever(SAVE_CAN_VERIFY_SPEED)
    return "chemical_spill:verify"
  }

  if STATE(":verify") {
    chemicalSpillVerifyAttempts += 1
    if chemicalSpillVerifyAttempts > SAVE_CAN_VERIFY_ATTEMPTS {
      chemicalSpillVerifyAttempts = 0
      go bot.motorRight.Stop()
      go bot.motorLeft.Stop()
      return "follow_line"
    }

    left, right := GetColors()
    if left == GREEN && right == GREEN {
      go bot.motorRight.RunForever(SAVE_CAN_ENTER_SPEED)
      go bot.motorLeft.RunForever(SAVE_CAN_ENTER_SPEED)
      return "chemical_spill:enter"
    }
  }

  if STATE(":enter") {
    chemicalSpillEnterCount += 1
    if chemicalSpillEnterCount > SAVE_CAN_ENTER_COUNT {
      chemicalSpillEnterCount = 0
      go bot.motorRight.Stop()
      go bot.motorLeft.Stop()
      return "chemical_spill:search"
    }
  }

  if STATE(":search") {
    // do sum searching
  }

  return BEHAVIOUR
}
