package main

import "time"
import "os/signal"
import "os"
// import "fmt"
import "strconv"

var log Logger = Logger{flag: "test", level: LOG_LEVEL}.new(":start")
var bot Bot

func main() {
  log.notice("program started")

  log.inc(":setup") // initialisation things
    log.debug("setting up io")
    bot = Bot{
      battery: Battery{}.new(),
      colorSensorR: ColorSensor{port: IN_1}.new(),
      colorSensorL: ColorSensor{port: IN_2}.new(),
      // ultrasonicSensor: UltrasonicSensor{port: IN_3}.new(),
      // gyroSensor: GyroSensor{port: IN_4}.new(),
      speaker: Speaker{playSound: true}.new(),

      motorL: Motor{port: OUT_D}.new(),
      motorR: Motor{port: OUT_B}.new(),

      button: Button{
        onKeypress: func(key int, state int) {
          if key == KEY_ESCAPE {
            end("escape")
          }
        },
      }.new(),
    }

    log.once(".sound")
    log.trace("playing startup sound")
    bot.speaker.volume(VOLUME)
    bot.speaker.song([]int{400, 400, 0, 500, 500}, 50, 1)
    time.Sleep(time.Millisecond * time.Duration(200))

    log.once(".interrupt")
    log.trace("setting up interrupts")
    setupInterrupt()
  log.dec()

  log.inc(":mode") // all mode sets and things
    log.trace("setting sensor modes")
    bot.colorSensorL.mode(bot.colorSensorL.REFLECT)
    bot.colorSensorR.mode(bot.colorSensorR.REFLECT)

    // bot.motorR.runForever(200)
    // bot.motorL.runForever(200)
  log.dec()

  log.inc(":status") // for checking status
    log.info("checking status")

    log.inc(".battery")
      checkBatteryVoltage()
    log.dec()

  log.dec()

  log.info("looping")
  log.rep("loop")
  loop()
}

func loop() {
  time.Sleep(time.Second / time.Duration(LOOP_SPEED))
  // log.trace("looping")
  // log.debug(strconv.Itoa(bot.gyroSensor.angle()))
  log.debug(strconv.Itoa(bot.colorSensorL.intensity()) + " | " + strconv.Itoa(bot.colorSensorR.intensity()))
  followLine()
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

func end(catch string) {
  log.set(":end")
  log.notice("caught " + catch)
  log.once(".sound")
  log.trace("playing exit sound")
  log.level = 0
  bot.speaker.song([]int{500, 500, 0, 400, 400}, 50, 1)

  bot.motorL.stop()
  bot.motorR.stop()

  os.Exit(0)

  bot.motorL.stop()
  bot.motorR.stop()
}
