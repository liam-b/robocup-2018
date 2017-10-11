package main

func init() {
  log.inc(":mock")
  log.notice("using mock!")

  log.debug("replacing driver paths")
  MOTOR_PATH = "mock/motor/"
  SENSOR_PATH = "mock/sensor/"
  BATTERY_PATH = "mock/battery/"
  SOUND_PATH = "mock/sound/"

  DO_KEYBOARD_EVENT = false
  // DO_KEYBOARD = false

  log.dec()
}