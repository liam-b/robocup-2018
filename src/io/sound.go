package main

import "strconv"
import "time"

type Speaker struct {
  path string
  device Device
}

func (speaker Speaker) init() Speaker {
  speaker.path = "/sys/devices/platform/snd-legoev3/"
  speaker.device = Device{path: speaker.path}
  return speaker
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

func (speaker Speaker) song(song []int, volume int) {
  if volume != 0 {
    speaker.device.set("volume", strconv.Itoa(volume))
  }

  for note := 0; note < len(song); note += 2 {
    speaker.device.set("tone", strconv.Itoa(song[note]) + " " + strconv.Itoa(song[note + 1]))
    time.Sleep(time.Nanosecond * time.Duration(song[note + 1] * 1000000))
  }
}

func (speaker Speaker) beep(volume int) {
  speaker.play(300, 200, volume)
}