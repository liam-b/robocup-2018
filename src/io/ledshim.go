package main

import "time"

const BEHAVIOUR_REGISTER = 0x00
const FRAME_REGISTER = 0x01
const AUTOPLAY1_REGISTER = 0x02
const AUTOPLAY2_REGISTER = 0x03
const BLINK_REGISTER = 0x05
const AUDIOSYNC_REGISTER = 0x06
const BREATH1_REGISTER = 0x08
const BREATH2_REGISTER = 0x09
const SHUTDOWN_REGISTER = 0x0a
const GAIN_REGISTER = 0x0b
const ADC_REGISTER = 0x0c

const CONFIG_BANK = 0x0b
const BANK_ADDRESS = 0xfd

const PICTURE_BEHAVIOUR = 0x00
const AUTOPLAY_BEHAVIOUR = 0x08
const AUDIOPLAY_BEHAVIOUR = 0x18

const ENABLE_OFFSET = 0x00
const BLINK_OFFSET = 0x12
const COLOR_OFFSET = 0x24

var LED_GAMMA = [...]int{
  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,  0,
  0,  0,  0,  0,  0,  0,  1,  1,  1,  1,  1,  1,  1,  2,  2,  2,
  2,  2,  2,  3,  3,  3,  3,  3,  4,  4,  4,  4,  5,  5,  5,  5,
  6,  6,  6,  7,  7,  7,  8,  8,  8,  9,  9,  9,  10, 10, 11, 11,
  11, 12, 12, 13, 13, 13, 14, 14, 15, 15, 16, 16, 17, 17, 18, 18,
  19, 19, 20, 21, 21, 22, 22, 23, 23, 24, 25, 25, 26, 27, 27, 28,
  29, 29, 30, 31, 31, 32, 33, 34, 34, 35, 36, 37, 37, 38, 39, 40,
  40, 41, 42, 43, 44, 45, 46, 46, 47, 48, 49, 50, 51, 52, 53, 54,
  55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70,
  71, 72, 73, 74, 76, 77, 78, 79, 80, 81, 83, 84, 85, 86, 88, 89,
  90, 91, 93, 94, 95, 96, 98, 99, 100,102,103,104,106,107,109,110,
  111,113,114,116,117,119,120,121,123,124,126,128,129,131,132,134,
  135,137,138,140,142,143,145,146,148,150,151,153,155,157,158,160,
  162,163,165,167,169,170,172,174,176,178,179,181,183,185,187,189,
  191,193,194,196,198,200,202,204,206,208,210,212,214,216,218,220,
  222,224,227,229,231,233,235,237,239,241,244,246,248,250,252,255}

type Ledshim struct {
  Address uint8

  width int
  height int
  brightness int

  i2cDevice *I2C

  gamma [256]int
  Buffer [28][4]int
}

func (ledshim Ledshim) New() Ledshim {
  ledshim.width = 28
  ledshim.height = 1
  ledshim.brightness = 0
  ledshim.gamma = LED_GAMMA
  ledshim.i2cDevice, _ = NewI2C(ledshim.Address, 1)

  ledshim.reset()
  ledshim.Show()
  ledshim.setup()

  return ledshim
}

func (ledshim *Ledshim) Clear() {
  for x := 0; x < ledshim.width; x++ {
    ledshim.BufferPixel(x, 0, 0, 0)
  }
  ledshim.Show()
}

func (ledshim *Ledshim) SetPixel(x int, color [3]int) {
  if ledshim.Buffer[x] != [4]int{color[0], color[1], color[2], 0} {
    ledshim.Buffer[x] = [4]int{color[0], color[1], color[2], 0}
    ledshim.ShowIndividual(x)
  }
}

func (ledshim *Ledshim) BufferPixel(x int, red int, green int, blue int) {
  ledshim.Buffer[x] = [4]int{red, green, blue, 0}
}

func (ledshim Ledshim) Show() {
  ledshim.i2cDevice.WriteRegU8(BANK_ADDRESS, 0x00)

  var output [144]uint8
  for x := 0; x < ledshim.width; x++ {
    red := ledshim.gamma[capInt(ledshim.Buffer[x][0])]
    green := ledshim.gamma[capInt(ledshim.Buffer[x][1])]
    blue := ledshim.gamma[capInt(ledshim.Buffer[x][2])]

    rgb := [3]uint8{uint8(red), uint8(green), uint8(blue)}
    for y := 0; y < 3; y++ {
      idx := ledshim.pixelAddress(x, y)
      output[idx] = rgb[y]
    }
  }

  for value := 0; value < len(output); value++ {
    ledshim.i2cDevice.WriteRegU8(uint8(COLOR_OFFSET + value), uint8(output[value]))
  }

  ledshim.i2cDevice.WriteRegU8(BANK_ADDRESS, CONFIG_BANK)
  ledshim.i2cDevice.WriteRegU8(FRAME_REGISTER, 0x00)
}

