package accessory

import (
	"github.com/twxstar/hap/service"
)

type DimmerLightbulb struct {
	*A
	Lightbulb *service.DimmerLightbulb
}

// NewLightbulb returns an light bulb accessory.
func NewDimmerLightbulb(info Info) *DimmerLightbulb {
	a := DimmerLightbulb{}
	a.A = New(info, TypeLightbulb)

	a.Lightbulb = service.NewDimmerLightbulb()
	a.AddS(a.Lightbulb.S)

	return &a
}
