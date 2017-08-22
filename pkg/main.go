package main

import "fmt"
// import "time"

func main() {
  dev := ColorSensor{port: IN_2}
  dev.init()

  // fmt.Println(dev.get("name"))
  dev.mode("COL-REFLECT")
  fmt.Println(dev.intensity())

  // fmt.Println("\x1b[32mhi\x1b[0m")
  // fmt.Println(time.Parse(time.RFC3339, "2010-02-04T21:00:57-08:00"))

  // fmt.Println(time.Minute(time.Now()))

  log := Logger{flag:"test"}
  log.init()

  log.inc("main")

  log.trace("trace")
  log.debug("debug")
  log.inc("test")
  log.info("info")
  log.notice("notice")
  // log.success("success")
  log.dec()
  log.warn("warn")
  // time.Sleep(1 * time.Second)
  log.error("error")
  log.set([]string{"test", "hey"})
  log.fatal("fatal")
}