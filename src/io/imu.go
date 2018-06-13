package io

type IMU struct {
  Address uint8
  i2cDevice *I2C
}

func (imu IMU) New() IMU {
  imu.i2cDevice, _ = NewI2C(imu.Address, 1)
  defer imu.i2cDevice.Close()
  return imu
}

func (imu IMU) GetByte(register byte) byte {
  result, _ := imu.i2cDevice.ReadRegU8(register)
  return result
}