package service

import (
	"github.com/twxstar/hap/characteristic"
)

type Heater_New struct {
	*S

	Active                      *characteristic.Active
	CurrentHeaterCoolerState    *characteristic.CurrentHeaterCoolerState
	TargetHeaterCoolerState     *characteristic.TargetHeaterCoolerState
	CurrentTemperature          *characteristic.CurrentTemperature
	HeatingThresholdTemperature *characteristic.HeatingThresholdTemperature
}

func NewHeater_New() *Heater_New {
	s := Heater_New{}
	s.S = New(TypeHeaterCooler)

	s.Active = characteristic.NewActive()
	s.AddC(s.Active.C)

	s.CurrentHeaterCoolerState = characteristic.NewCurrentHeaterCoolerState()
	s.CurrentHeaterCoolerState.ValidVals = []int{
		characteristic.CurrentHeaterCoolerStateInactive,
		characteristic.CurrentHeaterCoolerStateIdle,
		characteristic.CurrentHeaterCoolerStateHeating,
	}
	s.AddC(s.CurrentHeaterCoolerState.C)

	s.TargetHeaterCoolerState = characteristic.NewTargetHeaterCoolerState()
	s.TargetHeaterCoolerState.ValidVals = []int{
		characteristic.TargetHeaterCoolerStateHeat,
	}
	s.AddC(s.TargetHeaterCoolerState.C)

	s.CurrentTemperature = characteristic.NewCurrentTemperature()
	s.AddC(s.CurrentTemperature.C)

	s.HeatingThresholdTemperature = characteristic.NewHeatingThresholdTemperature()
	s.AddC(s.HeatingThresholdTemperature.C)

	return &s
}
