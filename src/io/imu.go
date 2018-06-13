package io

import "encoding/binary"

const GYRO_ZOUT_H = 0x47
const GYRO_ZOUT_L = 0x48

type IMU struct {
  Address uint8
  i2cDevice *I2C
}

func (imu IMU) New() IMU {
  imu.i2cDevice, _ = NewI2C(imu.Address, 1)
  return imu
}

func (imu IMU) ReadGyro() int {
  gyroHigh, _ := imu.i2cDevice.ReadRegU8(GYRO_ZOUT_H)
  gyroLow, _ := imu.i2cDevice.ReadRegU8(GYRO_ZOUT_L)

  return int(binary.BigEndian.Uint16([]byte{gyroHigh, gyroLow}))
}

func (imu IMU) Cleanup() {
  imu.i2cDevice.Close()
}