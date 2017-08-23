package main

import "time"

var log Logger = Logger{flag: "test", level: 7}
// var bot = map[string]interface{}{
//   "battery": "apple",
//   "b": 2,
// }

func main() {
  log.init("start")
  log.info("program started")

  // dev := ColorSensor{port: IN_2}
  // dev.init()

  // fmt.Println(dev.get("name"))
  // dev.mode("COL-REFLECT")
  // fmt.Println(dev.intensity())

  battery := Battery{}.init()

  log.info("voltage " + log.number(battery.voltage()))

  // speaker := Speaker{}.init()

  // speaker.song([]int{300, 100, 400, 100, 500, 100, 600, 100}, 1)

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