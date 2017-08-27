package main

import "time"
import "os/signal"
import "os"
import "strconv"

var log Logger = Logger{flag: "test", level: 7}.new(":start")
var bot Bot

func main() {
  log.notice("program started")

  log.info("setting up io")
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

  log.inc(":status")
  if bot.battery.voltage() > 72 {
    log.info("voltage is at " + log.value(bot.battery.voltageString() + "v"))
  } else {
    log.warn("voltage is at " + log.value(bot.battery.voltageString() + "v"))
  }
  log.dec()

  // go bot.speaker.song([]int{300, 400, 500, 600}, 100, 1)
  go bot.speaker.song([]int{300, 400, 500, 600, 0, 500, 600}, 100, 1)

  log.info("looping")
  log.rep("loop")

  setupInterrupt()
  loop()
}

func loop() {
  time.Sleep(time.Second / 10)
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
  log.level = 0
  bot.speaker.song([]int{600, 500, 400, 300}, 100, 1)
  os.Exit(0)
}