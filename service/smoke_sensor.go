// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/twxstar/hap/characteristic"
)

const TypeSmokeSensor = "87"

type SmokeSensor struct {
	*S

	SmokeDetected *characteristic.SmokeDetected
}

func NewSmokeSensor() *SmokeSensor {
	s := SmokeSensor{}
	s.S = New(TypeSmokeSensor)

	s.SmokeDetected = characteristic.NewSmokeDetected()
	s.AddC(s.SmokeDetected.C)

	return &s
}
