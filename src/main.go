package main

import "time"
import "os/signal"
import "os"
// import "strconv"
// import "fmt"

var log Logger = Logger{flag: "test", level: LOG_LEVEL}.New(":start")
var bot Bot

// TODO: add useful starting params (eg. can turn direction)

func main() {
  log.notice("program started")

  log.inc(":setup")
    log.debug("setting up io")
    bot = Bot{
      battery: Battery{}.New(),

      touchSensor: TouchSensor{Port: S1}.New(),
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
    bot.touchSensor.Mode(bot.touchSensor.TOUCH)
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

    log.inc(":behave")
      Behave()
    log.dec()

    // FollowLine(true, true)
    // log.debug(BEHAVIOUR + ", " + strconv.Itoa(int(float64(2550 - bot.ultrasonicSensor.Distance()) / 2.55)))
    // log.debug(BEHAVIOUR + ", " + strconv.Itoa(bot.imu.ReadGyro()))
    // log.debug(BEHAVIOUR + ", " + "l: " + strconv.Itoa(bot.colorSensorLeft.RgbIntensity()) + " r: " + strconv.Itoa(bot.colorSensorRight.RgbIntensity()))
    // log.debug(strconv.Itoa(bot.colorSensorLeft.RgbIntensity() - bot.colorSensorRight.RgbIntensity()))
    // PID()
    // log.debug(strconv.FormatBool())

    // leftCol, rightCol := GetColors()
    // log.debug(BEHAVIOUR + ", " + "l: " + leftCol + " r: " + rightCol)

    // _, leftGreen, _ := bot.colorSensorLeft.Rgb()
    // _, rightGreen, _ := bot.colorSensorRight.Rgb()
    // log.debug(BEHAVIOUR + ", " + "l: " + strconv.Itoa(leftGreen) + " r: " + strconv.Itoa(rightGreen))
    //
    // bot.ledshim.SetPixel(COLOR_LEFT_PIXEL, [3]int{leftRed * 4, leftGreen * 4, leftBlue * 4})
    // bot.ledshim.SetPixel(COLOR_RIGHT_PIXEL, [3]int{rightRed * 4, rightGreen * 4, rightBlue * 4})
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
  log.end()

  time.Sleep(time.Millisecond * time.Duration(END_DELAY))

  bot.motorLeft.Stop()
  bot.motorRight.Stop()
  bot.imu.Cleanup()
  bot.ledshim.Clear()

  os.Exit(0)
}
