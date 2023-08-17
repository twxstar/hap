// THIS FILE IS AUTO-GENERATED
// 20230808_yz 新增
package service

import (
	"github.com/twxstar/hap/characteristic"
)

const TypeCarbonMonoxideSensor = "7F"

type CarbonMonoxideSensor struct {
	*S

	CarbonMonoxideDetected *characteristic.CarbonMonoxideDetected
	CarbonMonoxideLevel    *characteristic.CarbonMonoxideLevel
}

func NewCarbonMonoxideSensor() *CarbonMonoxideSensor {
	s := CarbonMonoxideSensor{}
	s.S = New(TypeCarbonMonoxideSensor)

	s.CarbonMonoxideDetected = characteristic.NewCarbonMonoxideDetected()
	s.AddC(s.CarbonMonoxideDetected.C)

	s.CarbonMonoxideLevel = characteristic.NewCarbonMonoxideLevel()
	s.AddC(s.CarbonMonoxideLevel.C)

	return &s
}
