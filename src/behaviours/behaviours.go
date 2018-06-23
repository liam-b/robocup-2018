package main

import "strings"

const DO_GREEN_TURN = true
const DO_WATER_TOWER = true
const DO_CHEMICAL_SPILL = true

var BEHAVIOUR = "follow_line"

var behavioursFunctions = map[string]func()string{
  "chemical_spill": SaveCan,
  "water_tower": AvoidWaterTower,
  "turn_green": TurnOnGreen,
  "follow_line": FollowLine}

var behaviourLeds = map[string][3]int{
  "chemical_spill": COLOR_BLUE,
  "water_tower": COLOR_RED,
  "turn_green": COLOR_GREEN,
  "follow_line": COLOR_WHITE}

func Behave() {
  if BEHAVIOUR == "follow_line" {
    if DO_GREEN_TURN && DetectedGreen(LEFT) { BEHAVIOUR = "turn_green:left" }
    if DO_GREEN_TURN && DetectedGreen(RIGHT) { BEHAVIOUR = "turn_green:right" }
    if DO_WATER_TOWER && DetectedWaterTower(WATER_TOWER_DETECT_DISTANCE, WATER_TOWER_DETECT_COUNT) { BEHAVIOUR = "water_tower:start" }
    if DO_CHEMICAL_SPILL && DetectedSilver() { BEHAVIOUR = "chemical_spill:start" }
  }

  BEHAVIOUR = behavioursFunctions[strings.Split(BEHAVIOUR, ":")[0]]()
  go bot.ledshim.SetPixel(BEHAVIOUR_PIXEL, behaviourLeds[BEHAVIOUR])
}