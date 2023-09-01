// THIS FILE IS AUTO-GENERATED
// 20230831_yz 修改：em6增加PM2.5
package service

import (
	"github.com/twxstar/hap/characteristic"
)

const TypeAirQualitySensor = "8D"

type AirQualitySensor struct {
	*S

	AirQuality *characteristic.AirQuality

	//新增_PM2.5
	PM2_5Density *characteristic.PM2_5Density
}

func NewAirQualitySensor() *AirQualitySensor {
	s := AirQualitySensor{}
	s.S = New(TypeAirQualitySensor)

	s.AirQuality = characteristic.NewAirQuality()
	s.AddC(s.AirQuality.C)

	//新增_PM2.5
	s.PM2_5Density = characteristic.NewPM2_5Density()
	s.AddC(s.PM2_5Density.C)

	return &s
}
