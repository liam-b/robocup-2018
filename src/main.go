package main

import "time"
import "os/signal"
import "os"
import "strconv"

var log Logger = Logger{flag: "test", level: LOG_LEVEL}.new(":start")
var bot Bot

func main() {
  log.notice("program started")

  // setup the robot //
  log.inc(":setup")
    log.debug("setting up io")
    bot = Bot{
      battery: Battery{}.new(),
      colorSensorR: ColorSensor{port: S2}.new(),
      colorSensorL: ColorSensor{port: S3}.new(),
      ultrasonicSensor: UltrasonicSensor{port: S4}.new(),

      // motorL: Motor{port: MC}.new(),
      // motorR: Motor{port: MB}.new(),
    }

    log.once(".interrupt")
      log.trace("setting up interrupts")
      setupInterrupt()
  log.dec()

  // initial mode selections //
  log.inc(":mode")
    log.trace("setting sensor modes")
    bot.colorSensorL.mode(bot.colorSensorL.REFLECT)
    bot.colorSensorR.mode(bot.colorSensorR.REFLECT)
    bot.ultrasonicSensor.mode(bot.ultrasonicSensor.US_DIST_CM)
  log.dec()

  // status checks //
  log.inc(":status")
    log.info("checking status")
    batteryStatus()
  log.dec()

  log.info("looping")
  log.rep("loop")
  loop()
}

func loop() {
  time.Sleep(time.Second / time.Duration(LOOP_SPEED))
  // log.trace("looping")
  // log.debug(strconv.Itoa(bot.gyroSensor.angle()))
  log.debug("col left: " + strconv.Itoa(bot.colorSensorL.intensity()) + ", col right: " + strconv.Itoa(bot.colorSensorR.intensity()) + ", ultra dist: " + strconv.Itoa(bot.ultrasonicSensor.distance()))
  // followLine()
  // printStatusWindow()
  loop()
}

func setupInterrupt() {
  stop := make(chan os.Signal, 1)
  signal.Notify(stop, os.Interrupt)
  go func() {
    <-stop
    end("ctrl-c")
  }()
}

// exit function //
func end(catch string) {
  log.set(":end")
  log.notice("caught " + catch)
  log.once(".sound")
  log.trace("playing exit sound")
  log.level = 0

  bot.motorL.stop()
  bot.motorR.stop()

  os.Exit(0)
}
