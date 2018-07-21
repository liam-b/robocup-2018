package main

import "strings"

func TOP_LEVEL() string {
  return strings.Split(strings.Split(BEHAVIOUR, ".")[0], ":")[0]
}

func STATE(comparison string) bool {
  splitBehaviour := strings.Split(BEHAVIOUR, ".")
  return ":" + strings.Join(strings.Split(splitBehaviour[0], ":")[1:], ":") == comparison
}

func PARAM(param string) bool {
  return strings.Contains(BEHAVIOUR, param)
}

func PARAMS() string {
  splitBehaviour := strings.Split(BEHAVIOUR, ":")
  return "." + strings.Join(strings.Split(splitBehaviour[len(splitBehaviour) - 1], ".")[1:], ".")
}

func STRIP() string {
  return strings.Split(BEHAVIOUR, ".")[0]
}

func BehaviourDebug(text string) {
  log.debug(log.state("[" + STRIP() + "] ") + text)
}

func BehaviourTrace(text string) {
  log.trace(log.state("[" + STRIP() + "] ") + text)
}

func SpeedRatio(speed int, ratio float64, sign int) int {
  return speed + int(float64(speed) * ratio * float64(sign))
}

func ResetHelpers() {
  liftedMatches = 0
  waterTowerMatches = 0
  canMatches = 0
}

func min(x, y int) int {
  if x < y { return x }
  return y
}

func max(x, y int) int {
  if x > y { return x }
  return y
}

func contains(s []string, e string) bool {
  for _, a := range s {
    if a == e {
      return true
    }
  }
  return false
}
