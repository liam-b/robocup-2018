package main

import "strings"

const DO_GREEN_TURN = false
const DO_WATER_TOWER = false
const DO_CHEMICAL_SPILL = false

var BEHAVIOUR = "follow_line"

var behavioursFunctions = map[string]func()string{
  "chemical_spill": SaveCan,
  "water_tower": AvoidWaterTower,
  "turn_green": TurnOnGreen,
  "follow_line": PID,
  "lifted": Lifted}

var behaviourLeds = map[string][3]int{
  "chemical_spill": COLOR_BLUE,
  "water_tower": COLOR_RED,
  "turn_green": COLOR_GREEN,
  "follow_line": COLOR_WHITE,
  "lifted": COLOR_YELLOW}

func Behave() {
  if BEHAVIOUR == "follow_line" {
    if DO_GREEN_TURN && DetectedGreen(LEFT) { BEHAVIOUR = "turn_green:start.left" }
    if DO_GREEN_TURN && DetectedGreen(RIGHT) { BEHAVIOUR = "turn_green:start.right" }
    if DO_WATER_TOWER && DetectedWaterTower(WATER_TOWER_DETECT_DISTANCE, WATER_TOWER_DETECT_COUNT) { BEHAVIOUR = "water_tower:start" }
    if DO_CHEMICAL_SPILL && DetectedSilver() { BEHAVIOUR = "chemical_spill:start" }
  }
  if BotLifted(LIFTED_DETECT_COUNT) { BEHAVIOUR = "lifted:start" }

  BEHAVIOUR = behavioursFunctions[TOP_LEVEL()]()
  go bot.ledshim.SetPixel(BEHAVIOUR_PIXEL, behaviourLeds[TOP_LEVEL()])
}

func TOP_LEVEL() string {
  splitBehaviour := strings.Split(BEHAVIOUR, ".")
  return strings.Split(splitBehaviour[0], ":")[0]
}

func STATE(comparison string) bool {
  splitBehaviour := strings.Split(BEHAVIOUR, ".")
  return ":" + strings.Split(splitBehaviour[0], ":")[1] == comparison
}

func PARAM(param string) bool {
  return strings.Contains(BEHAVIOUR, param)
}

func PARAMS() string {
  splitBehaviour := strings.Split(BEHAVIOUR, ":")
  return "." + strings.Join(strings.Split(splitBehaviour[len(splitBehaviour) - 1], ".")[1:], ".")
}