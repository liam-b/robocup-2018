# Robocup Junior 2018 - I am Groot
**Liam Brennan** and **Parsa Cheeky Boi** from Box Hill High School

## Planning
As with all projects we started with some plans for how the end robot should turn out. While we haven't followed this plan exactly throughout this project, much of what we decided at the beginning has made it through all the way to the end.

### Hardware
Hardware was the first big choice. Last time we attempted rescue, we discovered and used something called [ev3dev](www.ev3dev.org). In short ev3dev lets you expose the underlying linux operating system of the ev3s which allows you to use many more languages with much more control over the robot. While this is much better than a standard ev3, it is still limited by the processing power of the ev3's hardware and as such is a slow and restrictive control unit.

While still in the planning phase we also researched into something called a [brickpi](https://www.dexterindustries.com/brickpi/) which in essence lets you use a raspberry pi to control lego sensors and motors (and much more) in the same way an ev3 would. This option was quite pricy but with the speed and computation power benifits of a raspberry pi, this seemed like an great option.

In the end, between an ev3 running ev3dev and a raspberry pi with a brickpi shield we chose the brickpi for its speed and flexibility. As for motors and sensors, for the most part we chose to use standard ev3 parts with the exception of a four-in-one i2c [IMU sensor](https://www.pakronics.com.au/products/grove-imu-9dof-v2-0-ss101020080) (inertial measurement unit sensor) which includes a gyroscope, accelerometer, compass and temperature sensor.

### Software
Another major choice made at the start of this project was what programming language to use. Previously we had used python (rescue) and javascript (soccer) but with the idea of speed in mind we decided to have a go with using a compiled language. This not only lead to our program running many times faster than what we could have ever hope of before, but was also a great learning experience. We chose to use [golang](https://golang.org/), an open source language made by Google. While it was tough at first having to learn all the different nuances of go, in the end it was well worth it as now we have a great program and much more knowledge of compiled languages than when we started.
