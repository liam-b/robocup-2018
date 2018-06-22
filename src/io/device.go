package main

import "strings"
import "errors"

type Device struct {
  path string
}

func (device Device) get(attribute string) string {
  return read(device.path + attribute)
}

func (device Device) set(attribute string, data string) {
  write(device.path + attribute, data)
}

type IndexedDevice struct {
  port string
  path string
}

func (device *IndexedDevice) findDeviceFromPort() {
  files := list(device.path)
  foundDevice := false

  for _, file := range files {
    if strings.Contains(read(device.path + file + "/address"), device.port) {
      device.path = device.path + file + "/"
      foundDevice = true
      break
    }
  }

  if !foundDevice {
    check(errors.New("no device in address " + device.port))
  }
}

func (device IndexedDevice) get(attribute string) string {
  return read(device.path + attribute)
}

func (device IndexedDevice) set(attribute string, data string) {
  write(device.path + attribute, data)
}

type ManualDevice struct {
  port string
  path string

  connection string
  name string
}

func (device *ManualDevice) findDeviceFromPort() {
  files := list(device.path)
  foundInitialDevice := false

  for _, file := range files {
    if strings.Contains(read(device.path + file + "/address"), device.port) {
      device.path = device.path + file + "/"
      foundInitialDevice = true
      break
    }
  }

  if !foundInitialDevice {
    portPath := LEGO_PORT + PORT[device.port] + "/"
    write(portPath + "mode", device.connection)
    write(portPath + "set_device", device.name)

    files := list(device.path)
    foundDevice := false

    for _, file := range files {
      if strings.Contains(read(device.path + file + "/address"), device.port) {
        device.path = device.path + file + "/"
        foundDevice = true
        break
      }
    }

    if !foundDevice {
      check(errors.New("no device in address " + device.port))
    }
  }
}

func (device ManualDevice) get(attribute string) string {
  return read(device.path + attribute)
}

func (device ManualDevice) set(attribute string, data string) {
  write(device.path + attribute, data)
}