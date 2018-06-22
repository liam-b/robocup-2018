package main

import "time"
import "os/signal"
import "os"
import "strconv"
// import "fmt"

import "./io"

var log Logger = Logger{flag: "test", level: LOG_LEVEL}.New(":start")
var bot Bot

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

      motorLeft: io.Motor{Port: io.MC}.New(),
      motorRight: io.Motor{Port: io.MB}.New(),

      ledshim: io.Ledshim{Address: 0x75}.New(),
    }

    bot.ledshim.SetPixel(io.ENABLED_PIXEL, io.GREEN)
    bot.ledshim.SetPixel(io.SCOPE_PIXEL, io.BLUE)

    log.once(".interrupt")
      log.trace("setting up interrupts")
      SetupInterrupts()
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

  time.Sleep(time.Millisecond * time.Duration(START_LOOP_DELAY))
  bot.ledshim.SetPixel(io.SCOPE_PIXEL, io.GREEN)
  log.info("looping")
  log.rep("loop")
  loop()
}

func loop() {
  if LOOPING {
    bot.ResetAllCaches()
    time.Sleep(time.Second / time.Duration(LOOP_SPEED))

    Behave()
    // FollowLine(true, true)
    // log.debug(BEHAVIOUR + ", " + strconv.Itoa(int(float64(2550 - bot.ultrasonicSensor.Distance()) / 2.55)))
    // log.debug(BEHAVIOUR + ", " + strconv.Itoa(totalAngle))
    log.debug(BEHAVIOUR + ", " + "l: " + strconv.Itoa(bot.colorSensorLeft.RgbIntensity()) + " r: " + strconv.Itoa(bot.colorSensorRight.RgbIntensity()))
    // log.debug(strconv.FormatBool(DetectedGreen(LEFT)))
  }

  loop()
}

func SetupInterrupts() {
  stop := make(chan os.Signal, 1)
  signal.Notify(stop, os.Interrupt)
  go func() {
    <-stop
    LOOPING = false
    end("ctrl-c")
  }()
}

func end(catch string) {
  bot.ledshim.SetPixel(io.SCOPE_PIXEL, io.RED)
  log.set(":end")
  log.trace("caught " + log.value(catch))
  log.notice("exiting program")
  log.level = 0

  time.Sleep(time.Millisecond * time.Duration(END_DELAY))

  bot.motorLeft.Stop()
  bot.motorRight.Stop()
  bot.imu.Cleanup()
  bot.ledshim.Clear()

  os.Exit(0)
}