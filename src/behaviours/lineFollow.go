package main

func followLine() {
  intensityL := bot.colorSensorL.intensity()
  intensityR := bot.colorSensorR.intensity()

  if (intensityL < 72) {
    go bot.motorR.runForever(440)
    go bot.motorL.runForever(230)
  }

  if (intensityR < 72) {
    go bot.motorR.runForever(230)
    go bot.motorL.runForever(440)
  }

  if (intensityL < 16) {
    go bot.motorR.runForever(500)
    go bot.motorL.runForever(80)
  }

  if (intensityR < 16) {
    go bot.motorR.runForever(80)
    go bot.motorL.runForever(500)
  }

  if (intensityR > 60 && intensityL > 60) {
    go bot.motorR.runForever(300)
    go bot.motorL.runForever(300)
  }
}