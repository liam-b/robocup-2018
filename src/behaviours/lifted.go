package main

const LIFTED_DETECT_COUNT = 3

func Lifted() string {
  if STATE(":start") {
    BehaviourDebug("bot has been picked up and is moving to " + log.state(":wait") + " until placed down")
    ResetHelpers()
    go bot.motorRight.Stop()
    go bot.motorLeft.Stop()

    return "lifted:wait"
  }

  if STATE(":wait") {
    BehaviourTrace("waiting for bot to be placed down")
    if BotPlacedDown() && bot.touchSensor.Pressed() {
      BehaviourDebug("bot has been placed down, returning to " + log.state("follow_line"))
      ResetPID()
      return "follow_line"
    }
  }

  return BEHAVIOUR
}
