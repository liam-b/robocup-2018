package main

import "time"
import "os/signal"
import "os"
// import "strconv"
import "fmt"

import "./io"

var log Logger = Logger{flag: "test", level: LOG_LEVEL}.New(":start")
var bot Bot

var gyroVal int = 0

func main() {
  log.notice("program started")

  log.inc(":setup")
    log.debug("setting up io")
    bot = Bot{
      battery: io.Battery{}.New(),
      colorSensorRight: io.ColorSensor{Port: io.S2}.New(),
      colorSensorLeft: io.ColorSensor{Port: io.S3}.New(),
      ultrasonicSensor: io.UltrasonicSensor{Port: io.S4}.New(),
      imu: io.IMU{Address: 0x68}.New(),

      // motorLeft: Motor{Port: io.MC, Logger: log}.New(),
      // motorRight: Motor{Port: io.MB, Logger: log}.New(),

      ledshim: io.Ledshim{Address: 0x75}.New(),
    }

    bot.ledshim.SetPixel(io.ENABLED_PIXEL, io.COLOR_GREEN)

    log.once(".interrupt")
      log.trace("setting up interrupts")
      setupInterrupt()
  log.dec()

  time.Sleep(time.Millisecond * time.Duration(SENSOR_INIT_DELAY))

  log.inc(":mode")
    log.trace("setting sensor modes")
    bot.colorSensorLeft.Mode(bot.colorSensorLeft.RGB)
    bot.colorSensorRight.Mode(bot.colorSensorRight.RGB)
    bot.ultrasonicSensor.Mode(bot.ultrasonicSensor.DISTANCE)
  log.dec()

  log.inc(":status")
    log.info("checking status")
    batteryStatus()
  log.dec()

  bot.ledshim.SetPixel(io.SCOPE_STATUS_PIXEL, io.COLOR_BLUE)
  time.Sleep(time.Millisecond * time.Duration(START_LOOP_DELAY))

  bot.ledshim.SetPixel(io.SCOPE_STATUS_PIXEL, io.COLOR_GREEN)
  log.info("looping")
  log.rep("loop")
  loop()
}

func loop() {
  bot.ResetAllCaches()
  time.Sleep(time.Second / time.Duration(LOOP_SPEED))

  leftColor, rightColor := findColors()
  fmt.Println("left: " + leftColor + ", right: " + rightColor)

  if LOOPING { loop() }
}

func setupInterrupt() {
  stop := make(chan os.Signal, 1)
  signal.Notify(stop, os.Interrupt)
  go func() {
    <- stop
    end("ctrl-c")
  }()
}

func end(catch string) {
  LOOPING = false
  bot.ledshim.SetPixel(io.SCOPE_STATUS_PIXEL, io.COLOR_RED)
  log.set(":end")
  log.trace("caught " + catch)
  log.notice("exiting program")
  log.level = 0

  // bot.motorLeft.Stop()
  // bot.motorRight.Stop()
  bot.imu.Cleanup()
  bot.ledshim.Clear()

  time.Sleep(time.Millisecond * time.Duration(END_DELAY))
  os.Exit(0)
}