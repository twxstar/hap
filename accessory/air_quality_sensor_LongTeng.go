package accessory

// 20230831_yz 新增：龙腾空气质量

import (
	"github.com/twxstar/hap/service"
)

type AirQualitySensor_LongTeng struct {
	*A
	AirQualitySensor *service.AirQualitySensor_LongTeng
}

// NewAirQualitySensor_LongTeng returns a AirQualitySensor_LongTeng which implements model.AirQualitySensor_LongTeng.
func NewAirQualitySensor_LongTeng(info Info) *AirQualitySensor_LongTeng {
	a := AirQualitySensor_LongTeng{}
	a.A = New(info, TypeSensor)

	a.AirQualitySensor = service.NewAirQualitySensor_LongTeng()
	a.AddS(a.AirQualitySensor.S)

	return &a
}
