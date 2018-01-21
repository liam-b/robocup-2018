package main

import "fmt"

func printStatusWindow() {
  fmt.Println(`  ---------------------------------------------------------
  |[Sensors]                                               |
  | colorL:  ` + progressBar(float64(bot.colorSensorL.intensity()), 11, 10, false) + ` [w]                               |
  | colorR:  ` + progressBar(float64(bot.colorSensorR.intensity()), 11, 10, false) + ` [b]       ⇡⇡     [Extra]          |
  | dist:    ` + progressBar(float64(1000 - bot.ultrasonicSensor.distance()), 11, 1000, false) + ` [w]      ┌__┐     tile:  line     |
  |                             ⇠ ╭──╮ ⇢   stage: transit  |
  |                              ╭│__│╮                    |
  |[Motors]                      ╰└──┘╯                    |
  | driveL:  ` + progressBar(float64(bot.gyroSensor.angle()), 5, 10, true) + `|` + progressBar(float64(bot.gyroSensor.angle()), 5, 10, false) + `           ⇣⇣                      |
  | driveR:  ` + progressBar(float64(bot.gyroSensor.angle()), 5, 10, true) + `|` + progressBar(float64(bot.gyroSensor.angle()), 5, 10, false) + `                                   |
  | claw:    ░░░░░|░░░░░                                   |
  ---------------------------------------------------------`)
}

func progressBar(percent float64, length int, number int, reversed bool) string {
	output := ""
  for i := 0; i < length; i++ {
    if reversed {
      if i < int(percent * float64(length) / float64(number)) {
        output += "░"
      } else {
        output += "█"
      }
    } else {
      if i < int(percent * float64(length) / float64(number)) {
        output += "█"
      } else {
        output += "░"
      }
    }
  }
	return output
}

func replaceArrow(direction int, enabled bool) string {
  if enabled {
    if direction == 0 { return "↑" }
    if direction == 1 { return "→" }
    if direction == 2 { return "↓" }
    if direction == 3 { return "←" }
  } else {
    if direction == 0 { return "⇡" }
    if direction == 1 { return "⇢" }
    if direction == 2 { return "⇣" }
    if direction == 3 { return "⇠" }
  }
  return " "
}
