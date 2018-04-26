
// Initial values change as you test

var kp = 2.5;
var ki = 0.05;
var kd = 5;
var tp = 25;
var offset = 45; // average
var integral = 0;
var lasterror = 0;
var derivative = 0;

// right motor => B, left motor => D
// right sensor => 1, left sensor => 2

// INFITIE LOOP


while(button) {

  var Lightvalue = (bot.colorSensorR.intensity() + bot.colorSensorL.intensity()) / 2;
  var error = Lightvalue - offset;
  integral = integral + error;
  derivative = error - lasterror;
  var turn = (kp * error) + (ki * integral) + (kd * derivative);

  var powerB = tp + turn;
  var powerC = tp - turn;

  lasterror = error;

  bot.motorR.runForever(25) // same with right and left
  bot.motorL.runForever(25) // same with right and left

  // INFITIE for now
  if (bot.colorSensorR.intensity() = 90 && bot.colorSensorL.intensity() = 90) {
    break;
  };
}
