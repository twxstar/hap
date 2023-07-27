package rtp

import (
	"github.com/twxstar/hap/characteristic"
	"github.com/twxstar/hap/tlv8"
	"testing"
)

func TestSetupEndpoints(t *testing.T) {
	c := characteristic.NewSetupEndpoints()
	c.Val = "ARBz21VuCupGZre3A62biD8XAxkBAQACDDE5Mi4xNjguMC4xMwMC+OUEAurRBCUCEPpLBUWQEzkfFiGd1qkieqoDDi8cIMO0Vl1+kegzGgnpAQEABSUCEJQ27Ze9EEmuxcIVPhDEs68DDlaHwww6f6d5+NSClT7TAQEA"

	b := c.Value()

	var setup SetupEndpoints
	err := tlv8.Unmarshal(b, &setup)
	if err != nil {
		t.Fatal(err)
	}
}
