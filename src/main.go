package main

// import "fmt"
import "time"

var log Logger = Logger{flag: "test", level: 7}

func main() {
  log.init()
  log.inc("start")
  log.trace("logger ready")
  // dev := ColorSensor{port: IN_2}
  // dev.init()

  // fmt.Println(dev.get("name"))
  // dev.mode("COL-REFLECT")
  // fmt.Println(dev.intensity())

  dev := Speaker{}
  dev.init()

  dev.play(300, 200, 20)

  log.trace("starting loop")
  log.dec()
  log.inc("loop")
  loop()
}

func loop() {
  time.Sleep(time.Second / 10)
  log.trace("looping")
  loop()
}