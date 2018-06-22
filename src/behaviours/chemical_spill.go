package main

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