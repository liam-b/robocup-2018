package main

import "strconv"

type Speaker struct {
  path string
  device Device
}

func (speaker *Speaker) init() {
  speaker.path = "/sys/devices/platform/snd-legoev3/"
  speaker.device = Device{path: speaker.path}
}

func (speaker Speaker) play(tone int, time int, volume int) {
  if volume != 0 {
    speaker.device.set("volume", strconv.Itoa(volume))
  }
  speaker.device.set("tone", strconv.Itoa(tone) + " " + strconv.Itoa(time))
}

func (speaker Speaker) volume(volume int) {
  speaker.device.set("volume", strconv.Itoa(volume))
}