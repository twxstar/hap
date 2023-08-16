// 20230808_yz 新增
package accessory

import (
	"github.com/twxstar/hap/service"
)

type CarbonMonoxideSensor struct {
	*A
	CarbonMonoxideSensor *service.CarbonMonoxideSensor
}

// NewCarbonMonoxideSensor returns a CarbonMonoxideSensor which implements model.CarbonMonoxideSensor.
func NewCarbonMonoxideSensor(info Info) *CarbonMonoxideSensor {
	a := CarbonMonoxideSensor{}
	a.A = New(info, TypeSensor)

	a.CarbonMonoxideSensor = service.NewCarbonMonoxideSensor()
	a.AddS(a.CarbonMonoxideSensor.S)

	return &a
}
