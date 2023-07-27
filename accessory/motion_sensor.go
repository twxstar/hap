package accessory

import (
	"github.com/twxstar/hap/service"
)

type MotionSensor struct {
	*A
	MotionSensor *service.MotionSensor
}

// NewMotionSensor returns a MotionSensor which implements model.MotionSensor.
func NewMotionSensor(info Info) *MotionSensor {
	a := MotionSensor{}
	a.A = New(info, TypeSensor)

	a.MotionSensor = service.NewMotionSensor()
	a.AddS(a.MotionSensor.S)

	return &a
}
