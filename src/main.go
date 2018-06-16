package main

import "time"
import "os/signal"
import "os"
import "strconv"
// import "fmt"

import "./io"

var log Logger = Logger{flag: "test", level: LOG_LEVEL}.New(":start")
var bot Bot

var gyroVal int = 0

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
      ledshim: io.Ledshim{Address: 0x75}.New(),

      // motorL: Motor{Port: io.MC, Logger: log}.New(),
      // motorR: Motor{Port: io.MB, Logger: log}.New(),
    }

    log.once(".interrupt")
      log.trace("setting up interrupts")
      setupInterrupt()
  log.dec()

  time.Sleep(time.Millisecond * time.Duration(SENSOR_INIT_DELAY))

  // initial mode selections //
  log.inc(":mode")
    log.trace("setting sensor modes")
    bot.colorSensorL.Mode(bot.colorSensorL.RGB)
    bot.colorSensorR.Mode(bot.colorSensorR.RGB)
    bot.ultrasonicSensor.Mode(bot.ultrasonicSensor.DISTANCE)
  log.dec()

  // status checks //
  log.inc(":status")
    log.info("checking status")
    batteryStatus()
  log.dec()

  for index := 0; index < 28; index++ {
    bot.ledshim.SetPixel(index, index * 5, 0, 0)
    bot.ledshim.Show()
    time.Sleep(time.Millisecond * time.Duration(80))
  }
  // bot.ledshim.Buffer[0] = [4]int{150, 150, 150, 0}

  log.info("looping")
  log.rep("loop")
  loop()
}

func loop() {
  time.Sleep(time.Second / time.Duration(LOOP_SPEED))
  // log.debug("col left: " + strconv.Itoa(bot.colorSensorL.Intensity()) + ", col right: " + strconv.Itoa(bot.colorSensorR.Intensity()) + ", ultra dist: " + strconv.Itoa(bot.ultrasonicSensor.Distance()))
  // followLine()
  // gyroVal += bot.imu.ReadGyro()
  // log.trace(strconv.Itoa(gyroVal))
  // log.debug("t: " + strconv.Itoa(total) + ", r: " + strconv.Itoa(red) + ", g: " + strconv.Itoa(green) + ", b: " + strconv.Itoa(blue))

  findColor()
  // strconv.Itoa(bot.colorSensorL.Intensity())

  loop()
}

func findColor() {
  total := bot.colorSensorL.RgbIntensity()
  red, green, blue := bot.colorSensorL.Rgb()
  color := NONE

  if (total > 45) {
    color = SILVER
  } else if (total > 20) {
    color = WHITE
  } else if (green > blue + 6 && green > red + 6) {
    color = GREEN
  } else if (total < 6) {
    color = BLACK
  } else if (total == 0) {
    color = NONE
  }

  log.debug(color)
  log.debug("t: " + strconv.Itoa(total) + ", r: " + strconv.Itoa(red) + ", g: " + strconv.Itoa(green) + ", b: " + strconv.Itoa(blue))
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

  // bot.motorL.Stop()
  // bot.motorR.Stop()
  bot.imu.Cleanup()

  os.Exit(0)
}