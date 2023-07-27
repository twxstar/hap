// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/twxstar/hap/characteristic"
)

const TypeContactSensor = "80"

type ContactSensor struct {
	*S

	ContactSensorState *characteristic.ContactSensorState
}

func NewContactSensor() *ContactSensor {
	s := ContactSensor{}
	s.S = New(TypeContactSensor)

	s.ContactSensorState = characteristic.NewContactSensorState()
	s.AddC(s.ContactSensorState.C)

	return &s
}
