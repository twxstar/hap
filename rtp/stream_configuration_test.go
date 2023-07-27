package rtp

import (
	"fmt"
	"github.com/twxstar/hap/characteristic"
	"github.com/twxstar/hap/tlv8"
	"testing"
)

func TestSelectedStreamConfiguration(t *testing.T) {
	c := characteristic.NewSelectedStreamConfiguration()
	c.Val = "ARUCAQABEHW8tiJ9E0F4tLlvOURdFCc="

	b := c.Value()

	var cfg StreamConfiguration
	err := tlv8.Unmarshal(b, &cfg)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%+v", cfg)
}
