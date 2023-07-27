// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/twxstar/hap/characteristic"
)

const TypeCarbonDioxideSensor = "97"

type CarbonDioxideSensor struct {
	*S

	CarbonDioxideDetected *characteristic.CarbonDioxideDetected
	CarbonDioxideLevel    *characteristic.CarbonDioxideLevel
}

func NewCarbonDioxideSensor() *CarbonDioxideSensor {
	s := CarbonDioxideSensor{}
	s.S = New(TypeCarbonDioxideSensor)

	s.CarbonDioxideDetected = characteristic.NewCarbonDioxideDetected()
	s.AddC(s.CarbonDioxideDetected.C)

	s.CarbonDioxideLevel = characteristic.NewCarbonDioxideLevel()
	s.AddC(s.CarbonDioxideLevel.C)

	return &s
}
