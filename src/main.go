package main

import "time"
import "os/signal"
import "os"
import "strconv"

var log Logger = Logger{flag: "test", level: 7}.new(":start")
var bot Bot

func main() {
  log.notice("program started")

  log.inc(":setup")
    // initialisation things
    log.debug("setting up io")
    bot = Bot{
      battery: Battery{}.new(),
      // colorSensor: ColorSensor{port: IN_3}.new(),
      speaker: Speaker{playSound: true}.new(),
      touchSensor: TouchSensor{port: IN_1}.new(),
      gyroSensor: GyroSensor{port: IN_2}.new(),

      button: Button{
        onKeypress: func (key int, state int) {
          if key == KEY_ESCAPE {
            end("escape")
          }
        },
      }.new(),
    }

    log.once(".sound")
    log.trace("playing startup sound")
    go bot.speaker.song([]int{400, 400, 0, 500, 500}, 50, 1)

    log.once(".interrupt")
    log.trace("setting up interrupts")
    setupInterrupt()
  log.dec()

  log.inc(":mode")
    // all mode sets and things
  log.dec()

  log.inc(":status")
    // for checking status
    log.info("checking status")

    log.inc(".battery")
      log.debug("voltage is at " + log.value(bot.battery.voltageString() + "v"))
      if (bot.battery.voltage() < 72) {
        log.warn("battery needs replacing now")
      } else if (bot.battery.voltage() < 75) {
        log.warn("current voltage is not fit for comp")
      }
    log.dec()

  log.dec()

  log.info("looping")
  log.rep("loop")
  loop()
}

func loop() {
  time.Sleep(time.Second / 5)
  // log.trace("looping")
  // log.trace(strconv.FormatBool(bot.touchSensor.pressed()))
  log.trace(strconv.Itoa(bot.gyroSensor.angle()))
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
  os.Exit(0)
}