package main

import "io/ioutil"
import "os"

const S1 = "spi0.1:S1"
const S2 = "spi0.1:S2"
const S3 = "spi0.1:S3"
const S4 = "spi0.1:S4"

const MA = "spi0.1:MA"
const MB = "spi0.1:MB"
const MC = "spi0.1:MC"
const MD = "spi0.1:MD"

var PORT = map[string]string{
  "spi0.1:S1": "port0",
  "spi0.1:S2": "port1",
  "spi0.1:S3": "port2",
  "spi0.1:S4": "port3",

  "spi0.1:MA": "port4",
  "spi0.1:MB": "port5",
  "spi0.1:MC": "port6",
  "spi0.1:MD": "port7"}

const MOTOR_PATH = "/sys/class/tacho-motor/"
const SENSOR_PATH = "/sys/class/lego-sensor/"
const BATTERY_PATH = "/sys/class/power_supply/brickpi3-battery/"
const LEGO_PORT = "/sys/class/lego-port/"

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

func exists(path string) bool {
  _, err := os.Stat(path)
  return os.IsNotExist(err)
}