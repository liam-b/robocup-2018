package main

import "time"

var log Logger = Logger{flag: "test", level: 7}
var bot Bot = Bot{
  battery: Battery{}.init(),
  colorSensor: ColorSensor{port: IN_2}.init(),
}

func main() {
  log.init("start")
  log.info("program started")

  log.info("voltage is at" + log.number(bot.battery.voltage()))

  speaker := Speaker{}.init()

  speaker.song([]int{300, 100, 400, 100, 500, 100, 600, 100}, 1)

  log.trace("starting loop")
  log.info("looping")
  log.rep("loop")
  loop()
}

func loop() {
  time.Sleep(time.Second / 10)
  log.trace("looping")
  loop()
}