package main

func FollowLine() {
  intensityL := bot.colorSensorL.Intensity()
  intensityR := bot.colorSensorR.Intensity()

  if (intensityL < 72) {
    go bot.motorR.RunForever(440)
    go bot.motorL.RunForever(230)
  }

  if (intensityR < 72) {
    go bot.motorR.RunForever(230)
    go bot.motorL.RunForever(440)
  }

  if (intensityL < 16) {
    go bot.motorR.RunForever(500)
    go bot.motorL.RunForever(80)
  }

  if (intensityR < 16) {
    go bot.motorR.RunForever(80)
    go bot.motorL.RunForever(500)
  }

  if (intensityR > 60 && intensityL > 60) {
    go bot.motorR.RunForever(300)
    go bot.motorL.RunForever(300)
  }
}