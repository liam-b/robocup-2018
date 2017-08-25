package main

import "io/ioutil"

const IN_1 = "in1"
const IN_2 = "in2"
const IN_3 = "in3"
const IN_4 = "in4"

const OUT_A = "outA"
const OUT_B = "outB"
const OUT_C = "outC"
const OUT_D = "outD"

var MOTOR_PATH string = "/sys/class/tacho-motor/"
var SENSOR_PATH string = "/sys/class/lego-sensor/"
var BATTERY_PATH string = "/sys/class/power_supply/legoev3-battery/"
var SOUND_PATH string = "/sys/devices/platform/snd-legoev3/"

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func read(path string) string {
  dat, err := ioutil.ReadFile(path)
  check(err)
  return string(dat)
}

func readBytes(path string) []byte {
  dat, err := ioutil.ReadFile(path)
  check(err)
  return dat
}

func write(path string, data string) {
  dat := []byte(data)
  err := ioutil.WriteFile(path, dat, 0644)
  check(err)
}

func list(path string) []string {
  files, _ := ioutil.ReadDir(path)
  stringFiles := make([]string, len(files))
  for i, file := range files {
    stringFiles[i] = file.Name()
  }
  return stringFiles
}
