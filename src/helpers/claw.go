package main

const CLAW_TOTAL_DISTANCE = 90
const CLAW_SPEED = 100

const CLAW_OPEN = 2
const CLAW_HALF = 1
const CLAW_CLOSED = 0

var CLAW_POSITION = 0

func SetupClaw() {
  // bot.motorClaw.Stop()
  bot.motorClaw.SetPosition(0)
  bot.motorClaw.StopAction("brake")
  HalfOpenClaw()
}

func CleanupClaw() {
  bot.motorClaw.RunToAbsolutePosition(0, 100)
}

func OpenClaw() {
  if CLAW_POSITION == CLAW_CLOSED {
    go bot.motorClaw.RunToPosition(CLAW_TOTAL_DISTANCE, CLAW_SPEED)
    CLAW_POSITION = CLAW_OPEN
  } else if CLAW_POSITION == CLAW_HALF {
    go bot.motorClaw.RunToPosition(CLAW_TOTAL_DISTANCE / 2, CLAW_SPEED)
    CLAW_POSITION = CLAW_OPEN
  }
}

func CloseClaw() {
  if CLAW_POSITION == CLAW_OPEN {
    go bot.motorClaw.RunToPosition(-CLAW_TOTAL_DISTANCE, CLAW_SPEED)
    CLAW_POSITION = CLAW_CLOSED
  } else if CLAW_POSITION == CLAW_HALF {
    go bot.motorClaw.RunToPosition(-CLAW_TOTAL_DISTANCE / 2, CLAW_SPEED)
    CLAW_POSITION = CLAW_CLOSED
  }
}

func HalfOpenClaw() {
  if CLAW_POSITION != CLAW_OPEN {
    go bot.motorClaw.RunToPosition(CLAW_TOTAL_DISTANCE / 2, CLAW_SPEED)
    CLAW_POSITION += CLAW_HALF
  }
}

func HalfCloseClaw() {
  if CLAW_POSITION != CLAW_CLOSED {
    go bot.motorClaw.RunToPosition(-CLAW_TOTAL_DISTANCE / 2, CLAW_SPEED)
    CLAW_POSITION -= CLAW_HALF
  }
}