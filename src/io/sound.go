package main

import "strconv"
import "time"

type Speaker struct {
  playSound bool

  path string
  device Device
}

func (speaker Speaker) new() Speaker {
  speaker.path = SOUND_PATH
  speaker.device = Device{path: speaker.path}
  return speaker
}

func (speaker Speaker) play(tone int, time int, volume int) {
  if speaker.playSound {
    if volume != 0 {
      speaker.device.set("volume", strconv.Itoa(volume))
    }
    speaker.device.set("tone", strconv.Itoa(tone) + " " + strconv.Itoa(time))
  }
}

func (speaker Speaker) volume(volume int) {
  if speaker.playSound {
    speaker.device.set("volume", strconv.Itoa(volume))
  }
}

func (speaker Speaker) song(song []int, delay int, volume int) {
  if speaker.playSound {
    if volume != 0 {
      speaker.device.set("volume", strconv.Itoa(volume))
    }

    for note := 0; note < len(song); note ++ {
      speaker.device.set("tone", strconv.Itoa(song[note]))
      time.Sleep(time.Millisecond * time.Duration(delay))
    }
    speaker.device.set("tone", "0")
  }
}

func (speaker Speaker) beep(volume int) {
  if speaker.playSound {
    speaker.play(300, 200, volume)
  }
}