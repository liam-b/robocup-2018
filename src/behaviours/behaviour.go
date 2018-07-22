package main

const DO_BOT_LIFTED = true
const DO_GREEN_TURN = true
const DO_WATER_TOWER = true
const DO_CHEMICAL_SPILL = true

var BEHAVIOUR = "lifted:start"
// var BEHAVIOUR = "turn_green:start.left"
// var BEHAVIOUR = "follow_line"

var behaviourFunctions = map[string]func()string{
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
  if TOP_LEVEL() == "follow_line" {
    if DO_GREEN_TURN && DetectedGreen(LEFT) { BEHAVIOUR = "turn_green:start.left" }
    if DO_GREEN_TURN && DetectedGreen(RIGHT) { BEHAVIOUR = "turn_green:start.right" }
    if DO_WATER_TOWER && DetectedWaterTower(WATER_TOWER_DETECT_DISTANCE, WATER_TOWER_DETECT_COUNT) { BEHAVIOUR = "water_tower:start" }
    if DO_CHEMICAL_SPILL && DetectedSilver() { BEHAVIOUR = "chemical_spill:start" }
  }
  if DO_BOT_LIFTED && BotLifted(LIFTED_DETECT_COUNT) && TOP_LEVEL() != "lifted" { BEHAVIOUR = "lifted:start" }

  BEHAVIOUR = behaviourFunctions[TOP_LEVEL()]()
  go bot.ledshim.SetPixel(BEHAVIOUR_PIXEL, behaviourLeds[TOP_LEVEL()])
}