func (ledshim Ledshim) ShowIndividual(id int) {
  ledshim.i2cDevice.WriteRegU8(BANK_ADDRESS, 0x00)

  red := ledshim.gamma[capInt(ledshim.Buffer[id][0])]
  green := ledshim.gamma[capInt(ledshim.Buffer[id][1])]
  blue := ledshim.gamma[capInt(ledshim.Buffer[id][2])]

  rgb := [3]uint8{uint8(red), uint8(green), uint8(blue)}

  for y := 0; y < 3; y++ {
    ledshim.i2cDevice.WriteRegU8(uint8(COLOR_OFFSET + ledshim.pixelAddress(id, y)), rgb[y])
  }

  ledshim.i2cDevice.WriteRegU8(BANK_ADDRESS, CONFIG_BANK)
  ledshim.i2cDevice.WriteRegU8(FRAME_REGISTER, 0x00)
}

func (ledshim Ledshim) pixelAddress(x int, y int) int {
  lookup := [28][3]int{
    {118, 69, 85},
    {117, 68, 101},
    {116, 84, 100},
    {115, 83, 99},
    {114, 82, 98},
    {113, 81, 97},
    {112, 80, 96},
    {134, 21, 37},
    {133, 20, 36},
    {132, 19, 35},
    {131, 18, 34},
    {130, 17, 50},
    {129, 33, 49},
    {128, 32, 48},

    {127, 47, 63},
    {121, 41, 57},
    {122, 25, 58},
    {123, 26, 42},
    {124, 27, 43},
    {125, 28, 44},
    {126, 29, 45},
    {15, 95, 111},
    {8, 89, 105},
    {9, 90, 106},
    {10, 91, 107},
    {11, 92, 108},
    {12, 76, 109},
    {13, 77, 93}}

    return lookup[x][y]
}

func (ledshim Ledshim) setup() {
  ledshim.i2cDevice.WriteRegU8(BANK_ADDRESS, CONFIG_BANK)
  ledshim.i2cDevice.WriteRegU8(BEHAVIOUR_REGISTER, PICTURE_BEHAVIOUR)
  ledshim.i2cDevice.WriteRegU8(AUDIOSYNC_REGISTER, 0x00)
  ledshim.i2cDevice.WriteRegU8(BANK_ADDRESS, 0x01)
  ledshim.enableLeds()
  ledshim.i2cDevice.WriteRegU8(BANK_ADDRESS, 0x00)
  ledshim.enableLeds()
}

func (ledshim Ledshim) reset() {
  ledshim.i2cDevice.WriteRegU8(BANK_ADDRESS, CONFIG_BANK)
  ledshim.i2cDevice.WriteRegU8(SHUTDOWN_REGISTER, 0x00)
  time.Sleep(time.Millisecond * time.Duration(50))
  ledshim.i2cDevice.WriteRegU8(BANK_ADDRESS, CONFIG_BANK)
  ledshim.i2cDevice.WriteRegU8(SHUTDOWN_REGISTER, 0x01)
}

func (ledshim Ledshim) enableLeds() {
  ledshim.i2cDevice.WriteRegU8(0x00, 0x00)
  ledshim.i2cDevice.WriteRegU8(0x01, 0xbf)
  ledshim.i2cDevice.WriteRegU8(0x02, 0x3e)
  ledshim.i2cDevice.WriteRegU8(0x03, 0x3e)
  ledshim.i2cDevice.WriteRegU8(0x04, 0x3f)
  ledshim.i2cDevice.WriteRegU8(0x05, 0xb3)
  ledshim.i2cDevice.WriteRegU8(0x06, 0x07)
  ledshim.i2cDevice.WriteRegU8(0x07, 0x86)
  ledshim.i2cDevice.WriteRegU8(0x08, 0x30)
  ledshim.i2cDevice.WriteRegU8(0x09, 0x30)
  ledshim.i2cDevice.WriteRegU8(0x0a, 0x3f)
  ledshim.i2cDevice.WriteRegU8(0x0b, 0xb3)
  ledshim.i2cDevice.WriteRegU8(0x0c, 0x3f)
  ledshim.i2cDevice.WriteRegU8(0x0d, 0xb3)
  ledshim.i2cDevice.WriteRegU8(0x0e, 0x7f)
  ledshim.i2cDevice.WriteRegU8(0x0f, 0xfe)
  ledshim.i2cDevice.WriteRegU8(0x10, 0x7f)
  ledshim.i2cDevice.WriteRegU8(0x11, 0x00)
}

func capInt(x int) uint8 {
  if x > 255 { x = 255 }
  return uint8(x)
}