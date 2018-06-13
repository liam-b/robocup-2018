package main

import "time"
import "os/signal"
import "os"
import "strconv"

import "./io"

var log Logger = Logger{flag: "test", level: LOG_LEVEL}.New(":start")
var bot Bot

func main() {
  log.notice("program started")

  // setup the robot //
  log.inc(":setup")
    log.debug("setting up io")
    bot = Bot{
      battery: io.Battery{}.New(),
      colorSensorR: io.ColorSensor{Port: io.S2}.New(),
      colorSensorL: io.ColorSensor{Port: io.S3}.New(),
      ultrasonicSensor: io.UltrasonicSensor{Port: io.S4}.New(),

      imu: io.IMU{Address: 0x68}.New(),

      // motorL: Motor{Port: io.MC, Logger: log}.New(),
      // motorR: Motor{Port: io.MB, Logger: log}.New(),
    }

    log.once(".interrupt")
      log.trace("setting up interrupts")
      setupInterrupt()
  log.dec()

  // initial mode selections //
  log.inc(":mode")
    log.trace("setting sensor modes")
    bot.colorSensorL.Mode(bot.colorSensorL.REFLECT)
    bot.colorSensorR.Mode(bot.colorSensorR.REFLECT)
    bot.ultrasonicSensor.Mode(bot.ultrasonicSensor.US_DIST_CM)
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
  // log.debug("col left: " + strconv.Itoa(bot.colorSensorL.intensity()) + ", col right: " + strconv.Itoa(bot.colorSensorR.intensity()) + ", ultra dist: " + strconv.Itoa(bot.ultrasonicSensor.distance()))
  // followLine()
  // printStatusWindow()
  log.trace(strconv.Itoa(bot.imu.ReadGyro()))
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
  log.trace("caught " + catch)
  log.notice("exiting program")
  log.level = 0

  bot.motorL.Stop()
  bot.motorR.Stop()
  bot.imu.Cleanup()

  os.Exit(0)
}
