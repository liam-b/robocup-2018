package io

import "encoding/binary"

const GYRO_ZOUT_H = 0x47
const GYRO_ZOUT_L = 0x48

type IMU struct {
  Address uint8
  i2cDevice *I2C

  cachedValue int
  hasCached bool
}

func (imu IMU) New() IMU {
  imu.i2cDevice, _ = NewI2C(imu.Address, 1)
  return imu
}

func (imu *IMU) ResetCache() {
  imu.cachedValue = 0
  imu.hasCached = false
}

func (imu *IMU) ReadGyro() int {
  if imu.hasCached { return imu.cachedValue }
  result := imu.getGyroValue()
  imu.hasCached = true
  imu.cachedValue = result

  return result
}

func (imu IMU) Cleanup() {
  imu.i2cDevice.Close()
}

func (imu IMU) getGyroValue() int {
  gyroHigh, _ := imu.i2cDevice.ReadRegU8(GYRO_ZOUT_H)
  gyroLow, _ := imu.i2cDevice.ReadRegU8(GYRO_ZOUT_L)
  gyroValue := int16(binary.BigEndian.Uint16([]byte{gyroHigh, gyroLow}))

  return int(float64(gyroValue) / 300.0)
}