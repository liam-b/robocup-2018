package main

import "time"
import "os"
import "os/signal"

var log Logger = Logger{flag: "test", level: 7}
var bot Bot = Bot{
  battery: Battery{}.init(),
  colorSensor: ColorSensor{port: IN_2}.init(),
}

func main() {
  log.init("start")
  log.info("program started")

  log.info("voltage is at " + log.number(bot.battery.voltage()))

  speaker := Speaker{}.init()

  go speaker.song([]int{300, 400, 500, 600}, 100, 1)

  log.trace("starting loop")
  log.info("looping")
  log.rep("loop")
  setupInterrupt()
  loop()
}

func loop() {
  time.Sleep(time.Second / 10)
  log.trace("looping")
  loop()
}

func setupInterrupt() {
  stop := make(chan os.Signal, 1)
  signal.Notify(stop, os.Interrupt)
  go func() {
    <-stop
    log.notice("caught ctrl-c")
    os.Exit(0)
  }()
}