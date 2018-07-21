package main

import (
	"io/ioutil"
  "strconv"
  "strings"
)

const (
	I2C_SLAVE = 0x0703
)

type I2C struct {
	address uint8
}

func NewI2C(addr uint8, bus int) (*I2C, error) {
	v := &I2C{address: addr}
	return v, nil
}

func (v *I2C) Close() error {
  return nil
}

func (v *I2C) ReadRegU8(reg byte) (byte, error) {
	dat, _ := ioutil.ReadFile("mock/fs/i2c/" + strconv.Itoa(int(v.address)))
  lines := strings.Split(string(dat), "\n")
  result, _ := strconv.Atoi(lines[reg])
	return byte(result), nil
}

func (v *I2C) WriteRegU8(reg byte, value byte) error {
	dat, _ := ioutil.ReadFile("mock/fs/i2c/" + strconv.Itoa(int(v.address)))
  lines := strings.Split(string(dat), "\n")
  lines[reg] = strconv.Itoa(int(value))
  output := strings.Join(lines, "\n")
  ioutil.WriteFile("mock/fs/i2c/" + strconv.Itoa(int(v.address)), []byte(output), 0644)
	return nil
}