package main

const LIFTED_DETECT_COUNT = 3

func Lifted() string {
  if BEHAVIOUR == "lifted:start" {
    ResetHelpers()
    go bot.motorRight.Stop()
    go bot.motorLeft.Stop()

    return "lifted:wait"
  }

  if BEHAVIOUR == "lifted:wait" {
    if BotPlacedDown() { return "follow_line" }
  }

  return BEHAVIOUR
}