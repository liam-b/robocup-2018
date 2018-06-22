package io

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

const ENABLED_PIXEL = 0x00
const BATTERY_PIXEL = 0x01
const SCOPE_PIXEL = 0x02
const BEHAVIOUR_PIXEL = 0x04

const COLOR_RIGHT_PIXEL = 0x1a
const COLOR_LEFT_PIXEL = 0x1b

var BLACK = [3]int{0, 0, 0}
var WHITE = [3]int{120, 120, 120}

var RED = [3]int{150, 0, 0}
var GREEN = [3]int{0, 120, 0}
var BLUE = [3]int{0, 0, 150}

var YELLOW = [3]int{250, 200, 0}
var PURPLE = [3]int{150, 0, 150}
var CYAN = [3]int{0, 150, 150}