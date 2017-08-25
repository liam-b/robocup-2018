package main

import "time"
import "os"
import "os/signal"
import "fmt"

var log Logger = Logger{flag: "test", level: 7}.new("start")
var bot Bot

func main() {
  log.info("program started")

  log.info("setting up io")
  bot = Bot{
    battery: Battery{}.new(),
    colorSensor: ColorSensor{port: IN_2}.new(),
    speaker: Speaker{}.new(),
  }

  log.inc(":status")
  log.info("voltage is at " + log.number(bot.battery.voltage()))
  log.dec()

  // go bot.speaker.song([]int{300, 400, 500, 600}, 100, 1)

  f, err := os.Open("/dev/input/by-path/platform-gpio-keys.0-event")
  defer f.Close()
  check(err)

  bytes := make([]byte, 32)

	f.Read(bytes)

  fmt.Println("%d bytes", bytes)

  log.trace("starting loop")
  log.info("looping")
  log.rep("loop")

  // setupInterrupt()
  // loop()
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