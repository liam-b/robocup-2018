package main

const CLAW_SIZE_POSITION = 100
const CLAW_SPEED = 100

func OpenClaw() {
  bot.motorClaw.SetPosition(0)
  bot.motorClaw.RunToPosition(CLAW_SIZE_POSITION, CLAW_SPEED)
}

func CloseClaw() {
  bot.motorClaw.SetPosition(0)
  bot.motorClaw.RunToPosition(-CLAW_SIZE_POSITION, CLAW_SPEED)
}

func HalfOpenClaw() {
  bot.motorClaw.SetPosition(0)
  bot.motorClaw.RunToPosition(CLAW_SIZE_POSITION / 2, CLAW_SPEED)
}

func HalfCloseClaw() {
  bot.motorClaw.SetPosition(0)
  bot.motorClaw.RunToPosition(-CLAW_SIZE_POSITION / 2, CLAW_SPEED)
}