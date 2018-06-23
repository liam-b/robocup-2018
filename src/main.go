package main

import "time"
import "os/signal"
import "os"
import "strconv"
// import "fmt"

var log Logger = Logger{flag: "test", level: LOG_LEVEL}.New(":start")
var bot Bot

func main() {
  log.notice("program started")

  log.inc(":setup")
    log.debug("setting up io")
    bot = Bot{
      battery: Battery{}.New(),
      colorSensorRight: ColorSensor{Port: S2}.New(),
      colorSensorLeft: ColorSensor{Port: S3}.New(),
      ultrasonicSensor: UltrasonicSensor{Port: S4}.New(),
      imu: IMU{Address: 0x68}.New(),

      motorLeft: Motor{Port: MC}.New(),
      motorRight: Motor{Port: MB}.New(),

      ledshim: Ledshim{Address: 0x75}.New(),
    }

    bot.ledshim.SetPixel(ENABLED_PIXEL, COLOR_GREEN)
    bot.ledshim.SetPixel(SCOPE_PIXEL, COLOR_BLUE)

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
    BatteryStatus()
  log.dec()

  time.Sleep(time.Millisecond * time.Duration(START_LOOP_DELAY))
  bot.ledshim.SetPixel(SCOPE_PIXEL, COLOR_GREEN)
  log.info("looping")
  log.rep("loop")
  loop()
}

func loop() {
  if LOOPING {
    bot.ResetAllCaches()
    time.Sleep(time.Second / time.Duration(LOOP_SPEED))

    // Behave()
    // FollowLine(true, true)
    // log.debug(BEHAVIOUR + ", " + strconv.Itoa(int(float64(2550 - bot.ultrasonicSensor.Distance()) / 2.55)))
    // log.debug(BEHAVIOUR + ", " + strconv.Itoa(totalAngle))
    // log.debug(BEHAVIOUR + ", " + "l: " + strconv.Itoa(bot.colorSensorLeft.RgbIntensity()) + " r: " + strconv.Itoa(bot.colorSensorRight.RgbIntensity()))
    log.debug(strconv.Itoa(bot.colorSensorLeft.RgbIntensity() - bot.colorSensorRight.RgbIntensity()))
    PID()
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
  bot.ledshim.SetPixel(SCOPE_PIXEL, COLOR_RED)
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