// THIS FILE IS AUTO-GENERATED
// 20230831_yz 新增：龙腾空气质量
package service

import (
	"github.com/twxstar/hap/characteristic"
)

const TypeAirQualitySensor_LongTeng = "8D"

type AirQualitySensor_LongTeng struct {
	*S

	AirQuality *characteristic.AirQuality
	//新增_挥发性有机物浓度
	VOCDensity *characteristic.VOCDensity
	//新增_PM2.5
	PM2_5Density *characteristic.PM2_5Density
}

func NewAirQualitySensor_LongTeng() *AirQualitySensor_LongTeng {
	s := AirQualitySensor_LongTeng{}
	s.S = New(TypeAirQualitySensor_LongTeng)

	s.AirQuality = characteristic.NewAirQuality()
	s.AddC(s.AirQuality.C)

	//新增_挥发性有机物浓度
	s.VOCDensity = characteristic.NewVOCDensity()
	s.AddC(s.VOCDensity.C)

	//新增_PM2.5
	s.PM2_5Density = characteristic.NewPM2_5Density()
	s.AddC(s.PM2_5Density.C)

	return &s
}
