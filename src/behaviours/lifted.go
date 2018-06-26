package main

const LIFTED_DETECT_COUNT = 3

func Lifted() string {
  if STATE(":start") {
    ResetHelpers()
    go bot.motorRight.Stop()
    go bot.motorLeft.Stop()

    return "lifted:wait"
  }

  if STATE(":wait") {
    if BotPlacedDown() { return "follow_line" }
  }

  return BEHAVIOUR
}