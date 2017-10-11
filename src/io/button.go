package main

import "os"
// import "fmt"
// import "strconv"
// import "reflect"

const PRESSED = 1
const RELEASED = 0

const KEY_UP = 103
const KEY_DOWN = 108
const KEY_LEFT = 105
const KEY_RIGHT = 106

const KEY_ENTER = 28
const KEY_ESCAPE = 14

type Button struct {
  onKeypress func (key int, state int)
  file *os.File
}

func (button Button) new() Button {
  if DO_KEYBOARD_EVENT {
    f, err := os.Open("/dev/input/by-path/platform-gpio-keys.0-event")
    check(err)

    button.file = f
    go button.loopKeypressRead()
  }
  return button
}

func (button Button) loopKeypressRead() {
  bytes := make([]byte, 32)
  button.file.Read(bytes)
  button.onKeypress(int(bytes[10]), int(bytes[12]))
  button.loopKeypressRead()
}
